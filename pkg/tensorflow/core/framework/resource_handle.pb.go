// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/framework/resource_handle.proto

package framework

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Protocol buffer representing a handle to a tensorflow resource. Handles are
// not valid across executions, but can be serialized back and forth from within
// a single run.
type ResourceHandle struct {
	// Unique name for the device containing the resource.
	Device string `protobuf:"bytes,1,opt,name=device" json:"device,omitempty"`
	// Container in which this resource is placed.
	Container string `protobuf:"bytes,2,opt,name=container" json:"container,omitempty"`
	// Unique name of this resource.
	Name string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	// Hash code for the type of the resource. Is only valid in the same device
	// and in the same execution.
	HashCode uint64 `protobuf:"varint,4,opt,name=hash_code,json=hashCode" json:"hash_code,omitempty"`
	// For debug-only, the name of the type pointed to by this handle, if
	// available.
	MaybeTypeName string `protobuf:"bytes,5,opt,name=maybe_type_name,json=maybeTypeName" json:"maybe_type_name,omitempty"`
}

func (m *ResourceHandle) Reset()                    { *m = ResourceHandle{} }
func (m *ResourceHandle) String() string            { return proto.CompactTextString(m) }
func (*ResourceHandle) ProtoMessage()               {}
func (*ResourceHandle) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{0} }

func (m *ResourceHandle) GetDevice() string {
	if m != nil {
		return m.Device
	}
	return ""
}

func (m *ResourceHandle) GetContainer() string {
	if m != nil {
		return m.Container
	}
	return ""
}

func (m *ResourceHandle) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ResourceHandle) GetHashCode() uint64 {
	if m != nil {
		return m.HashCode
	}
	return 0
}

func (m *ResourceHandle) GetMaybeTypeName() string {
	if m != nil {
		return m.MaybeTypeName
	}
	return ""
}

func init() {
	proto.RegisterType((*ResourceHandle)(nil), "tensorflow.ResourceHandle")
}

func init() { proto.RegisterFile("tensorflow/core/framework/resource_handle.proto", fileDescriptor14) }

var fileDescriptor14 = []byte{
	// 256 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x50, 0x4d, 0x4b, 0x03, 0x31,
	0x10, 0x25, 0x5a, 0x8b, 0x1b, 0x50, 0x21, 0x82, 0x04, 0xf4, 0x50, 0x3c, 0x48, 0x2f, 0xdd, 0x1c,
	0xfc, 0x07, 0xd5, 0x83, 0x27, 0x29, 0x8b, 0x27, 0x2f, 0x4b, 0x36, 0x99, 0x66, 0x97, 0x6d, 0x32,
	0xcb, 0x6c, 0xda, 0xb2, 0xfe, 0x19, 0xff, 0xa6, 0x47, 0x69, 0x28, 0xae, 0x1e, 0xbc, 0xcd, 0xbc,
	0x0f, 0x78, 0xef, 0x71, 0x15, 0x21, 0xf4, 0x48, 0xeb, 0x0d, 0xee, 0x95, 0x41, 0x02, 0xb5, 0x26,
	0xed, 0x61, 0x8f, 0xd4, 0x2a, 0x82, 0x1e, 0xb7, 0x64, 0xa0, 0xac, 0x75, 0xb0, 0x1b, 0xc8, 0x3b,
	0xc2, 0x88, 0x82, 0x8f, 0x86, 0xfb, 0x4f, 0xc6, 0x2f, 0x8b, 0xa3, 0xea, 0x25, 0x89, 0xc4, 0x0d,
	0x9f, 0x5a, 0xd8, 0x35, 0x06, 0x24, 0x9b, 0xb1, 0x79, 0x56, 0x1c, 0x3f, 0x71, 0xc7, 0x33, 0x83,
	0x21, 0xea, 0x26, 0x00, 0xc9, 0x93, 0x44, 0x8d, 0x80, 0x10, 0x7c, 0x12, 0xb4, 0x07, 0x79, 0x9a,
	0x88, 0x74, 0x8b, 0x5b, 0x9e, 0xd5, 0xba, 0xaf, 0x4b, 0x83, 0x16, 0xe4, 0x64, 0xc6, 0xe6, 0x93,
	0xe2, 0xfc, 0x00, 0x3c, 0xa1, 0x05, 0xf1, 0xc0, 0xaf, 0xbc, 0x1e, 0x2a, 0x28, 0xe3, 0xd0, 0x41,
	0x99, 0xbc, 0x67, 0xc9, 0x7b, 0x91, 0xe0, 0xb7, 0xa1, 0x83, 0x57, 0xed, 0x61, 0xf9, 0xc1, 0x25,
	0x92, 0xcb, 0xc7, 0xcc, 0xf9, 0x4f, 0xbf, 0xe5, 0xf5, 0xdf, 0xe8, 0xab, 0x43, 0xbd, 0x15, 0x7b,
	0x7f, 0x76, 0x4d, 0xac, 0xb7, 0x55, 0x6e, 0xd0, 0x2b, 0x4b, 0x30, 0xb4, 0xbf, 0x26, 0x5a, 0xf4,
	0x40, 0xbb, 0x26, 0xb8, 0x85, 0x43, 0xd5, 0xb5, 0xee, 0xff, 0xf1, 0xbe, 0x18, 0xab, 0xa6, 0x69,
	0xb0, 0xc7, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4c, 0xd3, 0x46, 0xa7, 0x63, 0x01, 0x00, 0x00,
}