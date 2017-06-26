// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/tools/tfprof/tfprof_log.proto

/*
Package tfprof is a generated protocol buffer package.

It is generated from these files:
	tensorflow/tools/tfprof/tfprof_log.proto
	tensorflow/tools/tfprof/tfprof_options.proto
	tensorflow/tools/tfprof/tfprof_output.proto

It has these top-level messages:
	CodeDef
	OpLogEntry
	OpLog
	OptionsProto
	TFProfTensorProto
	TFGraphNodeProto
	TFMultiGraphNodeProto
*/
package tfprof

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

// It specifies the Python callstack that creates an op.
type CodeDef struct {
	Traces           []*CodeDef_Trace `protobuf:"bytes,1,rep,name=traces" json:"traces,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *CodeDef) Reset()                    { *m = CodeDef{} }
func (m *CodeDef) String() string            { return proto.CompactTextString(m) }
func (*CodeDef) ProtoMessage()               {}
func (*CodeDef) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CodeDef) GetTraces() []*CodeDef_Trace {
	if m != nil {
		return m.Traces
	}
	return nil
}

type CodeDef_Trace struct {
	File             *string `protobuf:"bytes,1,opt,name=file" json:"file,omitempty"`
	Lineno           *int32  `protobuf:"varint,2,opt,name=lineno" json:"lineno,omitempty"`
	Function         *string `protobuf:"bytes,3,opt,name=function" json:"function,omitempty"`
	Line             *string `protobuf:"bytes,4,opt,name=line" json:"line,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CodeDef_Trace) Reset()                    { *m = CodeDef_Trace{} }
func (m *CodeDef_Trace) String() string            { return proto.CompactTextString(m) }
func (*CodeDef_Trace) ProtoMessage()               {}
func (*CodeDef_Trace) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *CodeDef_Trace) GetFile() string {
	if m != nil && m.File != nil {
		return *m.File
	}
	return ""
}

func (m *CodeDef_Trace) GetLineno() int32 {
	if m != nil && m.Lineno != nil {
		return *m.Lineno
	}
	return 0
}

func (m *CodeDef_Trace) GetFunction() string {
	if m != nil && m.Function != nil {
		return *m.Function
	}
	return ""
}

func (m *CodeDef_Trace) GetLine() string {
	if m != nil && m.Line != nil {
		return *m.Line
	}
	return ""
}

type OpLogEntry struct {
	// op name.
	Name *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// float_ops is filled by tfprof Python API when called. It requires the
	// op has RegisterStatistics defined. Currently, Conv2D, MatMul, etc, are
	// implemented.
	FloatOps *int64 `protobuf:"varint,2,opt,name=float_ops,json=floatOps" json:"float_ops,omitempty"`
	// User can define extra op type information for an op. This allows the user
	// to select a group of ops precisely using op_type as a key.
	Types []string `protobuf:"bytes,3,rep,name=types" json:"types,omitempty"`
	// Used to support tfprof "code" view.
	CodeDef          *CodeDef `protobuf:"bytes,4,opt,name=code_def,json=codeDef" json:"code_def,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *OpLogEntry) Reset()                    { *m = OpLogEntry{} }
func (m *OpLogEntry) String() string            { return proto.CompactTextString(m) }
func (*OpLogEntry) ProtoMessage()               {}
func (*OpLogEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *OpLogEntry) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *OpLogEntry) GetFloatOps() int64 {
	if m != nil && m.FloatOps != nil {
		return *m.FloatOps
	}
	return 0
}

func (m *OpLogEntry) GetTypes() []string {
	if m != nil {
		return m.Types
	}
	return nil
}

func (m *OpLogEntry) GetCodeDef() *CodeDef {
	if m != nil {
		return m.CodeDef
	}
	return nil
}

type OpLog struct {
	LogEntries       []*OpLogEntry `protobuf:"bytes,1,rep,name=log_entries,json=logEntries" json:"log_entries,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *OpLog) Reset()                    { *m = OpLog{} }
func (m *OpLog) String() string            { return proto.CompactTextString(m) }
func (*OpLog) ProtoMessage()               {}
func (*OpLog) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *OpLog) GetLogEntries() []*OpLogEntry {
	if m != nil {
		return m.LogEntries
	}
	return nil
}

func init() {
	proto.RegisterType((*CodeDef)(nil), "tensorflow.tfprof.CodeDef")
	proto.RegisterType((*CodeDef_Trace)(nil), "tensorflow.tfprof.CodeDef.Trace")
	proto.RegisterType((*OpLogEntry)(nil), "tensorflow.tfprof.OpLogEntry")
	proto.RegisterType((*OpLog)(nil), "tensorflow.tfprof.OpLog")
}

func init() { proto.RegisterFile("tensorflow/tools/tfprof/tfprof_log.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 324 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xbf, 0x6a, 0xeb, 0x30,
	0x14, 0xc6, 0xf1, 0x75, 0x9c, 0x3f, 0x27, 0xd3, 0x15, 0x97, 0x8b, 0x49, 0x29, 0x98, 0x4c, 0x5e,
	0x62, 0x43, 0xa0, 0xd0, 0xa9, 0x43, 0x9a, 0xd2, 0xa5, 0x10, 0x10, 0x9d, 0xba, 0x98, 0xd4, 0x96,
	0x54, 0x11, 0x45, 0x47, 0x48, 0x4a, 0x4b, 0x5e, 0xa1, 0x4f, 0xd2, 0xc7, 0x2c, 0x96, 0x4d, 0x32,
	0xb4, 0x99, 0xf4, 0x9d, 0xa3, 0x9f, 0xbe, 0xef, 0x48, 0x82, 0xdc, 0x33, 0xed, 0xd0, 0x72, 0x85,
	0x1f, 0xa5, 0x47, 0x54, 0xae, 0xf4, 0xdc, 0x58, 0xe4, 0xfd, 0x52, 0x29, 0x14, 0x85, 0xb1, 0xe8,
	0x91, 0xfc, 0x3d, 0x93, 0x45, 0xb7, 0x39, 0xff, 0x8a, 0x60, 0x74, 0x8f, 0x0d, 0x5b, 0x33, 0x4e,
	0x6e, 0x61, 0xe8, 0xed, 0xb6, 0x66, 0x2e, 0x8d, 0xb2, 0x38, 0x9f, 0x2e, 0xb3, 0xe2, 0x07, 0x5f,
	0xf4, 0x6c, 0xf1, 0xdc, 0x82, 0xb4, 0xe7, 0x67, 0x35, 0x24, 0xa1, 0x41, 0x08, 0x0c, 0xb8, 0x54,
	0x2c, 0x8d, 0xb2, 0x28, 0x9f, 0xd0, 0xa0, 0xc9, 0x7f, 0x18, 0x2a, 0xa9, 0x99, 0xc6, 0xf4, 0x4f,
	0x16, 0xe5, 0x09, 0xed, 0x2b, 0x32, 0x83, 0x31, 0x3f, 0xe8, 0xda, 0x4b, 0xd4, 0x69, 0x1c, 0xf8,
	0x53, 0xdd, 0xfa, 0xb4, 0x54, 0x3a, 0xe8, 0x7c, 0x5a, 0x3d, 0xff, 0x8c, 0x00, 0x36, 0xe6, 0x09,
	0xc5, 0x83, 0xf6, 0xf6, 0xd8, 0x22, 0x7a, 0xbb, 0x3f, 0x45, 0xb5, 0x9a, 0x5c, 0xc1, 0x84, 0x2b,
	0xdc, 0xfa, 0x0a, 0x8d, 0x0b, 0x69, 0x31, 0x1d, 0x87, 0xc6, 0xc6, 0x38, 0xf2, 0x0f, 0x12, 0x7f,
	0x34, 0xcc, 0xa5, 0x71, 0x16, 0xe7, 0x13, 0xda, 0x15, 0xe4, 0x06, 0xc6, 0x35, 0x36, 0xac, 0x6a,
	0x18, 0x0f, 0x69, 0xd3, 0xe5, 0xec, 0xf2, 0xb5, 0xe9, 0xa8, 0xee, 0xc4, 0xfc, 0x11, 0x92, 0x30,
	0x0b, 0xb9, 0x83, 0xa9, 0x42, 0x51, 0x31, 0xed, 0xad, 0x3c, 0xbd, 0xdc, 0xf5, 0x2f, 0x16, 0xe7,
	0xd1, 0x29, 0xa8, 0x4e, 0x49, 0xe6, 0x56, 0xeb, 0x97, 0x95, 0x90, 0xfe, 0xed, 0xf0, 0x5a, 0xd4,
	0xb8, 0x2f, 0x1b, 0xcb, 0x8e, 0xbb, 0xf2, 0x7c, 0x78, 0xe1, 0x98, 0x7d, 0x97, 0x5a, 0x2c, 0x04,
	0x96, 0x66, 0x27, 0xca, 0x0b, 0x5f, 0xfd, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x1c, 0x34, 0x04, 0xc3,
	0x04, 0x02, 0x00, 0x00,
}