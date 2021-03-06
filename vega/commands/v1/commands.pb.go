// Code generated by protoc-gen-go. DO NOT EDIT.
// source: vega/commands/v1/commands.proto

package v1

import (
	fmt "fmt"
	math "math"

	vega "code.vegaprotocol.io/sdk-golang/vega"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/mwitkow/go-proto-validators"
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

type UndelegateSubmission_Method int32

const (
	UndelegateSubmission_METHOD_UNSPECIFIED     UndelegateSubmission_Method = 0
	UndelegateSubmission_METHOD_NOW             UndelegateSubmission_Method = 1
	UndelegateSubmission_METHOD_AT_END_OF_EPOCH UndelegateSubmission_Method = 2
	UndelegateSubmission_METHOD_IN_ANGER        UndelegateSubmission_Method = 3
)

var UndelegateSubmission_Method_name = map[int32]string{
	0: "METHOD_UNSPECIFIED",
	1: "METHOD_NOW",
	2: "METHOD_AT_END_OF_EPOCH",
	3: "METHOD_IN_ANGER",
}

var UndelegateSubmission_Method_value = map[string]int32{
	"METHOD_UNSPECIFIED":     0,
	"METHOD_NOW":             1,
	"METHOD_AT_END_OF_EPOCH": 2,
	"METHOD_IN_ANGER":        3,
}

func (x UndelegateSubmission_Method) String() string {
	return proto.EnumName(UndelegateSubmission_Method_name, int32(x))
}

func (UndelegateSubmission_Method) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9c3a53c66a940c51, []int{8, 0}
}

// An order submission is a request to submit or create a new order on Vega
type OrderSubmission struct {
	// Market identifier for the order, required field
	MarketId string `protobuf:"bytes,1,opt,name=market_id,json=marketId,proto3" json:"market_id,omitempty"`
	// Price for the order, the price is an integer, for example `123456` is a correctly
	// formatted price of `1.23456` assuming market configured to 5 decimal places,
	// , required field for limit orders, however it is not required for market orders
	Price string `protobuf:"bytes,2,opt,name=price,proto3" json:"price,omitempty"`
	// Size for the order, for example, in a futures market the size equals the number of contracts, cannot be negative
	Size uint64 `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	// Side for the order, e.g. SIDE_BUY or SIDE_SELL, required field - See [`Side`](#vega.Side)
	Side vega.Side `protobuf:"varint,4,opt,name=side,proto3,enum=vega.Side" json:"side,omitempty"`
	// Time in force indicates how long an order will remain active before it is executed or expires, required field
	// - See [`Order.TimeInForce`](#vega.Order.TimeInForce)
	TimeInForce vega.Order_TimeInForce `protobuf:"varint,5,opt,name=time_in_force,json=timeInForce,proto3,enum=vega.Order_TimeInForce" json:"time_in_force,omitempty"`
	// Timestamp for when the order will expire, in nanoseconds since the epoch,
	// required field only for [`Order.TimeInForce`](#vega.Order.TimeInForce)`.TIME_IN_FORCE_GTT`
	// - See [`VegaTimeResponse`](#api.VegaTimeResponse).`timestamp`
	ExpiresAt int64 `protobuf:"varint,6,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	// Type for the order, required field - See [`Order.Type`](#vega.Order.Type)
	Type vega.Order_Type `protobuf:"varint,7,opt,name=type,proto3,enum=vega.Order_Type" json:"type,omitempty"`
	// Reference given for the order, this is typically used to retrieve an order submitted through consensus, currently
	// set internally by the node to return a unique reference identifier for the order submission
	Reference string `protobuf:"bytes,8,opt,name=reference,proto3" json:"reference,omitempty"`
	// Used to specify the details for a pegged order
	// - See [`PeggedOrder`](#vega.PeggedOrder)
	PeggedOrder          *vega.PeggedOrder `protobuf:"bytes,9,opt,name=pegged_order,json=peggedOrder,proto3" json:"pegged_order,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *OrderSubmission) Reset()         { *m = OrderSubmission{} }
func (m *OrderSubmission) String() string { return proto.CompactTextString(m) }
func (*OrderSubmission) ProtoMessage()    {}
func (*OrderSubmission) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c3a53c66a940c51, []int{0}
}

func (m *OrderSubmission) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderSubmission.Unmarshal(m, b)
}
func (m *OrderSubmission) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderSubmission.Marshal(b, m, deterministic)
}
func (m *OrderSubmission) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderSubmission.Merge(m, src)
}
func (m *OrderSubmission) XXX_Size() int {
	return xxx_messageInfo_OrderSubmission.Size(m)
}
func (m *OrderSubmission) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderSubmission.DiscardUnknown(m)
}

var xxx_messageInfo_OrderSubmission proto.InternalMessageInfo

func (m *OrderSubmission) GetMarketId() string {
	if m != nil {
		return m.MarketId
	}
	return ""
}

func (m *OrderSubmission) GetPrice() string {
	if m != nil {
		return m.Price
	}
	return ""
}

func (m *OrderSubmission) GetSize() uint64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *OrderSubmission) GetSide() vega.Side {
	if m != nil {
		return m.Side
	}
	return vega.Side_SIDE_UNSPECIFIED
}

func (m *OrderSubmission) GetTimeInForce() vega.Order_TimeInForce {
	if m != nil {
		return m.TimeInForce
	}
	return vega.Order_TIME_IN_FORCE_UNSPECIFIED
}

func (m *OrderSubmission) GetExpiresAt() int64 {
	if m != nil {
		return m.ExpiresAt
	}
	return 0
}

func (m *OrderSubmission) GetType() vega.Order_Type {
	if m != nil {
		return m.Type
	}
	return vega.Order_TYPE_UNSPECIFIED
}

func (m *OrderSubmission) GetReference() string {
	if m != nil {
		return m.Reference
	}
	return ""
}

func (m *OrderSubmission) GetPeggedOrder() *vega.PeggedOrder {
	if m != nil {
		return m.PeggedOrder
	}
	return nil
}

// An order cancellation is a request to cancel an existing order on Vega
type OrderCancellation struct {
	// Unique identifier for the order (set by the system after consensus), required field
	OrderId string `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	// Market identifier for the order, required field
	MarketId             string   `protobuf:"bytes,2,opt,name=market_id,json=marketId,proto3" json:"market_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderCancellation) Reset()         { *m = OrderCancellation{} }
func (m *OrderCancellation) String() string { return proto.CompactTextString(m) }
func (*OrderCancellation) ProtoMessage()    {}
func (*OrderCancellation) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c3a53c66a940c51, []int{1}
}

func (m *OrderCancellation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderCancellation.Unmarshal(m, b)
}
func (m *OrderCancellation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderCancellation.Marshal(b, m, deterministic)
}
func (m *OrderCancellation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderCancellation.Merge(m, src)
}
func (m *OrderCancellation) XXX_Size() int {
	return xxx_messageInfo_OrderCancellation.Size(m)
}
func (m *OrderCancellation) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderCancellation.DiscardUnknown(m)
}

var xxx_messageInfo_OrderCancellation proto.InternalMessageInfo

func (m *OrderCancellation) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

func (m *OrderCancellation) GetMarketId() string {
	if m != nil {
		return m.MarketId
	}
	return ""
}

// An order amendment is a request to amend or update an existing order on Vega
type OrderAmendment struct {
	// Order identifier, this is required to find the order and will not be updated, required field
	OrderId string `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	// Market identifier, this is required to find the order and will not be updated
	MarketId string `protobuf:"bytes,2,opt,name=market_id,json=marketId,proto3" json:"market_id,omitempty"`
	// Amend the price for the order, if the Price value is set, otherwise price will remain unchanged - See [`Price`](#vega.Price)
	Price *vega.Price `protobuf:"bytes,3,opt,name=price,proto3" json:"price,omitempty"`
	// Amend the size for the order by the delta specified:
	// - To reduce the size from the current value set a negative integer value
	// - To increase the size from the current value, set a positive integer value
	// - To leave the size unchanged set a value of zero
	SizeDelta int64 `protobuf:"varint,4,opt,name=size_delta,json=sizeDelta,proto3" json:"size_delta,omitempty"`
	// Amend the expiry time for the order, if the Timestamp value is set, otherwise expiry time will remain unchanged
	// - See [`VegaTimeResponse`](#api.VegaTimeResponse).`timestamp`
	ExpiresAt *vega.Timestamp `protobuf:"bytes,5,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	// Amend the time in force for the order, set to TIME_IN_FORCE_UNSPECIFIED to remain unchanged
	// - See [`TimeInForce`](#api.VegaTimeResponse).`timestamp`
	TimeInForce vega.Order_TimeInForce `protobuf:"varint,6,opt,name=time_in_force,json=timeInForce,proto3,enum=vega.Order_TimeInForce" json:"time_in_force,omitempty"`
	// Amend the pegged order offset for the order
	PeggedOffset *wrappers.Int64Value `protobuf:"bytes,7,opt,name=pegged_offset,json=peggedOffset,proto3" json:"pegged_offset,omitempty"`
	// Amend the pegged order reference for the order
	// - See [`PeggedReference`](#vega.PeggedReference)
	PeggedReference      vega.PeggedReference `protobuf:"varint,8,opt,name=pegged_reference,json=peggedReference,proto3,enum=vega.PeggedReference" json:"pegged_reference,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *OrderAmendment) Reset()         { *m = OrderAmendment{} }
func (m *OrderAmendment) String() string { return proto.CompactTextString(m) }
func (*OrderAmendment) ProtoMessage()    {}
func (*OrderAmendment) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c3a53c66a940c51, []int{2}
}

func (m *OrderAmendment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderAmendment.Unmarshal(m, b)
}
func (m *OrderAmendment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderAmendment.Marshal(b, m, deterministic)
}
func (m *OrderAmendment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderAmendment.Merge(m, src)
}
func (m *OrderAmendment) XXX_Size() int {
	return xxx_messageInfo_OrderAmendment.Size(m)
}
func (m *OrderAmendment) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderAmendment.DiscardUnknown(m)
}

var xxx_messageInfo_OrderAmendment proto.InternalMessageInfo

func (m *OrderAmendment) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

func (m *OrderAmendment) GetMarketId() string {
	if m != nil {
		return m.MarketId
	}
	return ""
}

func (m *OrderAmendment) GetPrice() *vega.Price {
	if m != nil {
		return m.Price
	}
	return nil
}

func (m *OrderAmendment) GetSizeDelta() int64 {
	if m != nil {
		return m.SizeDelta
	}
	return 0
}

func (m *OrderAmendment) GetExpiresAt() *vega.Timestamp {
	if m != nil {
		return m.ExpiresAt
	}
	return nil
}

func (m *OrderAmendment) GetTimeInForce() vega.Order_TimeInForce {
	if m != nil {
		return m.TimeInForce
	}
	return vega.Order_TIME_IN_FORCE_UNSPECIFIED
}

func (m *OrderAmendment) GetPeggedOffset() *wrappers.Int64Value {
	if m != nil {
		return m.PeggedOffset
	}
	return nil
}

func (m *OrderAmendment) GetPeggedReference() vega.PeggedReference {
	if m != nil {
		return m.PeggedReference
	}
	return vega.PeggedReference_PEGGED_REFERENCE_UNSPECIFIED
}

// A liquidity provision submitted for a given market
type LiquidityProvisionSubmission struct {
	// Market identifier for the order, required field
	MarketId string `protobuf:"bytes,1,opt,name=market_id,json=marketId,proto3" json:"market_id,omitempty"`
	// Specified as a unitless number that represents the amount of settlement asset of the market
	CommitmentAmount string `protobuf:"bytes,2,opt,name=commitment_amount,json=commitmentAmount,proto3" json:"commitment_amount,omitempty"`
	// Nominated liquidity fee factor, which is an input to the calculation of taker fees on the market, as per seeting fees and rewarding liquidity providers
	Fee string `protobuf:"bytes,3,opt,name=fee,proto3" json:"fee,omitempty"`
	// A set of liquidity sell orders to meet the liquidity provision obligation
	Sells []*vega.LiquidityOrder `protobuf:"bytes,4,rep,name=sells,proto3" json:"sells,omitempty"`
	// A set of liquidity buy orders to meet the liquidity provision obligation
	Buys []*vega.LiquidityOrder `protobuf:"bytes,5,rep,name=buys,proto3" json:"buys,omitempty"`
	// A reference to be added to every order created out of this liquidityProvisionSubmission
	Reference            string   `protobuf:"bytes,6,opt,name=reference,proto3" json:"reference,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LiquidityProvisionSubmission) Reset()         { *m = LiquidityProvisionSubmission{} }
func (m *LiquidityProvisionSubmission) String() string { return proto.CompactTextString(m) }
func (*LiquidityProvisionSubmission) ProtoMessage()    {}
func (*LiquidityProvisionSubmission) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c3a53c66a940c51, []int{3}
}

func (m *LiquidityProvisionSubmission) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LiquidityProvisionSubmission.Unmarshal(m, b)
}
func (m *LiquidityProvisionSubmission) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LiquidityProvisionSubmission.Marshal(b, m, deterministic)
}
func (m *LiquidityProvisionSubmission) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LiquidityProvisionSubmission.Merge(m, src)
}
func (m *LiquidityProvisionSubmission) XXX_Size() int {
	return xxx_messageInfo_LiquidityProvisionSubmission.Size(m)
}
func (m *LiquidityProvisionSubmission) XXX_DiscardUnknown() {
	xxx_messageInfo_LiquidityProvisionSubmission.DiscardUnknown(m)
}

var xxx_messageInfo_LiquidityProvisionSubmission proto.InternalMessageInfo

func (m *LiquidityProvisionSubmission) GetMarketId() string {
	if m != nil {
		return m.MarketId
	}
	return ""
}

func (m *LiquidityProvisionSubmission) GetCommitmentAmount() string {
	if m != nil {
		return m.CommitmentAmount
	}
	return ""
}

func (m *LiquidityProvisionSubmission) GetFee() string {
	if m != nil {
		return m.Fee
	}
	return ""
}

func (m *LiquidityProvisionSubmission) GetSells() []*vega.LiquidityOrder {
	if m != nil {
		return m.Sells
	}
	return nil
}

func (m *LiquidityProvisionSubmission) GetBuys() []*vega.LiquidityOrder {
	if m != nil {
		return m.Buys
	}
	return nil
}

func (m *LiquidityProvisionSubmission) GetReference() string {
	if m != nil {
		return m.Reference
	}
	return ""
}

// Represents the submission request to withdraw funds for a party on Vega
type WithdrawSubmission struct {
	// The amount to be withdrawn
	Amount string `protobuf:"bytes,1,opt,name=amount,proto3" json:"amount,omitempty"`
	// The asset we want to withdraw
	Asset string `protobuf:"bytes,2,opt,name=asset,proto3" json:"asset,omitempty"`
	// Foreign chain specifics
	Ext                  *vega.WithdrawExt `protobuf:"bytes,3,opt,name=ext,proto3" json:"ext,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *WithdrawSubmission) Reset()         { *m = WithdrawSubmission{} }
func (m *WithdrawSubmission) String() string { return proto.CompactTextString(m) }
func (*WithdrawSubmission) ProtoMessage()    {}
func (*WithdrawSubmission) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c3a53c66a940c51, []int{4}
}

func (m *WithdrawSubmission) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WithdrawSubmission.Unmarshal(m, b)
}
func (m *WithdrawSubmission) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WithdrawSubmission.Marshal(b, m, deterministic)
}
func (m *WithdrawSubmission) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WithdrawSubmission.Merge(m, src)
}
func (m *WithdrawSubmission) XXX_Size() int {
	return xxx_messageInfo_WithdrawSubmission.Size(m)
}
func (m *WithdrawSubmission) XXX_DiscardUnknown() {
	xxx_messageInfo_WithdrawSubmission.DiscardUnknown(m)
}

var xxx_messageInfo_WithdrawSubmission proto.InternalMessageInfo

func (m *WithdrawSubmission) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

func (m *WithdrawSubmission) GetAsset() string {
	if m != nil {
		return m.Asset
	}
	return ""
}

func (m *WithdrawSubmission) GetExt() *vega.WithdrawExt {
	if m != nil {
		return m.Ext
	}
	return nil
}

// A command to submit a new proposal for the
// vega network governance
type ProposalSubmission struct {
	// Proposal reference
	Reference string `protobuf:"bytes,1,opt,name=reference,proto3" json:"reference,omitempty"`
	// Proposal configuration and the actual change that is meant to be executed when proposal is enacted
	Terms                *vega.ProposalTerms `protobuf:"bytes,2,opt,name=terms,proto3" json:"terms,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ProposalSubmission) Reset()         { *m = ProposalSubmission{} }
func (m *ProposalSubmission) String() string { return proto.CompactTextString(m) }
func (*ProposalSubmission) ProtoMessage()    {}
func (*ProposalSubmission) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c3a53c66a940c51, []int{5}
}

func (m *ProposalSubmission) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProposalSubmission.Unmarshal(m, b)
}
func (m *ProposalSubmission) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProposalSubmission.Marshal(b, m, deterministic)
}
func (m *ProposalSubmission) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProposalSubmission.Merge(m, src)
}
func (m *ProposalSubmission) XXX_Size() int {
	return xxx_messageInfo_ProposalSubmission.Size(m)
}
func (m *ProposalSubmission) XXX_DiscardUnknown() {
	xxx_messageInfo_ProposalSubmission.DiscardUnknown(m)
}

var xxx_messageInfo_ProposalSubmission proto.InternalMessageInfo

func (m *ProposalSubmission) GetReference() string {
	if m != nil {
		return m.Reference
	}
	return ""
}

func (m *ProposalSubmission) GetTerms() *vega.ProposalTerms {
	if m != nil {
		return m.Terms
	}
	return nil
}

// A command to submit a new vote for a governance
// proposal.
type VoteSubmission struct {
	// The ID of the proposal to vote for.
	ProposalId string `protobuf:"bytes,1,opt,name=proposal_id,json=proposalId,proto3" json:"proposal_id,omitempty"`
	// The actual value of the vote
	Value                vega.Vote_Value `protobuf:"varint,2,opt,name=value,proto3,enum=vega.Vote_Value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *VoteSubmission) Reset()         { *m = VoteSubmission{} }
func (m *VoteSubmission) String() string { return proto.CompactTextString(m) }
func (*VoteSubmission) ProtoMessage()    {}
func (*VoteSubmission) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c3a53c66a940c51, []int{6}
}

func (m *VoteSubmission) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VoteSubmission.Unmarshal(m, b)
}
func (m *VoteSubmission) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VoteSubmission.Marshal(b, m, deterministic)
}
func (m *VoteSubmission) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VoteSubmission.Merge(m, src)
}
func (m *VoteSubmission) XXX_Size() int {
	return xxx_messageInfo_VoteSubmission.Size(m)
}
func (m *VoteSubmission) XXX_DiscardUnknown() {
	xxx_messageInfo_VoteSubmission.DiscardUnknown(m)
}

var xxx_messageInfo_VoteSubmission proto.InternalMessageInfo

func (m *VoteSubmission) GetProposalId() string {
	if m != nil {
		return m.ProposalId
	}
	return ""
}

func (m *VoteSubmission) GetValue() vega.Vote_Value {
	if m != nil {
		return m.Value
	}
	return vega.Vote_VALUE_UNSPECIFIED
}

// A command to submit an instruction to delegate some stake to a node
type DelegateSubmission struct {
	// The ID for the node to delegate to
	NodeId string `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	// The amount of stake to delegate
	Amount               string   `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DelegateSubmission) Reset()         { *m = DelegateSubmission{} }
func (m *DelegateSubmission) String() string { return proto.CompactTextString(m) }
func (*DelegateSubmission) ProtoMessage()    {}
func (*DelegateSubmission) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c3a53c66a940c51, []int{7}
}

func (m *DelegateSubmission) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DelegateSubmission.Unmarshal(m, b)
}
func (m *DelegateSubmission) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DelegateSubmission.Marshal(b, m, deterministic)
}
func (m *DelegateSubmission) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelegateSubmission.Merge(m, src)
}
func (m *DelegateSubmission) XXX_Size() int {
	return xxx_messageInfo_DelegateSubmission.Size(m)
}
func (m *DelegateSubmission) XXX_DiscardUnknown() {
	xxx_messageInfo_DelegateSubmission.DiscardUnknown(m)
}

var xxx_messageInfo_DelegateSubmission proto.InternalMessageInfo

func (m *DelegateSubmission) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *DelegateSubmission) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

type UndelegateSubmission struct {
	NodeId string `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	// optional, if not specified = ALL
	Amount               string                      `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Method               UndelegateSubmission_Method `protobuf:"varint,3,opt,name=method,proto3,enum=vega.commands.v1.UndelegateSubmission_Method" json:"method,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *UndelegateSubmission) Reset()         { *m = UndelegateSubmission{} }
func (m *UndelegateSubmission) String() string { return proto.CompactTextString(m) }
func (*UndelegateSubmission) ProtoMessage()    {}
func (*UndelegateSubmission) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c3a53c66a940c51, []int{8}
}

func (m *UndelegateSubmission) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UndelegateSubmission.Unmarshal(m, b)
}
func (m *UndelegateSubmission) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UndelegateSubmission.Marshal(b, m, deterministic)
}
func (m *UndelegateSubmission) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UndelegateSubmission.Merge(m, src)
}
func (m *UndelegateSubmission) XXX_Size() int {
	return xxx_messageInfo_UndelegateSubmission.Size(m)
}
func (m *UndelegateSubmission) XXX_DiscardUnknown() {
	xxx_messageInfo_UndelegateSubmission.DiscardUnknown(m)
}

var xxx_messageInfo_UndelegateSubmission proto.InternalMessageInfo

func (m *UndelegateSubmission) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *UndelegateSubmission) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

func (m *UndelegateSubmission) GetMethod() UndelegateSubmission_Method {
	if m != nil {
		return m.Method
	}
	return UndelegateSubmission_METHOD_UNSPECIFIED
}

// A command that loads the state from a given checkpoint
type RestoreSnapshot struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RestoreSnapshot) Reset()         { *m = RestoreSnapshot{} }
func (m *RestoreSnapshot) String() string { return proto.CompactTextString(m) }
func (*RestoreSnapshot) ProtoMessage()    {}
func (*RestoreSnapshot) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c3a53c66a940c51, []int{9}
}

func (m *RestoreSnapshot) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RestoreSnapshot.Unmarshal(m, b)
}
func (m *RestoreSnapshot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RestoreSnapshot.Marshal(b, m, deterministic)
}
func (m *RestoreSnapshot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RestoreSnapshot.Merge(m, src)
}
func (m *RestoreSnapshot) XXX_Size() int {
	return xxx_messageInfo_RestoreSnapshot.Size(m)
}
func (m *RestoreSnapshot) XXX_DiscardUnknown() {
	xxx_messageInfo_RestoreSnapshot.DiscardUnknown(m)
}

var xxx_messageInfo_RestoreSnapshot proto.InternalMessageInfo

func (m *RestoreSnapshot) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterEnum("vega.commands.v1.UndelegateSubmission_Method", UndelegateSubmission_Method_name, UndelegateSubmission_Method_value)
	proto.RegisterType((*OrderSubmission)(nil), "vega.commands.v1.OrderSubmission")
	proto.RegisterType((*OrderCancellation)(nil), "vega.commands.v1.OrderCancellation")
	proto.RegisterType((*OrderAmendment)(nil), "vega.commands.v1.OrderAmendment")
	proto.RegisterType((*LiquidityProvisionSubmission)(nil), "vega.commands.v1.LiquidityProvisionSubmission")
	proto.RegisterType((*WithdrawSubmission)(nil), "vega.commands.v1.WithdrawSubmission")
	proto.RegisterType((*ProposalSubmission)(nil), "vega.commands.v1.ProposalSubmission")
	proto.RegisterType((*VoteSubmission)(nil), "vega.commands.v1.VoteSubmission")
	proto.RegisterType((*DelegateSubmission)(nil), "vega.commands.v1.DelegateSubmission")
	proto.RegisterType((*UndelegateSubmission)(nil), "vega.commands.v1.UndelegateSubmission")
	proto.RegisterType((*RestoreSnapshot)(nil), "vega.commands.v1.RestoreSnapshot")
}

func init() { proto.RegisterFile("vega/commands/v1/commands.proto", fileDescriptor_9c3a53c66a940c51) }

var fileDescriptor_9c3a53c66a940c51 = []byte{
	// 953 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0x4d, 0x6f, 0xdb, 0x46,
	0x13, 0x0e, 0xad, 0x0f, 0x5b, 0xa3, 0x44, 0xa2, 0x37, 0x8e, 0x5f, 0xbe, 0xce, 0x87, 0x15, 0xa6,
	0x1f, 0x42, 0x0b, 0x53, 0xb5, 0x1a, 0xe4, 0xd2, 0x4b, 0x1c, 0x5b, 0x6e, 0x84, 0xd6, 0x96, 0x41,
	0x3b, 0x49, 0xd1, 0x0b, 0xb1, 0xd6, 0x8e, 0xe8, 0x85, 0x49, 0x2e, 0x4b, 0xae, 0x64, 0xbb, 0xd7,
	0xfe, 0xcf, 0x16, 0xe8, 0x8f, 0x28, 0xd0, 0x5b, 0xb1, 0xbb, 0x94, 0x2c, 0xc9, 0x45, 0xd1, 0x00,
	0xbd, 0x71, 0x9e, 0x99, 0x79, 0x66, 0x67, 0x9e, 0x9d, 0x25, 0x6c, 0x4f, 0x30, 0xa4, 0x9d, 0xa1,
	0x88, 0x63, 0x9a, 0xb0, 0xbc, 0x33, 0xd9, 0x9d, 0x7d, 0x7b, 0x69, 0x26, 0xa4, 0x20, 0xb6, 0x0a,
	0xf0, 0x66, 0xe0, 0x64, 0x77, 0xeb, 0x91, 0x4e, 0x09, 0xc5, 0x04, 0xb3, 0x84, 0x26, 0x43, 0x34,
	0x81, 0x5b, 0x4d, 0x0d, 0xeb, 0x68, 0x03, 0x3c, 0x0b, 0x85, 0x08, 0x23, 0xec, 0x68, 0xeb, 0x7c,
	0x3c, 0xea, 0x5c, 0x65, 0x34, 0x4d, 0x31, 0x2b, 0x98, 0xb7, 0x5e, 0x85, 0x5c, 0x5e, 0x8c, 0xcf,
	0x15, 0x77, 0x27, 0xbe, 0xe2, 0xf2, 0x52, 0x5c, 0x75, 0x42, 0xb1, 0xa3, 0x9d, 0x3b, 0x13, 0x1a,
	0x71, 0x46, 0xa5, 0xc8, 0xf2, 0xce, 0xec, 0xd3, 0xe4, 0xb9, 0xbf, 0xae, 0x40, 0x73, 0x90, 0x31,
	0xcc, 0x4e, 0xc7, 0xe7, 0x31, 0xcf, 0x73, 0x2e, 0x12, 0xf2, 0x02, 0x6a, 0x31, 0xcd, 0x2e, 0x51,
	0x06, 0x9c, 0x39, 0x56, 0xcb, 0x6a, 0xd7, 0xde, 0x54, 0x7f, 0xff, 0x6d, 0x7b, 0xe5, 0x07, 0xcb,
	0x5f, 0x33, 0x8e, 0x3e, 0x23, 0x1b, 0x50, 0x49, 0x33, 0x3e, 0x44, 0x67, 0x45, 0x05, 0xf8, 0xc6,
	0x20, 0x5b, 0x50, 0xce, 0xf9, 0xcf, 0xe8, 0x94, 0x5a, 0x56, 0xbb, 0x6c, 0xb2, 0xec, 0x7b, 0xbe,
	0xc6, 0xc8, 0x33, 0xe5, 0x63, 0xe8, 0x94, 0x5b, 0x56, 0xbb, 0xd1, 0x05, 0x4f, 0x77, 0x77, 0xca,
	0x19, 0xfa, 0x1a, 0x27, 0xdf, 0xc0, 0x03, 0xc9, 0x63, 0x0c, 0x78, 0x12, 0x8c, 0x44, 0x36, 0x44,
	0xa7, 0xa2, 0x03, 0xff, 0x67, 0x02, 0xf5, 0x21, 0xbd, 0x33, 0x1e, 0x63, 0x3f, 0x39, 0x54, 0x6e,
	0xbf, 0x2e, 0x6f, 0x0d, 0xf2, 0x14, 0x00, 0xaf, 0x53, 0x9e, 0x61, 0x1e, 0x50, 0xe9, 0x54, 0x5b,
	0x56, 0xbb, 0xe4, 0xd7, 0x0a, 0x64, 0x4f, 0x92, 0x4f, 0xa0, 0x2c, 0x6f, 0x52, 0x74, 0x56, 0x35,
	0xa5, 0xbd, 0x40, 0x79, 0x93, 0xa2, 0xaf, 0xbd, 0xe4, 0x09, 0xd4, 0x32, 0x1c, 0x61, 0x86, 0xc9,
	0x10, 0x9d, 0x35, 0xdd, 0xd7, 0x2d, 0x40, 0x5e, 0xc2, 0xfd, 0x14, 0xc3, 0x10, 0x59, 0x20, 0x54,
	0xa2, 0x53, 0x6b, 0x59, 0xed, 0x7a, 0x77, 0xdd, 0x70, 0x9d, 0x68, 0x8f, 0x66, 0xf4, 0xeb, 0xe9,
	0xad, 0xe1, 0x7e, 0x07, 0xeb, 0xfa, 0x63, 0x5f, 0xa9, 0x1b, 0x45, 0x54, 0xaa, 0x09, 0xff, 0x1f,
	0xd6, 0x34, 0xc7, 0x6c, 0xc0, 0xfe, 0xaa, 0xb6, 0xfb, 0x8c, 0x3c, 0x9e, 0x1f, 0xbe, 0x99, 0xed,
	0x6c, 0xe8, 0xee, 0x2f, 0x25, 0x68, 0x68, 0xb6, 0xbd, 0x18, 0x13, 0x16, 0x63, 0x22, 0xc9, 0xf3,
	0x65, 0xaa, 0x99, 0x56, 0xff, 0x8a, 0x92, 0x3c, 0x9f, 0xea, 0x58, 0xd2, 0xed, 0xd4, 0x8b, 0x76,
	0x14, 0x34, 0x15, 0xf5, 0x29, 0x80, 0x12, 0x30, 0x60, 0x18, 0x49, 0xaa, 0xe5, 0x2b, 0xf9, 0x35,
	0x85, 0x1c, 0x28, 0x80, 0x78, 0x0b, 0xa3, 0xaf, 0x68, 0x9a, 0xa6, 0xa1, 0x51, 0x72, 0xe5, 0x92,
	0xc6, 0xe9, 0xbc, 0x16, 0x77, 0x74, 0xae, 0x7e, 0x84, 0xce, 0xaf, 0xe1, 0xc1, 0x54, 0x84, 0xd1,
	0x28, 0x47, 0xa9, 0x15, 0xad, 0x77, 0x1f, 0x7b, 0x66, 0x3f, 0xbc, 0xe9, 0x7e, 0x78, 0xfd, 0x44,
	0xbe, 0x7a, 0xf9, 0x9e, 0x46, 0x63, 0xf4, 0x0b, 0xd9, 0x06, 0x3a, 0x81, 0xbc, 0x06, 0xbb, 0x60,
	0x58, 0xd4, 0xba, 0xd1, 0x7d, 0x34, 0x2f, 0xa5, 0x3f, 0x75, 0xfa, 0xcd, 0x74, 0x11, 0x70, 0xff,
	0xb0, 0xe0, 0xc9, 0xf7, 0xfc, 0xa7, 0x31, 0x67, 0x5c, 0xde, 0x9c, 0x64, 0x62, 0xc2, 0xd5, 0xda,
	0x7c, 0xec, 0x02, 0x7d, 0x09, 0xeb, 0xea, 0x21, 0xe0, 0x52, 0xc9, 0x18, 0xd0, 0x58, 0x8c, 0x13,
	0x59, 0xa8, 0x63, 0xdf, 0x3a, 0xf6, 0x34, 0x4e, 0x6c, 0x28, 0x8d, 0xd0, 0x68, 0x54, 0xf3, 0xd5,
	0x27, 0xf9, 0x02, 0x2a, 0x39, 0x46, 0x51, 0xee, 0x94, 0x5b, 0xa5, 0x76, 0xbd, 0xbb, 0x61, 0xce,
	0x3e, 0x3b, 0x96, 0xb9, 0x89, 0x26, 0x84, 0xb4, 0xa1, 0x7c, 0x3e, 0xbe, 0xc9, 0x9d, 0xca, 0x3f,
	0x84, 0xea, 0x88, 0xc5, 0x0d, 0xa8, 0x2e, 0x6d, 0x80, 0x1b, 0x02, 0xf9, 0xc0, 0xe5, 0x05, 0xcb,
	0xe8, 0xd5, 0x5c, 0xb7, 0x9b, 0x50, 0x2d, 0x4e, 0x6f, 0xae, 0x72, 0x61, 0xa9, 0x17, 0x82, 0xe6,
	0x4a, 0xa2, 0xe2, 0x85, 0xd0, 0x06, 0x79, 0x01, 0x25, 0xbc, 0x96, 0xc5, 0x6d, 0x2b, 0x96, 0x67,
	0x4a, 0xda, 0xbb, 0x96, 0xbe, 0xf2, 0xba, 0x08, 0xe4, 0x24, 0x13, 0xa9, 0xc8, 0x69, 0x34, 0x57,
	0x68, 0xe1, 0x70, 0xd6, 0xf2, 0x7a, 0xee, 0x42, 0x45, 0x62, 0x16, 0xe7, 0xba, 0x5c, 0xbd, 0xfb,
	0x70, 0x7a, 0x91, 0x0d, 0xcd, 0x99, 0x72, 0x19, 0x15, 0x5a, 0x96, 0x6f, 0x22, 0x5d, 0x0a, 0x8d,
	0xf7, 0x42, 0xe2, 0x5c, 0x89, 0xcf, 0xa1, 0x9e, 0x16, 0x19, 0x77, 0xb5, 0x83, 0xa9, 0xab, 0xcf,
	0xc8, 0x67, 0x50, 0x99, 0xa8, 0xcb, 0xa5, 0xab, 0xcd, 0x5e, 0x14, 0xc5, 0xe6, 0x99, 0x4b, 0x67,
	0xdc, 0xee, 0x11, 0x90, 0x03, 0x8c, 0x30, 0xa4, 0x0b, 0x65, 0xb6, 0x61, 0x35, 0x11, 0x0c, 0xef,
	0x96, 0xa8, 0x2a, 0xb8, 0xcf, 0xe6, 0x66, 0xba, 0x32, 0x3f, 0x53, 0xf7, 0x4f, 0x0b, 0x36, 0xde,
	0x25, 0xec, 0xbf, 0x63, 0x24, 0x3d, 0xa8, 0xc6, 0x28, 0x2f, 0x04, 0xd3, 0x92, 0x34, 0xba, 0x3b,
	0xde, 0xf2, 0x3f, 0xca, 0xfb, 0xbb, 0x82, 0xde, 0x91, 0x4e, 0xf2, 0x8b, 0x64, 0x97, 0x42, 0xd5,
	0x20, 0x64, 0x13, 0xc8, 0x51, 0xef, 0xec, 0xed, 0xe0, 0x20, 0x78, 0x77, 0x7c, 0x7a, 0xd2, 0xdb,
	0xef, 0x1f, 0xf6, 0x7b, 0x07, 0xf6, 0x3d, 0xd2, 0x00, 0x28, 0xf0, 0xe3, 0xc1, 0x07, 0xdb, 0x22,
	0x5b, 0xb0, 0x59, 0xd8, 0x7b, 0x67, 0x41, 0xef, 0xf8, 0x20, 0x18, 0x1c, 0x06, 0xbd, 0x93, 0xc1,
	0xfe, 0x5b, 0x7b, 0x85, 0x3c, 0x84, 0x66, 0xe1, 0xeb, 0x1f, 0x07, 0x7b, 0xc7, 0xdf, 0xf6, 0x7c,
	0xbb, 0xe4, 0x7e, 0x0a, 0x4d, 0x1f, 0x73, 0x29, 0x32, 0x3c, 0x4d, 0x68, 0x9a, 0x5f, 0x08, 0x49,
	0x08, 0x94, 0x19, 0x95, 0x54, 0xb7, 0x7c, 0xdf, 0xd7, 0xdf, 0x6f, 0xba, 0x3f, 0x7e, 0x35, 0x14,
	0x0c, 0x75, 0x1b, 0xfa, 0x35, 0x18, 0x8a, 0xc8, 0xe3, 0xa2, 0x93, 0xb3, 0xcb, 0x9d, 0x50, 0x44,
	0x34, 0x09, 0x3b, 0xcb, 0xbf, 0xe9, 0xf3, 0xaa, 0x0e, 0xfc, 0xfa, 0xaf, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x28, 0xb5, 0x1a, 0xd0, 0xc1, 0x07, 0x00, 0x00,
}
