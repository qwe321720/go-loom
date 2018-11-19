// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/loomnetwork/go-loom/builtin/types/dposv2/dposv2.proto

/*
Package dposv2 is a generated protocol buffer package.

It is generated from these files:
	github.com/loomnetwork/go-loom/builtin/types/dposv2/dposv2.proto

It has these top-level messages:
	ParamsV2
	StateV2
	DposValidator
	CandidateV2
	CandidateListV2
	DelegationV2
	DelegationListV2
	DistributionV2
	DistributionListV2
	DPOSInitRequestV2
	DelegateRequestV2
	RedelegateRequestV2
	CheckDelegationRequestV2
	CheckDelegationResponseV2
	ClaimDistributionRequestV2
	ClaimDistributionResponseV2
	UnbondRequestV2
	RegisterCandidateRequestV2
	UnregisterCandidateRequestV2
	ElectDelegationRequestV2
	ListValidatorsRequestV2
	ListValidatorsResponseV2
	ListCandidateRequestV2
	ListCandidateResponseV2
*/
package dposv2

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import types "github.com/loomnetwork/go-loom/types"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type ParamsV2 struct {
	ValidatorCount      uint64         `protobuf:"varint,1,opt,name=validator_count,json=validatorCount,proto3" json:"validator_count,omitempty"`
	ElectionCycleLength int64          `protobuf:"varint,2,opt,name=election_cycle_length,json=electionCycleLength,proto3" json:"election_cycle_length,omitempty"`
	CoinContractAddress *types.Address `protobuf:"bytes,3,opt,name=coin_contract_address,json=coinContractAddress" json:"coin_contract_address,omitempty"`
}

func (m *ParamsV2) Reset()                    { *m = ParamsV2{} }
func (m *ParamsV2) String() string            { return proto.CompactTextString(m) }
func (*ParamsV2) ProtoMessage()               {}
func (*ParamsV2) Descriptor() ([]byte, []int) { return fileDescriptorDposv2, []int{0} }

func (m *ParamsV2) GetValidatorCount() uint64 {
	if m != nil {
		return m.ValidatorCount
	}
	return 0
}

func (m *ParamsV2) GetElectionCycleLength() int64 {
	if m != nil {
		return m.ElectionCycleLength
	}
	return 0
}

func (m *ParamsV2) GetCoinContractAddress() *types.Address {
	if m != nil {
		return m.CoinContractAddress
	}
	return nil
}

type StateV2 struct {
	Params           *ParamsV2        `protobuf:"bytes,1,opt,name=params" json:"params,omitempty"`
	Validators       []*DposValidator `protobuf:"bytes,2,rep,name=validators" json:"validators,omitempty"`
	LastElectionTime int64            `protobuf:"varint,3,opt,name=last_election_time,json=lastElectionTime,proto3" json:"last_election_time,omitempty"`
}

func (m *StateV2) Reset()                    { *m = StateV2{} }
func (m *StateV2) String() string            { return proto.CompactTextString(m) }
func (*StateV2) ProtoMessage()               {}
func (*StateV2) Descriptor() ([]byte, []int) { return fileDescriptorDposv2, []int{1} }

func (m *StateV2) GetParams() *ParamsV2 {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *StateV2) GetValidators() []*DposValidator {
	if m != nil {
		return m.Validators
	}
	return nil
}

func (m *StateV2) GetLastElectionTime() int64 {
	if m != nil {
		return m.LastElectionTime
	}
	return 0
}

type DposValidator struct {
	PubKey            []byte         `protobuf:"bytes,1,opt,name=pub_key,json=pubKey,proto3" json:"pub_key,omitempty"`
	Power             int64          `protobuf:"varint,2,opt,name=power,proto3" json:"power,omitempty"`
	UpblockCount      uint64         `protobuf:"varint,3,opt,name=upblock_count,json=upblockCount,proto3" json:"upblock_count,omitempty"`
	DistributionTotal *types.BigUInt `protobuf:"bytes,4,opt,name=distribution_total,json=distributionTotal" json:"distribution_total,omitempty"`
	DelegationTotal   *types.BigUInt `protobuf:"bytes,5,opt,name=delegation_total,json=delegationTotal" json:"delegation_total,omitempty"`
}

func (m *DposValidator) Reset()                    { *m = DposValidator{} }
func (m *DposValidator) String() string            { return proto.CompactTextString(m) }
func (*DposValidator) ProtoMessage()               {}
func (*DposValidator) Descriptor() ([]byte, []int) { return fileDescriptorDposv2, []int{2} }

func (m *DposValidator) GetPubKey() []byte {
	if m != nil {
		return m.PubKey
	}
	return nil
}

func (m *DposValidator) GetPower() int64 {
	if m != nil {
		return m.Power
	}
	return 0
}

func (m *DposValidator) GetUpblockCount() uint64 {
	if m != nil {
		return m.UpblockCount
	}
	return 0
}

func (m *DposValidator) GetDistributionTotal() *types.BigUInt {
	if m != nil {
		return m.DistributionTotal
	}
	return nil
}

func (m *DposValidator) GetDelegationTotal() *types.BigUInt {
	if m != nil {
		return m.DelegationTotal
	}
	return nil
}

type CandidateV2 struct {
	Address *types.Address `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	PubKey  []byte         `protobuf:"bytes,2,opt,name=pub_key,json=pubKey,proto3" json:"pub_key,omitempty"`
	Fee     uint64         `protobuf:"varint,3,opt,name=fee,proto3" json:"fee,omitempty"`
}

func (m *CandidateV2) Reset()                    { *m = CandidateV2{} }
func (m *CandidateV2) String() string            { return proto.CompactTextString(m) }
func (*CandidateV2) ProtoMessage()               {}
func (*CandidateV2) Descriptor() ([]byte, []int) { return fileDescriptorDposv2, []int{3} }

func (m *CandidateV2) GetAddress() *types.Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *CandidateV2) GetPubKey() []byte {
	if m != nil {
		return m.PubKey
	}
	return nil
}

func (m *CandidateV2) GetFee() uint64 {
	if m != nil {
		return m.Fee
	}
	return 0
}

type CandidateListV2 struct {
	Candidates []*CandidateV2 `protobuf:"bytes,1,rep,name=candidates" json:"candidates,omitempty"`
}

func (m *CandidateListV2) Reset()                    { *m = CandidateListV2{} }
func (m *CandidateListV2) String() string            { return proto.CompactTextString(m) }
func (*CandidateListV2) ProtoMessage()               {}
func (*CandidateListV2) Descriptor() ([]byte, []int) { return fileDescriptorDposv2, []int{4} }

func (m *CandidateListV2) GetCandidates() []*CandidateV2 {
	if m != nil {
		return m.Candidates
	}
	return nil
}

type DelegationV2 struct {
	Validator *types.Address `protobuf:"bytes,1,opt,name=validator" json:"validator,omitempty"`
	Delegator *types.Address `protobuf:"bytes,2,opt,name=delegator" json:"delegator,omitempty"`
	Height    uint64         `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	Amount    *types.BigUInt `protobuf:"bytes,4,opt,name=amount" json:"amount,omitempty"`
}

func (m *DelegationV2) Reset()                    { *m = DelegationV2{} }
func (m *DelegationV2) String() string            { return proto.CompactTextString(m) }
func (*DelegationV2) ProtoMessage()               {}
func (*DelegationV2) Descriptor() ([]byte, []int) { return fileDescriptorDposv2, []int{5} }

func (m *DelegationV2) GetValidator() *types.Address {
	if m != nil {
		return m.Validator
	}
	return nil
}

func (m *DelegationV2) GetDelegator() *types.Address {
	if m != nil {
		return m.Delegator
	}
	return nil
}

func (m *DelegationV2) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *DelegationV2) GetAmount() *types.BigUInt {
	if m != nil {
		return m.Amount
	}
	return nil
}

type DelegationListV2 struct {
	Delegations []*DelegationV2 `protobuf:"bytes,1,rep,name=delegations" json:"delegations,omitempty"`
}

func (m *DelegationListV2) Reset()                    { *m = DelegationListV2{} }
func (m *DelegationListV2) String() string            { return proto.CompactTextString(m) }
func (*DelegationListV2) ProtoMessage()               {}
func (*DelegationListV2) Descriptor() ([]byte, []int) { return fileDescriptorDposv2, []int{6} }

func (m *DelegationListV2) GetDelegations() []*DelegationV2 {
	if m != nil {
		return m.Delegations
	}
	return nil
}

type DistributionV2 struct {
	Address *types.Address `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	Amount  *types.BigUInt `protobuf:"bytes,2,opt,name=amount" json:"amount,omitempty"`
}

func (m *DistributionV2) Reset()                    { *m = DistributionV2{} }
func (m *DistributionV2) String() string            { return proto.CompactTextString(m) }
func (*DistributionV2) ProtoMessage()               {}
func (*DistributionV2) Descriptor() ([]byte, []int) { return fileDescriptorDposv2, []int{7} }

func (m *DistributionV2) GetAddress() *types.Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *DistributionV2) GetAmount() *types.BigUInt {
	if m != nil {
		return m.Amount
	}
	return nil
}

type DistributionListV2 struct {
	Distributions []*DistributionV2 `protobuf:"bytes,1,rep,name=distributions" json:"distributions,omitempty"`
}

func (m *DistributionListV2) Reset()                    { *m = DistributionListV2{} }
func (m *DistributionListV2) String() string            { return proto.CompactTextString(m) }
func (*DistributionListV2) ProtoMessage()               {}
func (*DistributionListV2) Descriptor() ([]byte, []int) { return fileDescriptorDposv2, []int{8} }

func (m *DistributionListV2) GetDistributions() []*DistributionV2 {
	if m != nil {
		return m.Distributions
	}
	return nil
}

type DPOSInitRequestV2 struct {
	Params     *ParamsV2          `protobuf:"bytes,1,opt,name=params" json:"params,omitempty"`
	Validators []*types.Validator `protobuf:"bytes,2,rep,name=validators" json:"validators,omitempty"`
}

func (m *DPOSInitRequestV2) Reset()                    { *m = DPOSInitRequestV2{} }
func (m *DPOSInitRequestV2) String() string            { return proto.CompactTextString(m) }
func (*DPOSInitRequestV2) ProtoMessage()               {}
func (*DPOSInitRequestV2) Descriptor() ([]byte, []int) { return fileDescriptorDposv2, []int{9} }

func (m *DPOSInitRequestV2) GetParams() *ParamsV2 {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *DPOSInitRequestV2) GetValidators() []*types.Validator {
	if m != nil {
		return m.Validators
	}
	return nil
}

type DelegateRequestV2 struct {
	ValidatorAddress *types.Address `protobuf:"bytes,1,opt,name=validator_address,json=validatorAddress" json:"validator_address,omitempty"`
	Amount           *types.BigUInt `protobuf:"bytes,2,opt,name=amount" json:"amount,omitempty"`
}

func (m *DelegateRequestV2) Reset()                    { *m = DelegateRequestV2{} }
func (m *DelegateRequestV2) String() string            { return proto.CompactTextString(m) }
func (*DelegateRequestV2) ProtoMessage()               {}
func (*DelegateRequestV2) Descriptor() ([]byte, []int) { return fileDescriptorDposv2, []int{10} }

func (m *DelegateRequestV2) GetValidatorAddress() *types.Address {
	if m != nil {
		return m.ValidatorAddress
	}
	return nil
}

func (m *DelegateRequestV2) GetAmount() *types.BigUInt {
	if m != nil {
		return m.Amount
	}
	return nil
}

type RedelegateRequestV2 struct {
	ValidatorAddress *types.Address `protobuf:"bytes,1,opt,name=validator_address,json=validatorAddress" json:"validator_address,omitempty"`
	Amount           *types.BigUInt `protobuf:"bytes,2,opt,name=amount" json:"amount,omitempty"`
}

func (m *RedelegateRequestV2) Reset()                    { *m = RedelegateRequestV2{} }
func (m *RedelegateRequestV2) String() string            { return proto.CompactTextString(m) }
func (*RedelegateRequestV2) ProtoMessage()               {}
func (*RedelegateRequestV2) Descriptor() ([]byte, []int) { return fileDescriptorDposv2, []int{11} }

func (m *RedelegateRequestV2) GetValidatorAddress() *types.Address {
	if m != nil {
		return m.ValidatorAddress
	}
	return nil
}

func (m *RedelegateRequestV2) GetAmount() *types.BigUInt {
	if m != nil {
		return m.Amount
	}
	return nil
}

type CheckDelegationRequestV2 struct {
	ValidatorAddress *types.Address `protobuf:"bytes,1,opt,name=validator_address,json=validatorAddress" json:"validator_address,omitempty"`
	DelegatorAddress *types.Address `protobuf:"bytes,2,opt,name=delegator_address,json=delegatorAddress" json:"delegator_address,omitempty"`
}

func (m *CheckDelegationRequestV2) Reset()                    { *m = CheckDelegationRequestV2{} }
func (m *CheckDelegationRequestV2) String() string            { return proto.CompactTextString(m) }
func (*CheckDelegationRequestV2) ProtoMessage()               {}
func (*CheckDelegationRequestV2) Descriptor() ([]byte, []int) { return fileDescriptorDposv2, []int{12} }

func (m *CheckDelegationRequestV2) GetValidatorAddress() *types.Address {
	if m != nil {
		return m.ValidatorAddress
	}
	return nil
}

func (m *CheckDelegationRequestV2) GetDelegatorAddress() *types.Address {
	if m != nil {
		return m.DelegatorAddress
	}
	return nil
}

type CheckDelegationResponseV2 struct {
	Delegation *DelegationV2 `protobuf:"bytes,1,opt,name=delegation" json:"delegation,omitempty"`
}

func (m *CheckDelegationResponseV2) Reset()                    { *m = CheckDelegationResponseV2{} }
func (m *CheckDelegationResponseV2) String() string            { return proto.CompactTextString(m) }
func (*CheckDelegationResponseV2) ProtoMessage()               {}
func (*CheckDelegationResponseV2) Descriptor() ([]byte, []int) { return fileDescriptorDposv2, []int{13} }

func (m *CheckDelegationResponseV2) GetDelegation() *DelegationV2 {
	if m != nil {
		return m.Delegation
	}
	return nil
}

type ClaimDistributionRequestV2 struct {
	WithdrawalAddress *types.Address `protobuf:"bytes,1,opt,name=withdrawal_address,json=withdrawalAddress" json:"withdrawal_address,omitempty"`
}

func (m *ClaimDistributionRequestV2) Reset()         { *m = ClaimDistributionRequestV2{} }
func (m *ClaimDistributionRequestV2) String() string { return proto.CompactTextString(m) }
func (*ClaimDistributionRequestV2) ProtoMessage()    {}
func (*ClaimDistributionRequestV2) Descriptor() ([]byte, []int) {
	return fileDescriptorDposv2, []int{14}
}

func (m *ClaimDistributionRequestV2) GetWithdrawalAddress() *types.Address {
	if m != nil {
		return m.WithdrawalAddress
	}
	return nil
}

type ClaimDistributionResponseV2 struct {
	Amount *types.BigUInt `protobuf:"bytes,1,opt,name=amount" json:"amount,omitempty"`
}

func (m *ClaimDistributionResponseV2) Reset()         { *m = ClaimDistributionResponseV2{} }
func (m *ClaimDistributionResponseV2) String() string { return proto.CompactTextString(m) }
func (*ClaimDistributionResponseV2) ProtoMessage()    {}
func (*ClaimDistributionResponseV2) Descriptor() ([]byte, []int) {
	return fileDescriptorDposv2, []int{15}
}

func (m *ClaimDistributionResponseV2) GetAmount() *types.BigUInt {
	if m != nil {
		return m.Amount
	}
	return nil
}

type UnbondRequestV2 struct {
	ValidatorAddress *types.Address `protobuf:"bytes,1,opt,name=validator_address,json=validatorAddress" json:"validator_address,omitempty"`
	Amount           *types.BigUInt `protobuf:"bytes,2,opt,name=amount" json:"amount,omitempty"`
}

func (m *UnbondRequestV2) Reset()                    { *m = UnbondRequestV2{} }
func (m *UnbondRequestV2) String() string            { return proto.CompactTextString(m) }
func (*UnbondRequestV2) ProtoMessage()               {}
func (*UnbondRequestV2) Descriptor() ([]byte, []int) { return fileDescriptorDposv2, []int{16} }

func (m *UnbondRequestV2) GetValidatorAddress() *types.Address {
	if m != nil {
		return m.ValidatorAddress
	}
	return nil
}

func (m *UnbondRequestV2) GetAmount() *types.BigUInt {
	if m != nil {
		return m.Amount
	}
	return nil
}

type RegisterCandidateRequestV2 struct {
	PubKey []byte `protobuf:"bytes,1,opt,name=pub_key,json=pubKey,proto3" json:"pub_key,omitempty"`
	Fee    uint64 `protobuf:"varint,2,opt,name=fee,proto3" json:"fee,omitempty"`
}

func (m *RegisterCandidateRequestV2) Reset()         { *m = RegisterCandidateRequestV2{} }
func (m *RegisterCandidateRequestV2) String() string { return proto.CompactTextString(m) }
func (*RegisterCandidateRequestV2) ProtoMessage()    {}
func (*RegisterCandidateRequestV2) Descriptor() ([]byte, []int) {
	return fileDescriptorDposv2, []int{17}
}

func (m *RegisterCandidateRequestV2) GetPubKey() []byte {
	if m != nil {
		return m.PubKey
	}
	return nil
}

func (m *RegisterCandidateRequestV2) GetFee() uint64 {
	if m != nil {
		return m.Fee
	}
	return 0
}

type UnregisterCandidateRequestV2 struct {
}

func (m *UnregisterCandidateRequestV2) Reset()         { *m = UnregisterCandidateRequestV2{} }
func (m *UnregisterCandidateRequestV2) String() string { return proto.CompactTextString(m) }
func (*UnregisterCandidateRequestV2) ProtoMessage()    {}
func (*UnregisterCandidateRequestV2) Descriptor() ([]byte, []int) {
	return fileDescriptorDposv2, []int{18}
}

type ElectDelegationRequestV2 struct {
}

func (m *ElectDelegationRequestV2) Reset()                    { *m = ElectDelegationRequestV2{} }
func (m *ElectDelegationRequestV2) String() string            { return proto.CompactTextString(m) }
func (*ElectDelegationRequestV2) ProtoMessage()               {}
func (*ElectDelegationRequestV2) Descriptor() ([]byte, []int) { return fileDescriptorDposv2, []int{19} }

type ListValidatorsRequestV2 struct {
}

func (m *ListValidatorsRequestV2) Reset()                    { *m = ListValidatorsRequestV2{} }
func (m *ListValidatorsRequestV2) String() string            { return proto.CompactTextString(m) }
func (*ListValidatorsRequestV2) ProtoMessage()               {}
func (*ListValidatorsRequestV2) Descriptor() ([]byte, []int) { return fileDescriptorDposv2, []int{20} }

type ListValidatorsResponseV2 struct {
	Validators []*DposValidator `protobuf:"bytes,1,rep,name=validators" json:"validators,omitempty"`
}

func (m *ListValidatorsResponseV2) Reset()                    { *m = ListValidatorsResponseV2{} }
func (m *ListValidatorsResponseV2) String() string            { return proto.CompactTextString(m) }
func (*ListValidatorsResponseV2) ProtoMessage()               {}
func (*ListValidatorsResponseV2) Descriptor() ([]byte, []int) { return fileDescriptorDposv2, []int{21} }

func (m *ListValidatorsResponseV2) GetValidators() []*DposValidator {
	if m != nil {
		return m.Validators
	}
	return nil
}

type ListCandidateRequestV2 struct {
}

func (m *ListCandidateRequestV2) Reset()                    { *m = ListCandidateRequestV2{} }
func (m *ListCandidateRequestV2) String() string            { return proto.CompactTextString(m) }
func (*ListCandidateRequestV2) ProtoMessage()               {}
func (*ListCandidateRequestV2) Descriptor() ([]byte, []int) { return fileDescriptorDposv2, []int{22} }

type ListCandidateResponseV2 struct {
	Candidates []*CandidateV2 `protobuf:"bytes,1,rep,name=candidates" json:"candidates,omitempty"`
}

func (m *ListCandidateResponseV2) Reset()                    { *m = ListCandidateResponseV2{} }
func (m *ListCandidateResponseV2) String() string            { return proto.CompactTextString(m) }
func (*ListCandidateResponseV2) ProtoMessage()               {}
func (*ListCandidateResponseV2) Descriptor() ([]byte, []int) { return fileDescriptorDposv2, []int{23} }

func (m *ListCandidateResponseV2) GetCandidates() []*CandidateV2 {
	if m != nil {
		return m.Candidates
	}
	return nil
}

func init() {
	proto.RegisterType((*ParamsV2)(nil), "ParamsV2")
	proto.RegisterType((*StateV2)(nil), "StateV2")
	proto.RegisterType((*DposValidator)(nil), "DposValidator")
	proto.RegisterType((*CandidateV2)(nil), "CandidateV2")
	proto.RegisterType((*CandidateListV2)(nil), "CandidateListV2")
	proto.RegisterType((*DelegationV2)(nil), "DelegationV2")
	proto.RegisterType((*DelegationListV2)(nil), "DelegationListV2")
	proto.RegisterType((*DistributionV2)(nil), "DistributionV2")
	proto.RegisterType((*DistributionListV2)(nil), "DistributionListV2")
	proto.RegisterType((*DPOSInitRequestV2)(nil), "DPOSInitRequestV2")
	proto.RegisterType((*DelegateRequestV2)(nil), "DelegateRequestV2")
	proto.RegisterType((*RedelegateRequestV2)(nil), "RedelegateRequestV2")
	proto.RegisterType((*CheckDelegationRequestV2)(nil), "CheckDelegationRequestV2")
	proto.RegisterType((*CheckDelegationResponseV2)(nil), "CheckDelegationResponseV2")
	proto.RegisterType((*ClaimDistributionRequestV2)(nil), "ClaimDistributionRequestV2")
	proto.RegisterType((*ClaimDistributionResponseV2)(nil), "ClaimDistributionResponseV2")
	proto.RegisterType((*UnbondRequestV2)(nil), "UnbondRequestV2")
	proto.RegisterType((*RegisterCandidateRequestV2)(nil), "RegisterCandidateRequestV2")
	proto.RegisterType((*UnregisterCandidateRequestV2)(nil), "UnregisterCandidateRequestV2")
	proto.RegisterType((*ElectDelegationRequestV2)(nil), "ElectDelegationRequestV2")
	proto.RegisterType((*ListValidatorsRequestV2)(nil), "ListValidatorsRequestV2")
	proto.RegisterType((*ListValidatorsResponseV2)(nil), "ListValidatorsResponseV2")
	proto.RegisterType((*ListCandidateRequestV2)(nil), "ListCandidateRequestV2")
	proto.RegisterType((*ListCandidateResponseV2)(nil), "ListCandidateResponseV2")
}

func init() {
	proto.RegisterFile("github.com/loomnetwork/go-loom/builtin/types/dposv2/dposv2.proto", fileDescriptorDposv2)
}

var fileDescriptorDposv2 = []byte{
	// 816 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0xed, 0x6e, 0xe2, 0x46,
	0x14, 0x95, 0x21, 0x21, 0xc9, 0x05, 0x02, 0x4c, 0x9a, 0xc4, 0xa1, 0x55, 0x45, 0x5d, 0xa9, 0x45,
	0x55, 0x62, 0x2a, 0x47, 0x51, 0xfe, 0x54, 0x4a, 0x53, 0xa8, 0xa2, 0x7c, 0x48, 0x8d, 0x9c, 0xc0,
	0xaf, 0x4a, 0xc8, 0x1f, 0x53, 0x98, 0x62, 0x3c, 0x5e, 0x7b, 0x1c, 0xc4, 0x1b, 0xac, 0xf6, 0x09,
	0xf6, 0x05, 0xf6, 0x81, 0xf6, 0x8d, 0x56, 0xb6, 0xc7, 0x63, 0x43, 0x40, 0xc9, 0x6a, 0xa5, 0xfc,
	0x01, 0x66, 0xce, 0x99, 0x3b, 0xe7, 0x9e, 0x7b, 0x67, 0x06, 0xf8, 0x73, 0x44, 0xd8, 0x38, 0x34,
	0x55, 0x8b, 0x4e, 0x3b, 0x0e, 0xa5, 0x53, 0x17, 0xb3, 0x19, 0xf5, 0x27, 0x9d, 0x11, 0x3d, 0x89,
	0x86, 0x1d, 0x33, 0x24, 0x0e, 0x23, 0x6e, 0x87, 0xcd, 0x3d, 0x1c, 0x74, 0x6c, 0x8f, 0x06, 0x4f,
	0x1a, 0xff, 0x52, 0x3d, 0x9f, 0x32, 0xda, 0xfc, 0xfd, 0x85, 0x08, 0xc9, 0xca, 0xf8, 0x33, 0x59,
	0xa1, 0x7c, 0x92, 0x60, 0xfb, 0xde, 0xf0, 0x8d, 0x69, 0x30, 0xd0, 0xd0, 0xaf, 0x50, 0x7b, 0x32,
	0x1c, 0x62, 0x1b, 0x8c, 0xfa, 0x43, 0x8b, 0x86, 0x2e, 0x93, 0xa5, 0x96, 0xd4, 0xde, 0xd0, 0x77,
	0xc5, 0x74, 0x37, 0x9a, 0x45, 0x1a, 0xec, 0x63, 0x07, 0x5b, 0x8c, 0x50, 0x77, 0x68, 0xcd, 0x2d,
	0x07, 0x0f, 0x1d, 0xec, 0x8e, 0xd8, 0x58, 0x2e, 0xb4, 0xa4, 0x76, 0x51, 0xdf, 0x4b, 0xc1, 0x6e,
	0x84, 0xdd, 0xc5, 0x10, 0xfa, 0x03, 0xf6, 0x2d, 0x4a, 0xdc, 0xa1, 0x45, 0x5d, 0xe6, 0x1b, 0x16,
	0x1b, 0x1a, 0xb6, 0xed, 0xe3, 0x20, 0x90, 0x8b, 0x2d, 0xa9, 0x5d, 0xd6, 0xb6, 0xd5, 0xcb, 0x64,
	0xac, 0xef, 0x45, 0xb4, 0x2e, 0x67, 0xf1, 0x49, 0xe5, 0x83, 0x04, 0x5b, 0x0f, 0xcc, 0x60, 0x78,
	0xa0, 0xa1, 0x9f, 0xa0, 0xe4, 0xc5, 0x92, 0x63, 0x75, 0x65, 0x6d, 0x47, 0x4d, 0x33, 0xd0, 0x39,
	0x80, 0x54, 0x00, 0x21, 0x39, 0x90, 0x0b, 0xad, 0x62, 0xbb, 0xac, 0xed, 0xaa, 0x3d, 0x8f, 0x06,
	0x83, 0x74, 0x5a, 0xcf, 0x31, 0xd0, 0x31, 0x20, 0xc7, 0x08, 0xd8, 0x50, 0x64, 0xc5, 0xc8, 0x14,
	0xc7, 0xca, 0x8a, 0x7a, 0x3d, 0x42, 0xfe, 0xe6, 0xc0, 0x23, 0x99, 0x62, 0xe5, 0xb3, 0x04, 0xd5,
	0x85, 0x58, 0xe8, 0x10, 0xb6, 0xbc, 0xd0, 0x1c, 0x4e, 0xf0, 0x3c, 0xd6, 0x54, 0xd1, 0x4b, 0x5e,
	0x68, 0xde, 0xe2, 0x39, 0xfa, 0x0e, 0x36, 0x3d, 0x3a, 0xc3, 0x3e, 0x77, 0x26, 0x19, 0xa0, 0x9f,
	0xa1, 0x1a, 0x7a, 0xa6, 0x43, 0xad, 0x09, 0xb7, 0xb9, 0x18, 0xdb, 0x5c, 0xe1, 0x93, 0x89, 0xc9,
	0xe7, 0x80, 0x6c, 0x12, 0x30, 0x9f, 0x98, 0x61, 0x22, 0x89, 0x32, 0xc3, 0x91, 0x37, 0xb8, 0x5b,
	0x7f, 0x91, 0x51, 0xff, 0xda, 0x65, 0x7a, 0x23, 0xcf, 0x79, 0x8c, 0x28, 0xe8, 0x14, 0xea, 0x36,
	0x76, 0xf0, 0xc8, 0xc8, 0x2d, 0xdb, 0x5c, 0x5a, 0x56, 0xcb, 0x18, 0xf1, 0x22, 0xe5, 0x5f, 0x28,
	0x77, 0x0d, 0xd7, 0x8e, 0xf2, 0x89, 0x3c, 0x56, 0x60, 0x2b, 0xad, 0x8f, 0xb4, 0x54, 0x9f, 0x14,
	0xc8, 0x27, 0x5d, 0x58, 0x48, 0xba, 0x0e, 0xc5, 0xff, 0x30, 0xe6, 0x49, 0x45, 0x3f, 0x95, 0x0b,
	0xa8, 0x89, 0xe8, 0x77, 0x24, 0x60, 0x03, 0x0d, 0x1d, 0x03, 0x58, 0xe9, 0x54, 0xb4, 0x49, 0x54,
	0xa2, 0x8a, 0x9a, 0xd3, 0xa0, 0xe7, 0x70, 0xe5, 0xa3, 0x04, 0x95, 0x9e, 0x90, 0x3c, 0xd0, 0xd0,
	0x2f, 0xb0, 0x23, 0xea, 0xf7, 0x4c, 0x62, 0x06, 0x45, 0x3c, 0x9e, 0x2a, 0x4d, 0x8a, 0xb0, 0xc0,
	0x13, 0x10, 0x3a, 0x80, 0xd2, 0x18, 0x93, 0xd1, 0x38, 0xad, 0x05, 0x1f, 0xa1, 0x16, 0x94, 0x8c,
	0x69, 0x5c, 0xa3, 0x65, 0xe7, 0xf9, 0xbc, 0xd2, 0x85, 0x7a, 0xa6, 0x8c, 0x27, 0xd7, 0x81, 0x72,
	0x66, 0x70, 0x9a, 0x5d, 0x55, 0xcd, 0x67, 0xa0, 0xe7, 0x19, 0xca, 0x00, 0x76, 0x7b, 0xb9, 0x42,
	0xbe, 0xb2, 0x02, 0x99, 0xb8, 0xc2, 0x1a, 0x71, 0xb7, 0x80, 0xf2, 0x71, 0xb9, 0xbc, 0x33, 0xa8,
	0xe6, 0xdb, 0x26, 0x15, 0x58, 0x53, 0x17, 0x35, 0xe8, 0x8b, 0x2c, 0xc5, 0x84, 0x46, 0xef, 0xfe,
	0x9f, 0x87, 0x6b, 0x97, 0x30, 0x1d, 0xbf, 0x0b, 0x71, 0x1c, 0xeb, 0x15, 0xa7, 0xf1, 0xb7, 0x15,
	0xa7, 0x11, 0xd4, 0x95, 0x27, 0x51, 0x71, 0xa0, 0xc1, 0x5d, 0xc2, 0xd9, 0x1e, 0x67, 0xd0, 0xc8,
	0x2e, 0xa6, 0x75, 0xae, 0xd4, 0x05, 0xe5, 0xf2, 0xd5, 0xf6, 0xb8, 0xb0, 0xa7, 0x63, 0xfb, 0xed,
	0xf6, 0x7b, 0x2f, 0x81, 0xdc, 0x1d, 0x63, 0x6b, 0x92, 0x75, 0xc2, 0x37, 0xef, 0x7a, 0x06, 0x0d,
	0xd1, 0xc6, 0x62, 0xd9, 0x72, 0xa7, 0xd7, 0x05, 0x25, 0xbd, 0x51, 0x6f, 0xe0, 0xe8, 0x99, 0x92,
	0xc0, 0xa3, 0x6e, 0x10, 0x1d, 0xff, 0x13, 0x80, 0xac, 0x3b, 0xb9, 0x86, 0xa5, 0xf6, 0xcd, 0x11,
	0x94, 0x3e, 0x34, 0xbb, 0x8e, 0x41, 0xa6, 0xf9, 0xf6, 0xc9, 0xf2, 0x3a, 0x07, 0x34, 0x23, 0x6c,
	0x6c, 0xfb, 0xc6, 0xcc, 0x70, 0xd6, 0x26, 0xd6, 0xc8, 0x38, 0xa9, 0xc4, 0x0b, 0xf8, 0x7e, 0x45,
	0x58, 0x21, 0x32, 0xb3, 0x5b, 0x5a, 0x63, 0xf7, 0xff, 0x50, 0xeb, 0xbb, 0x26, 0x75, 0xed, 0x37,
	0x28, 0xed, 0x15, 0x34, 0x75, 0x3c, 0x22, 0x01, 0xc3, 0xbe, 0xb8, 0xc4, 0xb2, 0x6d, 0xd7, 0x3e,
	0x10, 0xfc, 0xae, 0x2c, 0x64, 0x77, 0xe5, 0x8f, 0xf0, 0x43, 0xdf, 0xf5, 0xd7, 0x86, 0x52, 0x9a,
	0x20, 0xc7, 0xaf, 0xd1, 0x8a, 0x16, 0x52, 0x8e, 0xe0, 0x30, 0x3e, 0xe2, 0xe2, 0x3c, 0x65, 0xd0,
	0x0d, 0xc8, 0xcb, 0x90, 0x70, 0x72, 0xf1, 0xb9, 0x94, 0x5e, 0x7a, 0x2e, 0x15, 0x19, 0x0e, 0xa2,
	0x58, 0x2b, 0xc4, 0x5d, 0x25, 0x02, 0x72, 0x88, 0xd8, 0xe4, 0xab, 0x2e, 0x7c, 0xb3, 0x14, 0xff,
	0x3f, 0x39, 0xfd, 0x12, 0x00, 0x00, 0xff, 0xff, 0x2d, 0x93, 0x4f, 0x0d, 0x15, 0x09, 0x00, 0x00,
}
