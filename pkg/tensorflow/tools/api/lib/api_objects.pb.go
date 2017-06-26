// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/tools/api/lib/api_objects.proto

/*
Package lib is a generated protocol buffer package.

It is generated from these files:
	tensorflow/tools/api/lib/api_objects.proto

It has these top-level messages:
	TFAPIMember
	TFAPIMethod
	TFAPIModule
	TFAPIClass
	TFAPIObject
*/
package lib

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

type TFAPIMember struct {
	Name             *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Mtype            *string `protobuf:"bytes,2,opt,name=mtype" json:"mtype,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TFAPIMember) Reset()                    { *m = TFAPIMember{} }
func (m *TFAPIMember) String() string            { return proto.CompactTextString(m) }
func (*TFAPIMember) ProtoMessage()               {}
func (*TFAPIMember) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TFAPIMember) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *TFAPIMember) GetMtype() string {
	if m != nil && m.Mtype != nil {
		return *m.Mtype
	}
	return ""
}

type TFAPIMethod struct {
	Name             *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Path             *string `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
	Argspec          *string `protobuf:"bytes,3,opt,name=argspec" json:"argspec,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TFAPIMethod) Reset()                    { *m = TFAPIMethod{} }
func (m *TFAPIMethod) String() string            { return proto.CompactTextString(m) }
func (*TFAPIMethod) ProtoMessage()               {}
func (*TFAPIMethod) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TFAPIMethod) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *TFAPIMethod) GetPath() string {
	if m != nil && m.Path != nil {
		return *m.Path
	}
	return ""
}

func (m *TFAPIMethod) GetArgspec() string {
	if m != nil && m.Argspec != nil {
		return *m.Argspec
	}
	return ""
}

type TFAPIModule struct {
	Member           []*TFAPIMember `protobuf:"bytes,1,rep,name=member" json:"member,omitempty"`
	MemberMethod     []*TFAPIMethod `protobuf:"bytes,2,rep,name=member_method,json=memberMethod" json:"member_method,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *TFAPIModule) Reset()                    { *m = TFAPIModule{} }
func (m *TFAPIModule) String() string            { return proto.CompactTextString(m) }
func (*TFAPIModule) ProtoMessage()               {}
func (*TFAPIModule) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *TFAPIModule) GetMember() []*TFAPIMember {
	if m != nil {
		return m.Member
	}
	return nil
}

func (m *TFAPIModule) GetMemberMethod() []*TFAPIMethod {
	if m != nil {
		return m.MemberMethod
	}
	return nil
}

type TFAPIClass struct {
	IsInstance       []string       `protobuf:"bytes,1,rep,name=is_instance,json=isInstance" json:"is_instance,omitempty"`
	Member           []*TFAPIMember `protobuf:"bytes,2,rep,name=member" json:"member,omitempty"`
	MemberMethod     []*TFAPIMethod `protobuf:"bytes,3,rep,name=member_method,json=memberMethod" json:"member_method,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *TFAPIClass) Reset()                    { *m = TFAPIClass{} }
func (m *TFAPIClass) String() string            { return proto.CompactTextString(m) }
func (*TFAPIClass) ProtoMessage()               {}
func (*TFAPIClass) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *TFAPIClass) GetIsInstance() []string {
	if m != nil {
		return m.IsInstance
	}
	return nil
}

func (m *TFAPIClass) GetMember() []*TFAPIMember {
	if m != nil {
		return m.Member
	}
	return nil
}

func (m *TFAPIClass) GetMemberMethod() []*TFAPIMethod {
	if m != nil {
		return m.MemberMethod
	}
	return nil
}

type TFAPIObject struct {
	Path             *string      `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
	TfModule         *TFAPIModule `protobuf:"bytes,2,opt,name=tf_module,json=tfModule" json:"tf_module,omitempty"`
	TfClass          *TFAPIClass  `protobuf:"bytes,3,opt,name=tf_class,json=tfClass" json:"tf_class,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *TFAPIObject) Reset()                    { *m = TFAPIObject{} }
func (m *TFAPIObject) String() string            { return proto.CompactTextString(m) }
func (*TFAPIObject) ProtoMessage()               {}
func (*TFAPIObject) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *TFAPIObject) GetPath() string {
	if m != nil && m.Path != nil {
		return *m.Path
	}
	return ""
}

func (m *TFAPIObject) GetTfModule() *TFAPIModule {
	if m != nil {
		return m.TfModule
	}
	return nil
}

func (m *TFAPIObject) GetTfClass() *TFAPIClass {
	if m != nil {
		return m.TfClass
	}
	return nil
}

func init() {
	proto.RegisterType((*TFAPIMember)(nil), "third_party.tensorflow.tools.api.TFAPIMember")
	proto.RegisterType((*TFAPIMethod)(nil), "third_party.tensorflow.tools.api.TFAPIMethod")
	proto.RegisterType((*TFAPIModule)(nil), "third_party.tensorflow.tools.api.TFAPIModule")
	proto.RegisterType((*TFAPIClass)(nil), "third_party.tensorflow.tools.api.TFAPIClass")
	proto.RegisterType((*TFAPIObject)(nil), "third_party.tensorflow.tools.api.TFAPIObject")
}

func init() { proto.RegisterFile("tensorflow/tools/api/lib/api_objects.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 363 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xcf, 0x4e, 0xbb, 0x40,
	0x10, 0xc7, 0x43, 0xdb, 0xdf, 0xaf, 0xed, 0xa2, 0x97, 0x8d, 0x07, 0x6e, 0x36, 0x9c, 0x1a, 0x63,
	0x21, 0xe9, 0xc5, 0xb3, 0x36, 0xd5, 0xd4, 0xc4, 0xd4, 0x10, 0x4f, 0x5e, 0xc8, 0x02, 0x0b, 0xac,
	0x05, 0x76, 0xb3, 0x3b, 0xd5, 0xf4, 0x8d, 0x7c, 0x02, 0x9f, 0xc1, 0xc7, 0x32, 0x1d, 0xe8, 0x9f,
	0x83, 0x46, 0x1b, 0x3d, 0x31, 0x33, 0xe4, 0x33, 0xcc, 0x77, 0xbe, 0x03, 0x39, 0x03, 0x5e, 0x19,
	0xa9, 0xd3, 0x42, 0xbe, 0xf8, 0x20, 0x65, 0x61, 0x7c, 0xa6, 0x84, 0x5f, 0x88, 0x68, 0xfd, 0x0c,
	0x65, 0xf4, 0xc4, 0x63, 0x30, 0x9e, 0xd2, 0x12, 0x24, 0x1d, 0x40, 0x2e, 0x74, 0x12, 0x2a, 0xa6,
	0x61, 0xe5, 0xed, 0x38, 0x0f, 0x39, 0x8f, 0x29, 0xe1, 0x5e, 0x10, 0xfb, 0xe1, 0xfa, 0xf2, 0x7e,
	0x76, 0xc7, 0xcb, 0x88, 0x6b, 0x4a, 0x49, 0xa7, 0x62, 0x25, 0x77, 0xac, 0x81, 0x35, 0xec, 0x07,
	0x18, 0xd3, 0x13, 0xf2, 0xaf, 0x84, 0x95, 0xe2, 0x4e, 0x0b, 0x8b, 0x75, 0xe2, 0xce, 0xb7, 0x20,
	0xe4, 0x32, 0xf9, 0x14, 0xa4, 0xa4, 0xa3, 0x18, 0xe4, 0x0d, 0x87, 0x31, 0x75, 0x48, 0x97, 0xe9,
	0xcc, 0x28, 0x1e, 0x3b, 0x6d, 0x2c, 0x6f, 0x52, 0xf7, 0xd5, 0xda, 0x74, 0x94, 0xc9, 0xb2, 0xe0,
	0x74, 0x4a, 0xfe, 0x97, 0x38, 0x94, 0x63, 0x0d, 0xda, 0x43, 0x7b, 0x3c, 0xf2, 0xbe, 0x13, 0xe3,
	0xed, 0x29, 0x09, 0x1a, 0x98, 0x06, 0xe4, 0xb8, 0x8e, 0xc2, 0x12, 0x27, 0x75, 0x5a, 0x07, 0x76,
	0x5b, 0x43, 0xc1, 0x51, 0xdd, 0xa3, 0xce, 0xdc, 0x77, 0x8b, 0x10, 0x7c, 0x3b, 0x29, 0x98, 0x31,
	0xf4, 0x94, 0xd8, 0xc2, 0x84, 0xa2, 0x32, 0xc0, 0xaa, 0x98, 0xe3, 0xb8, 0xfd, 0x80, 0x08, 0x33,
	0x6b, 0x2a, 0x7b, 0x52, 0x5a, 0x7f, 0x2a, 0xa5, 0xfd, 0x7b, 0x29, 0x6f, 0x9b, 0xad, 0xcf, 0xf1,
	0x70, 0xb6, 0x9e, 0x59, 0x7b, 0x9e, 0xdd, 0x92, 0x3e, 0xa4, 0x61, 0x89, 0xb6, 0xa0, 0x99, 0x07,
	0x7c, 0x13, 0xa1, 0xa0, 0x07, 0x69, 0xe3, 0xea, 0x0d, 0xe9, 0x41, 0x1a, 0xc6, 0xeb, 0xbd, 0xe1,
	0x01, 0xd8, 0xe3, 0xf3, 0x1f, 0xb6, 0xc2, 0x5d, 0x07, 0x5d, 0x48, 0x31, 0xb8, 0x9a, 0x3e, 0x4e,
	0x32, 0x01, 0xf9, 0x32, 0xf2, 0x62, 0x59, 0xfa, 0x89, 0xe6, 0xab, 0x85, 0xbf, 0x83, 0x47, 0x86,
	0xeb, 0x67, 0x51, 0x65, 0xa3, 0x4c, 0xfa, 0x6a, 0x91, 0xf9, 0x5f, 0xfd, 0x33, 0x1f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x29, 0x03, 0x4b, 0xa2, 0x4e, 0x03, 0x00, 0x00,
}
