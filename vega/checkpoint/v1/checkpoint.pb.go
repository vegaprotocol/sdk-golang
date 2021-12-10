// Code generated by protoc-gen-go. DO NOT EDIT.
// source: vega/checkpoint/v1/checkpoint.proto

package v1

import (
	fmt "fmt"
	math "math"

	vega "code.vegaprotocol.io/sdk-golang/vega"
	proto "github.com/golang/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// CheckpointState is the entire checkpoint serialised (basically serialised the Checkpoint message + hash)
type CheckpointState struct {
	Hash                 []byte   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	State                []byte   `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckpointState) Reset()         { *m = CheckpointState{} }
func (m *CheckpointState) String() string { return proto.CompactTextString(m) }
func (*CheckpointState) ProtoMessage()    {}
func (*CheckpointState) Descriptor() ([]byte, []int) {
	return fileDescriptor_692b4d714b180f82, []int{0}
}

func (m *CheckpointState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckpointState.Unmarshal(m, b)
}
func (m *CheckpointState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckpointState.Marshal(b, m, deterministic)
}
func (m *CheckpointState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckpointState.Merge(m, src)
}
func (m *CheckpointState) XXX_Size() int {
	return xxx_messageInfo_CheckpointState.Size(m)
}
func (m *CheckpointState) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckpointState.DiscardUnknown(m)
}

var xxx_messageInfo_CheckpointState proto.InternalMessageInfo

func (m *CheckpointState) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *CheckpointState) GetState() []byte {
	if m != nil {
		return m.State
	}
	return nil
}

// Checkpoint aggregates the various engine snapshots
type Checkpoint struct {
	Governance           []byte   `protobuf:"bytes,1,opt,name=governance,proto3" json:"governance,omitempty"`
	Assets               []byte   `protobuf:"bytes,2,opt,name=assets,proto3" json:"assets,omitempty"`
	Collateral           []byte   `protobuf:"bytes,3,opt,name=collateral,proto3" json:"collateral,omitempty"`
	NetworkParameters    []byte   `protobuf:"bytes,4,opt,name=network_parameters,json=networkParameters,proto3" json:"network_parameters,omitempty"`
	Delegation           []byte   `protobuf:"bytes,5,opt,name=delegation,proto3" json:"delegation,omitempty"`
	Epoch                []byte   `protobuf:"bytes,6,opt,name=epoch,proto3" json:"epoch,omitempty"`
	Block                []byte   `protobuf:"bytes,7,opt,name=block,proto3" json:"block,omitempty"`
	Rewards              []byte   `protobuf:"bytes,8,opt,name=rewards,proto3" json:"rewards,omitempty"`
	KeyRotations         []byte   `protobuf:"bytes,9,opt,name=key_rotations,json=keyRotations,proto3" json:"key_rotations,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Checkpoint) Reset()         { *m = Checkpoint{} }
func (m *Checkpoint) String() string { return proto.CompactTextString(m) }
func (*Checkpoint) ProtoMessage()    {}
func (*Checkpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_692b4d714b180f82, []int{1}
}

func (m *Checkpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Checkpoint.Unmarshal(m, b)
}
func (m *Checkpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Checkpoint.Marshal(b, m, deterministic)
}
func (m *Checkpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Checkpoint.Merge(m, src)
}
func (m *Checkpoint) XXX_Size() int {
	return xxx_messageInfo_Checkpoint.Size(m)
}
func (m *Checkpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_Checkpoint.DiscardUnknown(m)
}

var xxx_messageInfo_Checkpoint proto.InternalMessageInfo

func (m *Checkpoint) GetGovernance() []byte {
	if m != nil {
		return m.Governance
	}
	return nil
}

func (m *Checkpoint) GetAssets() []byte {
	if m != nil {
		return m.Assets
	}
	return nil
}

func (m *Checkpoint) GetCollateral() []byte {
	if m != nil {
		return m.Collateral
	}
	return nil
}

func (m *Checkpoint) GetNetworkParameters() []byte {
	if m != nil {
		return m.NetworkParameters
	}
	return nil
}

func (m *Checkpoint) GetDelegation() []byte {
	if m != nil {
		return m.Delegation
	}
	return nil
}

func (m *Checkpoint) GetEpoch() []byte {
	if m != nil {
		return m.Epoch
	}
	return nil
}

func (m *Checkpoint) GetBlock() []byte {
	if m != nil {
		return m.Block
	}
	return nil
}

func (m *Checkpoint) GetRewards() []byte {
	if m != nil {
		return m.Rewards
	}
	return nil
}

func (m *Checkpoint) GetKeyRotations() []byte {
	if m != nil {
		return m.KeyRotations
	}
	return nil
}

// AssetEntrty is a single (enabled) asset
type AssetEntry struct {
	Id                   string             `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	AssetDetails         *vega.AssetDetails `protobuf:"bytes,2,opt,name=asset_details,json=assetDetails,proto3" json:"asset_details,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *AssetEntry) Reset()         { *m = AssetEntry{} }
func (m *AssetEntry) String() string { return proto.CompactTextString(m) }
func (*AssetEntry) ProtoMessage()    {}
func (*AssetEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_692b4d714b180f82, []int{2}
}

func (m *AssetEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AssetEntry.Unmarshal(m, b)
}
func (m *AssetEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AssetEntry.Marshal(b, m, deterministic)
}
func (m *AssetEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssetEntry.Merge(m, src)
}
func (m *AssetEntry) XXX_Size() int {
	return xxx_messageInfo_AssetEntry.Size(m)
}
func (m *AssetEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_AssetEntry.DiscardUnknown(m)
}

var xxx_messageInfo_AssetEntry proto.InternalMessageInfo

func (m *AssetEntry) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *AssetEntry) GetAssetDetails() *vega.AssetDetails {
	if m != nil {
		return m.AssetDetails
	}
	return nil
}

// Assets contains all the enabled assets as AssetEntries
type Assets struct {
	Assets               []*AssetEntry `protobuf:"bytes,1,rep,name=assets,proto3" json:"assets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Assets) Reset()         { *m = Assets{} }
func (m *Assets) String() string { return proto.CompactTextString(m) }
func (*Assets) ProtoMessage()    {}
func (*Assets) Descriptor() ([]byte, []int) {
	return fileDescriptor_692b4d714b180f82, []int{3}
}

func (m *Assets) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Assets.Unmarshal(m, b)
}
func (m *Assets) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Assets.Marshal(b, m, deterministic)
}
func (m *Assets) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Assets.Merge(m, src)
}
func (m *Assets) XXX_Size() int {
	return xxx_messageInfo_Assets.Size(m)
}
func (m *Assets) XXX_DiscardUnknown() {
	xxx_messageInfo_Assets.DiscardUnknown(m)
}

var xxx_messageInfo_Assets proto.InternalMessageInfo

func (m *Assets) GetAssets() []*AssetEntry {
	if m != nil {
		return m.Assets
	}
	return nil
}

// AssetBalance represents the total balance of a given asset for a party
type AssetBalance struct {
	Party                string   `protobuf:"bytes,1,opt,name=party,proto3" json:"party,omitempty"`
	Asset                string   `protobuf:"bytes,2,opt,name=asset,proto3" json:"asset,omitempty"`
	Balance              string   `protobuf:"bytes,3,opt,name=balance,proto3" json:"balance,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AssetBalance) Reset()         { *m = AssetBalance{} }
func (m *AssetBalance) String() string { return proto.CompactTextString(m) }
func (*AssetBalance) ProtoMessage()    {}
func (*AssetBalance) Descriptor() ([]byte, []int) {
	return fileDescriptor_692b4d714b180f82, []int{4}
}

func (m *AssetBalance) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AssetBalance.Unmarshal(m, b)
}
func (m *AssetBalance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AssetBalance.Marshal(b, m, deterministic)
}
func (m *AssetBalance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssetBalance.Merge(m, src)
}
func (m *AssetBalance) XXX_Size() int {
	return xxx_messageInfo_AssetBalance.Size(m)
}
func (m *AssetBalance) XXX_DiscardUnknown() {
	xxx_messageInfo_AssetBalance.DiscardUnknown(m)
}

var xxx_messageInfo_AssetBalance proto.InternalMessageInfo

func (m *AssetBalance) GetParty() string {
	if m != nil {
		return m.Party
	}
	return ""
}

func (m *AssetBalance) GetAsset() string {
	if m != nil {
		return m.Asset
	}
	return ""
}

func (m *AssetBalance) GetBalance() string {
	if m != nil {
		return m.Balance
	}
	return ""
}

// Collateral contains the balances per party
type Collateral struct {
	Balances             []*AssetBalance `protobuf:"bytes,1,rep,name=balances,proto3" json:"balances,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Collateral) Reset()         { *m = Collateral{} }
func (m *Collateral) String() string { return proto.CompactTextString(m) }
func (*Collateral) ProtoMessage()    {}
func (*Collateral) Descriptor() ([]byte, []int) {
	return fileDescriptor_692b4d714b180f82, []int{5}
}

func (m *Collateral) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Collateral.Unmarshal(m, b)
}
func (m *Collateral) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Collateral.Marshal(b, m, deterministic)
}
func (m *Collateral) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Collateral.Merge(m, src)
}
func (m *Collateral) XXX_Size() int {
	return xxx_messageInfo_Collateral.Size(m)
}
func (m *Collateral) XXX_DiscardUnknown() {
	xxx_messageInfo_Collateral.DiscardUnknown(m)
}

var xxx_messageInfo_Collateral proto.InternalMessageInfo

func (m *Collateral) GetBalances() []*AssetBalance {
	if m != nil {
		return m.Balances
	}
	return nil
}

// NetParams contains all network parameters
type NetParams struct {
	Params               []*vega.NetworkParameter `protobuf:"bytes,1,rep,name=params,proto3" json:"params,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *NetParams) Reset()         { *m = NetParams{} }
func (m *NetParams) String() string { return proto.CompactTextString(m) }
func (*NetParams) ProtoMessage()    {}
func (*NetParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_692b4d714b180f82, []int{6}
}

func (m *NetParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetParams.Unmarshal(m, b)
}
func (m *NetParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetParams.Marshal(b, m, deterministic)
}
func (m *NetParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetParams.Merge(m, src)
}
func (m *NetParams) XXX_Size() int {
	return xxx_messageInfo_NetParams.Size(m)
}
func (m *NetParams) XXX_DiscardUnknown() {
	xxx_messageInfo_NetParams.DiscardUnknown(m)
}

var xxx_messageInfo_NetParams proto.InternalMessageInfo

func (m *NetParams) GetParams() []*vega.NetworkParameter {
	if m != nil {
		return m.Params
	}
	return nil
}

// Proposals will contain all accepted proposals
type Proposals struct {
	Proposals            []*vega.Proposal `protobuf:"bytes,1,rep,name=proposals,proto3" json:"proposals,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Proposals) Reset()         { *m = Proposals{} }
func (m *Proposals) String() string { return proto.CompactTextString(m) }
func (*Proposals) ProtoMessage()    {}
func (*Proposals) Descriptor() ([]byte, []int) {
	return fileDescriptor_692b4d714b180f82, []int{7}
}

func (m *Proposals) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Proposals.Unmarshal(m, b)
}
func (m *Proposals) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Proposals.Marshal(b, m, deterministic)
}
func (m *Proposals) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Proposals.Merge(m, src)
}
func (m *Proposals) XXX_Size() int {
	return xxx_messageInfo_Proposals.Size(m)
}
func (m *Proposals) XXX_DiscardUnknown() {
	xxx_messageInfo_Proposals.DiscardUnknown(m)
}

var xxx_messageInfo_Proposals proto.InternalMessageInfo

func (m *Proposals) GetProposals() []*vega.Proposal {
	if m != nil {
		return m.Proposals
	}
	return nil
}

// Delegated amounts for party/node
// undelegate and epoch seq are only relevant for pending entries
type DelegateEntry struct {
	Party                string   `protobuf:"bytes,1,opt,name=party,proto3" json:"party,omitempty"`
	Node                 string   `protobuf:"bytes,2,opt,name=node,proto3" json:"node,omitempty"`
	Amount               string   `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Undelegate           bool     `protobuf:"varint,4,opt,name=undelegate,proto3" json:"undelegate,omitempty"`
	EpochSeq             uint64   `protobuf:"varint,5,opt,name=epoch_seq,json=epochSeq,proto3" json:"epoch_seq,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DelegateEntry) Reset()         { *m = DelegateEntry{} }
func (m *DelegateEntry) String() string { return proto.CompactTextString(m) }
func (*DelegateEntry) ProtoMessage()    {}
func (*DelegateEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_692b4d714b180f82, []int{8}
}

func (m *DelegateEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DelegateEntry.Unmarshal(m, b)
}
func (m *DelegateEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DelegateEntry.Marshal(b, m, deterministic)
}
func (m *DelegateEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelegateEntry.Merge(m, src)
}
func (m *DelegateEntry) XXX_Size() int {
	return xxx_messageInfo_DelegateEntry.Size(m)
}
func (m *DelegateEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_DelegateEntry.DiscardUnknown(m)
}

var xxx_messageInfo_DelegateEntry proto.InternalMessageInfo

func (m *DelegateEntry) GetParty() string {
	if m != nil {
		return m.Party
	}
	return ""
}

func (m *DelegateEntry) GetNode() string {
	if m != nil {
		return m.Node
	}
	return ""
}

func (m *DelegateEntry) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

func (m *DelegateEntry) GetUndelegate() bool {
	if m != nil {
		return m.Undelegate
	}
	return false
}

func (m *DelegateEntry) GetEpochSeq() uint64 {
	if m != nil {
		return m.EpochSeq
	}
	return 0
}

// Delegate contains all entries for a checkpoint
type Delegate struct {
	Active               []*DelegateEntry `protobuf:"bytes,1,rep,name=active,proto3" json:"active,omitempty"`
	Pending              []*DelegateEntry `protobuf:"bytes,2,rep,name=pending,proto3" json:"pending,omitempty"`
	AutoDelegation       []string         `protobuf:"bytes,3,rep,name=auto_delegation,json=autoDelegation,proto3" json:"auto_delegation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Delegate) Reset()         { *m = Delegate{} }
func (m *Delegate) String() string { return proto.CompactTextString(m) }
func (*Delegate) ProtoMessage()    {}
func (*Delegate) Descriptor() ([]byte, []int) {
	return fileDescriptor_692b4d714b180f82, []int{9}
}

func (m *Delegate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Delegate.Unmarshal(m, b)
}
func (m *Delegate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Delegate.Marshal(b, m, deterministic)
}
func (m *Delegate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Delegate.Merge(m, src)
}
func (m *Delegate) XXX_Size() int {
	return xxx_messageInfo_Delegate.Size(m)
}
func (m *Delegate) XXX_DiscardUnknown() {
	xxx_messageInfo_Delegate.DiscardUnknown(m)
}

var xxx_messageInfo_Delegate proto.InternalMessageInfo

func (m *Delegate) GetActive() []*DelegateEntry {
	if m != nil {
		return m.Active
	}
	return nil
}

func (m *Delegate) GetPending() []*DelegateEntry {
	if m != nil {
		return m.Pending
	}
	return nil
}

func (m *Delegate) GetAutoDelegation() []string {
	if m != nil {
		return m.AutoDelegation
	}
	return nil
}

// Block message contains data related to block at which the checkpoint
// was created (ie block height)
type Block struct {
	Height               int64    `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Block) Reset()         { *m = Block{} }
func (m *Block) String() string { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()    {}
func (*Block) Descriptor() ([]byte, []int) {
	return fileDescriptor_692b4d714b180f82, []int{10}
}

func (m *Block) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Block.Unmarshal(m, b)
}
func (m *Block) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Block.Marshal(b, m, deterministic)
}
func (m *Block) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Block.Merge(m, src)
}
func (m *Block) XXX_Size() int {
	return xxx_messageInfo_Block.Size(m)
}
func (m *Block) XXX_DiscardUnknown() {
	xxx_messageInfo_Block.DiscardUnknown(m)
}

var xxx_messageInfo_Block proto.InternalMessageInfo

func (m *Block) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

type Rewards struct {
	Rewards              []*RewardPayout `protobuf:"bytes,1,rep,name=rewards,proto3" json:"rewards,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Rewards) Reset()         { *m = Rewards{} }
func (m *Rewards) String() string { return proto.CompactTextString(m) }
func (*Rewards) ProtoMessage()    {}
func (*Rewards) Descriptor() ([]byte, []int) {
	return fileDescriptor_692b4d714b180f82, []int{11}
}

func (m *Rewards) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Rewards.Unmarshal(m, b)
}
func (m *Rewards) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Rewards.Marshal(b, m, deterministic)
}
func (m *Rewards) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Rewards.Merge(m, src)
}
func (m *Rewards) XXX_Size() int {
	return xxx_messageInfo_Rewards.Size(m)
}
func (m *Rewards) XXX_DiscardUnknown() {
	xxx_messageInfo_Rewards.DiscardUnknown(m)
}

var xxx_messageInfo_Rewards proto.InternalMessageInfo

func (m *Rewards) GetRewards() []*RewardPayout {
	if m != nil {
		return m.Rewards
	}
	return nil
}

type RewardPayout struct {
	PayoutTime           int64                  `protobuf:"varint,1,opt,name=payout_time,json=payoutTime,proto3" json:"payout_time,omitempty"`
	RewardsPayout        []*PendingRewardPayout `protobuf:"bytes,2,rep,name=rewards_payout,json=rewardsPayout,proto3" json:"rewards_payout,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *RewardPayout) Reset()         { *m = RewardPayout{} }
func (m *RewardPayout) String() string { return proto.CompactTextString(m) }
func (*RewardPayout) ProtoMessage()    {}
func (*RewardPayout) Descriptor() ([]byte, []int) {
	return fileDescriptor_692b4d714b180f82, []int{12}
}

func (m *RewardPayout) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RewardPayout.Unmarshal(m, b)
}
func (m *RewardPayout) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RewardPayout.Marshal(b, m, deterministic)
}
func (m *RewardPayout) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RewardPayout.Merge(m, src)
}
func (m *RewardPayout) XXX_Size() int {
	return xxx_messageInfo_RewardPayout.Size(m)
}
func (m *RewardPayout) XXX_DiscardUnknown() {
	xxx_messageInfo_RewardPayout.DiscardUnknown(m)
}

var xxx_messageInfo_RewardPayout proto.InternalMessageInfo

func (m *RewardPayout) GetPayoutTime() int64 {
	if m != nil {
		return m.PayoutTime
	}
	return 0
}

func (m *RewardPayout) GetRewardsPayout() []*PendingRewardPayout {
	if m != nil {
		return m.RewardsPayout
	}
	return nil
}

type PendingRewardPayout struct {
	FromAccount          string         `protobuf:"bytes,1,opt,name=from_account,json=fromAccount,proto3" json:"from_account,omitempty"`
	Asset                string         `protobuf:"bytes,2,opt,name=asset,proto3" json:"asset,omitempty"`
	PartyAmount          []*PartyAmount `protobuf:"bytes,3,rep,name=party_amount,json=partyAmount,proto3" json:"party_amount,omitempty"`
	TotalReward          string         `protobuf:"bytes,4,opt,name=total_reward,json=totalReward,proto3" json:"total_reward,omitempty"`
	EpochSeq             string         `protobuf:"bytes,5,opt,name=epoch_seq,json=epochSeq,proto3" json:"epoch_seq,omitempty"`
	Timestamp            int64          `protobuf:"varint,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *PendingRewardPayout) Reset()         { *m = PendingRewardPayout{} }
func (m *PendingRewardPayout) String() string { return proto.CompactTextString(m) }
func (*PendingRewardPayout) ProtoMessage()    {}
func (*PendingRewardPayout) Descriptor() ([]byte, []int) {
	return fileDescriptor_692b4d714b180f82, []int{13}
}

func (m *PendingRewardPayout) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PendingRewardPayout.Unmarshal(m, b)
}
func (m *PendingRewardPayout) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PendingRewardPayout.Marshal(b, m, deterministic)
}
func (m *PendingRewardPayout) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PendingRewardPayout.Merge(m, src)
}
func (m *PendingRewardPayout) XXX_Size() int {
	return xxx_messageInfo_PendingRewardPayout.Size(m)
}
func (m *PendingRewardPayout) XXX_DiscardUnknown() {
	xxx_messageInfo_PendingRewardPayout.DiscardUnknown(m)
}

var xxx_messageInfo_PendingRewardPayout proto.InternalMessageInfo

func (m *PendingRewardPayout) GetFromAccount() string {
	if m != nil {
		return m.FromAccount
	}
	return ""
}

func (m *PendingRewardPayout) GetAsset() string {
	if m != nil {
		return m.Asset
	}
	return ""
}

func (m *PendingRewardPayout) GetPartyAmount() []*PartyAmount {
	if m != nil {
		return m.PartyAmount
	}
	return nil
}

func (m *PendingRewardPayout) GetTotalReward() string {
	if m != nil {
		return m.TotalReward
	}
	return ""
}

func (m *PendingRewardPayout) GetEpochSeq() string {
	if m != nil {
		return m.EpochSeq
	}
	return ""
}

func (m *PendingRewardPayout) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type PartyAmount struct {
	Party                string   `protobuf:"bytes,1,opt,name=party,proto3" json:"party,omitempty"`
	Amount               string   `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PartyAmount) Reset()         { *m = PartyAmount{} }
func (m *PartyAmount) String() string { return proto.CompactTextString(m) }
func (*PartyAmount) ProtoMessage()    {}
func (*PartyAmount) Descriptor() ([]byte, []int) {
	return fileDescriptor_692b4d714b180f82, []int{14}
}

func (m *PartyAmount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PartyAmount.Unmarshal(m, b)
}
func (m *PartyAmount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PartyAmount.Marshal(b, m, deterministic)
}
func (m *PartyAmount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PartyAmount.Merge(m, src)
}
func (m *PartyAmount) XXX_Size() int {
	return xxx_messageInfo_PartyAmount.Size(m)
}
func (m *PartyAmount) XXX_DiscardUnknown() {
	xxx_messageInfo_PartyAmount.DiscardUnknown(m)
}

var xxx_messageInfo_PartyAmount proto.InternalMessageInfo

func (m *PartyAmount) GetParty() string {
	if m != nil {
		return m.Party
	}
	return ""
}

func (m *PartyAmount) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

type PendingKeyRotation struct {
	// Relative target block height is: target block height - current block height.
	// Useful for cross blockchain compatibility.
	RelativeTargetBlockHeight uint64   `protobuf:"varint,1,opt,name=relative_target_block_height,json=relativeTargetBlockHeight,proto3" json:"relative_target_block_height,omitempty"`
	NodeId                    string   `protobuf:"bytes,2,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	NewPubKey                 string   `protobuf:"bytes,3,opt,name=new_pub_key,json=newPubKey,proto3" json:"new_pub_key,omitempty"`
	NewPubKeyIndex            uint32   `protobuf:"varint,4,opt,name=new_pub_key_index,json=newPubKeyIndex,proto3" json:"new_pub_key_index,omitempty"`
	XXX_NoUnkeyedLiteral      struct{} `json:"-"`
	XXX_unrecognized          []byte   `json:"-"`
	XXX_sizecache             int32    `json:"-"`
}

func (m *PendingKeyRotation) Reset()         { *m = PendingKeyRotation{} }
func (m *PendingKeyRotation) String() string { return proto.CompactTextString(m) }
func (*PendingKeyRotation) ProtoMessage()    {}
func (*PendingKeyRotation) Descriptor() ([]byte, []int) {
	return fileDescriptor_692b4d714b180f82, []int{15}
}

func (m *PendingKeyRotation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PendingKeyRotation.Unmarshal(m, b)
}
func (m *PendingKeyRotation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PendingKeyRotation.Marshal(b, m, deterministic)
}
func (m *PendingKeyRotation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PendingKeyRotation.Merge(m, src)
}
func (m *PendingKeyRotation) XXX_Size() int {
	return xxx_messageInfo_PendingKeyRotation.Size(m)
}
func (m *PendingKeyRotation) XXX_DiscardUnknown() {
	xxx_messageInfo_PendingKeyRotation.DiscardUnknown(m)
}

var xxx_messageInfo_PendingKeyRotation proto.InternalMessageInfo

func (m *PendingKeyRotation) GetRelativeTargetBlockHeight() uint64 {
	if m != nil {
		return m.RelativeTargetBlockHeight
	}
	return 0
}

func (m *PendingKeyRotation) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *PendingKeyRotation) GetNewPubKey() string {
	if m != nil {
		return m.NewPubKey
	}
	return ""
}

func (m *PendingKeyRotation) GetNewPubKeyIndex() uint32 {
	if m != nil {
		return m.NewPubKeyIndex
	}
	return 0
}

type KeyRotations struct {
	PendingKeyRotations  []*PendingKeyRotation `protobuf:"bytes,1,rep,name=pending_key_rotations,json=pendingKeyRotations,proto3" json:"pending_key_rotations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *KeyRotations) Reset()         { *m = KeyRotations{} }
func (m *KeyRotations) String() string { return proto.CompactTextString(m) }
func (*KeyRotations) ProtoMessage()    {}
func (*KeyRotations) Descriptor() ([]byte, []int) {
	return fileDescriptor_692b4d714b180f82, []int{16}
}

func (m *KeyRotations) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyRotations.Unmarshal(m, b)
}
func (m *KeyRotations) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyRotations.Marshal(b, m, deterministic)
}
func (m *KeyRotations) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyRotations.Merge(m, src)
}
func (m *KeyRotations) XXX_Size() int {
	return xxx_messageInfo_KeyRotations.Size(m)
}
func (m *KeyRotations) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyRotations.DiscardUnknown(m)
}

var xxx_messageInfo_KeyRotations proto.InternalMessageInfo

func (m *KeyRotations) GetPendingKeyRotations() []*PendingKeyRotation {
	if m != nil {
		return m.PendingKeyRotations
	}
	return nil
}

func init() {
	proto.RegisterType((*CheckpointState)(nil), "vega.checkpoint.v1.CheckpointState")
	proto.RegisterType((*Checkpoint)(nil), "vega.checkpoint.v1.Checkpoint")
	proto.RegisterType((*AssetEntry)(nil), "vega.checkpoint.v1.AssetEntry")
	proto.RegisterType((*Assets)(nil), "vega.checkpoint.v1.Assets")
	proto.RegisterType((*AssetBalance)(nil), "vega.checkpoint.v1.AssetBalance")
	proto.RegisterType((*Collateral)(nil), "vega.checkpoint.v1.Collateral")
	proto.RegisterType((*NetParams)(nil), "vega.checkpoint.v1.NetParams")
	proto.RegisterType((*Proposals)(nil), "vega.checkpoint.v1.Proposals")
	proto.RegisterType((*DelegateEntry)(nil), "vega.checkpoint.v1.DelegateEntry")
	proto.RegisterType((*Delegate)(nil), "vega.checkpoint.v1.Delegate")
	proto.RegisterType((*Block)(nil), "vega.checkpoint.v1.Block")
	proto.RegisterType((*Rewards)(nil), "vega.checkpoint.v1.Rewards")
	proto.RegisterType((*RewardPayout)(nil), "vega.checkpoint.v1.RewardPayout")
	proto.RegisterType((*PendingRewardPayout)(nil), "vega.checkpoint.v1.PendingRewardPayout")
	proto.RegisterType((*PartyAmount)(nil), "vega.checkpoint.v1.PartyAmount")
	proto.RegisterType((*PendingKeyRotation)(nil), "vega.checkpoint.v1.PendingKeyRotation")
	proto.RegisterType((*KeyRotations)(nil), "vega.checkpoint.v1.KeyRotations")
}

func init() {
	proto.RegisterFile("vega/checkpoint/v1/checkpoint.proto", fileDescriptor_692b4d714b180f82)
}

var fileDescriptor_692b4d714b180f82 = []byte{
	// 930 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x56, 0x5d, 0x6e, 0xdb, 0x46,
	0x10, 0x86, 0x7e, 0x2c, 0x89, 0xa3, 0x1f, 0xc3, 0x9b, 0xc4, 0x65, 0xd3, 0xc0, 0x76, 0x18, 0xa0,
	0x71, 0x81, 0x46, 0x46, 0xdc, 0xa2, 0x45, 0xea, 0x02, 0xad, 0x1d, 0x07, 0x68, 0x1a, 0xc0, 0x10,
	0x36, 0xee, 0x4b, 0x5e, 0x88, 0x15, 0x39, 0x95, 0x58, 0x51, 0x5c, 0x86, 0x5c, 0xc9, 0xd5, 0x53,
	0x6f, 0xd0, 0x23, 0xf4, 0x02, 0x3d, 0x41, 0x6f, 0xd5, 0x23, 0x14, 0x3b, 0xbb, 0x14, 0xe9, 0x4a,
	0x0e, 0xf2, 0xc6, 0xf9, 0xe6, 0x9b, 0xd9, 0xe1, 0xb7, 0xdf, 0x50, 0x82, 0x27, 0x4b, 0x9c, 0x88,
	0x93, 0x60, 0x8a, 0xc1, 0x2c, 0x95, 0x51, 0xa2, 0x4e, 0x96, 0xcf, 0x2b, 0xd1, 0x30, 0xcd, 0xa4,
	0x92, 0x8c, 0x69, 0xd2, 0xb0, 0x02, 0x2f, 0x9f, 0x3f, 0xdc, 0xa5, 0x42, 0x4a, 0x10, 0xe9, 0xe1,
	0x1e, 0x01, 0x22, 0xcf, 0x51, 0xe5, 0x16, 0x7a, 0x40, 0xd0, 0x44, 0x2e, 0x31, 0x4b, 0x44, 0x12,
	0xa0, 0x81, 0xbd, 0x33, 0xd8, 0x7d, 0xb9, 0xee, 0xf5, 0x56, 0x09, 0x85, 0x8c, 0x41, 0x73, 0x2a,
	0xf2, 0xa9, 0x5b, 0x3b, 0xaa, 0x1d, 0xf7, 0x38, 0x3d, 0xb3, 0xfb, 0xb0, 0x93, 0xeb, 0xa4, 0x5b,
	0x27, 0xd0, 0x04, 0xde, 0x5f, 0x75, 0x80, 0xb2, 0x9a, 0x1d, 0x00, 0x94, 0xfd, 0x6d, 0x79, 0x05,
	0x61, 0xfb, 0xd0, 0x32, 0x23, 0xd9, 0x2e, 0x36, 0xd2, 0x75, 0x81, 0x8c, 0x63, 0xa1, 0x30, 0x13,
	0xb1, 0xdb, 0x30, 0x75, 0x25, 0xc2, 0x9e, 0x01, 0x4b, 0x50, 0xdd, 0xc8, 0x6c, 0xe6, 0xa7, 0x22,
	0x13, 0x73, 0x54, 0x98, 0xe5, 0x6e, 0x93, 0x78, 0x7b, 0x36, 0x33, 0x5a, 0x27, 0x74, 0xbb, 0x10,
	0x63, 0x9c, 0x08, 0x15, 0xc9, 0xc4, 0xdd, 0x31, 0xed, 0x4a, 0x44, 0xbf, 0x0b, 0xa6, 0x32, 0x98,
	0xba, 0x2d, 0xf3, 0x2e, 0x14, 0x68, 0x74, 0x1c, 0xcb, 0x60, 0xe6, 0xb6, 0x0d, 0x4a, 0x01, 0x73,
	0xa1, 0x9d, 0xe1, 0x8d, 0xc8, 0xc2, 0xdc, 0xed, 0x10, 0x5e, 0x84, 0xec, 0x09, 0xf4, 0x67, 0xb8,
	0xf2, 0x33, 0xa9, 0xa8, 0x6b, 0xee, 0x3a, 0x94, 0xef, 0xcd, 0x70, 0xc5, 0x0b, 0xcc, 0xfb, 0x05,
	0xe0, 0x5c, 0xbf, 0xe3, 0xab, 0x44, 0x65, 0x2b, 0x36, 0x80, 0x7a, 0x14, 0x92, 0x2e, 0x0e, 0xaf,
	0x47, 0x21, 0xfb, 0x16, 0xfa, 0xa4, 0x80, 0x1f, 0xa2, 0x12, 0x51, 0x6c, 0x64, 0xe9, 0x9e, 0xb2,
	0x21, 0xdd, 0x24, 0x15, 0x5e, 0x9a, 0x0c, 0xef, 0x89, 0x4a, 0xe4, 0xfd, 0x08, 0xad, 0x73, 0x23,
	0xdd, 0x37, 0x6b, 0x49, 0x6b, 0x47, 0x8d, 0xe3, 0xee, 0xe9, 0xc1, 0x70, 0xd3, 0x1e, 0xc3, 0x72,
	0x84, 0x42, 0x72, 0xef, 0x1a, 0x7a, 0x84, 0x5e, 0x88, 0x98, 0xae, 0xe6, 0x3e, 0xec, 0xa4, 0x22,
	0x53, 0x2b, 0x3b, 0x9d, 0x09, 0x34, 0x4a, 0x7c, 0x1a, 0xcc, 0xe1, 0x26, 0xd0, 0x9a, 0x8c, 0x4d,
	0x19, 0xdd, 0x95, 0xc3, 0x8b, 0xd0, 0xfb, 0x19, 0xe0, 0x65, 0x79, 0x6d, 0xdf, 0x43, 0xc7, 0x26,
	0x8a, 0xe9, 0x8e, 0xee, 0x9c, 0xce, 0xce, 0xc1, 0xd7, 0x15, 0xde, 0x19, 0x38, 0x57, 0xa8, 0xe8,
	0x5a, 0x73, 0x36, 0x84, 0x16, 0xdd, 0x7c, 0xd1, 0x68, 0xdf, 0x34, 0xba, 0xfa, 0xdf, 0xdd, 0x73,
	0xcb, 0xf2, 0x5e, 0x80, 0x33, 0xca, 0x64, 0x2a, 0x73, 0x11, 0xe7, 0xec, 0x4b, 0x70, 0xd2, 0x22,
	0xb0, 0xf5, 0x03, 0x53, 0x5f, 0x70, 0x78, 0x49, 0xf0, 0xfe, 0xac, 0x41, 0xff, 0xd2, 0x98, 0x05,
	0xcd, 0xb5, 0x6d, 0xd7, 0x86, 0x41, 0x33, 0x91, 0x21, 0x5a, 0x69, 0xe8, 0x99, 0x0c, 0x3e, 0x97,
	0x8b, 0x44, 0x59, 0x61, 0x6c, 0xa4, 0x1d, 0xb9, 0x48, 0xac, 0x03, 0x91, 0x8c, 0xdb, 0xe1, 0x15,
	0x84, 0x7d, 0x06, 0x0e, 0x99, 0xd0, 0xcf, 0xf1, 0x3d, 0x19, 0xb6, 0xc9, 0x3b, 0x04, 0xbc, 0xc5,
	0xf7, 0xde, 0xdf, 0x35, 0xe8, 0x14, 0x03, 0xb1, 0x17, 0xd0, 0x12, 0x81, 0x8a, 0x96, 0x68, 0x5f,
	0xe4, 0xf1, 0x36, 0x45, 0x6f, 0x8d, 0xcf, 0x6d, 0x01, 0x3b, 0x83, 0x76, 0x8a, 0x49, 0x18, 0x25,
	0x13, 0xb7, 0xfe, 0xb1, 0xb5, 0x45, 0x05, 0x7b, 0x0a, 0xbb, 0x62, 0xa1, 0xa4, 0x5f, 0x59, 0xac,
	0xc6, 0x51, 0xe3, 0xd8, 0xe1, 0x03, 0x0d, 0x5f, 0xae, 0x51, 0xef, 0x10, 0x76, 0x2e, 0x68, 0x73,
	0xf6, 0xa1, 0x35, 0xc5, 0x68, 0x32, 0x55, 0x24, 0x5b, 0x83, 0xdb, 0xc8, 0x7b, 0x05, 0x6d, 0x6e,
	0x57, 0xe8, 0xbb, 0x72, 0xb9, 0x3e, 0xe0, 0x0f, 0xc3, 0x1e, 0x89, 0x95, 0x5c, 0xa8, 0xf5, 0xfa,
	0x79, 0x7f, 0x40, 0xaf, 0x9a, 0x60, 0x87, 0xd0, 0x4d, 0xe9, 0xc9, 0x57, 0xd1, 0x1c, 0xed, 0x99,
	0x60, 0xa0, 0xeb, 0x68, 0x8e, 0xec, 0x0a, 0x06, 0xb6, 0xd6, 0x37, 0xa8, 0x55, 0xe1, 0xe9, 0xb6,
	0x33, 0x47, 0xe6, 0xb5, 0x6f, 0x1d, 0xdd, 0xb7, 0xe5, 0x26, 0xf4, 0xfe, 0xad, 0xc1, 0xbd, 0x2d,
	0x34, 0xf6, 0x18, 0x7a, 0xbf, 0x66, 0x72, 0xee, 0x8b, 0x20, 0x20, 0x27, 0x18, 0xd3, 0x74, 0x35,
	0x76, 0x6e, 0xa0, 0x3b, 0xd6, 0xea, 0x02, 0x7a, 0xe4, 0x2c, 0x7f, 0x6d, 0x21, 0x3d, 0xde, 0xe1,
	0xd6, 0xf1, 0x34, 0xef, 0x9c, 0x68, 0xbc, 0x9b, 0x96, 0x81, 0x3e, 0x5c, 0x49, 0x25, 0x62, 0xdf,
	0xcc, 0x4a, 0x56, 0x73, 0x78, 0x97, 0x30, 0x33, 0xe5, 0xa6, 0xd7, 0x9c, 0xd2, 0x6b, 0xec, 0x11,
	0x38, 0x5a, 0xbe, 0x5c, 0x89, 0x79, 0x4a, 0x9f, 0xc7, 0x06, 0x2f, 0x01, 0xef, 0x0c, 0xba, 0x95,
	0x93, 0xef, 0xd8, 0x8b, 0x72, 0x07, 0xea, 0xd5, 0x1d, 0xf0, 0xfe, 0xa9, 0x01, 0xb3, 0x7a, 0xbd,
	0x29, 0x3f, 0x91, 0xec, 0x07, 0x78, 0x94, 0x61, 0x2c, 0xb4, 0x43, 0x7d, 0x25, 0xb2, 0x09, 0x2a,
	0x9f, 0xbe, 0xbc, 0x7e, 0xc5, 0x3c, 0x4d, 0xfe, 0x69, 0xc1, 0xb9, 0x26, 0x0a, 0x39, 0xec, 0x27,
	0x22, 0xb0, 0x4f, 0xa0, 0xad, 0x77, 0xcf, 0x8f, 0xc2, 0xe2, 0x40, 0x1d, 0xbe, 0x0e, 0xd9, 0x01,
	0x74, 0x13, 0xbc, 0xf1, 0xd3, 0xc5, 0xd8, 0x9f, 0xe1, 0xca, 0x6e, 0xa4, 0x93, 0xe0, 0xcd, 0x68,
	0x31, 0x7e, 0x83, 0x2b, 0xf6, 0x05, 0xec, 0x55, 0xf2, 0x7e, 0x94, 0x84, 0xf8, 0x3b, 0x09, 0xd6,
	0xe7, 0x83, 0x35, 0xeb, 0xb5, 0x46, 0xbd, 0xdf, 0xa0, 0x57, 0x99, 0x39, 0x67, 0xef, 0xe0, 0x81,
	0x5d, 0x0c, 0xff, 0xf6, 0x6f, 0x80, 0xb1, 0xf1, 0xe7, 0x1f, 0xb0, 0x54, 0xa5, 0x0f, 0xbf, 0x97,
	0x6e, 0x60, 0xf9, 0xc5, 0xd7, 0xef, 0x4e, 0x03, 0x19, 0x22, 0xb5, 0xa0, 0x9f, 0xe8, 0x40, 0xc6,
	0xc3, 0x48, 0x9e, 0xe4, 0xe1, 0xec, 0xd9, 0x44, 0xc6, 0x22, 0x99, 0x9c, 0x6c, 0xfe, 0x4d, 0x18,
	0xb7, 0x88, 0xfa, 0xd5, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x48, 0x06, 0x5b, 0x55, 0x43, 0x08,
	0x00, 0x00,
}
