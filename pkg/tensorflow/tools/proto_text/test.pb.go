// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/tools/proto_text/test.proto

/*
Package proto_text is a generated protocol buffer package.

It is generated from these files:
	tensorflow/tools/proto_text/test.proto

It has these top-level messages:
	TestAllTypes
	NestedTestAllTypes
	ForeignMessage
	TestEmptyMessage
*/
package proto_text

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ForeignEnum int32

const (
	ForeignEnum_FOREIGN_ZERO ForeignEnum = 0
	ForeignEnum_FOREIGN_FOO  ForeignEnum = 4
	ForeignEnum_FOREIGN_BAR  ForeignEnum = 5
	ForeignEnum_FOREIGN_BAZ  ForeignEnum = 6
)

var ForeignEnum_name = map[int32]string{
	0: "FOREIGN_ZERO",
	4: "FOREIGN_FOO",
	5: "FOREIGN_BAR",
	6: "FOREIGN_BAZ",
}
var ForeignEnum_value = map[string]int32{
	"FOREIGN_ZERO": 0,
	"FOREIGN_FOO":  4,
	"FOREIGN_BAR":  5,
	"FOREIGN_BAZ":  6,
}

func (x ForeignEnum) String() string {
	return proto.EnumName(ForeignEnum_name, int32(x))
}
func (ForeignEnum) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type TestAllTypes_NestedEnum int32

const (
	TestAllTypes_ZERO TestAllTypes_NestedEnum = 0
	TestAllTypes_FOO  TestAllTypes_NestedEnum = 1
	TestAllTypes_BAR  TestAllTypes_NestedEnum = 2
	TestAllTypes_BAZ  TestAllTypes_NestedEnum = 3
	TestAllTypes_NEG  TestAllTypes_NestedEnum = -1
)

var TestAllTypes_NestedEnum_name = map[int32]string{
	0:  "ZERO",
	1:  "FOO",
	2:  "BAR",
	3:  "BAZ",
	-1: "NEG",
}
var TestAllTypes_NestedEnum_value = map[string]int32{
	"ZERO": 0,
	"FOO":  1,
	"BAR":  2,
	"BAZ":  3,
	"NEG":  -1,
}

func (x TestAllTypes_NestedEnum) String() string {
	return proto.EnumName(TestAllTypes_NestedEnum_name, int32(x))
}
func (TestAllTypes_NestedEnum) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type TestAllTypes struct {
	// Singular
	OptionalInt32          int32                       `protobuf:"varint,1000,opt,name=optional_int32,json=optionalInt32" json:"optional_int32,omitempty"`
	OptionalInt64          int64                       `protobuf:"varint,2,opt,name=optional_int64,json=optionalInt64" json:"optional_int64,omitempty"`
	OptionalUint32         uint32                      `protobuf:"varint,3,opt,name=optional_uint32,json=optionalUint32" json:"optional_uint32,omitempty"`
	OptionalUint64         uint64                      `protobuf:"varint,999,opt,name=optional_uint64,json=optionalUint64" json:"optional_uint64,omitempty"`
	OptionalSint32         int32                       `protobuf:"zigzag32,5,opt,name=optional_sint32,json=optionalSint32" json:"optional_sint32,omitempty"`
	OptionalSint64         int64                       `protobuf:"zigzag64,6,opt,name=optional_sint64,json=optionalSint64" json:"optional_sint64,omitempty"`
	OptionalFixed32        uint32                      `protobuf:"fixed32,7,opt,name=optional_fixed32,json=optionalFixed32" json:"optional_fixed32,omitempty"`
	OptionalFixed64        uint64                      `protobuf:"fixed64,8,opt,name=optional_fixed64,json=optionalFixed64" json:"optional_fixed64,omitempty"`
	OptionalSfixed32       int32                       `protobuf:"fixed32,9,opt,name=optional_sfixed32,json=optionalSfixed32" json:"optional_sfixed32,omitempty"`
	OptionalSfixed64       int64                       `protobuf:"fixed64,10,opt,name=optional_sfixed64,json=optionalSfixed64" json:"optional_sfixed64,omitempty"`
	OptionalFloat          float32                     `protobuf:"fixed32,11,opt,name=optional_float,json=optionalFloat" json:"optional_float,omitempty"`
	OptionalDouble         float64                     `protobuf:"fixed64,12,opt,name=optional_double,json=optionalDouble" json:"optional_double,omitempty"`
	OptionalBool           bool                        `protobuf:"varint,13,opt,name=optional_bool,json=optionalBool" json:"optional_bool,omitempty"`
	OptionalString         string                      `protobuf:"bytes,14,opt,name=optional_string,json=optionalString" json:"optional_string,omitempty"`
	OptionalBytes          []byte                      `protobuf:"bytes,15,opt,name=optional_bytes,json=optionalBytes,proto3" json:"optional_bytes,omitempty"`
	OptionalNestedMessage  *TestAllTypes_NestedMessage `protobuf:"bytes,18,opt,name=optional_nested_message,json=optionalNestedMessage" json:"optional_nested_message,omitempty"`
	OptionalForeignMessage *ForeignMessage             `protobuf:"bytes,19,opt,name=optional_foreign_message,json=optionalForeignMessage" json:"optional_foreign_message,omitempty"`
	OptionalNestedEnum     TestAllTypes_NestedEnum     `protobuf:"varint,21,opt,name=optional_nested_enum,json=optionalNestedEnum,enum=tensorflow.test.TestAllTypes_NestedEnum" json:"optional_nested_enum,omitempty"`
	OptionalForeignEnum    ForeignEnum                 `protobuf:"varint,22,opt,name=optional_foreign_enum,json=optionalForeignEnum,enum=tensorflow.test.ForeignEnum" json:"optional_foreign_enum,omitempty"`
	OptionalCord           string                      `protobuf:"bytes,25,opt,name=optional_cord,json=optionalCord" json:"optional_cord,omitempty"`
	// Repeated
	RepeatedInt32         []int32                       `protobuf:"varint,31,rep,packed,name=repeated_int32,json=repeatedInt32" json:"repeated_int32,omitempty"`
	RepeatedInt64         []int64                       `protobuf:"varint,32,rep,packed,name=repeated_int64,json=repeatedInt64" json:"repeated_int64,omitempty"`
	RepeatedUint32        []uint32                      `protobuf:"varint,33,rep,packed,name=repeated_uint32,json=repeatedUint32" json:"repeated_uint32,omitempty"`
	RepeatedUint64        []uint64                      `protobuf:"varint,34,rep,packed,name=repeated_uint64,json=repeatedUint64" json:"repeated_uint64,omitempty"`
	RepeatedSint32        []int32                       `protobuf:"zigzag32,35,rep,packed,name=repeated_sint32,json=repeatedSint32" json:"repeated_sint32,omitempty"`
	RepeatedSint64        []int64                       `protobuf:"zigzag64,36,rep,packed,name=repeated_sint64,json=repeatedSint64" json:"repeated_sint64,omitempty"`
	RepeatedFixed32       []uint32                      `protobuf:"fixed32,37,rep,packed,name=repeated_fixed32,json=repeatedFixed32" json:"repeated_fixed32,omitempty"`
	RepeatedFixed64       []uint64                      `protobuf:"fixed64,38,rep,packed,name=repeated_fixed64,json=repeatedFixed64" json:"repeated_fixed64,omitempty"`
	RepeatedSfixed32      []int32                       `protobuf:"fixed32,39,rep,packed,name=repeated_sfixed32,json=repeatedSfixed32" json:"repeated_sfixed32,omitempty"`
	RepeatedSfixed64      []int64                       `protobuf:"fixed64,40,rep,packed,name=repeated_sfixed64,json=repeatedSfixed64" json:"repeated_sfixed64,omitempty"`
	RepeatedFloat         []float32                     `protobuf:"fixed32,41,rep,packed,name=repeated_float,json=repeatedFloat" json:"repeated_float,omitempty"`
	RepeatedDouble        []float64                     `protobuf:"fixed64,42,rep,packed,name=repeated_double,json=repeatedDouble" json:"repeated_double,omitempty"`
	RepeatedBool          []bool                        `protobuf:"varint,43,rep,packed,name=repeated_bool,json=repeatedBool" json:"repeated_bool,omitempty"`
	RepeatedString        []string                      `protobuf:"bytes,44,rep,name=repeated_string,json=repeatedString" json:"repeated_string,omitempty"`
	RepeatedBytes         [][]byte                      `protobuf:"bytes,45,rep,name=repeated_bytes,json=repeatedBytes,proto3" json:"repeated_bytes,omitempty"`
	RepeatedNestedMessage []*TestAllTypes_NestedMessage `protobuf:"bytes,48,rep,name=repeated_nested_message,json=repeatedNestedMessage" json:"repeated_nested_message,omitempty"`
	RepeatedNestedEnum    []TestAllTypes_NestedEnum     `protobuf:"varint,51,rep,packed,name=repeated_nested_enum,json=repeatedNestedEnum,enum=tensorflow.test.TestAllTypes_NestedEnum" json:"repeated_nested_enum,omitempty"`
	RepeatedCord          []string                      `protobuf:"bytes,55,rep,name=repeated_cord,json=repeatedCord" json:"repeated_cord,omitempty"`
	// Types that are valid to be assigned to OneofField:
	//	*TestAllTypes_OneofUint32
	//	*TestAllTypes_OneofNestedMessage
	//	*TestAllTypes_OneofString
	//	*TestAllTypes_OneofBytes
	//	*TestAllTypes_OneofEnum
	OneofField                isTestAllTypes_OneofField              `protobuf_oneof:"oneof_field"`
	MapStringToMessage        map[string]*TestAllTypes_NestedMessage `protobuf:"bytes,58,rep,name=map_string_to_message,json=mapStringToMessage" json:"map_string_to_message,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	MapInt32ToMessage         map[int32]*TestAllTypes_NestedMessage  `protobuf:"bytes,59,rep,name=map_int32_to_message,json=mapInt32ToMessage" json:"map_int32_to_message,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	MapInt64ToMessage         map[int64]*TestAllTypes_NestedMessage  `protobuf:"bytes,60,rep,name=map_int64_to_message,json=mapInt64ToMessage" json:"map_int64_to_message,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	MapBoolToMessage          map[bool]*TestAllTypes_NestedMessage   `protobuf:"bytes,61,rep,name=map_bool_to_message,json=mapBoolToMessage" json:"map_bool_to_message,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	MapStringToInt64          map[string]int64                       `protobuf:"bytes,62,rep,name=map_string_to_int64,json=mapStringToInt64" json:"map_string_to_int64,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	MapInt64ToString          map[int64]string                       `protobuf:"bytes,63,rep,name=map_int64_to_string,json=mapInt64ToString" json:"map_int64_to_string,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	AnotherMapStringToMessage map[string]*TestAllTypes_NestedMessage `protobuf:"bytes,65,rep,name=another_map_string_to_message,json=anotherMapStringToMessage" json:"another_map_string_to_message,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	PackedRepeatedInt64       []int64                                `protobuf:"varint,64,rep,packed,name=packed_repeated_int64,json=packedRepeatedInt64" json:"packed_repeated_int64,omitempty"`
}

func (m *TestAllTypes) Reset()                    { *m = TestAllTypes{} }
func (m *TestAllTypes) String() string            { return proto.CompactTextString(m) }
func (*TestAllTypes) ProtoMessage()               {}
func (*TestAllTypes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type isTestAllTypes_OneofField interface {
	isTestAllTypes_OneofField()
}

type TestAllTypes_OneofUint32 struct {
	OneofUint32 uint32 `protobuf:"varint,111,opt,name=oneof_uint32,json=oneofUint32,oneof"`
}
type TestAllTypes_OneofNestedMessage struct {
	OneofNestedMessage *TestAllTypes_NestedMessage `protobuf:"bytes,112,opt,name=oneof_nested_message,json=oneofNestedMessage,oneof"`
}
type TestAllTypes_OneofString struct {
	OneofString string `protobuf:"bytes,113,opt,name=oneof_string,json=oneofString,oneof"`
}
type TestAllTypes_OneofBytes struct {
	OneofBytes []byte `protobuf:"bytes,114,opt,name=oneof_bytes,json=oneofBytes,proto3,oneof"`
}
type TestAllTypes_OneofEnum struct {
	OneofEnum TestAllTypes_NestedEnum `protobuf:"varint,100,opt,name=oneof_enum,json=oneofEnum,enum=tensorflow.test.TestAllTypes_NestedEnum,oneof"`
}

func (*TestAllTypes_OneofUint32) isTestAllTypes_OneofField()        {}
func (*TestAllTypes_OneofNestedMessage) isTestAllTypes_OneofField() {}
func (*TestAllTypes_OneofString) isTestAllTypes_OneofField()        {}
func (*TestAllTypes_OneofBytes) isTestAllTypes_OneofField()         {}
func (*TestAllTypes_OneofEnum) isTestAllTypes_OneofField()          {}

func (m *TestAllTypes) GetOneofField() isTestAllTypes_OneofField {
	if m != nil {
		return m.OneofField
	}
	return nil
}

func (m *TestAllTypes) GetOptionalInt32() int32 {
	if m != nil {
		return m.OptionalInt32
	}
	return 0
}

func (m *TestAllTypes) GetOptionalInt64() int64 {
	if m != nil {
		return m.OptionalInt64
	}
	return 0
}

func (m *TestAllTypes) GetOptionalUint32() uint32 {
	if m != nil {
		return m.OptionalUint32
	}
	return 0
}

func (m *TestAllTypes) GetOptionalUint64() uint64 {
	if m != nil {
		return m.OptionalUint64
	}
	return 0
}

func (m *TestAllTypes) GetOptionalSint32() int32 {
	if m != nil {
		return m.OptionalSint32
	}
	return 0
}

func (m *TestAllTypes) GetOptionalSint64() int64 {
	if m != nil {
		return m.OptionalSint64
	}
	return 0
}

func (m *TestAllTypes) GetOptionalFixed32() uint32 {
	if m != nil {
		return m.OptionalFixed32
	}
	return 0
}

func (m *TestAllTypes) GetOptionalFixed64() uint64 {
	if m != nil {
		return m.OptionalFixed64
	}
	return 0
}

func (m *TestAllTypes) GetOptionalSfixed32() int32 {
	if m != nil {
		return m.OptionalSfixed32
	}
	return 0
}

func (m *TestAllTypes) GetOptionalSfixed64() int64 {
	if m != nil {
		return m.OptionalSfixed64
	}
	return 0
}

func (m *TestAllTypes) GetOptionalFloat() float32 {
	if m != nil {
		return m.OptionalFloat
	}
	return 0
}

func (m *TestAllTypes) GetOptionalDouble() float64 {
	if m != nil {
		return m.OptionalDouble
	}
	return 0
}

func (m *TestAllTypes) GetOptionalBool() bool {
	if m != nil {
		return m.OptionalBool
	}
	return false
}

func (m *TestAllTypes) GetOptionalString() string {
	if m != nil {
		return m.OptionalString
	}
	return ""
}

func (m *TestAllTypes) GetOptionalBytes() []byte {
	if m != nil {
		return m.OptionalBytes
	}
	return nil
}

func (m *TestAllTypes) GetOptionalNestedMessage() *TestAllTypes_NestedMessage {
	if m != nil {
		return m.OptionalNestedMessage
	}
	return nil
}

func (m *TestAllTypes) GetOptionalForeignMessage() *ForeignMessage {
	if m != nil {
		return m.OptionalForeignMessage
	}
	return nil
}

func (m *TestAllTypes) GetOptionalNestedEnum() TestAllTypes_NestedEnum {
	if m != nil {
		return m.OptionalNestedEnum
	}
	return TestAllTypes_ZERO
}

func (m *TestAllTypes) GetOptionalForeignEnum() ForeignEnum {
	if m != nil {
		return m.OptionalForeignEnum
	}
	return ForeignEnum_FOREIGN_ZERO
}

func (m *TestAllTypes) GetOptionalCord() string {
	if m != nil {
		return m.OptionalCord
	}
	return ""
}

func (m *TestAllTypes) GetRepeatedInt32() []int32 {
	if m != nil {
		return m.RepeatedInt32
	}
	return nil
}

func (m *TestAllTypes) GetRepeatedInt64() []int64 {
	if m != nil {
		return m.RepeatedInt64
	}
	return nil
}

func (m *TestAllTypes) GetRepeatedUint32() []uint32 {
	if m != nil {
		return m.RepeatedUint32
	}
	return nil
}

func (m *TestAllTypes) GetRepeatedUint64() []uint64 {
	if m != nil {
		return m.RepeatedUint64
	}
	return nil
}

func (m *TestAllTypes) GetRepeatedSint32() []int32 {
	if m != nil {
		return m.RepeatedSint32
	}
	return nil
}

func (m *TestAllTypes) GetRepeatedSint64() []int64 {
	if m != nil {
		return m.RepeatedSint64
	}
	return nil
}

func (m *TestAllTypes) GetRepeatedFixed32() []uint32 {
	if m != nil {
		return m.RepeatedFixed32
	}
	return nil
}

func (m *TestAllTypes) GetRepeatedFixed64() []uint64 {
	if m != nil {
		return m.RepeatedFixed64
	}
	return nil
}

func (m *TestAllTypes) GetRepeatedSfixed32() []int32 {
	if m != nil {
		return m.RepeatedSfixed32
	}
	return nil
}

func (m *TestAllTypes) GetRepeatedSfixed64() []int64 {
	if m != nil {
		return m.RepeatedSfixed64
	}
	return nil
}

func (m *TestAllTypes) GetRepeatedFloat() []float32 {
	if m != nil {
		return m.RepeatedFloat
	}
	return nil
}

func (m *TestAllTypes) GetRepeatedDouble() []float64 {
	if m != nil {
		return m.RepeatedDouble
	}
	return nil
}

func (m *TestAllTypes) GetRepeatedBool() []bool {
	if m != nil {
		return m.RepeatedBool
	}
	return nil
}

func (m *TestAllTypes) GetRepeatedString() []string {
	if m != nil {
		return m.RepeatedString
	}
	return nil
}

func (m *TestAllTypes) GetRepeatedBytes() [][]byte {
	if m != nil {
		return m.RepeatedBytes
	}
	return nil
}

func (m *TestAllTypes) GetRepeatedNestedMessage() []*TestAllTypes_NestedMessage {
	if m != nil {
		return m.RepeatedNestedMessage
	}
	return nil
}

func (m *TestAllTypes) GetRepeatedNestedEnum() []TestAllTypes_NestedEnum {
	if m != nil {
		return m.RepeatedNestedEnum
	}
	return nil
}

func (m *TestAllTypes) GetRepeatedCord() []string {
	if m != nil {
		return m.RepeatedCord
	}
	return nil
}

func (m *TestAllTypes) GetOneofUint32() uint32 {
	if x, ok := m.GetOneofField().(*TestAllTypes_OneofUint32); ok {
		return x.OneofUint32
	}
	return 0
}

func (m *TestAllTypes) GetOneofNestedMessage() *TestAllTypes_NestedMessage {
	if x, ok := m.GetOneofField().(*TestAllTypes_OneofNestedMessage); ok {
		return x.OneofNestedMessage
	}
	return nil
}

func (m *TestAllTypes) GetOneofString() string {
	if x, ok := m.GetOneofField().(*TestAllTypes_OneofString); ok {
		return x.OneofString
	}
	return ""
}

func (m *TestAllTypes) GetOneofBytes() []byte {
	if x, ok := m.GetOneofField().(*TestAllTypes_OneofBytes); ok {
		return x.OneofBytes
	}
	return nil
}

func (m *TestAllTypes) GetOneofEnum() TestAllTypes_NestedEnum {
	if x, ok := m.GetOneofField().(*TestAllTypes_OneofEnum); ok {
		return x.OneofEnum
	}
	return TestAllTypes_ZERO
}

func (m *TestAllTypes) GetMapStringToMessage() map[string]*TestAllTypes_NestedMessage {
	if m != nil {
		return m.MapStringToMessage
	}
	return nil
}

func (m *TestAllTypes) GetMapInt32ToMessage() map[int32]*TestAllTypes_NestedMessage {
	if m != nil {
		return m.MapInt32ToMessage
	}
	return nil
}

func (m *TestAllTypes) GetMapInt64ToMessage() map[int64]*TestAllTypes_NestedMessage {
	if m != nil {
		return m.MapInt64ToMessage
	}
	return nil
}

func (m *TestAllTypes) GetMapBoolToMessage() map[bool]*TestAllTypes_NestedMessage {
	if m != nil {
		return m.MapBoolToMessage
	}
	return nil
}

func (m *TestAllTypes) GetMapStringToInt64() map[string]int64 {
	if m != nil {
		return m.MapStringToInt64
	}
	return nil
}

func (m *TestAllTypes) GetMapInt64ToString() map[int64]string {
	if m != nil {
		return m.MapInt64ToString
	}
	return nil
}

func (m *TestAllTypes) GetAnotherMapStringToMessage() map[string]*TestAllTypes_NestedMessage {
	if m != nil {
		return m.AnotherMapStringToMessage
	}
	return nil
}

func (m *TestAllTypes) GetPackedRepeatedInt64() []int64 {
	if m != nil {
		return m.PackedRepeatedInt64
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*TestAllTypes) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _TestAllTypes_OneofMarshaler, _TestAllTypes_OneofUnmarshaler, _TestAllTypes_OneofSizer, []interface{}{
		(*TestAllTypes_OneofUint32)(nil),
		(*TestAllTypes_OneofNestedMessage)(nil),
		(*TestAllTypes_OneofString)(nil),
		(*TestAllTypes_OneofBytes)(nil),
		(*TestAllTypes_OneofEnum)(nil),
	}
}

func _TestAllTypes_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*TestAllTypes)
	// oneof_field
	switch x := m.OneofField.(type) {
	case *TestAllTypes_OneofUint32:
		b.EncodeVarint(111<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.OneofUint32))
	case *TestAllTypes_OneofNestedMessage:
		b.EncodeVarint(112<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.OneofNestedMessage); err != nil {
			return err
		}
	case *TestAllTypes_OneofString:
		b.EncodeVarint(113<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.OneofString)
	case *TestAllTypes_OneofBytes:
		b.EncodeVarint(114<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.OneofBytes)
	case *TestAllTypes_OneofEnum:
		b.EncodeVarint(100<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.OneofEnum))
	case nil:
	default:
		return fmt.Errorf("TestAllTypes.OneofField has unexpected type %T", x)
	}
	return nil
}

func _TestAllTypes_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*TestAllTypes)
	switch tag {
	case 111: // oneof_field.oneof_uint32
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.OneofField = &TestAllTypes_OneofUint32{uint32(x)}
		return true, err
	case 112: // oneof_field.oneof_nested_message
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TestAllTypes_NestedMessage)
		err := b.DecodeMessage(msg)
		m.OneofField = &TestAllTypes_OneofNestedMessage{msg}
		return true, err
	case 113: // oneof_field.oneof_string
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.OneofField = &TestAllTypes_OneofString{x}
		return true, err
	case 114: // oneof_field.oneof_bytes
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.OneofField = &TestAllTypes_OneofBytes{x}
		return true, err
	case 100: // oneof_field.oneof_enum
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.OneofField = &TestAllTypes_OneofEnum{TestAllTypes_NestedEnum(x)}
		return true, err
	default:
		return false, nil
	}
}

func _TestAllTypes_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*TestAllTypes)
	// oneof_field
	switch x := m.OneofField.(type) {
	case *TestAllTypes_OneofUint32:
		n += proto.SizeVarint(111<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.OneofUint32))
	case *TestAllTypes_OneofNestedMessage:
		s := proto.Size(x.OneofNestedMessage)
		n += proto.SizeVarint(112<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *TestAllTypes_OneofString:
		n += proto.SizeVarint(113<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.OneofString)))
		n += len(x.OneofString)
	case *TestAllTypes_OneofBytes:
		n += proto.SizeVarint(114<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.OneofBytes)))
		n += len(x.OneofBytes)
	case *TestAllTypes_OneofEnum:
		n += proto.SizeVarint(100<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.OneofEnum))
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type TestAllTypes_NestedMessage struct {
	OptionalInt32 int32                                           `protobuf:"varint,1,opt,name=optional_int32,json=optionalInt32" json:"optional_int32,omitempty"`
	RepeatedInt32 []int32                                         `protobuf:"varint,2,rep,packed,name=repeated_int32,json=repeatedInt32" json:"repeated_int32,omitempty"`
	Msg           *TestAllTypes_NestedMessage_DoubleNestedMessage `protobuf:"bytes,3,opt,name=msg" json:"msg,omitempty"`
	OptionalInt64 int64                                           `protobuf:"varint,4,opt,name=optional_int64,json=optionalInt64" json:"optional_int64,omitempty"`
}

func (m *TestAllTypes_NestedMessage) Reset()                    { *m = TestAllTypes_NestedMessage{} }
func (m *TestAllTypes_NestedMessage) String() string            { return proto.CompactTextString(m) }
func (*TestAllTypes_NestedMessage) ProtoMessage()               {}
func (*TestAllTypes_NestedMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *TestAllTypes_NestedMessage) GetOptionalInt32() int32 {
	if m != nil {
		return m.OptionalInt32
	}
	return 0
}

func (m *TestAllTypes_NestedMessage) GetRepeatedInt32() []int32 {
	if m != nil {
		return m.RepeatedInt32
	}
	return nil
}

func (m *TestAllTypes_NestedMessage) GetMsg() *TestAllTypes_NestedMessage_DoubleNestedMessage {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (m *TestAllTypes_NestedMessage) GetOptionalInt64() int64 {
	if m != nil {
		return m.OptionalInt64
	}
	return 0
}

type TestAllTypes_NestedMessage_DoubleNestedMessage struct {
	OptionalString string `protobuf:"bytes,1,opt,name=optional_string,json=optionalString" json:"optional_string,omitempty"`
}

func (m *TestAllTypes_NestedMessage_DoubleNestedMessage) Reset() {
	*m = TestAllTypes_NestedMessage_DoubleNestedMessage{}
}
func (m *TestAllTypes_NestedMessage_DoubleNestedMessage) String() string {
	return proto.CompactTextString(m)
}
func (*TestAllTypes_NestedMessage_DoubleNestedMessage) ProtoMessage() {}
func (*TestAllTypes_NestedMessage_DoubleNestedMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0, 0}
}

func (m *TestAllTypes_NestedMessage_DoubleNestedMessage) GetOptionalString() string {
	if m != nil {
		return m.OptionalString
	}
	return ""
}

// A recursive message.
type NestedTestAllTypes struct {
	Child            *NestedTestAllTypes `protobuf:"bytes,1,opt,name=child" json:"child,omitempty"`
	Payload          *TestAllTypes       `protobuf:"bytes,2,opt,name=payload" json:"payload,omitempty"`
	MapStringToInt64 map[string]int64    `protobuf:"bytes,3,rep,name=map_string_to_int64,json=mapStringToInt64" json:"map_string_to_int64,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
}

func (m *NestedTestAllTypes) Reset()                    { *m = NestedTestAllTypes{} }
func (m *NestedTestAllTypes) String() string            { return proto.CompactTextString(m) }
func (*NestedTestAllTypes) ProtoMessage()               {}
func (*NestedTestAllTypes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *NestedTestAllTypes) GetChild() *NestedTestAllTypes {
	if m != nil {
		return m.Child
	}
	return nil
}

func (m *NestedTestAllTypes) GetPayload() *TestAllTypes {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *NestedTestAllTypes) GetMapStringToInt64() map[string]int64 {
	if m != nil {
		return m.MapStringToInt64
	}
	return nil
}

type ForeignMessage struct {
	C int32 `protobuf:"varint,1,opt,name=c" json:"c,omitempty"`
}

func (m *ForeignMessage) Reset()                    { *m = ForeignMessage{} }
func (m *ForeignMessage) String() string            { return proto.CompactTextString(m) }
func (*ForeignMessage) ProtoMessage()               {}
func (*ForeignMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ForeignMessage) GetC() int32 {
	if m != nil {
		return m.C
	}
	return 0
}

type TestEmptyMessage struct {
}

func (m *TestEmptyMessage) Reset()                    { *m = TestEmptyMessage{} }
func (m *TestEmptyMessage) String() string            { return proto.CompactTextString(m) }
func (*TestEmptyMessage) ProtoMessage()               {}
func (*TestEmptyMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func init() {
	proto.RegisterType((*TestAllTypes)(nil), "tensorflow.test.TestAllTypes")
	proto.RegisterType((*TestAllTypes_NestedMessage)(nil), "tensorflow.test.TestAllTypes.NestedMessage")
	proto.RegisterType((*TestAllTypes_NestedMessage_DoubleNestedMessage)(nil), "tensorflow.test.TestAllTypes.NestedMessage.DoubleNestedMessage")
	proto.RegisterType((*NestedTestAllTypes)(nil), "tensorflow.test.NestedTestAllTypes")
	proto.RegisterType((*ForeignMessage)(nil), "tensorflow.test.ForeignMessage")
	proto.RegisterType((*TestEmptyMessage)(nil), "tensorflow.test.TestEmptyMessage")
	proto.RegisterEnum("tensorflow.test.ForeignEnum", ForeignEnum_name, ForeignEnum_value)
	proto.RegisterEnum("tensorflow.test.TestAllTypes_NestedEnum", TestAllTypes_NestedEnum_name, TestAllTypes_NestedEnum_value)
}

func init() { proto.RegisterFile("tensorflow/tools/proto_text/test.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1358 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x58, 0xdf, 0x76, 0xd3, 0x36,
	0x1c, 0xae, 0xa2, 0xfe, 0x55, 0xd2, 0xd6, 0xa8, 0x2d, 0x98, 0x9e, 0x01, 0xa2, 0x1d, 0x60, 0x60,
	0x24, 0x3b, 0xad, 0x67, 0x06, 0x63, 0xb0, 0x96, 0xb5, 0x94, 0x0b, 0xe8, 0x66, 0xe0, 0x62, 0xbd,
	0xc9, 0x49, 0x13, 0x35, 0xcd, 0xa9, 0x13, 0x19, 0xdb, 0x65, 0xe4, 0x66, 0xaf, 0xb1, 0x47, 0xd9,
	0x43, 0xec, 0x41, 0xb6, 0xb7, 0xd8, 0x8e, 0x24, 0x4b, 0xb6, 0x1c, 0xb7, 0x34, 0xe7, 0x00, 0x57,
	0xce, 0xd7, 0x4f, 0xdf, 0xf7, 0xd3, 0x4f, 0xfa, 0x24, 0x1b, 0x74, 0x3b, 0xa1, 0x83, 0x98, 0x45,
	0x47, 0x01, 0xfb, 0xbd, 0x91, 0x30, 0x16, 0xc4, 0x8d, 0x30, 0x62, 0x09, 0x6b, 0x26, 0xf4, 0x63,
	0xd2, 0x48, 0x68, 0x9c, 0xd4, 0xc5, 0x6f, 0xbc, 0x98, 0xf1, 0xea, 0x1c, 0x5e, 0xfb, 0xfb, 0x06,
	0xaa, 0xbd, 0xa5, 0x71, 0xb2, 0x15, 0x04, 0x6f, 0x87, 0x21, 0x8d, 0xf1, 0x6d, 0xb4, 0xc0, 0xc2,
	0xa4, 0xc7, 0x06, 0xad, 0xa0, 0xd9, 0x1b, 0x24, 0x9b, 0x1b, 0xf6, 0xbf, 0x33, 0x04, 0x38, 0x53,
	0xfe, 0xbc, 0x82, 0x5f, 0x72, 0x14, 0xdf, 0x32, 0x79, 0x9e, 0x6b, 0x57, 0x08, 0x70, 0xa0, 0x41,
	0xf3, 0x5c, 0x7c, 0x07, 0x2d, 0x6a, 0xda, 0xa9, 0xd4, 0x83, 0x04, 0x38, 0xf3, 0xbe, 0x1e, 0xfd,
	0x4e, 0xa0, 0xd8, 0x29, 0x10, 0x3d, 0xd7, 0xfe, 0x87, 0x1b, 0x4f, 0x9a, 0xcc, 0x82, 0x64, 0x2c,
	0x25, 0xa7, 0x08, 0x70, 0x2e, 0x65, 0xc4, 0x37, 0x52, 0xb2, 0x48, 0xf4, 0x5c, 0x7b, 0x9a, 0x00,
	0x07, 0x9b, 0x44, 0xcf, 0xc5, 0x77, 0x91, 0xa5, 0x89, 0x47, 0xbd, 0x8f, 0xb4, 0xb3, 0xb9, 0x61,
	0x73, 0xef, 0x19, 0x5f, 0x0b, 0xec, 0x4a, 0x78, 0x94, 0xea, 0xb9, 0xf6, 0x2c, 0x01, 0xce, 0x74,
	0x81, 0xea, 0xb9, 0xf8, 0x3e, 0xba, 0x94, 0xd9, 0x2b, 0xd9, 0x39, 0x02, 0x9c, 0x45, 0x5f, 0x6b,
	0xbc, 0x49, 0xf1, 0x12, 0xb2, 0xe7, 0xda, 0x88, 0x00, 0xc7, 0x2a, 0x92, 0x3d, 0xd7, 0xe8, 0xfd,
	0x51, 0xc0, 0x5a, 0x89, 0x5d, 0x25, 0xc0, 0xa9, 0x64, 0xbd, 0xdf, 0xe5, 0xa0, 0x31, 0xff, 0x0e,
	0x3b, 0x3d, 0x0c, 0xa8, 0x5d, 0x23, 0xc0, 0x01, 0xd9, 0xfc, 0x7f, 0x16, 0x28, 0x5e, 0x47, 0x7a,
	0x64, 0xf3, 0x90, 0xb1, 0xc0, 0x9e, 0x27, 0xc0, 0x99, 0xf5, 0x6b, 0x0a, 0xdc, 0x66, 0x2c, 0x30,
	0xbb, 0x99, 0x44, 0xbd, 0x41, 0xd7, 0x5e, 0x20, 0xc0, 0x99, 0xcb, 0x75, 0x53, 0xa0, 0x46, 0x75,
	0x87, 0xc3, 0x84, 0xc6, 0xf6, 0x22, 0x01, 0x4e, 0x2d, 0xab, 0x6e, 0x9b, 0x83, 0xb8, 0x8d, 0xae,
	0x68, 0xda, 0x80, 0xc6, 0x09, 0xed, 0x34, 0xfb, 0x34, 0x8e, 0x5b, 0x5d, 0x6a, 0x63, 0x02, 0x9c,
	0xea, 0xc6, 0xfd, 0x7a, 0x61, 0xb3, 0xd6, 0xf3, 0x1b, 0xb5, 0xfe, 0x5a, 0x8c, 0x79, 0x25, 0x87,
	0xf8, 0x2b, 0x4a, 0xcb, 0x80, 0xf1, 0x6f, 0xc8, 0xce, 0x3a, 0xc5, 0x22, 0xda, 0xeb, 0x0e, 0xb4,
	0xcb, 0x92, 0x70, 0xb9, 0x31, 0xe2, 0xb2, 0x2b, 0x79, 0x4a, 0xf9, 0xb2, 0x6e, 0xaa, 0x81, 0xe3,
	0x03, 0xb4, 0x5c, 0xac, 0x9f, 0x0e, 0x4e, 0xfb, 0xf6, 0x0a, 0x01, 0xce, 0xc2, 0x86, 0x73, 0x91,
	0xe2, 0x77, 0x06, 0xa7, 0x7d, 0x1f, 0x9b, 0x95, 0x73, 0x0c, 0xff, 0x82, 0x56, 0x46, 0xca, 0x16,
	0xe2, 0x97, 0x85, 0xf8, 0x57, 0x67, 0xd5, 0x2c, 0x04, 0x97, 0x0a, 0x05, 0x0b, 0xc5, 0xfc, 0x12,
	0xb7, 0x59, 0xd4, 0xb1, 0xaf, 0x8a, 0xb5, 0xd3, 0x4b, 0xfc, 0x9c, 0x45, 0x1d, 0xbe, 0x72, 0x11,
	0x0d, 0x69, 0x8b, 0xcf, 0x45, 0x06, 0xeb, 0x06, 0x81, 0x3c, 0xfa, 0x0a, 0xd5, 0xd1, 0xcf, 0xd3,
	0x3c, 0xd7, 0x26, 0x04, 0xf2, 0xe8, 0xe7, 0x68, 0x32, 0xa7, 0x9a, 0x96, 0x46, 0xff, 0x26, 0x81,
	0x3c, 0xfa, 0x0a, 0x7e, 0xa7, 0x73, 0x6a, 0x10, 0x3d, 0xd7, 0x5e, 0x23, 0x90, 0x27, 0x3f, 0x4f,
	0x2c, 0x28, 0xa6, 0xc9, 0x5f, 0x27, 0x90, 0x27, 0x5f, 0xc1, 0x6f, 0x46, 0x15, 0xd3, 0xe4, 0x7f,
	0x4d, 0x20, 0x4f, 0x7e, 0x9e, 0x28, 0x93, 0xaf, 0x89, 0x2a, 0xa2, 0xb7, 0x08, 0xe4, 0xc9, 0x57,
	0x78, 0x2e, 0xf9, 0x26, 0xd5, 0x73, 0xed, 0xdb, 0x04, 0xf2, 0xe4, 0x1b, 0x54, 0x99, 0xfc, 0xcc,
	0x5e, 0xc9, 0xde, 0x21, 0x90, 0x27, 0x5f, 0x17, 0x90, 0x4b, 0x7e, 0x81, 0xec, 0xb9, 0xb6, 0x43,
	0x20, 0x4f, 0xbe, 0x49, 0x96, 0xc9, 0xcf, 0x8a, 0x10, 0xc9, 0xbf, 0x4b, 0x20, 0x4f, 0xbe, 0x2e,
	0x41, 0x25, 0x5f, 0xd3, 0xd2, 0xe4, 0xdf, 0x23, 0x90, 0x27, 0x5f, 0xc1, 0x59, 0xf2, 0x35, 0x51,
	0x24, 0xff, 0x3e, 0x81, 0x3c, 0xf9, 0x0a, 0x54, 0xc9, 0xcf, 0x2a, 0x94, 0xc9, 0xff, 0x86, 0x40,
	0x9e, 0x7c, 0x5d, 0x9f, 0x4e, 0x7e, 0xa6, 0x26, 0x92, 0xff, 0x80, 0x40, 0x9e, 0x7c, 0x2d, 0xa7,
	0x92, 0xaf, 0x69, 0x85, 0xe4, 0x7f, 0x4b, 0xe0, 0xd8, 0xc9, 0x57, 0x5a, 0x66, 0xf2, 0x0f, 0xd0,
	0x72, 0xd1, 0x44, 0x24, 0x68, 0x93, 0xc0, 0xf1, 0xe2, 0x69, 0xca, 0xab, 0x30, 0x69, 0x6d, 0x11,
	0xa6, 0x87, 0xa2, 0x1d, 0xba, 0x6b, 0x22, 0x4c, 0xeb, 0xa8, 0xc6, 0x06, 0x94, 0x1d, 0xa9, 0xbd,
	0xcf, 0xf8, 0xb5, 0xb7, 0x37, 0xe1, 0x57, 0x05, 0x9a, 0x6e, 0xfd, 0x26, 0x5a, 0x96, 0xa4, 0x42,
	0x1f, 0xc2, 0xb1, 0x4f, 0xc0, 0xbd, 0x09, 0x1f, 0x0b, 0x29, 0xb3, 0x0d, 0xba, 0x8a, 0x74, 0xe1,
	0xde, 0xf3, 0xd8, 0xeb, 0x2a, 0xd2, 0x75, 0xbb, 0x89, 0xe4, 0xcf, 0x74, 0xd1, 0x22, 0x7e, 0x5c,
	0xef, 0x4d, 0xf8, 0x48, 0x80, 0x72, 0xcd, 0x5e, 0x22, 0xf9, 0x4b, 0x36, 0xb1, 0x33, 0xde, 0x19,
	0xb7, 0x37, 0xe1, 0xcf, 0x89, 0xd1, 0xa2, 0x7b, 0xc7, 0x68, 0xa5, 0xdf, 0x0a, 0xd3, 0x82, 0x9a,
	0x09, 0xd3, 0x93, 0x7e, 0x2c, 0x16, 0xff, 0xbb, 0xf3, 0x55, 0x5f, 0xb5, 0x42, 0x59, 0xf5, 0x5b,
	0x96, 0xce, 0x71, 0x67, 0x90, 0x44, 0x43, 0x1f, 0xf7, 0x47, 0xfe, 0x80, 0x29, 0x5a, 0xe6, 0x4e,
	0xa2, 0xd5, 0x79, 0xa3, 0x1f, 0x84, 0x91, 0xfb, 0x49, 0x23, 0x71, 0xdc, 0x15, 0x7c, 0x2e, 0xf5,
	0x8b, 0x78, 0xce, 0xc6, 0x73, 0xf3, 0x36, 0x4f, 0x2e, 0x6e, 0xe3, 0xb9, 0xe5, 0x36, 0x39, 0x1c,
	0x1f, 0xa2, 0x25, 0x6e, 0xc3, 0x63, 0x9a, 0x77, 0xf9, 0x51, 0xb8, 0x6c, 0x7e, 0xd2, 0x85, 0x47,
	0xb9, 0x60, 0x62, 0xf5, 0x0b, 0xb0, 0xf2, 0xc8, 0xd6, 0x46, 0x1e, 0x9e, 0x4f, 0x2f, 0xe8, 0xa1,
	0x16, 0x40, 0x54, 0x9e, 0x79, 0x18, 0xb0, 0xf2, 0xd0, 0xed, 0x4a, 0x77, 0xe6, 0xb3, 0x0b, 0x7a,
	0xa4, 0x5d, 0x91, 0x9a, 0x99, 0x87, 0x01, 0xe3, 0x3f, 0xd0, 0xb5, 0xd6, 0x80, 0x25, 0xc7, 0x34,
	0x6a, 0x96, 0xef, 0xb5, 0x2d, 0xe1, 0xf6, 0xe4, 0x7c, 0xb7, 0x2d, 0x29, 0x71, 0xd6, 0x96, 0xbb,
	0xda, 0x3a, 0xeb, 0xef, 0xd8, 0x43, 0x2b, 0x61, 0xab, 0x7d, 0x42, 0x3b, 0xcd, 0xc2, 0x4d, 0xf9,
	0x13, 0xbf, 0x29, 0xb7, 0x2b, 0x16, 0xf0, 0x97, 0x24, 0xc1, 0xcf, 0xdf, 0x99, 0xab, 0x7f, 0x56,
	0xd0, 0xbc, 0x19, 0xe0, 0x5b, 0x23, 0xef, 0xe3, 0xe0, 0x8c, 0xd7, 0xf1, 0xc2, 0xd5, 0x5d, 0x29,
	0xbb, 0xba, 0x7f, 0x45, 0xb0, 0x1f, 0x77, 0xc5, 0x2b, 0x78, 0x75, 0xe3, 0xd9, 0x18, 0xc7, 0x4b,
	0x5d, 0x5e, 0x18, 0xe6, 0xd1, 0xcb, 0xb5, 0x4a, 0x3e, 0x04, 0x26, 0x4b, 0x3e, 0x04, 0x56, 0x9f,
	0xa2, 0xa5, 0x12, 0x89, 0xb2, 0xb7, 0x4a, 0x50, 0xf6, 0x56, 0xb9, 0x1a, 0xa1, 0x2b, 0x67, 0xac,
	0x03, 0xb6, 0x10, 0x3c, 0xa1, 0xc3, 0x74, 0x1c, 0x7f, 0xc4, 0x5b, 0x68, 0xea, 0x43, 0x2b, 0x38,
	0xa5, 0xe2, 0x9b, 0x64, 0xcc, 0xfb, 0x44, 0x8e, 0x7c, 0x5c, 0xf9, 0x1e, 0xac, 0xbe, 0x47, 0x97,
	0xcb, 0x4f, 0x81, 0xbc, 0xe5, 0xd4, 0x97, 0xb1, 0x2c, 0x9e, 0x08, 0x79, 0x4b, 0xf8, 0x59, 0x2d,
	0x43, 0xb4, 0x52, 0x7a, 0x3c, 0xe4, 0x1d, 0x67, 0x3f, 0xab, 0xe3, 0x73, 0xe1, 0x38, 0x7a, 0x58,
	0x94, 0xac, 0xe4, 0x72, 0xde, 0x11, 0x8e, 0x8a, 0x8c, 0x9e, 0x06, 0x25, 0x8d, 0x32, 0x44, 0xe6,
	0xf2, 0x22, 0x43, 0x74, 0xfd, 0xfc, 0x90, 0x7f, 0xb1, 0xcd, 0xb5, 0xf6, 0x1c, 0xa1, 0xdc, 0x2b,
	0xc5, 0x2c, 0x9a, 0x3c, 0xd8, 0xf1, 0xf7, 0xad, 0x09, 0x3c, 0x83, 0xe0, 0xee, 0xfe, 0xbe, 0x05,
	0xf8, 0xc3, 0xf6, 0x96, 0x6f, 0x55, 0xe4, 0xc3, 0x81, 0x05, 0x79, 0x2d, 0xaf, 0x77, 0x5e, 0x58,
	0xff, 0xa9, 0x7f, 0x60, 0x7b, 0x5e, 0xdd, 0xdc, 0x47, 0x3d, 0x1a, 0x74, 0xd6, 0xfe, 0xaa, 0x20,
	0x2c, 0x45, 0x8d, 0x6f, 0xfa, 0x47, 0x68, 0xaa, 0x7d, 0xdc, 0x0b, 0x3a, 0x62, 0x16, 0xd5, 0x8d,
	0xf5, 0x91, 0x8a, 0x47, 0xc7, 0xf8, 0x72, 0x04, 0x7e, 0x88, 0x66, 0xc2, 0xd6, 0x30, 0x60, 0xad,
	0x4e, 0x3a, 0xdd, 0x6b, 0xe7, 0x4e, 0xd7, 0x57, 0x6c, 0x7c, 0x5c, 0x7e, 0x93, 0x40, 0x71, 0xee,
	0x3e, 0xba, 0x40, 0x05, 0x17, 0xbd, 0x4f, 0x3e, 0xcb, 0x6e, 0x5a, 0xbb, 0x8e, 0x16, 0x0a, 0xdf,
	0x77, 0x35, 0x04, 0xda, 0x69, 0xc0, 0x41, 0x7b, 0x0d, 0x23, 0x8b, 0x17, 0xb7, 0xd3, 0x0f, 0x93,
	0x61, 0xca, 0xb8, 0xe7, 0xa3, 0x6a, 0xfe, 0x13, 0xcb, 0x42, 0xb5, 0xdd, 0x7d, 0x7f, 0xe7, 0xe5,
	0x8b, 0xd7, 0xcd, 0x74, 0x29, 0x17, 0x51, 0x55, 0x21, 0x7c, 0x49, 0x27, 0xf3, 0x00, 0x5f, 0xda,
	0x29, 0x13, 0x38, 0xb0, 0xa6, 0xb7, 0xf7, 0x0e, 0x76, 0xbb, 0xbd, 0xe4, 0xf8, 0xf4, 0xb0, 0xde,
	0x66, 0xfd, 0x46, 0x27, 0xa2, 0xc3, 0x93, 0x46, 0xd6, 0xab, 0x07, 0x31, 0x8d, 0x3e, 0xf4, 0x06,
	0xdd, 0x07, 0x5d, 0xd6, 0x08, 0x4f, 0xba, 0x8d, 0x73, 0xfe, 0xd7, 0xe7, 0x70, 0x5a, 0x3c, 0x6f,
	0xfe, 0x1f, 0x00, 0x00, 0xff, 0xff, 0x39, 0xa4, 0x8c, 0x13, 0x1b, 0x12, 0x00, 0x00,
}
