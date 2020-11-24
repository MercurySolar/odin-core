package proof

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/store/rootmulti"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/gorilla/mux"
	"github.com/tendermint/iavl"
	tmbytes "github.com/tendermint/tendermint/libs/bytes"
	rpcclient "github.com/tendermint/tendermint/rpc/client"

	"github.com/bandprotocol/chain/pkg/obi"
	clientcmn "github.com/bandprotocol/chain/x/oracle/client/common"
	"github.com/bandprotocol/chain/x/oracle/types"
)

var (
	relayArguments       abi.Arguments
	verifyArguments      abi.Arguments
	verifyCountArguments abi.Arguments
)

const (
	RequestIDTag = "requestID"
)

func init() {
	err := json.Unmarshal(relayFormat, &relayArguments)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(verifyFormat, &verifyArguments)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(verifyCountFormat, &verifyCountArguments)
	if err != nil {
		panic(err)
	}
}

type BlockRelayProof struct {
	MultiStoreProof        MultiStoreProof        `json:"multiStoreProof"`
	BlockHeaderMerkleParts BlockHeaderMerkleParts `json:"blockHeaderMerkleParts"`
	Signatures             []TMSignature          `json:"signatures"`
}

func (blockRelay *BlockRelayProof) encodeToEthData() ([]byte, error) {
	parseSignatures := make([]TMSignatureEthereum, len(blockRelay.Signatures))
	for i, sig := range blockRelay.Signatures {
		parseSignatures[i] = sig.encodeToEthFormat()
	}
	return relayArguments.Pack(
		blockRelay.MultiStoreProof.encodeToEthFormat(),
		blockRelay.BlockHeaderMerkleParts.encodeToEthFormat(),
		parseSignatures,
	)
}

type OracleDataProof struct {
	RequestPacket  types.OracleRequestPacketData  `json:"requestPacket"`
	ResponsePacket types.OracleResponsePacketData `json:"responsePacket"`
	Version        uint64                         `json:"version"`
	MerklePaths    []IAVLMerklePath               `json:"merklePaths"`
}

func (o *OracleDataProof) encodeToEthData(blockHeight uint64) ([]byte, error) {
	parsePaths := make([]IAVLMerklePathEthereum, len(o.MerklePaths))
	for i, path := range o.MerklePaths {
		parsePaths[i] = path.encodeToEthFormat()
	}
	return verifyArguments.Pack(
		big.NewInt(int64(blockHeight)),
		transformRequestPacket(o.RequestPacket),
		transformResponsePacket(o.ResponsePacket),
		big.NewInt(int64(o.Version)),
		parsePaths,
	)
}

type RequestsCountProof struct {
	Count       uint64           `json:"count"`
	Version     uint64           `json:"version"`
	MerklePaths []IAVLMerklePath `json:"merklePaths"`
}

func (o *RequestsCountProof) encodeToEthData(blockHeight uint64) ([]byte, error) {
	parsePaths := make([]IAVLMerklePathEthereum, len(o.MerklePaths))
	for i, path := range o.MerklePaths {
		parsePaths[i] = path.encodeToEthFormat()
	}
	return verifyCountArguments.Pack(
		big.NewInt(int64(blockHeight)),
		big.NewInt(int64(o.Count)),
		big.NewInt(int64(o.Version)),
		parsePaths,
	)
}

type JsonProof struct {
	BlockHeight     uint64          `json:"blockHeight"`
	OracleDataProof OracleDataProof `json:"oracleDataProof"`
	BlockRelayProof BlockRelayProof `json:"blockRelayProof"`
}

type JsonMultiProof struct {
	BlockHeight          uint64            `json:"blockHeight"`
	OracleDataMultiProof []OracleDataProof `json:"oracleDataMultiProof"`
	BlockRelayProof      BlockRelayProof   `json:"blockRelayProof"`
}

type JsonRequestsCountProof struct {
	BlockHeight     uint64             `json:"blockHeigh"`
	CountProof      RequestsCountProof `json:"countProof"`
	BlockRelayProof BlockRelayProof    `json:"blockRelayProof"`
}

type Proof struct {
	JsonProof     JsonProof        `json:"jsonProof"`
	EVMProofBytes tmbytes.HexBytes `json:"evmProofBytes"`
}

type MultiProof struct {
	JsonProof     JsonMultiProof   `json:"jsonProof"`
	EVMProofBytes tmbytes.HexBytes `json:"evmProofBytes"`
}

type CountProof struct {
	JsonProof     JsonRequestsCountProof `json:"jsonProof"`
	EVMProofBytes tmbytes.HexBytes       `json:"evmProofBytes"`
}

func GetProofHandlerFn(cliCtx context.CLIContext, route string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}
		height := &ctx.Height
		if ctx.Height == 0 {
			height = nil
		}

		// Parse Request ID
		vars := mux.Vars(r)
		intRequestID, err := strconv.ParseUint(vars[RequestIDTag], 10, 64)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		requestID := types.RequestID(intRequestID)

		// Get Request and proof
		bz, _, err := ctx.Query(fmt.Sprintf("custom/%s/%s/%d", route, types.QueryRequests, requestID))
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		var qResult types.QueryResult
		if err := json.Unmarshal(bz, &qResult); err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		if qResult.Status != http.StatusOK {
			clientcmn.PostProcessQueryResponse(w, ctx, bz)
			return
		}
		var request types.QueryRequestResult
		if err := ctx.Codec.UnmarshalJSON(qResult.Result, &request); err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		if request.Result == nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, "Result has not been resolved")
			return
		}

		commit, err := ctx.Client.Commit(height)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		value, iavlProof, multiStoreProof, err := getProofsByKey(
			ctx,
			types.ResultStoreKey(requestID),
			rpcclient.ABCIQueryOptions{Height: commit.Height - 1, Prove: true},
		)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		if iavlProof.Proof == nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, "Proof has not been ready.")
			return
		}
		eventHeight := iavlProof.Proof.Leaves[0].Version
		signatures, err := GetSignaturesAndPrefix(&commit.SignedHeader)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		blockRelay := BlockRelayProof{
			MultiStoreProof:        GetMultiStoreProof(*multiStoreProof),
			BlockHeaderMerkleParts: GetBlockHeaderMerkleParts(ctx.Codec, commit.Header),
			Signatures:             signatures,
		}

		type result struct {
			Req types.OracleRequestPacketData
			Res types.OracleResponsePacketData
		}
		var rs result
		obi.MustDecode(value, &rs)

		oracleData := OracleDataProof{
			RequestPacket:  rs.Req,
			ResponsePacket: rs.Res,
			Version:        uint64(eventHeight),
			MerklePaths:    GetIAVLMerklePaths(iavlProof),
		}

		// Calculate byte for proofbytes
		var relayAndVerifyArguments abi.Arguments
		format := `[{"type":"bytes"},{"type":"bytes"}]`
		err = json.Unmarshal([]byte(format), &relayAndVerifyArguments)
		if err != nil {
			panic(err)
		}

		blockRelayBytes, err := blockRelay.encodeToEthData()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		oracleDataBytes, err := oracleData.encodeToEthData(uint64(commit.Height))
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		evmProofBytes, err := relayAndVerifyArguments.Pack(blockRelayBytes, oracleDataBytes)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		rest.PostProcessResponse(w, ctx, Proof{
			JsonProof: JsonProof{
				BlockHeight:     uint64(commit.Height),
				OracleDataProof: oracleData,
				BlockRelayProof: blockRelay,
			},
			EVMProofBytes: evmProofBytes,
		})
	}
}

func GetMutiProofHandlerFn(cliCtx context.CLIContext, route string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestIDs := r.URL.Query()["id"]
		ctx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}
		height := &ctx.Height
		if ctx.Height == 0 {
			height = nil
		}

		commit, err := ctx.Client.Commit(height)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		signatures, err := GetSignaturesAndPrefix(&commit.SignedHeader)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		blockRelay := BlockRelayProof{
			BlockHeaderMerkleParts: GetBlockHeaderMerkleParts(ctx.Codec, commit.Header),
			Signatures:             signatures,
		}

		oracleDataBytesList := make([][]byte, len(requestIDs))
		oracleDataList := make([]OracleDataProof, len(requestIDs))

		isFirstRequest := true
		for idx, requestID := range requestIDs {
			intRequestID, err := strconv.ParseUint(requestID, 10, 64)
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
				return
			}
			requestID := types.RequestID(intRequestID)
			bz, _, err := ctx.Query(fmt.Sprintf("custom/%s/%s/%d", route, types.QueryRequests, requestID))
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
				return
			}
			var qResult types.QueryResult
			if err := json.Unmarshal(bz, &qResult); err != nil {
				rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
				return
			}
			if qResult.Status != http.StatusOK {
				clientcmn.PostProcessQueryResponse(w, ctx, bz)
				return
			}
			var request types.QueryRequestResult
			if err := ctx.Codec.UnmarshalJSON(qResult.Result, &request); err != nil {
				rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
				return
			}
			if request.Result == nil {
				rest.WriteErrorResponse(w, http.StatusNotFound, "Result has not been resolved")
				return
			}

			resp, err := ctx.Client.ABCIQueryWithOptions(
				"/store/oracle/key",
				types.ResultStoreKey(requestID),
				rpcclient.ABCIQueryOptions{Height: commit.Height - 1, Prove: true},
			)
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
				return
			}

			proof := resp.Response.GetProof()
			if proof == nil {
				rest.WriteErrorResponse(w, http.StatusInternalServerError, "Proof not found")
				return
			}

			ops := proof.GetOps()
			if ops == nil {
				rest.WriteErrorResponse(w, http.StatusInternalServerError, "proof ops not found")
				return
			}

			var iavlProof iavl.ValueOp
			var multiStoreProof rootmulti.MultiStoreProofOp
			for _, op := range ops {
				switch op.GetType() {
				case "iavl:v":
					err := ctx.Codec.UnmarshalBinaryLengthPrefixed(op.GetData(), &iavlProof)
					if err != nil {
						rest.WriteErrorResponse(w, http.StatusInternalServerError,
							fmt.Sprintf("iavl: %s", err.Error()),
						)
						return
					}
					if iavlProof.Proof == nil {
						rest.WriteErrorResponse(w, http.StatusNotFound, "Proof has not been ready.")
						return
					}

					eventHeight := iavlProof.Proof.Leaves[0].Version
					resValue := resp.Response.GetValue()

					type result struct {
						Req types.OracleRequestPacketData
						Res types.OracleResponsePacketData
					}
					var rs result
					obi.MustDecode(resValue, &rs)

					oracleData := OracleDataProof{
						RequestPacket:  rs.Req,
						ResponsePacket: rs.Res,
						Version:        uint64(eventHeight),
						MerklePaths:    GetIAVLMerklePaths(&iavlProof),
					}
					oracleDataBytes, err := oracleData.encodeToEthData(uint64(commit.Height))
					if err != nil {
						rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
						return
					}
					// Append oracle data proof to list
					oracleDataBytesList[idx] = oracleDataBytes
					oracleDataList[idx] = oracleData
				case "multistore":
					if isFirstRequest {
						// Only create multi store proof in the first request.
						isFirstRequest = false
						mp, err := rootmulti.MultiStoreProofOpDecoder(op)
						multiStoreProof = mp.(rootmulti.MultiStoreProofOp)
						if err != nil {
							rest.WriteErrorResponse(w, http.StatusInternalServerError,
								fmt.Sprintf("multiStore: %s", err.Error()),
							)
							return
						}
						blockRelay.MultiStoreProof = GetMultiStoreProof(multiStoreProof)
					}
				case "iavl:a":
					rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf(
						"Proof of #%d is unavailable please wait on the next block", requestID,
					))
					return
				default:
					rest.WriteErrorResponse(w, http.StatusInternalServerError,
						fmt.Sprintf("Unknown proof type %s", op.GetType()),
					)
					return
				}
			}
		}

		blockRelayBytes, err := blockRelay.encodeToEthData()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		// Calculate byte for MultiProofbytes
		var relayAndVerifyArguments abi.Arguments
		format := `[{"type":"bytes"},{"type":"bytes[]"}]`
		err = json.Unmarshal([]byte(format), &relayAndVerifyArguments)
		if err != nil {
			panic(err)
		}

		evmProofBytes, err := relayAndVerifyArguments.Pack(blockRelayBytes, oracleDataBytesList)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		rest.PostProcessResponse(w, ctx, MultiProof{
			JsonProof: JsonMultiProof{
				BlockHeight:          uint64(commit.Height),
				OracleDataMultiProof: oracleDataList,
				BlockRelayProof:      blockRelay,
			},
			EVMProofBytes: evmProofBytes,
		})
	}
}

func GetRequestsCountProofHandlerFn(cliCtx context.CLIContext, route string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}
		height := &ctx.Height
		if ctx.Height == 0 {
			height = nil
		}

		commit, err := ctx.Client.Commit(height)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		value, iavlProof, multiStoreProof, err := getProofsByKey(
			ctx,
			types.RequestCountStoreKey,
			rpcclient.ABCIQueryOptions{Height: commit.Height - 1, Prove: true},
		)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		if iavlProof.Proof == nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, "Proof has not been ready.")
			return
		}

		eventHeight := iavlProof.Proof.Leaves[0].Version

		// Produce block relay proof
		signatures, err := GetSignaturesAndPrefix(&commit.SignedHeader)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		blockRelay := BlockRelayProof{
			MultiStoreProof:        GetMultiStoreProof(*multiStoreProof),
			BlockHeaderMerkleParts: GetBlockHeaderMerkleParts(ctx.Codec, commit.Header),
			Signatures:             signatures,
		}

		// Parse requests count
		var rs int64
		ctx.Codec.MustUnmarshalBinaryLengthPrefixed(value, &rs)

		requestsCountProof := RequestsCountProof{
			Count:       uint64(rs),
			Version:     uint64(eventHeight),
			MerklePaths: GetIAVLMerklePaths(iavlProof),
		}

		// Calculate byte for proofbytes
		var relayAndVerifyCountArguments abi.Arguments
		format := `[{"type":"bytes"},{"type":"bytes"}]`
		err = json.Unmarshal([]byte(format), &relayAndVerifyCountArguments)
		if err != nil {
			panic(err)
		}

		blockRelayBytes, err := blockRelay.encodeToEthData()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		requestsCountBytes, err := requestsCountProof.encodeToEthData(uint64(commit.Height))
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		evmProofBytes, err := relayAndVerifyCountArguments.Pack(blockRelayBytes, requestsCountBytes)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		rest.PostProcessResponse(w, ctx, CountProof{
			JsonProof: JsonRequestsCountProof{
				BlockHeight:     uint64(commit.Height),
				CountProof:      requestsCountProof,
				BlockRelayProof: blockRelay,
			},
			EVMProofBytes: evmProofBytes,
		})
	}
}

func getProofsByKey(ctx context.CLIContext, key []byte, queryOptions rpcclient.ABCIQueryOptions) ([]byte, *iavl.ValueOp, *rootmulti.MultiStoreProofOp, error) {
	resp, err := ctx.Client.ABCIQueryWithOptions(
		"/store/oracle/key",
		key,
		queryOptions,
	)
	if err != nil {
		return nil, nil, nil, err
	}

	proof := resp.Response.GetProof()
	if proof == nil {
		return nil, nil, nil, fmt.Errorf("Proof not found")
	}

	ops := proof.GetOps()
	if ops == nil {
		return nil, nil, nil, fmt.Errorf("Proof ops not found")
	}

	// Extract iavl proof and multi store proof
	var iavlProof iavl.ValueOp
	var multiStoreProof rootmulti.MultiStoreProofOp
	for _, op := range ops {

		opType := op.GetType()

		if opType == "iavl:v" {
			err := ctx.Codec.UnmarshalBinaryLengthPrefixed(op.GetData(), &iavlProof)
			if err != nil {
				return nil, nil, nil, fmt.Errorf("iavl: %s", err.Error())
			}
		} else if opType == "multistore" {
			mp, err := rootmulti.MultiStoreProofOpDecoder(op)

			multiStoreProof = mp.(rootmulti.MultiStoreProofOp)
			if err != nil {
				return nil, nil, nil, fmt.Errorf("multiStore: %s", err.Error())
			}
		}
	}

	return resp.Response.GetValue(), &iavlProof, &multiStoreProof, nil
}
