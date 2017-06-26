// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/contrib/verbs/verbs_service.proto

/*
Package verbs is a generated protocol buffer package.

It is generated from these files:
	tensorflow/contrib/verbs/verbs_service.proto

It has these top-level messages:
	Channel
	MemoryRegion
	GetRemoteAddressRequest
	GetRemoteAddressResponse
*/
package verbs

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Channel struct {
	Lid int32  `protobuf:"varint,1,opt,name=lid" json:"lid,omitempty"`
	Qpn int32  `protobuf:"varint,2,opt,name=qpn" json:"qpn,omitempty"`
	Psn int32  `protobuf:"varint,3,opt,name=psn" json:"psn,omitempty"`
	Snp uint64 `protobuf:"varint,4,opt,name=snp" json:"snp,omitempty"`
	Iid uint64 `protobuf:"varint,5,opt,name=iid" json:"iid,omitempty"`
}

func (m *Channel) Reset()                    { *m = Channel{} }
func (m *Channel) String() string            { return proto.CompactTextString(m) }
func (*Channel) ProtoMessage()               {}
func (*Channel) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Channel) GetLid() int32 {
	if m != nil {
		return m.Lid
	}
	return 0
}

func (m *Channel) GetQpn() int32 {
	if m != nil {
		return m.Qpn
	}
	return 0
}

func (m *Channel) GetPsn() int32 {
	if m != nil {
		return m.Psn
	}
	return 0
}

func (m *Channel) GetSnp() uint64 {
	if m != nil {
		return m.Snp
	}
	return 0
}

func (m *Channel) GetIid() uint64 {
	if m != nil {
		return m.Iid
	}
	return 0
}

type MemoryRegion struct {
	RemoteAddr uint64 `protobuf:"varint,1,opt,name=remote_addr,json=remoteAddr" json:"remote_addr,omitempty"`
	Rkey       uint32 `protobuf:"varint,2,opt,name=rkey" json:"rkey,omitempty"`
}

func (m *MemoryRegion) Reset()                    { *m = MemoryRegion{} }
func (m *MemoryRegion) String() string            { return proto.CompactTextString(m) }
func (*MemoryRegion) ProtoMessage()               {}
func (*MemoryRegion) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *MemoryRegion) GetRemoteAddr() uint64 {
	if m != nil {
		return m.RemoteAddr
	}
	return 0
}

func (m *MemoryRegion) GetRkey() uint32 {
	if m != nil {
		return m.Rkey
	}
	return 0
}

type GetRemoteAddressRequest struct {
	HostName string          `protobuf:"bytes,1,opt,name=host_name,json=hostName" json:"host_name,omitempty"`
	Channel  *Channel        `protobuf:"bytes,2,opt,name=channel" json:"channel,omitempty"`
	Mr       []*MemoryRegion `protobuf:"bytes,3,rep,name=mr" json:"mr,omitempty"`
}

func (m *GetRemoteAddressRequest) Reset()                    { *m = GetRemoteAddressRequest{} }
func (m *GetRemoteAddressRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRemoteAddressRequest) ProtoMessage()               {}
func (*GetRemoteAddressRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetRemoteAddressRequest) GetHostName() string {
	if m != nil {
		return m.HostName
	}
	return ""
}

func (m *GetRemoteAddressRequest) GetChannel() *Channel {
	if m != nil {
		return m.Channel
	}
	return nil
}

func (m *GetRemoteAddressRequest) GetMr() []*MemoryRegion {
	if m != nil {
		return m.Mr
	}
	return nil
}

type GetRemoteAddressResponse struct {
	HostName string          `protobuf:"bytes,1,opt,name=host_name,json=hostName" json:"host_name,omitempty"`
	Channel  *Channel        `protobuf:"bytes,2,opt,name=channel" json:"channel,omitempty"`
	Mr       []*MemoryRegion `protobuf:"bytes,3,rep,name=mr" json:"mr,omitempty"`
}

func (m *GetRemoteAddressResponse) Reset()                    { *m = GetRemoteAddressResponse{} }
func (m *GetRemoteAddressResponse) String() string            { return proto.CompactTextString(m) }
func (*GetRemoteAddressResponse) ProtoMessage()               {}
func (*GetRemoteAddressResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GetRemoteAddressResponse) GetHostName() string {
	if m != nil {
		return m.HostName
	}
	return ""
}

func (m *GetRemoteAddressResponse) GetChannel() *Channel {
	if m != nil {
		return m.Channel
	}
	return nil
}

func (m *GetRemoteAddressResponse) GetMr() []*MemoryRegion {
	if m != nil {
		return m.Mr
	}
	return nil
}

func init() {
	proto.RegisterType((*Channel)(nil), "tensorflow.Channel")
	proto.RegisterType((*MemoryRegion)(nil), "tensorflow.MemoryRegion")
	proto.RegisterType((*GetRemoteAddressRequest)(nil), "tensorflow.GetRemoteAddressRequest")
	proto.RegisterType((*GetRemoteAddressResponse)(nil), "tensorflow.GetRemoteAddressResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for VerbsService service

type VerbsServiceClient interface {
	GetRemoteAddress(ctx context.Context, in *GetRemoteAddressRequest, opts ...grpc.CallOption) (*GetRemoteAddressResponse, error)
}

type verbsServiceClient struct {
	cc *grpc.ClientConn
}

func NewVerbsServiceClient(cc *grpc.ClientConn) VerbsServiceClient {
	return &verbsServiceClient{cc}
}

func (c *verbsServiceClient) GetRemoteAddress(ctx context.Context, in *GetRemoteAddressRequest, opts ...grpc.CallOption) (*GetRemoteAddressResponse, error) {
	out := new(GetRemoteAddressResponse)
	err := grpc.Invoke(ctx, "/tensorflow.VerbsService/GetRemoteAddress", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for VerbsService service

type VerbsServiceServer interface {
	GetRemoteAddress(context.Context, *GetRemoteAddressRequest) (*GetRemoteAddressResponse, error)
}

func RegisterVerbsServiceServer(s *grpc.Server, srv VerbsServiceServer) {
	s.RegisterService(&_VerbsService_serviceDesc, srv)
}

func _VerbsService_GetRemoteAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRemoteAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VerbsServiceServer).GetRemoteAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tensorflow.VerbsService/GetRemoteAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VerbsServiceServer).GetRemoteAddress(ctx, req.(*GetRemoteAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _VerbsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tensorflow.VerbsService",
	HandlerType: (*VerbsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRemoteAddress",
			Handler:    _VerbsService_GetRemoteAddress_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tensorflow/contrib/verbs/verbs_service.proto",
}

func init() { proto.RegisterFile("tensorflow/contrib/verbs/verbs_service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 378 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x92, 0x4f, 0x8b, 0x13, 0x31,
	0x18, 0xc6, 0x99, 0xce, 0xac, 0xeb, 0xbe, 0x5d, 0x61, 0x89, 0x07, 0x83, 0x0a, 0x96, 0xd1, 0xc3,
	0x1c, 0xec, 0x0c, 0xd4, 0x4f, 0xe0, 0xf6, 0xe0, 0x49, 0x59, 0x22, 0x78, 0x10, 0xa4, 0xcc, 0x4c,
	0x5e, 0xa7, 0xa1, 0xcd, 0x9f, 0x26, 0xd9, 0x95, 0xf9, 0x14, 0xfa, 0x91, 0x25, 0xc9, 0x6a, 0x07,
	0xa5, 0x5e, 0xbd, 0x84, 0x97, 0x1f, 0x0f, 0x6f, 0x9e, 0x27, 0x79, 0xe0, 0xb5, 0x47, 0xe5, 0xb4,
	0xfd, 0xba, 0xd7, 0xdf, 0x9a, 0x5e, 0x2b, 0x6f, 0x45, 0xd7, 0xdc, 0xa1, 0xed, 0x5c, 0x3a, 0x37,
	0x0e, 0xed, 0x9d, 0xe8, 0xb1, 0x36, 0x56, 0x7b, 0x4d, 0xe0, 0xa8, 0x2e, 0x7b, 0x38, 0x5f, 0x6f,
	0x5b, 0xa5, 0x70, 0x4f, 0xae, 0x20, 0xdf, 0x0b, 0x4e, 0xb3, 0x45, 0x56, 0x9d, 0xb1, 0x30, 0x06,
	0x72, 0x30, 0x8a, 0xce, 0x12, 0x39, 0x18, 0x15, 0x88, 0x71, 0x8a, 0xe6, 0x89, 0x18, 0x17, 0x89,
	0x53, 0x86, 0x16, 0x8b, 0xac, 0x2a, 0x58, 0x18, 0x03, 0x11, 0x82, 0xd3, 0xb3, 0x44, 0x84, 0xe0,
	0xe5, 0x1a, 0x2e, 0xdf, 0xa3, 0xd4, 0x76, 0x64, 0x38, 0x08, 0xad, 0xc8, 0x0b, 0x98, 0x5b, 0x94,
	0xda, 0xe3, 0xa6, 0xe5, 0xdc, 0xc6, 0x1b, 0x0b, 0x06, 0x09, 0xbd, 0xe5, 0xdc, 0x12, 0x02, 0x85,
	0xdd, 0xe1, 0x18, 0x6f, 0x7e, 0xc4, 0xe2, 0x5c, 0x7e, 0xcf, 0xe0, 0xc9, 0x3b, 0xf4, 0xec, 0xb7,
	0x0a, 0x9d, 0x63, 0x78, 0xb8, 0x45, 0xe7, 0xc9, 0x33, 0xb8, 0xd8, 0x6a, 0xe7, 0x37, 0xaa, 0x95,
	0x18, 0xd7, 0x5d, 0xb0, 0x87, 0x01, 0x7c, 0x68, 0x25, 0x92, 0x25, 0x9c, 0xf7, 0x29, 0x62, 0xdc,
	0x37, 0x5f, 0x3d, 0xae, 0x8f, 0x0f, 0x50, 0xdf, 0xa7, 0x67, 0xbf, 0x34, 0xa4, 0x82, 0x99, 0xb4,
	0x34, 0x5f, 0xe4, 0xd5, 0x7c, 0x45, 0xa7, 0xca, 0x69, 0x04, 0x36, 0x93, 0xb6, 0xfc, 0x91, 0x01,
	0xfd, 0xdb, 0x91, 0x33, 0x5a, 0x39, 0xfc, 0x3f, 0x96, 0x56, 0x12, 0x2e, 0x3f, 0x85, 0x1f, 0xff,
	0x98, 0x3e, 0x9c, 0x7c, 0x81, 0xab, 0x3f, 0x1d, 0x92, 0x97, 0xd3, 0x0d, 0x27, 0x5e, 0xf4, 0xe9,
	0xab, 0x7f, 0x8b, 0x52, 0xc8, 0xeb, 0x11, 0x9e, 0x6b, 0x3b, 0x4c, 0xa5, 0xf7, 0xed, 0xab, 0x63,
	0xef, 0xae, 0xc9, 0xd4, 0xcc, 0x4d, 0x28, 0x9f, 0xbb, 0xc9, 0x3e, 0xaf, 0x07, 0xe1, 0xb7, 0xb7,
	0x5d, 0xdd, 0x6b, 0xd9, 0x70, 0x8b, 0xe3, 0xae, 0x39, 0x2e, 0x58, 0xc6, 0xaa, 0xaa, 0x61, 0x39,
	0xe8, 0xc6, 0xec, 0x86, 0xe6, 0x54, 0xb1, 0xbb, 0x07, 0xb1, 0xcb, 0x6f, 0x7e, 0x06, 0x00, 0x00,
	0xff, 0xff, 0xb9, 0x3a, 0x39, 0x72, 0xfb, 0x02, 0x00, 0x00,
}