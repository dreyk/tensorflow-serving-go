// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow_serving/apis/predict.proto

package apis

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import tensorflow5 "github.com/dreyk/tensorflow-serving-go/pkg/tensorflow/core/framework"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// PredictRequest specifies which TensorFlow model to run, as well as
// how inputs are mapped to tensors and how outputs are filtered before
// returning to user.
type PredictRequest struct {
	// Model Specification.
	ModelSpec *ModelSpec `protobuf:"bytes,1,opt,name=model_spec,json=modelSpec" json:"model_spec,omitempty"`
	// Input tensors.
	// Names of input tensor are alias names. The mapping from aliases to real
	// input tensor names is expected to be stored as named generic signature
	// under the key "inputs" in the model export.
	// Each alias listed in a generic signature named "inputs" should be provided
	// exactly once in order to run the prediction.
	Inputs map[string]*tensorflow5.TensorProto `protobuf:"bytes,2,rep,name=inputs" json:"inputs,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Output filter.
	// Names specified are alias names. The mapping from aliases to real output
	// tensor names is expected to be stored as named generic signature under
	// the key "outputs" in the model export.
	// Only tensors specified here will be run/fetched and returned, with the
	// exception that when none is specified, all tensors specified in the
	// named signature will be run/fetched and returned.
	OutputFilter []string `protobuf:"bytes,3,rep,name=output_filter,json=outputFilter" json:"output_filter,omitempty"`
}

func (m *PredictRequest) Reset()                    { *m = PredictRequest{} }
func (m *PredictRequest) String() string            { return proto.CompactTextString(m) }
func (*PredictRequest) ProtoMessage()               {}
func (*PredictRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *PredictRequest) GetModelSpec() *ModelSpec {
	if m != nil {
		return m.ModelSpec
	}
	return nil
}

func (m *PredictRequest) GetInputs() map[string]*tensorflow5.TensorProto {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *PredictRequest) GetOutputFilter() []string {
	if m != nil {
		return m.OutputFilter
	}
	return nil
}

// Response for PredictRequest on successful run.
type PredictResponse struct {
	// Output tensors.
	Outputs map[string]*tensorflow5.TensorProto `protobuf:"bytes,1,rep,name=outputs" json:"outputs,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *PredictResponse) Reset()                    { *m = PredictResponse{} }
func (m *PredictResponse) String() string            { return proto.CompactTextString(m) }
func (*PredictResponse) ProtoMessage()               {}
func (*PredictResponse) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

func (m *PredictResponse) GetOutputs() map[string]*tensorflow5.TensorProto {
	if m != nil {
		return m.Outputs
	}
	return nil
}

func init() {
	proto.RegisterType((*PredictRequest)(nil), "tensorflow.serving.PredictRequest")
	proto.RegisterType((*PredictResponse)(nil), "tensorflow.serving.PredictResponse")
}

func init() { proto.RegisterFile("tensorflow_serving/apis/predict.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 352 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xcf, 0x4a, 0xf3, 0x40,
	0x14, 0xc5, 0x99, 0x84, 0xaf, 0x1f, 0x9d, 0xf6, 0xfb, 0xc3, 0x6c, 0x0c, 0x01, 0xa1, 0xb4, 0x28,
	0xdd, 0x74, 0x46, 0xea, 0x46, 0xc4, 0x55, 0xc1, 0x8a, 0x82, 0x58, 0xa6, 0xae, 0xdc, 0x94, 0x36,
	0xbd, 0x8d, 0xa1, 0x49, 0x66, 0x9c, 0x99, 0xb4, 0xf4, 0x29, 0x7c, 0x17, 0x9f, 0xce, 0xa5, 0x64,
	0xa6, 0xb5, 0x15, 0xad, 0x2b, 0x77, 0x97, 0x3b, 0xe7, 0x97, 0x73, 0x4e, 0xb8, 0xf8, 0xc8, 0x40,
	0xae, 0x85, 0x9a, 0xa5, 0x62, 0x39, 0xd2, 0xa0, 0x16, 0x49, 0x1e, 0xb3, 0xb1, 0x4c, 0x34, 0x93,
	0x0a, 0xa6, 0x49, 0x64, 0xa8, 0x54, 0xc2, 0x08, 0x42, 0xb6, 0x32, 0xba, 0x96, 0x85, 0xc7, 0xdb,
	0x1d, 0x8b, 0x84, 0x02, 0x36, 0x53, 0xe3, 0x0c, 0x96, 0x42, 0xcd, 0x99, 0x7b, 0x71, 0x6c, 0xd8,
	0xda, 0x67, 0x91, 0x89, 0x29, 0xa4, 0x4e, 0xd4, 0x7c, 0xf6, 0xf0, 0xdf, 0x81, 0xb3, 0xe4, 0xf0,
	0x54, 0x80, 0x36, 0xe4, 0x02, 0x63, 0xab, 0x18, 0x69, 0x09, 0x51, 0x80, 0x1a, 0xa8, 0x5d, 0xeb,
	0x1e, 0xd2, 0xcf, 0x41, 0xe8, 0x6d, 0xa9, 0x1a, 0x4a, 0x88, 0x78, 0x35, 0xdb, 0x8c, 0xa4, 0x8f,
	0x2b, 0x49, 0x2e, 0x0b, 0xa3, 0x03, 0xaf, 0xe1, 0xb7, 0x6b, 0x5d, 0xfa, 0x15, 0xf9, 0xd1, 0x91,
	0x5e, 0x5b, 0xe0, 0x32, 0x37, 0x6a, 0xc5, 0xd7, 0x34, 0x69, 0xe1, 0x3f, 0xa2, 0x30, 0xb2, 0x30,
	0xa3, 0x59, 0x92, 0x1a, 0x50, 0x81, 0xdf, 0xf0, 0xdb, 0x55, 0x5e, 0x77, 0xcb, 0xbe, 0xdd, 0x85,
	0x1c, 0xd7, 0x76, 0x58, 0xf2, 0x1f, 0xfb, 0x73, 0x58, 0xd9, 0xc8, 0x55, 0x5e, 0x8e, 0xa4, 0x83,
	0x7f, 0x2d, 0xc6, 0x69, 0x01, 0x81, 0x67, 0x6b, 0x1c, 0xec, 0x86, 0xb9, 0xb7, 0xe3, 0xa0, 0xfc,
	0x0d, 0xdc, 0xa9, 0xce, 0xbd, 0x33, 0xd4, 0x7c, 0x41, 0xf8, 0xdf, 0x7b, 0x3e, 0x2d, 0x45, 0xae,
	0x81, 0xdc, 0xe0, 0xdf, 0xce, 0x57, 0x07, 0xc8, 0xb6, 0x3a, 0xf9, 0xb6, 0x95, 0xa3, 0xe8, 0x9d,
	0x43, 0x5c, 0xaf, 0xcd, 0x07, 0xc2, 0x21, 0xae, 0xef, 0x3e, 0xfc, 0x48, 0xe8, 0xde, 0xd5, 0x43,
	0x2f, 0x4e, 0xcc, 0x63, 0x31, 0xa1, 0x91, 0xc8, 0xd8, 0x54, 0xc1, 0x6a, 0x73, 0x0c, 0x25, 0xd5,
	0x59, 0x27, 0xec, 0xc4, 0x82, 0xc9, 0x79, 0xcc, 0xf6, 0x1c, 0xc6, 0x2b, 0x42, 0x93, 0x8a, 0x3d,
	0x8b, 0xd3, 0xb7, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0b, 0x2a, 0x53, 0xea, 0xa0, 0x02, 0x00, 0x00,
}
