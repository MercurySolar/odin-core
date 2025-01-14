// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mint/params.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Params holds parameters for the mint module.
type Params struct {
	// type of coin to mint
	MintDenom string `protobuf:"bytes,1,opt,name=mint_denom,json=mintDenom,proto3" json:"mint_denom,omitempty"`
	// maximum annual change in inflation rate
	InflationRateChange github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=inflation_rate_change,json=inflationRateChange,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"inflation_rate_change" yaml:"inflation_rate_change"`
	// maximum inflation rate
	InflationMax github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=inflation_max,json=inflationMax,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"inflation_max" yaml:"inflation_max"`
	// minimum inflation rate
	InflationMin github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=inflation_min,json=inflationMin,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"inflation_min" yaml:"inflation_min"`
	// goal of percent bonded atoms
	GoalBonded github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,5,opt,name=goal_bonded,json=goalBonded,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"goal_bonded" yaml:"goal_bonded"`
	// expected blocks per year
	BlocksPerYear uint64 `protobuf:"varint,6,opt,name=blocks_per_year,json=blocksPerYear,proto3" json:"blocks_per_year,omitempty" yaml:"blocks_per_year"`
	// max amount to withdraw per time
	MaxWithdrawalPerTime github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,7,rep,name=max_withdrawal_per_time,json=maxWithdrawalPerTime,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"max_withdrawal_per_time" yaml:"max_withdrawal_per_time"`
	// map with smart contracts addresses
	IntegrationAddresses map[string]string `protobuf:"bytes,8,rep,name=integration_addresses,json=integrationAddresses,proto3" json:"integration_addresses,omitempty" yaml:"integration_addresses" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// flag if minting from air
	MintAir bool `protobuf:"varint,9,opt,name=mint_air,json=mintAir,proto3" json:"mint_air,omitempty" yaml:"mint_air"`
	// eligible to withdraw accounts
	EligibleAccountsPool []string `protobuf:"bytes,10,rep,name=eligible_accounts_pool,json=eligibleAccountsPool,proto3" json:"eligible_accounts_pool,omitempty" yaml:"eligible_accounts_pool"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_04d03971f940ff2c, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetMintDenom() string {
	if m != nil {
		return m.MintDenom
	}
	return ""
}

func (m *Params) GetBlocksPerYear() uint64 {
	if m != nil {
		return m.BlocksPerYear
	}
	return 0
}

func (m *Params) GetMaxWithdrawalPerTime() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.MaxWithdrawalPerTime
	}
	return nil
}

func (m *Params) GetIntegrationAddresses() map[string]string {
	if m != nil {
		return m.IntegrationAddresses
	}
	return nil
}

func (m *Params) GetMintAir() bool {
	if m != nil {
		return m.MintAir
	}
	return false
}

func (m *Params) GetEligibleAccountsPool() []string {
	if m != nil {
		return m.EligibleAccountsPool
	}
	return nil
}

func init() {
	proto.RegisterType((*Params)(nil), "mint.Params")
	proto.RegisterMapType((map[string]string)(nil), "mint.Params.IntegrationAddressesEntry")
}

func init() { proto.RegisterFile("mint/params.proto", fileDescriptor_04d03971f940ff2c) }

var fileDescriptor_04d03971f940ff2c = []byte{
	// 651 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x3f, 0x6f, 0xd3, 0x4e,
	0x18, 0x8e, 0x9b, 0xfe, 0x49, 0xae, 0xbf, 0xaa, 0x3f, 0xdc, 0x50, 0xdc, 0x88, 0xda, 0xc1, 0x43,
	0x95, 0xa5, 0xb6, 0x0a, 0x0b, 0xea, 0x56, 0x37, 0x50, 0x55, 0x02, 0x14, 0x59, 0x48, 0x15, 0x2c,
	0xd6, 0xc5, 0x7e, 0x71, 0x4f, 0xb1, 0xef, 0xa2, 0xf3, 0xb5, 0x4d, 0x06, 0x16, 0x3e, 0x01, 0x23,
	0x63, 0x07, 0x26, 0x56, 0xbe, 0x44, 0xc7, 0x8e, 0x88, 0xc1, 0xa0, 0x76, 0x61, 0xce, 0x27, 0x40,
	0xe7, 0x73, 0xff, 0x50, 0xa5, 0x12, 0x95, 0x58, 0x92, 0x7b, 0x9f, 0xf7, 0xb9, 0xe7, 0xb9, 0xd7,
	0xf7, 0xbe, 0x87, 0xee, 0xa5, 0x84, 0x0a, 0x77, 0x80, 0x39, 0x4e, 0x33, 0x67, 0xc0, 0x99, 0x60,
	0xfa, 0xb4, 0x84, 0x9a, 0x8d, 0x98, 0xc5, 0xac, 0x00, 0x5c, 0xb9, 0x52, 0xb9, 0xa6, 0x19, 0xb2,
	0x2c, 0x65, 0x99, 0xdb, 0xc3, 0x19, 0xb8, 0x87, 0x1b, 0x3d, 0x10, 0x78, 0xc3, 0x0d, 0x19, 0xa1,
	0x65, 0x7e, 0xb1, 0x90, 0x93, 0x3f, 0x0a, 0xb0, 0xbf, 0xd6, 0xd0, 0x6c, 0xb7, 0x50, 0xd7, 0x57,
	0x11, 0x92, 0x89, 0x20, 0x02, 0xca, 0x52, 0x43, 0x6b, 0x69, 0xed, 0xba, 0x5f, 0x97, 0x48, 0x47,
	0x02, 0xfa, 0x07, 0x0d, 0xdd, 0x27, 0xf4, 0x5d, 0x82, 0x05, 0x61, 0x34, 0xe0, 0x58, 0x40, 0x10,
	0xee, 0x63, 0x1a, 0x83, 0x31, 0x25, 0xa9, 0xde, 0xab, 0x93, 0xdc, 0xaa, 0x7c, 0xcf, 0xad, 0xb5,
	0x98, 0x88, 0xfd, 0x83, 0x9e, 0x13, 0xb2, 0xd4, 0x2d, 0x4f, 0xa3, 0xfe, 0xd6, 0xb3, 0xa8, 0xef,
	0x8a, 0xd1, 0x00, 0x32, 0xa7, 0x03, 0xe1, 0x38, 0xb7, 0x1e, 0x8e, 0x70, 0x9a, 0x6c, 0xda, 0x13,
	0x45, 0x6d, 0x7f, 0xe9, 0x12, 0xf7, 0xb1, 0x80, 0xed, 0x02, 0xd5, 0xfb, 0x68, 0xe1, 0x8a, 0x9e,
	0xe2, 0xa1, 0x51, 0x2d, 0xbc, 0x9f, 0xdf, 0xd9, 0xbb, 0x71, 0xd3, 0x3b, 0xc5, 0x43, 0xdb, 0xff,
	0xef, 0x32, 0x7e, 0x89, 0x87, 0x37, 0xcc, 0x08, 0x35, 0xa6, 0xff, 0x99, 0x19, 0xa1, 0x7f, 0x98,
	0x11, 0xaa, 0x03, 0x9a, 0x8f, 0x19, 0x4e, 0x82, 0x1e, 0xa3, 0x11, 0x44, 0xc6, 0x4c, 0x61, 0xd5,
	0xb9, 0xb3, 0x95, 0xae, 0xac, 0xae, 0x49, 0xd9, 0x3e, 0x92, 0x91, 0x57, 0x04, 0xba, 0x87, 0x16,
	0x7b, 0x09, 0x0b, 0xfb, 0x59, 0x30, 0x00, 0x1e, 0x8c, 0x00, 0x73, 0x63, 0xb6, 0xa5, 0xb5, 0xa7,
	0xbd, 0xe6, 0x38, 0xb7, 0x96, 0xd5, 0xe6, 0x1b, 0x04, 0xdb, 0x5f, 0x50, 0x48, 0x17, 0xf8, 0x1b,
	0xc0, 0x5c, 0xff, 0xac, 0xa1, 0x07, 0x29, 0x1e, 0x06, 0x47, 0x44, 0xec, 0x47, 0x1c, 0x1f, 0xe1,
	0xa4, 0xe0, 0x0a, 0x92, 0x82, 0x31, 0xd7, 0xaa, 0xb6, 0xe7, 0x1f, 0xaf, 0x38, 0xea, 0x78, 0x8e,
	0xec, 0x43, 0xa7, 0xec, 0x43, 0x67, 0x9b, 0x11, 0xea, 0xf9, 0xb2, 0xa4, 0x71, 0x6e, 0x99, 0xca,
	0xeb, 0x16, 0x1d, 0xfb, 0xcb, 0x0f, 0xab, 0xfd, 0x17, 0x45, 0x4b, 0xc9, 0xcc, 0x6f, 0xa4, 0x78,
	0xb8, 0x77, 0x29, 0xd2, 0x05, 0xfe, 0x9a, 0xa4, 0xa0, 0xbf, 0x97, 0xfd, 0x2a, 0x20, 0xe6, 0xea,
	0x9b, 0xe3, 0x28, 0xe2, 0x90, 0x65, 0x90, 0x19, 0xb5, 0xe2, 0x8c, 0x6b, 0x4e, 0x31, 0x06, 0xaa,
	0xf9, 0x9d, 0xdd, 0x2b, 0xe6, 0xd6, 0x05, 0xf1, 0x19, 0x15, 0x7c, 0xe4, 0xb5, 0xae, 0x77, 0xea,
	0x04, 0x39, 0xdb, 0x6f, 0x90, 0x09, 0x9b, 0x75, 0x07, 0xd5, 0x8a, 0x71, 0xc2, 0x84, 0x1b, 0xf5,
	0x96, 0xd6, 0xae, 0x79, 0x4b, 0xe3, 0xdc, 0x5a, 0x2c, 0xcb, 0x2e, 0x33, 0xb6, 0x3f, 0x27, 0x97,
	0x5b, 0x84, 0xeb, 0x7b, 0x68, 0x19, 0x12, 0x12, 0x93, 0x5e, 0x02, 0x01, 0x0e, 0x43, 0x76, 0x40,
	0x45, 0x16, 0x0c, 0x18, 0x4b, 0x0c, 0xd4, 0xaa, 0xb6, 0xeb, 0xde, 0xa3, 0x71, 0x6e, 0xad, 0xaa,
	0xdd, 0x93, 0x79, 0xb6, 0xdf, 0xb8, 0x48, 0x6c, 0x95, 0x78, 0x97, 0xb1, 0xa4, 0xb9, 0x83, 0x56,
	0x6e, 0xad, 0x4e, 0xff, 0x1f, 0x55, 0xfb, 0x30, 0x2a, 0xa7, 0x5d, 0x2e, 0xf5, 0x06, 0x9a, 0x39,
	0xc4, 0xc9, 0x41, 0x39, 0xd6, 0xbe, 0x0a, 0x36, 0xa7, 0x9e, 0x6a, 0x9b, 0xb5, 0x4f, 0xc7, 0x56,
	0xe5, 0xd7, 0xb1, 0xa5, 0x79, 0xbb, 0x27, 0x67, 0xa6, 0x76, 0x7a, 0x66, 0x6a, 0x3f, 0xcf, 0x4c,
	0xed, 0xe3, 0xb9, 0x59, 0x39, 0x3d, 0x37, 0x2b, 0xdf, 0xce, 0xcd, 0xca, 0x5b, 0xf7, 0xda, 0xa5,
	0xed, 0x00, 0xeb, 0x78, 0xeb, 0x2f, 0x48, 0x4a, 0x04, 0x44, 0x2e, 0x8b, 0x08, 0x5d, 0x0f, 0x19,
	0x07, 0x77, 0x58, 0x3c, 0x40, 0xea, 0x06, 0x7b, 0xb3, 0xc5, 0x3b, 0xf4, 0xe4, 0x77, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xe8, 0xad, 0x96, 0x08, 0xe9, 0x04, 0x00, 0x00,
}

func (this *Params) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Params)
	if !ok {
		that2, ok := that.(Params)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.MintDenom != that1.MintDenom {
		return false
	}
	if !this.InflationRateChange.Equal(that1.InflationRateChange) {
		return false
	}
	if !this.InflationMax.Equal(that1.InflationMax) {
		return false
	}
	if !this.InflationMin.Equal(that1.InflationMin) {
		return false
	}
	if !this.GoalBonded.Equal(that1.GoalBonded) {
		return false
	}
	if this.BlocksPerYear != that1.BlocksPerYear {
		return false
	}
	if len(this.MaxWithdrawalPerTime) != len(that1.MaxWithdrawalPerTime) {
		return false
	}
	for i := range this.MaxWithdrawalPerTime {
		if !this.MaxWithdrawalPerTime[i].Equal(&that1.MaxWithdrawalPerTime[i]) {
			return false
		}
	}
	if len(this.IntegrationAddresses) != len(that1.IntegrationAddresses) {
		return false
	}
	for i := range this.IntegrationAddresses {
		if this.IntegrationAddresses[i] != that1.IntegrationAddresses[i] {
			return false
		}
	}
	if this.MintAir != that1.MintAir {
		return false
	}
	if len(this.EligibleAccountsPool) != len(that1.EligibleAccountsPool) {
		return false
	}
	for i := range this.EligibleAccountsPool {
		if this.EligibleAccountsPool[i] != that1.EligibleAccountsPool[i] {
			return false
		}
	}
	return true
}
func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.EligibleAccountsPool) > 0 {
		for iNdEx := len(m.EligibleAccountsPool) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.EligibleAccountsPool[iNdEx])
			copy(dAtA[i:], m.EligibleAccountsPool[iNdEx])
			i = encodeVarintParams(dAtA, i, uint64(len(m.EligibleAccountsPool[iNdEx])))
			i--
			dAtA[i] = 0x52
		}
	}
	if m.MintAir {
		i--
		if m.MintAir {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x48
	}
	if len(m.IntegrationAddresses) > 0 {
		for k := range m.IntegrationAddresses {
			v := m.IntegrationAddresses[k]
			baseI := i
			i -= len(v)
			copy(dAtA[i:], v)
			i = encodeVarintParams(dAtA, i, uint64(len(v)))
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintParams(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintParams(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.MaxWithdrawalPerTime) > 0 {
		for iNdEx := len(m.MaxWithdrawalPerTime) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MaxWithdrawalPerTime[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if m.BlocksPerYear != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.BlocksPerYear))
		i--
		dAtA[i] = 0x30
	}
	{
		size := m.GoalBonded.Size()
		i -= size
		if _, err := m.GoalBonded.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.InflationMin.Size()
		i -= size
		if _, err := m.InflationMin.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.InflationMax.Size()
		i -= size
		if _, err := m.InflationMax.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.InflationRateChange.Size()
		i -= size
		if _, err := m.InflationRateChange.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.MintDenom) > 0 {
		i -= len(m.MintDenom)
		copy(dAtA[i:], m.MintDenom)
		i = encodeVarintParams(dAtA, i, uint64(len(m.MintDenom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.MintDenom)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = m.InflationRateChange.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.InflationMax.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.InflationMin.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.GoalBonded.Size()
	n += 1 + l + sovParams(uint64(l))
	if m.BlocksPerYear != 0 {
		n += 1 + sovParams(uint64(m.BlocksPerYear))
	}
	if len(m.MaxWithdrawalPerTime) > 0 {
		for _, e := range m.MaxWithdrawalPerTime {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	if len(m.IntegrationAddresses) > 0 {
		for k, v := range m.IntegrationAddresses {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovParams(uint64(len(k))) + 1 + len(v) + sovParams(uint64(len(v)))
			n += mapEntrySize + 1 + sovParams(uint64(mapEntrySize))
		}
	}
	if m.MintAir {
		n += 2
	}
	if len(m.EligibleAccountsPool) > 0 {
		for _, s := range m.EligibleAccountsPool {
			l = len(s)
			n += 1 + l + sovParams(uint64(l))
		}
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MintDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InflationRateChange", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InflationRateChange.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InflationMax", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InflationMax.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InflationMin", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InflationMin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GoalBonded", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.GoalBonded.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlocksPerYear", wireType)
			}
			m.BlocksPerYear = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlocksPerYear |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxWithdrawalPerTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MaxWithdrawalPerTime = append(m.MaxWithdrawalPerTime, types.Coin{})
			if err := m.MaxWithdrawalPerTime[len(m.MaxWithdrawalPerTime)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IntegrationAddresses", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.IntegrationAddresses == nil {
				m.IntegrationAddresses = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowParams
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowParams
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthParams
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthParams
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowParams
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthParams
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue < 0 {
						return ErrInvalidLengthParams
					}
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipParams(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthParams
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.IntegrationAddresses[mapkey] = mapvalue
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintAir", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.MintAir = bool(v != 0)
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EligibleAccountsPool", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EligibleAccountsPool = append(m.EligibleAccountsPool, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowParams
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowParams
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
