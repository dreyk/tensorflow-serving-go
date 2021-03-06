// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/framework/versions.proto

package framework

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Version information for a piece of serialized data
//
// There are different types of versions for each type of data
// (GraphDef, etc.), but they all have the same common shape
// described here.
//
// Each consumer has "consumer" and "min_producer" versions (specified
// elsewhere).  A consumer is allowed to consume this data if
//
//   producer >= min_producer
//   consumer >= min_consumer
//   consumer not in bad_consumers
//
type VersionDef struct {
	// The version of the code that produced this data.
	Producer int32 `protobuf:"varint,1,opt,name=producer" json:"producer,omitempty"`
	// Any consumer below this version is not allowed to consume this data.
	MinConsumer int32 `protobuf:"varint,2,opt,name=min_consumer,json=minConsumer" json:"min_consumer,omitempty"`
	// Specific consumer versions which are disallowed (e.g. due to bugs).
	BadConsumers []int32 `protobuf:"varint,3,rep,packed,name=bad_consumers,json=badConsumers" json:"bad_consumers,omitempty"`
}

func (m *VersionDef) Reset()                    { *m = VersionDef{} }
func (m *VersionDef) String() string            { return proto.CompactTextString(m) }
func (*VersionDef) ProtoMessage()               {}
func (*VersionDef) Descriptor() ([]byte, []int) { return fileDescriptor23, []int{0} }

func (m *VersionDef) GetProducer() int32 {
	if m != nil {
		return m.Producer
	}
	return 0
}

func (m *VersionDef) GetMinConsumer() int32 {
	if m != nil {
		return m.MinConsumer
	}
	return 0
}

func (m *VersionDef) GetBadConsumers() []int32 {
	if m != nil {
		return m.BadConsumers
	}
	return nil
}

func init() {
	proto.RegisterType((*VersionDef)(nil), "tensorflow.VersionDef")
}

func init() { proto.RegisterFile("tensorflow/core/framework/versions.proto", fileDescriptor23) }

var fileDescriptor23 = []byte{
	// 219 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x65, 0xaa, 0x22, 0x74, 0x14, 0x86, 0x4c, 0x11, 0x53, 0x81, 0x25, 0x4b, 0xe3, 0x81,
	0x37, 0x28, 0x7d, 0x80, 0xaa, 0x03, 0x03, 0x0b, 0x4a, 0xec, 0x8b, 0xb1, 0x82, 0x7d, 0xd6, 0x5d,
	0xdc, 0x8a, 0x37, 0x67, 0x44, 0x94, 0x90, 0x4c, 0x8c, 0x77, 0xff, 0x77, 0xd2, 0x77, 0x3f, 0x54,
	0x03, 0x46, 0x21, 0xee, 0x3e, 0xe8, 0xa4, 0x0d, 0x31, 0xea, 0x8e, 0x9b, 0x80, 0x27, 0xe2, 0x5e,
	0x1f, 0x91, 0xc5, 0x53, 0x94, 0x3a, 0x31, 0x0d, 0x54, 0xc0, 0x4c, 0x3e, 0x24, 0x80, 0x97, 0xdf,
	0x74, 0x87, 0x5d, 0x71, 0x07, 0x57, 0x89, 0xc9, 0x66, 0x83, 0x5c, 0xaa, 0xb5, 0xaa, 0x96, 0x87,
	0x69, 0x2e, 0xee, 0x61, 0x15, 0x7c, 0x7c, 0x33, 0x14, 0x25, 0x07, 0xe4, 0xf2, 0xe2, 0x9c, 0x5f,
	0x07, 0x1f, 0x9f, 0xc7, 0x55, 0xf1, 0x08, 0x37, 0x6d, 0x63, 0x27, 0x44, 0xca, 0xc5, 0x7a, 0x51,
	0x2d, 0x0f, 0xab, 0xb6, 0xb1, 0x7f, 0x8c, 0x6c, 0x33, 0x94, 0xc4, 0xae, 0x9e, 0x1d, 0xea, 0x49,
	0x74, 0x7b, 0x3b, 0xba, 0xc8, 0xfe, 0x47, 0x54, 0xf6, 0xea, 0x75, 0xe7, 0xfc, 0xf0, 0x9e, 0xdb,
	0xda, 0x50, 0xd0, 0x96, 0xf1, 0xb3, 0xd7, 0xf3, 0xe1, 0x46, 0x90, 0x8f, 0x3e, 0xba, 0x8d, 0x23,
	0x9d, 0x7a, 0xa7, 0xff, 0x2d, 0xe0, 0x4b, 0xa9, 0xf6, 0xf2, 0xfc, 0xfb, 0xd3, 0x77, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xae, 0x94, 0x76, 0x68, 0x27, 0x01, 0x00, 0x00,
}
