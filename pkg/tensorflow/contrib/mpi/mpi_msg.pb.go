// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/contrib/mpi/mpi_msg.proto

/*
Package mpi is a generated protocol buffer package.

It is generated from these files:
	tensorflow/contrib/mpi/mpi_msg.proto

It has these top-level messages:
	MPIRecvTensorResponse
*/
package mpi

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import tensorflow21 "github.com/dreyk/tensorflow-serving-go/pkg/tensorflow/core/protobuf"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type MPIRecvTensorResponse struct {
	Response   *tensorflow21.RecvTensorResponse `protobuf:"bytes,1,opt,name=response" json:"response,omitempty"`
	SingleSend bool                             `protobuf:"varint,2,opt,name=singleSend" json:"singleSend,omitempty"`
	Key        string                           `protobuf:"bytes,3,opt,name=key" json:"key,omitempty"`
	StepId     int64                            `protobuf:"varint,4,opt,name=step_id,json=stepId" json:"step_id,omitempty"`
	Checksum   uint64                           `protobuf:"varint,5,opt,name=checksum" json:"checksum,omitempty"`
}

func (m *MPIRecvTensorResponse) Reset()                    { *m = MPIRecvTensorResponse{} }
func (m *MPIRecvTensorResponse) String() string            { return proto.CompactTextString(m) }
func (*MPIRecvTensorResponse) ProtoMessage()               {}
func (*MPIRecvTensorResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *MPIRecvTensorResponse) GetResponse() *tensorflow21.RecvTensorResponse {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *MPIRecvTensorResponse) GetSingleSend() bool {
	if m != nil {
		return m.SingleSend
	}
	return false
}

func (m *MPIRecvTensorResponse) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *MPIRecvTensorResponse) GetStepId() int64 {
	if m != nil {
		return m.StepId
	}
	return 0
}

func (m *MPIRecvTensorResponse) GetChecksum() uint64 {
	if m != nil {
		return m.Checksum
	}
	return 0
}

func init() {
	proto.RegisterType((*MPIRecvTensorResponse)(nil), "tensorflow.MPIRecvTensorResponse")
}

func init() { proto.RegisterFile("tensorflow/contrib/mpi/mpi_msg.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 259 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x41, 0x4b, 0x03, 0x31,
	0x10, 0x85, 0x89, 0x5b, 0xeb, 0x1a, 0x2f, 0x12, 0x10, 0x97, 0x1e, 0x4a, 0x10, 0x85, 0x5c, 0xba,
	0x01, 0xbd, 0x79, 0xd3, 0x83, 0xd0, 0x83, 0x20, 0xd1, 0x93, 0x97, 0xc2, 0x66, 0xa7, 0x69, 0x48,
	0x37, 0x09, 0x49, 0xb6, 0xa5, 0x7f, 0xcd, 0x5f, 0xe6, 0x51, 0x76, 0x95, 0xae, 0x62, 0x0f, 0x03,
	0xef, 0x0d, 0xef, 0x1b, 0x86, 0x87, 0xaf, 0x13, 0xd8, 0xe8, 0xc2, 0x72, 0xed, 0xb6, 0x5c, 0x3a,
	0x9b, 0x82, 0xae, 0x78, 0xe3, 0x75, 0x37, 0x8b, 0x26, 0xaa, 0xd2, 0x07, 0x97, 0x1c, 0xc1, 0x43,
	0x6a, 0x72, 0xf3, 0x87, 0x08, 0xc0, 0xfb, 0x48, 0xd5, 0x2e, 0xf9, 0xd6, 0x05, 0x03, 0xe1, 0x1b,
	0xb9, 0xfa, 0x40, 0xf8, 0xe2, 0xf9, 0x65, 0x2e, 0x40, 0x6e, 0xde, 0x7a, 0x40, 0x40, 0xf4, 0xce,
	0x46, 0x20, 0xf7, 0x38, 0x0f, 0x3f, 0xba, 0x40, 0x14, 0xb1, 0xb3, 0xdb, 0x69, 0x39, 0xdc, 0x2c,
	0xff, 0x13, 0x62, 0x9f, 0x27, 0x53, 0x8c, 0xa3, 0xb6, 0x6a, 0x0d, 0xaf, 0x60, 0xeb, 0xe2, 0x88,
	0x22, 0x96, 0x8b, 0x5f, 0x1b, 0x72, 0x8e, 0x33, 0x03, 0xbb, 0x22, 0xa3, 0x88, 0x9d, 0x8a, 0x4e,
	0x92, 0x4b, 0x7c, 0x12, 0x13, 0xf8, 0x85, 0xae, 0x8b, 0x11, 0x45, 0x2c, 0x13, 0xe3, 0xce, 0xce,
	0x6b, 0x32, 0xc1, 0xb9, 0x5c, 0x81, 0x34, 0xb1, 0x6d, 0x8a, 0x63, 0x8a, 0xd8, 0x48, 0xec, 0xfd,
	0xe3, 0xd3, 0xfb, 0x83, 0xd2, 0x69, 0xd5, 0x56, 0xa5, 0x74, 0x0d, 0xaf, 0x03, 0xec, 0x0c, 0x1f,
	0x5e, 0x9c, 0x45, 0x08, 0x1b, 0x6d, 0xd5, 0x4c, 0x39, 0xee, 0x8d, 0xe2, 0x87, 0x2b, 0xfc, 0x44,
	0xa8, 0x1a, 0xf7, 0x5d, 0xdc, 0x7d, 0x05, 0x00, 0x00, 0xff, 0xff, 0xdd, 0x0e, 0x32, 0x26, 0x66,
	0x01, 0x00, 0x00,
}
