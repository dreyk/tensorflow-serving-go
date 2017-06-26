// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/protobuf/saver.proto

package protobuf

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// A version number that identifies a different on-disk checkpoint format.
// Usually, each subclass of BaseSaverBuilder works with a particular
// version/format.  However, it is possible that the same builder may be
// upgraded to support a newer checkpoint format in the future.
type SaverDef_CheckpointFormatVersion int32

const (
	// Internal legacy format.
	SaverDef_LEGACY SaverDef_CheckpointFormatVersion = 0
	// Current format: tf.Saver() which works with tensorflow::table::Table.
	SaverDef_V1 SaverDef_CheckpointFormatVersion = 1
	// Experimental format under development.
	SaverDef_V2 SaverDef_CheckpointFormatVersion = 2
)

var SaverDef_CheckpointFormatVersion_name = map[int32]string{
	0: "LEGACY",
	1: "V1",
	2: "V2",
}
var SaverDef_CheckpointFormatVersion_value = map[string]int32{
	"LEGACY": 0,
	"V1":     1,
	"V2":     2,
}

func (x SaverDef_CheckpointFormatVersion) String() string {
	return proto.EnumName(SaverDef_CheckpointFormatVersion_name, int32(x))
}
func (SaverDef_CheckpointFormatVersion) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor12, []int{0, 0}
}

// Protocol buffer representing the configuration of a Saver.
type SaverDef struct {
	// The name of the tensor in which to specify the filename when saving or
	// restoring a model checkpoint.
	FilenameTensorName string `protobuf:"bytes,1,opt,name=filename_tensor_name,json=filenameTensorName" json:"filename_tensor_name,omitempty"`
	// The operation to run when saving a model checkpoint.
	SaveTensorName string `protobuf:"bytes,2,opt,name=save_tensor_name,json=saveTensorName" json:"save_tensor_name,omitempty"`
	// The operation to run when restoring a model checkpoint.
	RestoreOpName string `protobuf:"bytes,3,opt,name=restore_op_name,json=restoreOpName" json:"restore_op_name,omitempty"`
	// Maximum number of checkpoints to keep.  If 0, no checkpoints are deleted.
	MaxToKeep int32 `protobuf:"varint,4,opt,name=max_to_keep,json=maxToKeep" json:"max_to_keep,omitempty"`
	// Shard the save files, one per device that has Variable nodes.
	Sharded bool `protobuf:"varint,5,opt,name=sharded" json:"sharded,omitempty"`
	// How often to keep an additional checkpoint. If not specified, only the last
	// "max_to_keep" checkpoints are kept; if specified, in addition to keeping
	// the last "max_to_keep" checkpoints, an additional checkpoint will be kept
	// for every n hours of training.
	KeepCheckpointEveryNHours float32                          `protobuf:"fixed32,6,opt,name=keep_checkpoint_every_n_hours,json=keepCheckpointEveryNHours" json:"keep_checkpoint_every_n_hours,omitempty"`
	Version                   SaverDef_CheckpointFormatVersion `protobuf:"varint,7,opt,name=version,enum=tensorflow.SaverDef_CheckpointFormatVersion" json:"version,omitempty"`
}

func (m *SaverDef) Reset()                    { *m = SaverDef{} }
func (m *SaverDef) String() string            { return proto.CompactTextString(m) }
func (*SaverDef) ProtoMessage()               {}
func (*SaverDef) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{0} }

func (m *SaverDef) GetFilenameTensorName() string {
	if m != nil {
		return m.FilenameTensorName
	}
	return ""
}

func (m *SaverDef) GetSaveTensorName() string {
	if m != nil {
		return m.SaveTensorName
	}
	return ""
}

func (m *SaverDef) GetRestoreOpName() string {
	if m != nil {
		return m.RestoreOpName
	}
	return ""
}

func (m *SaverDef) GetMaxToKeep() int32 {
	if m != nil {
		return m.MaxToKeep
	}
	return 0
}

func (m *SaverDef) GetSharded() bool {
	if m != nil {
		return m.Sharded
	}
	return false
}

func (m *SaverDef) GetKeepCheckpointEveryNHours() float32 {
	if m != nil {
		return m.KeepCheckpointEveryNHours
	}
	return 0
}

func (m *SaverDef) GetVersion() SaverDef_CheckpointFormatVersion {
	if m != nil {
		return m.Version
	}
	return SaverDef_LEGACY
}

func init() {
	proto.RegisterType((*SaverDef)(nil), "tensorflow.SaverDef")
	proto.RegisterEnum("tensorflow.SaverDef_CheckpointFormatVersion", SaverDef_CheckpointFormatVersion_name, SaverDef_CheckpointFormatVersion_value)
}

func init() { proto.RegisterFile("tensorflow/core/protobuf/saver.proto", fileDescriptor12) }

var fileDescriptor12 = []byte{
	// 374 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0xd2, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x07, 0x70, 0xd6, 0xa5, 0x4e, 0xbb, 0x15, 0xc5, 0x5a, 0x90, 0x30, 0x07, 0x90, 0x55, 0x21,
	0xe4, 0x03, 0xb5, 0xa1, 0x88, 0x3b, 0x34, 0xb4, 0x20, 0x81, 0x42, 0x64, 0xa2, 0x48, 0x70, 0x59,
	0x39, 0xce, 0xf8, 0x43, 0xfe, 0x18, 0x6b, 0x77, 0x6d, 0x92, 0x47, 0xe0, 0x8d, 0x39, 0xa2, 0xdd,
	0x60, 0x9c, 0x1e, 0x72, 0xb2, 0x67, 0xe7, 0xf7, 0xf7, 0x48, 0xe3, 0xa5, 0x2f, 0x14, 0x34, 0x12,
	0x45, 0x5a, 0xe1, 0xaf, 0x30, 0x41, 0x01, 0x61, 0x2b, 0x50, 0xe1, 0xaa, 0x4b, 0x43, 0x19, 0xf7,
	0x20, 0x02, 0x53, 0x32, 0x3a, 0xaa, 0x8b, 0xdf, 0x47, 0xf4, 0xe4, 0xbb, 0xee, 0x7d, 0x84, 0x94,
	0xbd, 0xa6, 0x8f, 0xd3, 0xa2, 0x82, 0x26, 0xae, 0x81, 0xef, 0x0c, 0xd7, 0xef, 0x2e, 0xf1, 0x88,
	0x7f, 0x1a, 0xb1, 0xa1, 0xb7, 0x30, 0xad, 0x59, 0x5c, 0x03, 0xf3, 0xa9, 0xa3, 0xbf, 0x7c, 0x47,
	0x5b, 0x46, 0x9f, 0xeb, 0xf3, 0x3d, 0xf9, 0x92, 0x3e, 0x14, 0x20, 0x15, 0x0a, 0xe0, 0xd8, 0xee,
	0xe0, 0x91, 0x81, 0x0f, 0xfe, 0x1d, 0x7f, 0x6b, 0x8d, 0x7b, 0x4e, 0xcf, 0xea, 0x78, 0xc3, 0x15,
	0xf2, 0x12, 0xa0, 0x75, 0xef, 0x7b, 0xc4, 0x3f, 0x8e, 0x4e, 0xeb, 0x78, 0xb3, 0xc0, 0x2f, 0x00,
	0x2d, 0x73, 0xe9, 0x44, 0xe6, 0xb1, 0x58, 0xc3, 0xda, 0x3d, 0xf6, 0x88, 0x7f, 0x12, 0x0d, 0x25,
	0x7b, 0x4f, 0x9f, 0xe9, 0x08, 0x4f, 0x72, 0x48, 0xca, 0x16, 0x8b, 0x46, 0x71, 0xe8, 0x41, 0x6c,
	0x79, 0xc3, 0x73, 0xec, 0x84, 0x74, 0x6d, 0x8f, 0xf8, 0x56, 0xf4, 0x54, 0xa3, 0xe9, 0x7f, 0x73,
	0xa3, 0xc9, 0xec, 0xb3, 0x06, 0xec, 0x96, 0x4e, 0x7a, 0x10, 0xb2, 0xc0, 0xc6, 0x9d, 0x78, 0xc4,
	0x3f, 0xbf, 0x7a, 0x15, 0x8c, 0xab, 0x0a, 0x86, 0x35, 0x05, 0x63, 0xf8, 0x16, 0x45, 0x1d, 0xab,
	0xe5, 0x2e, 0x13, 0x0d, 0xe1, 0x8b, 0x77, 0xf4, 0xc9, 0x01, 0xc3, 0x28, 0xb5, 0xbf, 0xde, 0x7c,
	0xfa, 0x30, 0xfd, 0xe1, 0xdc, 0x63, 0x36, 0xb5, 0x96, 0x6f, 0x1c, 0x62, 0x9e, 0x57, 0x8e, 0x75,
	0x5d, 0xd1, 0x47, 0x28, 0xb2, 0xfd, 0x91, 0x9d, 0x2a, 0xaa, 0xeb, 0x33, 0x33, 0x78, 0xae, 0x7f,
	0x9d, 0x9c, 0x93, 0x9f, 0xd3, 0xac, 0x50, 0x79, 0xb7, 0x0a, 0x12, 0xac, 0xc3, 0xb5, 0x80, 0x6d,
	0x19, 0x8e, 0x81, 0x4b, 0x09, 0xa2, 0x2f, 0x9a, 0xec, 0x32, 0xc3, 0xb0, 0x2d, 0xb3, 0xf0, 0xd0,
	0x75, 0xf8, 0x43, 0xc8, 0xca, 0x36, 0xc5, 0xdb, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x4f, 0xd3,
	0x0e, 0xf4, 0x34, 0x02, 0x00, 0x00,
}
