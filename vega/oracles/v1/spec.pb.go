// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: vega/oracles/v1/spec.proto

package v1

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Status describe the status of the oracle spec
type OracleSpec_Status int32

const (
	// The default value.
	OracleSpec_STATUS_UNSPECIFIED OracleSpec_Status = 0
	// STATUS_ACTIVE describes an active oracle spec.
	OracleSpec_STATUS_ACTIVE OracleSpec_Status = 1
	// STATUS_DEACTIVATED describes an oracle spec that is not listening to data
	// anymore.
	OracleSpec_STATUS_DEACTIVATED OracleSpec_Status = 2
)

// Enum value maps for OracleSpec_Status.
var (
	OracleSpec_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "STATUS_ACTIVE",
		2: "STATUS_DEACTIVATED",
	}
	OracleSpec_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"STATUS_ACTIVE":      1,
		"STATUS_DEACTIVATED": 2,
	}
)

func (x OracleSpec_Status) Enum() *OracleSpec_Status {
	p := new(OracleSpec_Status)
	*p = x
	return p
}

func (x OracleSpec_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OracleSpec_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_vega_oracles_v1_spec_proto_enumTypes[0].Descriptor()
}

func (OracleSpec_Status) Type() protoreflect.EnumType {
	return &file_vega_oracles_v1_spec_proto_enumTypes[0]
}

func (x OracleSpec_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OracleSpec_Status.Descriptor instead.
func (OracleSpec_Status) EnumDescriptor() ([]byte, []int) {
	return file_vega_oracles_v1_spec_proto_rawDescGZIP(), []int{1, 0}
}

// Type describes the type of properties that are supported by the oracle
// engine.
type PropertyKey_Type int32

const (
	// The default value.
	PropertyKey_TYPE_UNSPECIFIED PropertyKey_Type = 0
	// Any type.
	PropertyKey_TYPE_EMPTY PropertyKey_Type = 1
	// Integer type.
	PropertyKey_TYPE_INTEGER PropertyKey_Type = 2
	// String type.
	PropertyKey_TYPE_STRING PropertyKey_Type = 3
	// Boolean type.
	PropertyKey_TYPE_BOOLEAN PropertyKey_Type = 4
	// Any floating point decimal type.
	PropertyKey_TYPE_DECIMAL PropertyKey_Type = 5
	// Timestamp date type.
	PropertyKey_TYPE_TIMESTAMP PropertyKey_Type = 6
)

// Enum value maps for PropertyKey_Type.
var (
	PropertyKey_Type_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "TYPE_EMPTY",
		2: "TYPE_INTEGER",
		3: "TYPE_STRING",
		4: "TYPE_BOOLEAN",
		5: "TYPE_DECIMAL",
		6: "TYPE_TIMESTAMP",
	}
	PropertyKey_Type_value = map[string]int32{
		"TYPE_UNSPECIFIED": 0,
		"TYPE_EMPTY":       1,
		"TYPE_INTEGER":     2,
		"TYPE_STRING":      3,
		"TYPE_BOOLEAN":     4,
		"TYPE_DECIMAL":     5,
		"TYPE_TIMESTAMP":   6,
	}
)

func (x PropertyKey_Type) Enum() *PropertyKey_Type {
	p := new(PropertyKey_Type)
	*p = x
	return p
}

func (x PropertyKey_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PropertyKey_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_vega_oracles_v1_spec_proto_enumTypes[1].Descriptor()
}

func (PropertyKey_Type) Type() protoreflect.EnumType {
	return &file_vega_oracles_v1_spec_proto_enumTypes[1]
}

func (x PropertyKey_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PropertyKey_Type.Descriptor instead.
func (PropertyKey_Type) EnumDescriptor() ([]byte, []int) {
	return file_vega_oracles_v1_spec_proto_rawDescGZIP(), []int{3, 0}
}

// Comparator describes the type of comparison.
type Condition_Operator int32

const (
	// The default value
	Condition_OPERATOR_UNSPECIFIED Condition_Operator = 0
	// Verify if the property values are strictly equal or not.
	Condition_OPERATOR_EQUALS Condition_Operator = 1
	// Verify if the oracle data value is greater than the Condition value.
	Condition_OPERATOR_GREATER_THAN Condition_Operator = 2
	// Verify if the oracle data value is greater than or equal to the Condition
	// value.
	Condition_OPERATOR_GREATER_THAN_OR_EQUAL Condition_Operator = 3
	// Verify if the oracle data value is less than the Condition value.
	Condition_OPERATOR_LESS_THAN Condition_Operator = 4
	// Verify if the oracle data value is less or equal to than the Condition
	// value.
	Condition_OPERATOR_LESS_THAN_OR_EQUAL Condition_Operator = 5
)

// Enum value maps for Condition_Operator.
var (
	Condition_Operator_name = map[int32]string{
		0: "OPERATOR_UNSPECIFIED",
		1: "OPERATOR_EQUALS",
		2: "OPERATOR_GREATER_THAN",
		3: "OPERATOR_GREATER_THAN_OR_EQUAL",
		4: "OPERATOR_LESS_THAN",
		5: "OPERATOR_LESS_THAN_OR_EQUAL",
	}
	Condition_Operator_value = map[string]int32{
		"OPERATOR_UNSPECIFIED":           0,
		"OPERATOR_EQUALS":                1,
		"OPERATOR_GREATER_THAN":          2,
		"OPERATOR_GREATER_THAN_OR_EQUAL": 3,
		"OPERATOR_LESS_THAN":             4,
		"OPERATOR_LESS_THAN_OR_EQUAL":    5,
	}
)

func (x Condition_Operator) Enum() *Condition_Operator {
	p := new(Condition_Operator)
	*p = x
	return p
}

func (x Condition_Operator) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Condition_Operator) Descriptor() protoreflect.EnumDescriptor {
	return file_vega_oracles_v1_spec_proto_enumTypes[2].Descriptor()
}

func (Condition_Operator) Type() protoreflect.EnumType {
	return &file_vega_oracles_v1_spec_proto_enumTypes[2]
}

func (x Condition_Operator) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Condition_Operator.Descriptor instead.
func (Condition_Operator) EnumDescriptor() ([]byte, []int) {
	return file_vega_oracles_v1_spec_proto_rawDescGZIP(), []int{4, 0}
}

// An oracle spec describe the oracle data that a product (or a risk model)
// wants to get from the oracle engine.
type OracleSpecConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// pubKeys is the list of authorized public keys that signed the data for this
	// oracle. All the public keys in the oracle data should be contained in these
	// public keys.
	PubKeys []string `protobuf:"bytes,1,rep,name=pub_keys,json=pubKeys,proto3" json:"pub_keys,omitempty"`
	// filters describes which oracle data are considered of interest or not for
	// the product (or the risk model).
	Filters []*Filter `protobuf:"bytes,2,rep,name=filters,proto3" json:"filters,omitempty"`
}

func (x *OracleSpecConfiguration) Reset() {
	*x = OracleSpecConfiguration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vega_oracles_v1_spec_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OracleSpecConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OracleSpecConfiguration) ProtoMessage() {}

func (x *OracleSpecConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_vega_oracles_v1_spec_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OracleSpecConfiguration.ProtoReflect.Descriptor instead.
func (*OracleSpecConfiguration) Descriptor() ([]byte, []int) {
	return file_vega_oracles_v1_spec_proto_rawDescGZIP(), []int{0}
}

func (x *OracleSpecConfiguration) GetPubKeys() []string {
	if x != nil {
		return x.PubKeys
	}
	return nil
}

func (x *OracleSpecConfiguration) GetFilters() []*Filter {
	if x != nil {
		return x.Filters
	}
	return nil
}

// An oracle spec describe the oracle data that a product (or a risk model)
// wants to get from the oracle engine.
// This message contains additional information used by the API.
type OracleSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id is a hash generated from the OracleSpec data.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Creation Date time
	CreatedAt int64 `protobuf:"varint,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Last Updated timestamp
	UpdatedAt int64 `protobuf:"varint,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	// pubKeys is the list of authorized public keys that signed the data for this
	// oracle. All the public keys in the oracle data should be contained in these
	// public keys.
	PubKeys []string `protobuf:"bytes,4,rep,name=pub_keys,json=pubKeys,proto3" json:"pub_keys,omitempty"`
	// filters describes which oracle data are considered of interest or not for
	// the product (or the risk model).
	Filters []*Filter `protobuf:"bytes,5,rep,name=filters,proto3" json:"filters,omitempty"`
	// status describes the status of the oracle spec
	Status OracleSpec_Status `protobuf:"varint,6,opt,name=status,proto3,enum=oracles.v1.OracleSpec_Status" json:"status,omitempty"`
}

func (x *OracleSpec) Reset() {
	*x = OracleSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vega_oracles_v1_spec_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OracleSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OracleSpec) ProtoMessage() {}

func (x *OracleSpec) ProtoReflect() protoreflect.Message {
	mi := &file_vega_oracles_v1_spec_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OracleSpec.ProtoReflect.Descriptor instead.
func (*OracleSpec) Descriptor() ([]byte, []int) {
	return file_vega_oracles_v1_spec_proto_rawDescGZIP(), []int{1}
}

func (x *OracleSpec) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *OracleSpec) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *OracleSpec) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *OracleSpec) GetPubKeys() []string {
	if x != nil {
		return x.PubKeys
	}
	return nil
}

func (x *OracleSpec) GetFilters() []*Filter {
	if x != nil {
		return x.Filters
	}
	return nil
}

func (x *OracleSpec) GetStatus() OracleSpec_Status {
	if x != nil {
		return x.Status
	}
	return OracleSpec_STATUS_UNSPECIFIED
}

// Filter describes the conditions under which an oracle data is considered of
// interest or not.
type Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// key is the oracle data property key targeted by the filter.
	Key *PropertyKey `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// conditions are the conditions that should be matched by the data to be
	// considered of interest.
	Conditions []*Condition `protobuf:"bytes,2,rep,name=conditions,proto3" json:"conditions,omitempty"`
}

func (x *Filter) Reset() {
	*x = Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vega_oracles_v1_spec_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Filter) ProtoMessage() {}

func (x *Filter) ProtoReflect() protoreflect.Message {
	mi := &file_vega_oracles_v1_spec_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Filter.ProtoReflect.Descriptor instead.
func (*Filter) Descriptor() ([]byte, []int) {
	return file_vega_oracles_v1_spec_proto_rawDescGZIP(), []int{2}
}

func (x *Filter) GetKey() *PropertyKey {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *Filter) GetConditions() []*Condition {
	if x != nil {
		return x.Conditions
	}
	return nil
}

// PropertyKey describes the property key contained in an oracle data.
type PropertyKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name is the name of the property.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// type is the type of the property.
	Type PropertyKey_Type `protobuf:"varint,2,opt,name=type,proto3,enum=oracles.v1.PropertyKey_Type" json:"type,omitempty"`
}

func (x *PropertyKey) Reset() {
	*x = PropertyKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vega_oracles_v1_spec_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PropertyKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PropertyKey) ProtoMessage() {}

func (x *PropertyKey) ProtoReflect() protoreflect.Message {
	mi := &file_vega_oracles_v1_spec_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PropertyKey.ProtoReflect.Descriptor instead.
func (*PropertyKey) Descriptor() ([]byte, []int) {
	return file_vega_oracles_v1_spec_proto_rawDescGZIP(), []int{3}
}

func (x *PropertyKey) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PropertyKey) GetType() PropertyKey_Type {
	if x != nil {
		return x.Type
	}
	return PropertyKey_TYPE_UNSPECIFIED
}

// Condition describes the condition that must be validated by the
type Condition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// comparator is the type of comparison to make on the value.
	Operator Condition_Operator `protobuf:"varint,1,opt,name=operator,proto3,enum=oracles.v1.Condition_Operator" json:"operator,omitempty"`
	// value is used by the comparator.
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Condition) Reset() {
	*x = Condition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vega_oracles_v1_spec_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Condition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Condition) ProtoMessage() {}

func (x *Condition) ProtoReflect() protoreflect.Message {
	mi := &file_vega_oracles_v1_spec_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Condition.ProtoReflect.Descriptor instead.
func (*Condition) Descriptor() ([]byte, []int) {
	return file_vega_oracles_v1_spec_proto_rawDescGZIP(), []int{4}
}

func (x *Condition) GetOperator() Condition_Operator {
	if x != nil {
		return x.Operator
	}
	return Condition_OPERATOR_UNSPECIFIED
}

func (x *Condition) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_vega_oracles_v1_spec_proto protoreflect.FileDescriptor

var file_vega_oracles_v1_spec_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x76, 0x65, 0x67, 0x61, 0x2f, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6f, 0x72,
	0x61, 0x63, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x22, 0x62, 0x0a, 0x17, 0x4f, 0x72, 0x61, 0x63,
	0x6c, 0x65, 0x53, 0x70, 0x65, 0x63, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x75, 0x62, 0x5f, 0x6b, 0x65, 0x79, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x70, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x2c,
	0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x22, 0xa7, 0x02, 0x0a,
	0x0a, 0x4f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x53, 0x70, 0x65, 0x63, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x75, 0x62,
	0x5f, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x70, 0x75, 0x62,
	0x4b, 0x65, 0x79, 0x73, 0x12, 0x2c, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x73, 0x12, 0x35, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x4f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x4b, 0x0a, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x01, 0x12, 0x16,
	0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x44, 0x45, 0x41, 0x43, 0x54, 0x49, 0x56,
	0x41, 0x54, 0x45, 0x44, 0x10, 0x02, 0x22, 0x6a, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x12, 0x29, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x65,
	0x72, 0x74, 0x79, 0x4b, 0x65, 0x79, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x35, 0x0a, 0x0a, 0x63,
	0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x22, 0xdd, 0x01, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x4b,
	0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x4b, 0x65, 0x79, 0x2e, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x87, 0x01, 0x0a, 0x04, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x45, 0x4d, 0x50, 0x54, 0x59, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x49, 0x4e, 0x54, 0x45, 0x47, 0x45, 0x52, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x42, 0x4f, 0x4f, 0x4c, 0x45, 0x41, 0x4e, 0x10, 0x04, 0x12, 0x10, 0x0a, 0x0c,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x45, 0x43, 0x49, 0x4d, 0x41, 0x4c, 0x10, 0x05, 0x12, 0x12,
	0x0a, 0x0e, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x53, 0x54, 0x41, 0x4d, 0x50,
	0x10, 0x06, 0x22, 0x91, 0x02, 0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x3a, 0x0a, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x22, 0xb1, 0x01, 0x0a, 0x08, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12,
	0x18, 0x0a, 0x14, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x4f, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x4f, 0x50, 0x45,
	0x52, 0x41, 0x54, 0x4f, 0x52, 0x5f, 0x45, 0x51, 0x55, 0x41, 0x4c, 0x53, 0x10, 0x01, 0x12, 0x19,
	0x0a, 0x15, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x4f, 0x52, 0x5f, 0x47, 0x52, 0x45, 0x41, 0x54,
	0x45, 0x52, 0x5f, 0x54, 0x48, 0x41, 0x4e, 0x10, 0x02, 0x12, 0x22, 0x0a, 0x1e, 0x4f, 0x50, 0x45,
	0x52, 0x41, 0x54, 0x4f, 0x52, 0x5f, 0x47, 0x52, 0x45, 0x41, 0x54, 0x45, 0x52, 0x5f, 0x54, 0x48,
	0x41, 0x4e, 0x5f, 0x4f, 0x52, 0x5f, 0x45, 0x51, 0x55, 0x41, 0x4c, 0x10, 0x03, 0x12, 0x16, 0x0a,
	0x12, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x4f, 0x52, 0x5f, 0x4c, 0x45, 0x53, 0x53, 0x5f, 0x54,
	0x48, 0x41, 0x4e, 0x10, 0x04, 0x12, 0x1f, 0x0a, 0x1b, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x4f,
	0x52, 0x5f, 0x4c, 0x45, 0x53, 0x53, 0x5f, 0x54, 0x48, 0x41, 0x4e, 0x5f, 0x4f, 0x52, 0x5f, 0x45,
	0x51, 0x55, 0x41, 0x4c, 0x10, 0x05, 0x42, 0x31, 0x5a, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x76,
	0x65, 0x67, 0x61, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x69, 0x6f, 0x2f, 0x73,
	0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x76, 0x65, 0x67, 0x61, 0x2f, 0x6f,
	0x72, 0x61, 0x63, 0x6c, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_vega_oracles_v1_spec_proto_rawDescOnce sync.Once
	file_vega_oracles_v1_spec_proto_rawDescData = file_vega_oracles_v1_spec_proto_rawDesc
)

func file_vega_oracles_v1_spec_proto_rawDescGZIP() []byte {
	file_vega_oracles_v1_spec_proto_rawDescOnce.Do(func() {
		file_vega_oracles_v1_spec_proto_rawDescData = protoimpl.X.CompressGZIP(file_vega_oracles_v1_spec_proto_rawDescData)
	})
	return file_vega_oracles_v1_spec_proto_rawDescData
}

var file_vega_oracles_v1_spec_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_vega_oracles_v1_spec_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_vega_oracles_v1_spec_proto_goTypes = []interface{}{
	(OracleSpec_Status)(0),          // 0: oracles.v1.OracleSpec.Status
	(PropertyKey_Type)(0),           // 1: oracles.v1.PropertyKey.Type
	(Condition_Operator)(0),         // 2: oracles.v1.Condition.Operator
	(*OracleSpecConfiguration)(nil), // 3: oracles.v1.OracleSpecConfiguration
	(*OracleSpec)(nil),              // 4: oracles.v1.OracleSpec
	(*Filter)(nil),                  // 5: oracles.v1.Filter
	(*PropertyKey)(nil),             // 6: oracles.v1.PropertyKey
	(*Condition)(nil),               // 7: oracles.v1.Condition
}
var file_vega_oracles_v1_spec_proto_depIdxs = []int32{
	5, // 0: oracles.v1.OracleSpecConfiguration.filters:type_name -> oracles.v1.Filter
	5, // 1: oracles.v1.OracleSpec.filters:type_name -> oracles.v1.Filter
	0, // 2: oracles.v1.OracleSpec.status:type_name -> oracles.v1.OracleSpec.Status
	6, // 3: oracles.v1.Filter.key:type_name -> oracles.v1.PropertyKey
	7, // 4: oracles.v1.Filter.conditions:type_name -> oracles.v1.Condition
	1, // 5: oracles.v1.PropertyKey.type:type_name -> oracles.v1.PropertyKey.Type
	2, // 6: oracles.v1.Condition.operator:type_name -> oracles.v1.Condition.Operator
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_vega_oracles_v1_spec_proto_init() }
func file_vega_oracles_v1_spec_proto_init() {
	if File_vega_oracles_v1_spec_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_vega_oracles_v1_spec_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OracleSpecConfiguration); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_vega_oracles_v1_spec_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OracleSpec); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_vega_oracles_v1_spec_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Filter); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_vega_oracles_v1_spec_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PropertyKey); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_vega_oracles_v1_spec_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Condition); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_vega_oracles_v1_spec_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_vega_oracles_v1_spec_proto_goTypes,
		DependencyIndexes: file_vega_oracles_v1_spec_proto_depIdxs,
		EnumInfos:         file_vega_oracles_v1_spec_proto_enumTypes,
		MessageInfos:      file_vega_oracles_v1_spec_proto_msgTypes,
	}.Build()
	File_vega_oracles_v1_spec_proto = out.File
	file_vega_oracles_v1_spec_proto_rawDesc = nil
	file_vega_oracles_v1_spec_proto_goTypes = nil
	file_vega_oracles_v1_spec_proto_depIdxs = nil
}
