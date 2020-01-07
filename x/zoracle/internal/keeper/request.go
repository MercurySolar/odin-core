package keeper

import (
	"github.com/bandprotocol/d3n/chain/x/zoracle/internal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetRequest is a function to save request to the given ID.
func (k Keeper) SetRequest(ctx sdk.Context, id uint64, request types.Request) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.RequestStoreKey(id), k.cdc.MustMarshalBinaryBare(request))
}

// GetRequest returns the entire Request metadata struct.
func (k Keeper) GetRequest(ctx sdk.Context, id uint64) (types.Request, sdk.Error) {
	store := ctx.KVStore(k.storeKey)
	if !k.CheckRequestExists(ctx, id) {
		return types.Request{}, types.ErrRequestNotFound(types.DefaultCodespace)
	}

	bz := store.Get(types.RequestStoreKey(id))
	var request types.Request
	k.cdc.MustUnmarshalBinaryBare(bz, &request)
	return request, nil
}

// CheckRequestExists checks if the request at this id is present in the store or not.
func (k Keeper) CheckRequestExists(ctx sdk.Context, id uint64) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.RequestStoreKey(id))
}

// uniqueReqIDs is used to create array with all elements being unique (deduplicated).
func uniqueReqIDs(intSlice []uint64) []uint64 {
	keys := make(map[uint64]bool)
	list := []uint64{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

// SetPending saves the list of request in pending period.
func (k Keeper) SetPending(ctx sdk.Context, reqIDs []uint64) {
	store := ctx.KVStore(k.storeKey)
	urIDs := uniqueReqIDs(reqIDs)
	encoded := k.cdc.MustMarshalBinaryBare(urIDs)
	if encoded == nil {
		encoded = []byte{}
	}
	store.Set(types.PendingListStoreKey, encoded)
}

// GetPending returns the list of request IDs in pending period.
func (k Keeper) GetPending(ctx sdk.Context) []uint64 {
	store := ctx.KVStore(k.storeKey)
	reqIDsBytes := store.Get(types.PendingListStoreKey)

	// If the state is empty
	if len(reqIDsBytes) == 0 {
		return []uint64{}
	}

	var reqIDs []uint64
	k.cdc.MustUnmarshalBinaryBare(reqIDsBytes, &reqIDs)

	return reqIDs
}