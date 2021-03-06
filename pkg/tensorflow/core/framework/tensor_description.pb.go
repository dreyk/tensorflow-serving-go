// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/framework/tensor_description.proto

package framework

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type TensorDescription struct {
	// Data type of tensor elements
	Dtype DataType `protobuf:"varint,1,opt,name=dtype,enum=tensorflow.DataType" json:"dtype,omitempty"`
	// Shape of the tensor.
	Shape *TensorShapeProto `protobuf:"bytes,2,opt,name=shape" json:"shape,omitempty"`
	// Information about the size and allocator used for the data
	AllocationDescription *AllocationDescription `protobuf:"bytes,4,opt,name=allocation_description,json=allocationDescription" json:"allocation_description,omitempty"`
}

func (m *TensorDescription) Reset()                    { *m = TensorDescription{} }
func (m *TensorDescription) String() string            { return proto.CompactTextString(m) }
func (*TensorDescription) ProtoMessage()               {}
func (*TensorDescription) Descriptor() ([]byte, []int) { return fileDescriptor18, []int{0} }

func (m *TensorDescription) GetDtype() DataType {
	if m != nil {
		return m.Dtype
	}
	return DataType_DT_INVALID
}

func (m *TensorDescription) GetShape() *TensorShapeProto {
	if m != nil {
		return m.Shape
	}
	return nil
}

func (m *TensorDescription) GetAllocationDescription() *AllocationDescription {
	if m != nil {
		return m.AllocationDescription
	}
	return nil
}

func init() {
	proto.RegisterType((*TensorDescription)(nil), "tensorflow.TensorDescription")
}

func init() {
	proto.RegisterFile("tensorflow/core/framework/tensor_description.proto", fileDescriptor18)
}

var fileDescriptor18 = []byte{
	// 274 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x4d, 0x4b, 0xc3, 0x30,
	0x18, 0xc7, 0x89, 0x38, 0x0f, 0x11, 0x04, 0x8b, 0x2f, 0x65, 0x78, 0x98, 0x82, 0x30, 0xc4, 0x35,
	0x50, 0xc1, 0xbb, 0xa3, 0x1f, 0x60, 0xd4, 0x1d, 0xc4, 0x8b, 0x64, 0x6d, 0x96, 0x95, 0x75, 0x7d,
	0xc2, 0x93, 0xe8, 0xe8, 0xc5, 0x6f, 0xe8, 0xf7, 0xf1, 0x28, 0x49, 0x74, 0x2d, 0x6c, 0x75, 0xd7,
	0xfc, 0x5f, 0x9e, 0x5f, 0xfe, 0x34, 0x36, 0xa2, 0xd2, 0x80, 0xf3, 0x12, 0xd6, 0x2c, 0x03, 0x14,
	0x6c, 0x8e, 0x7c, 0x25, 0xd6, 0x80, 0x4b, 0xe6, 0x95, 0xb7, 0x5c, 0xe8, 0x0c, 0x0b, 0x65, 0x0a,
	0xa8, 0x22, 0x85, 0x60, 0x20, 0xa0, 0x4d, 0xa6, 0x7f, 0xfb, 0x4f, 0xbe, 0x56, 0x42, 0xfb, 0x48,
	0xff, 0x7e, 0xef, 0x19, 0xbd, 0xe0, 0x4a, 0xfc, 0xba, 0x1f, 0xbb, 0xdd, 0xbc, 0x2c, 0x21, 0xe3,
	0x16, 0x66, 0x1b, 0xec, 0xe6, 0x8b, 0xd0, 0xd3, 0xa9, 0x8b, 0x26, 0x8d, 0x16, 0xdc, 0xd1, 0x5e,
	0x6e, 0x59, 0x42, 0x32, 0x20, 0xc3, 0x93, 0xf8, 0x2c, 0x6a, 0xda, 0xa3, 0x84, 0x1b, 0x3e, 0xad,
	0x95, 0x48, 0xbd, 0x25, 0x88, 0x69, 0xcf, 0x81, 0x84, 0x07, 0x03, 0x32, 0x3c, 0x8e, 0xaf, 0xda,
	0x5e, 0xdf, 0xfc, 0x6c, 0xe5, 0x89, 0x3d, 0x97, 0x7a, 0x6b, 0xf0, 0x42, 0x2f, 0x76, 0x53, 0x85,
	0x87, 0xae, 0xe4, 0xba, 0x5d, 0xf2, 0xb4, 0x71, 0xb6, 0x10, 0xd3, 0x73, 0xbe, 0xeb, 0x79, 0xfc,
	0x49, 0x43, 0x40, 0xd9, 0x8e, 0x6f, 0x86, 0x18, 0x5f, 0x6e, 0x7d, 0xd4, 0x41, 0xe9, 0x09, 0x79,
	0x4d, 0x64, 0x61, 0x16, 0xef, 0xb3, 0x28, 0x83, 0x15, 0xcb, 0x51, 0xd4, 0x7f, 0x1b, 0xdb, 0x86,
	0x91, 0x16, 0xf8, 0x51, 0x54, 0x72, 0x24, 0x81, 0xa9, 0xa5, 0x64, 0x9d, 0x4b, 0x7f, 0x13, 0x32,
	0x3b, 0x72, 0xb3, 0x3e, 0xfc, 0x04, 0x00, 0x00, 0xff, 0xff, 0x0d, 0xec, 0xdb, 0x61, 0x25, 0x02,
	0x00, 0x00,
}
