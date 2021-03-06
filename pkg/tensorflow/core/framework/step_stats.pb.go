// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/framework/step_stats.proto

package framework

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type AllocatorMemoryUsed struct {
	AllocatorName string `protobuf:"bytes,1,opt,name=allocator_name,json=allocatorName" json:"allocator_name,omitempty"`
	// These are per-node allocator memory stats.
	TotalBytes int64 `protobuf:"varint,2,opt,name=total_bytes,json=totalBytes" json:"total_bytes,omitempty"`
	PeakBytes  int64 `protobuf:"varint,3,opt,name=peak_bytes,json=peakBytes" json:"peak_bytes,omitempty"`
	// The bytes that are not deallocated.
	LiveBytes int64 `protobuf:"varint,4,opt,name=live_bytes,json=liveBytes" json:"live_bytes,omitempty"`
	// These are snapshots of the overall allocator memory stats.
	// The number of live bytes currently allocated by the allocator.
	AllocatorBytesInUse int64 `protobuf:"varint,5,opt,name=allocator_bytes_in_use,json=allocatorBytesInUse" json:"allocator_bytes_in_use,omitempty"`
}

func (m *AllocatorMemoryUsed) Reset()                    { *m = AllocatorMemoryUsed{} }
func (m *AllocatorMemoryUsed) String() string            { return proto.CompactTextString(m) }
func (*AllocatorMemoryUsed) ProtoMessage()               {}
func (*AllocatorMemoryUsed) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{0} }

func (m *AllocatorMemoryUsed) GetAllocatorName() string {
	if m != nil {
		return m.AllocatorName
	}
	return ""
}

func (m *AllocatorMemoryUsed) GetTotalBytes() int64 {
	if m != nil {
		return m.TotalBytes
	}
	return 0
}

func (m *AllocatorMemoryUsed) GetPeakBytes() int64 {
	if m != nil {
		return m.PeakBytes
	}
	return 0
}

func (m *AllocatorMemoryUsed) GetLiveBytes() int64 {
	if m != nil {
		return m.LiveBytes
	}
	return 0
}

func (m *AllocatorMemoryUsed) GetAllocatorBytesInUse() int64 {
	if m != nil {
		return m.AllocatorBytesInUse
	}
	return 0
}

// Output sizes recorded for a single execution of a graph node.
type NodeOutput struct {
	Slot              int32              `protobuf:"varint,1,opt,name=slot" json:"slot,omitempty"`
	TensorDescription *TensorDescription `protobuf:"bytes,3,opt,name=tensor_description,json=tensorDescription" json:"tensor_description,omitempty"`
}

func (m *NodeOutput) Reset()                    { *m = NodeOutput{} }
func (m *NodeOutput) String() string            { return proto.CompactTextString(m) }
func (*NodeOutput) ProtoMessage()               {}
func (*NodeOutput) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{1} }

func (m *NodeOutput) GetSlot() int32 {
	if m != nil {
		return m.Slot
	}
	return 0
}

func (m *NodeOutput) GetTensorDescription() *TensorDescription {
	if m != nil {
		return m.TensorDescription
	}
	return nil
}

// For memory tracking.
type MemoryStats struct {
	HostTempMemorySize             int64   `protobuf:"varint,1,opt,name=host_temp_memory_size,json=hostTempMemorySize" json:"host_temp_memory_size,omitempty"`
	DeviceTempMemorySize           int64   `protobuf:"varint,2,opt,name=device_temp_memory_size,json=deviceTempMemorySize" json:"device_temp_memory_size,omitempty"`
	HostPersistentMemorySize       int64   `protobuf:"varint,3,opt,name=host_persistent_memory_size,json=hostPersistentMemorySize" json:"host_persistent_memory_size,omitempty"`
	DevicePersistentMemorySize     int64   `protobuf:"varint,4,opt,name=device_persistent_memory_size,json=devicePersistentMemorySize" json:"device_persistent_memory_size,omitempty"`
	HostPersistentTensorAllocIds   []int64 `protobuf:"varint,5,rep,packed,name=host_persistent_tensor_alloc_ids,json=hostPersistentTensorAllocIds" json:"host_persistent_tensor_alloc_ids,omitempty"`
	DevicePersistentTensorAllocIds []int64 `protobuf:"varint,6,rep,packed,name=device_persistent_tensor_alloc_ids,json=devicePersistentTensorAllocIds" json:"device_persistent_tensor_alloc_ids,omitempty"`
}

func (m *MemoryStats) Reset()                    { *m = MemoryStats{} }
func (m *MemoryStats) String() string            { return proto.CompactTextString(m) }
func (*MemoryStats) ProtoMessage()               {}
func (*MemoryStats) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{2} }

func (m *MemoryStats) GetHostTempMemorySize() int64 {
	if m != nil {
		return m.HostTempMemorySize
	}
	return 0
}

func (m *MemoryStats) GetDeviceTempMemorySize() int64 {
	if m != nil {
		return m.DeviceTempMemorySize
	}
	return 0
}

func (m *MemoryStats) GetHostPersistentMemorySize() int64 {
	if m != nil {
		return m.HostPersistentMemorySize
	}
	return 0
}

func (m *MemoryStats) GetDevicePersistentMemorySize() int64 {
	if m != nil {
		return m.DevicePersistentMemorySize
	}
	return 0
}

func (m *MemoryStats) GetHostPersistentTensorAllocIds() []int64 {
	if m != nil {
		return m.HostPersistentTensorAllocIds
	}
	return nil
}

func (m *MemoryStats) GetDevicePersistentTensorAllocIds() []int64 {
	if m != nil {
		return m.DevicePersistentTensorAllocIds
	}
	return nil
}

// Time/size stats recorded for a single execution of a graph node.
type NodeExecStats struct {
	// TODO(tucker): Use some more compact form of node identity than
	// the full string name.  Either all processes should agree on a
	// global id (cost_id?) for each node, or we should use a hash of
	// the name.
	NodeName         string                   `protobuf:"bytes,1,opt,name=node_name,json=nodeName" json:"node_name,omitempty"`
	AllStartMicros   int64                    `protobuf:"varint,2,opt,name=all_start_micros,json=allStartMicros" json:"all_start_micros,omitempty"`
	OpStartRelMicros int64                    `protobuf:"varint,3,opt,name=op_start_rel_micros,json=opStartRelMicros" json:"op_start_rel_micros,omitempty"`
	OpEndRelMicros   int64                    `protobuf:"varint,4,opt,name=op_end_rel_micros,json=opEndRelMicros" json:"op_end_rel_micros,omitempty"`
	AllEndRelMicros  int64                    `protobuf:"varint,5,opt,name=all_end_rel_micros,json=allEndRelMicros" json:"all_end_rel_micros,omitempty"`
	Memory           []*AllocatorMemoryUsed   `protobuf:"bytes,6,rep,name=memory" json:"memory,omitempty"`
	Output           []*NodeOutput            `protobuf:"bytes,7,rep,name=output" json:"output,omitempty"`
	TimelineLabel    string                   `protobuf:"bytes,8,opt,name=timeline_label,json=timelineLabel" json:"timeline_label,omitempty"`
	ScheduledMicros  int64                    `protobuf:"varint,9,opt,name=scheduled_micros,json=scheduledMicros" json:"scheduled_micros,omitempty"`
	ThreadId         uint32                   `protobuf:"varint,10,opt,name=thread_id,json=threadId" json:"thread_id,omitempty"`
	ReferencedTensor []*AllocationDescription `protobuf:"bytes,11,rep,name=referenced_tensor,json=referencedTensor" json:"referenced_tensor,omitempty"`
	MemoryStats      *MemoryStats             `protobuf:"bytes,12,opt,name=memory_stats,json=memoryStats" json:"memory_stats,omitempty"`
}

func (m *NodeExecStats) Reset()                    { *m = NodeExecStats{} }
func (m *NodeExecStats) String() string            { return proto.CompactTextString(m) }
func (*NodeExecStats) ProtoMessage()               {}
func (*NodeExecStats) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{3} }

func (m *NodeExecStats) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *NodeExecStats) GetAllStartMicros() int64 {
	if m != nil {
		return m.AllStartMicros
	}
	return 0
}

func (m *NodeExecStats) GetOpStartRelMicros() int64 {
	if m != nil {
		return m.OpStartRelMicros
	}
	return 0
}

func (m *NodeExecStats) GetOpEndRelMicros() int64 {
	if m != nil {
		return m.OpEndRelMicros
	}
	return 0
}

func (m *NodeExecStats) GetAllEndRelMicros() int64 {
	if m != nil {
		return m.AllEndRelMicros
	}
	return 0
}

func (m *NodeExecStats) GetMemory() []*AllocatorMemoryUsed {
	if m != nil {
		return m.Memory
	}
	return nil
}

func (m *NodeExecStats) GetOutput() []*NodeOutput {
	if m != nil {
		return m.Output
	}
	return nil
}

func (m *NodeExecStats) GetTimelineLabel() string {
	if m != nil {
		return m.TimelineLabel
	}
	return ""
}

func (m *NodeExecStats) GetScheduledMicros() int64 {
	if m != nil {
		return m.ScheduledMicros
	}
	return 0
}

func (m *NodeExecStats) GetThreadId() uint32 {
	if m != nil {
		return m.ThreadId
	}
	return 0
}

func (m *NodeExecStats) GetReferencedTensor() []*AllocationDescription {
	if m != nil {
		return m.ReferencedTensor
	}
	return nil
}

func (m *NodeExecStats) GetMemoryStats() *MemoryStats {
	if m != nil {
		return m.MemoryStats
	}
	return nil
}

type DeviceStepStats struct {
	Device    string           `protobuf:"bytes,1,opt,name=device" json:"device,omitempty"`
	NodeStats []*NodeExecStats `protobuf:"bytes,2,rep,name=node_stats,json=nodeStats" json:"node_stats,omitempty"`
}

func (m *DeviceStepStats) Reset()                    { *m = DeviceStepStats{} }
func (m *DeviceStepStats) String() string            { return proto.CompactTextString(m) }
func (*DeviceStepStats) ProtoMessage()               {}
func (*DeviceStepStats) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{4} }

func (m *DeviceStepStats) GetDevice() string {
	if m != nil {
		return m.Device
	}
	return ""
}

func (m *DeviceStepStats) GetNodeStats() []*NodeExecStats {
	if m != nil {
		return m.NodeStats
	}
	return nil
}

type StepStats struct {
	DevStats []*DeviceStepStats `protobuf:"bytes,1,rep,name=dev_stats,json=devStats" json:"dev_stats,omitempty"`
}

func (m *StepStats) Reset()                    { *m = StepStats{} }
func (m *StepStats) String() string            { return proto.CompactTextString(m) }
func (*StepStats) ProtoMessage()               {}
func (*StepStats) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{5} }

func (m *StepStats) GetDevStats() []*DeviceStepStats {
	if m != nil {
		return m.DevStats
	}
	return nil
}

func init() {
	proto.RegisterType((*AllocatorMemoryUsed)(nil), "tensorflow.AllocatorMemoryUsed")
	proto.RegisterType((*NodeOutput)(nil), "tensorflow.NodeOutput")
	proto.RegisterType((*MemoryStats)(nil), "tensorflow.MemoryStats")
	proto.RegisterType((*NodeExecStats)(nil), "tensorflow.NodeExecStats")
	proto.RegisterType((*DeviceStepStats)(nil), "tensorflow.DeviceStepStats")
	proto.RegisterType((*StepStats)(nil), "tensorflow.StepStats")
}

func init() { proto.RegisterFile("tensorflow/core/framework/step_stats.proto", fileDescriptor15) }

var fileDescriptor15 = []byte{
	// 792 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x55, 0xd1, 0x4e, 0xeb, 0x46,
	0x10, 0x95, 0x6f, 0x20, 0x4d, 0x26, 0x97, 0x4b, 0x58, 0x5a, 0xae, 0x4b, 0x4a, 0x6f, 0x1a, 0xa9,
	0x52, 0x6e, 0x2b, 0x12, 0x95, 0xab, 0xb6, 0xa8, 0x52, 0x1f, 0x40, 0x50, 0x89, 0x0a, 0x28, 0x32,
	0xf0, 0xd2, 0x17, 0xcb, 0xf1, 0x0e, 0xc9, 0x8a, 0xb5, 0xd7, 0xda, 0xdd, 0x84, 0xc2, 0x3f, 0xf4,
	0xa3, 0xfa, 0x07, 0xfd, 0x9c, 0x3e, 0x56, 0xbb, 0xeb, 0xd8, 0x4e, 0x42, 0xde, 0x36, 0x67, 0xce,
	0xcc, 0x9e, 0xcc, 0x9c, 0x59, 0xc3, 0x77, 0x1a, 0x53, 0x25, 0xe4, 0x03, 0x17, 0x4f, 0xc3, 0x58,
	0x48, 0x1c, 0x3e, 0xc8, 0x28, 0xc1, 0x27, 0x21, 0x1f, 0x87, 0x4a, 0x63, 0x16, 0x2a, 0x1d, 0x69,
	0x35, 0xc8, 0xa4, 0xd0, 0x82, 0x40, 0xc9, 0xdd, 0xff, 0x69, 0x7d, 0x5e, 0xc4, 0xb9, 0x88, 0x23,
	0xcd, 0x44, 0x1a, 0x52, 0x54, 0xb1, 0x64, 0x99, 0x39, 0xbb, 0x1a, 0xfb, 0x47, 0xeb, 0xf3, 0x5c,
	0x64, 0x35, 0xa7, 0xf7, 0xaf, 0x07, 0xbb, 0x27, 0xae, 0xa8, 0x90, 0x57, 0x98, 0x08, 0xf9, 0x7c,
	0xaf, 0x90, 0x92, 0x6f, 0xe1, 0x5d, 0x34, 0x87, 0xc3, 0x34, 0x4a, 0xd0, 0xf7, 0xba, 0x5e, 0xbf,
	0x19, 0x6c, 0x15, 0xe8, 0x75, 0x94, 0x20, 0xf9, 0x00, 0x2d, 0x2d, 0x74, 0xc4, 0xc3, 0xd1, 0xb3,
	0x46, 0xe5, 0xbf, 0xe9, 0x7a, 0xfd, 0x5a, 0x00, 0x16, 0x3a, 0x35, 0x08, 0x39, 0x00, 0xc8, 0x30,
	0x7a, 0xcc, 0xe3, 0x35, 0x1b, 0x6f, 0x1a, 0xa4, 0x08, 0x73, 0x36, 0xc3, 0x3c, 0xbc, 0xe1, 0xc2,
	0x06, 0x71, 0xe1, 0x4f, 0xb0, 0x57, 0xaa, 0xb0, 0x9c, 0x90, 0xa5, 0xe1, 0x54, 0xa1, 0xbf, 0x69,
	0xa9, 0xbb, 0x45, 0xd4, 0xf2, 0x2f, 0xd2, 0x7b, 0x85, 0xbd, 0x14, 0xe0, 0x5a, 0x50, 0xfc, 0x63,
	0xaa, 0xb3, 0xa9, 0x26, 0x04, 0x36, 0x14, 0x17, 0xda, 0xca, 0xdf, 0x0c, 0xec, 0x99, 0x5c, 0x02,
	0x59, 0x6d, 0x88, 0x15, 0xd7, 0x3a, 0x3a, 0x18, 0x94, 0x5d, 0x1c, 0xdc, 0xd9, 0xe3, 0x59, 0x49,
	0x0a, 0x76, 0xf4, 0x32, 0xd4, 0xfb, 0xbb, 0x06, 0x2d, 0xd7, 0xb9, 0x5b, 0x33, 0x50, 0xf2, 0x03,
	0x7c, 0x31, 0x11, 0x4a, 0x87, 0x1a, 0x93, 0x2c, 0x4c, 0x6c, 0x20, 0x54, 0xec, 0xc5, 0x75, 0xb0,
	0x16, 0x10, 0x13, 0xbc, 0xc3, 0x24, 0xcb, 0x73, 0xd8, 0x0b, 0x92, 0x1f, 0xe1, 0x3d, 0xc5, 0x19,
	0x8b, 0x71, 0x35, 0xc9, 0xb5, 0xf4, 0x73, 0x17, 0x5e, 0x4a, 0xfb, 0x15, 0x3a, 0xf6, 0xa6, 0x0c,
	0xa5, 0x62, 0x4a, 0x63, 0xaa, 0x17, 0x52, 0x5d, 0xb7, 0x7d, 0x43, 0xb9, 0x29, 0x18, 0x95, 0xf4,
	0x13, 0x38, 0xc8, 0x6f, 0x5d, 0x53, 0xc0, 0xcd, 0x63, 0xdf, 0x91, 0x5e, 0x2d, 0xf1, 0x1b, 0x74,
	0x97, 0x15, 0xe4, 0x9d, 0xb5, 0x93, 0x09, 0x19, 0x55, 0xfe, 0x66, 0xb7, 0xd6, 0xaf, 0x05, 0x5f,
	0x2d, 0xca, 0x70, 0x9d, 0xb5, 0xce, 0xbb, 0xa0, 0x8a, 0xfc, 0x0e, 0xbd, 0x55, 0x29, 0x2b, 0x95,
	0xea, 0xb6, 0xd2, 0xd7, 0xcb, 0x7a, 0x16, 0x6b, 0xf5, 0xfe, 0xd9, 0x80, 0x2d, 0x63, 0x80, 0xf3,
	0xbf, 0x30, 0x76, 0x13, 0xe9, 0x40, 0x33, 0x15, 0x14, 0xab, 0x3e, 0x6e, 0x18, 0xc0, 0x5a, 0xb8,
	0x0f, 0xed, 0x88, 0x73, 0xb3, 0x8c, 0x52, 0x87, 0x09, 0x8b, 0xa5, 0x98, 0xfb, 0xd8, 0x6c, 0xc0,
	0xad, 0x81, 0xaf, 0x2c, 0x4a, 0x0e, 0x61, 0x57, 0x64, 0x39, 0x51, 0x22, 0x9f, 0x93, 0x5d, 0x9b,
	0xdb, 0x22, 0xb3, 0xdc, 0x00, 0x79, 0x4e, 0xff, 0x08, 0x3b, 0x22, 0x0b, 0x31, 0xa5, 0x55, 0xb2,
	0x6b, 0xe9, 0x3b, 0x91, 0x9d, 0xa7, 0xb4, 0xa4, 0x7e, 0x0f, 0xc4, 0x68, 0x58, 0xe2, 0x3a, 0x8f,
	0x6f, 0x47, 0x9c, 0x2f, 0x90, 0x7f, 0x86, 0xba, 0x1b, 0x92, 0xed, 0x47, 0xeb, 0xe8, 0x43, 0xd5,
	0xb1, 0xaf, 0xec, 0x72, 0x90, 0xd3, 0xc9, 0x00, 0xea, 0xc2, 0x2e, 0x85, 0xff, 0x99, 0x4d, 0xdc,
	0xab, 0x26, 0x96, 0x2b, 0x13, 0xe4, 0x2c, 0xf3, 0x06, 0x68, 0x96, 0x20, 0x67, 0x29, 0x86, 0x3c,
	0x1a, 0x21, 0xf7, 0x1b, 0xee, 0x0d, 0x98, 0xa3, 0x97, 0x06, 0x24, 0x1f, 0xa1, 0xad, 0xe2, 0x09,
	0xd2, 0x29, 0x47, 0x3a, 0x97, 0xde, 0x74, 0xd2, 0x0b, 0x3c, 0x97, 0xde, 0x81, 0xa6, 0x9e, 0x48,
	0x8c, 0x68, 0xc8, 0xa8, 0x0f, 0x5d, 0xaf, 0xbf, 0x15, 0x34, 0x1c, 0x70, 0x41, 0xc9, 0x35, 0xec,
	0x48, 0x7c, 0x40, 0x89, 0x69, 0x8c, 0x34, 0x1f, 0xbe, 0xdf, 0xb2, 0x4a, 0xbf, 0x79, 0xe5, 0x2f,
	0x32, 0x91, 0x56, 0x17, 0xb3, 0x5d, 0xe6, 0x3a, 0x3f, 0x90, 0x5f, 0xe0, 0xed, 0xdc, 0xcc, 0xc6,
	0x05, 0xfe, 0x5b, 0xbb, 0xdf, 0xef, 0xab, 0xa5, 0x2a, 0x6b, 0x1b, 0xb4, 0x92, 0xf2, 0x47, 0x2f,
	0x86, 0xed, 0x33, 0xeb, 0xb2, 0x5b, 0x8d, 0x99, 0x33, 0xd1, 0x1e, 0xd4, 0x9d, 0xf1, 0x72, 0x07,
	0xe5, 0xbf, 0xc8, 0x31, 0x80, 0x35, 0x97, 0xbb, 0xe4, 0x8d, 0xd5, 0xfb, 0xe5, 0x72, 0x67, 0x0b,
	0x2f, 0x06, 0xd6, 0x89, 0xee, 0x92, 0x73, 0x68, 0x96, 0xe5, 0x8f, 0xa1, 0x49, 0x71, 0x96, 0x57,
	0xf1, 0x6c, 0x95, 0x4e, 0xb5, 0xca, 0x92, 0x9c, 0xa0, 0x41, 0x71, 0x66, 0x4f, 0xa7, 0x33, 0xf0,
	0x85, 0x1c, 0x57, 0xb9, 0xc5, 0xbb, 0x7f, 0xba, 0x5d, 0x24, 0xdc, 0x98, 0xe7, 0x5e, 0xdd, 0x78,
	0x7f, 0x9e, 0x8d, 0x99, 0x9e, 0x4c, 0x47, 0x83, 0x58, 0x24, 0x43, 0x2a, 0xf1, 0x79, 0xfe, 0x71,
	0x30, 0x99, 0x87, 0x0a, 0xe5, 0x8c, 0xa5, 0xe3, 0xc3, 0xb1, 0x18, 0x66, 0x8f, 0xe3, 0xe1, 0xda,
	0x0f, 0xca, 0x7f, 0x9e, 0x37, 0xaa, 0xdb, 0x2f, 0xc8, 0xa7, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff,
	0xc6, 0x18, 0x89, 0xae, 0xe7, 0x06, 0x00, 0x00,
}
