// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/framework/device_attributes.proto

package framework

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type DeviceLocality struct {
	// Optional bus locality of device.  Default value of 0 means
	// no specific locality.  Specific localities are indexed from 1.
	BusId int32 `protobuf:"varint,1,opt,name=bus_id,json=busId" json:"bus_id,omitempty"`
}

func (m *DeviceLocality) Reset()                    { *m = DeviceLocality{} }
func (m *DeviceLocality) String() string            { return proto.CompactTextString(m) }
func (*DeviceLocality) ProtoMessage()               {}
func (*DeviceLocality) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *DeviceLocality) GetBusId() int32 {
	if m != nil {
		return m.BusId
	}
	return 0
}

type DeviceAttributes struct {
	// Fully specified name of the device within a cluster.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// String representation of device_type.
	DeviceType string `protobuf:"bytes,2,opt,name=device_type,json=deviceType" json:"device_type,omitempty"`
	// Memory capacity of device in bytes.
	MemoryLimit int64 `protobuf:"varint,4,opt,name=memory_limit,json=memoryLimit" json:"memory_limit,omitempty"`
	// Platform-specific data about device that may be useful
	// for supporting efficient data transfers.
	Locality *DeviceLocality `protobuf:"bytes,5,opt,name=locality" json:"locality,omitempty"`
	// A device is assigned a global unique number each time it is
	// initialized. "incarnation" should never be 0.
	Incarnation uint64 `protobuf:"fixed64,6,opt,name=incarnation" json:"incarnation,omitempty"`
	// String representation of the physical device that this device maps to.
	PhysicalDeviceDesc string `protobuf:"bytes,7,opt,name=physical_device_desc,json=physicalDeviceDesc" json:"physical_device_desc,omitempty"`
}

func (m *DeviceAttributes) Reset()                    { *m = DeviceAttributes{} }
func (m *DeviceAttributes) String() string            { return proto.CompactTextString(m) }
func (*DeviceAttributes) ProtoMessage()               {}
func (*DeviceAttributes) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *DeviceAttributes) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DeviceAttributes) GetDeviceType() string {
	if m != nil {
		return m.DeviceType
	}
	return ""
}

func (m *DeviceAttributes) GetMemoryLimit() int64 {
	if m != nil {
		return m.MemoryLimit
	}
	return 0
}

func (m *DeviceAttributes) GetLocality() *DeviceLocality {
	if m != nil {
		return m.Locality
	}
	return nil
}

func (m *DeviceAttributes) GetIncarnation() uint64 {
	if m != nil {
		return m.Incarnation
	}
	return 0
}

func (m *DeviceAttributes) GetPhysicalDeviceDesc() string {
	if m != nil {
		return m.PhysicalDeviceDesc
	}
	return ""
}

func init() {
	proto.RegisterType((*DeviceLocality)(nil), "tensorflow.DeviceLocality")
	proto.RegisterType((*DeviceAttributes)(nil), "tensorflow.DeviceAttributes")
}

func init() { proto.RegisterFile("tensorflow/core/framework/device_attributes.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 322 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x59, 0x6d, 0xa3, 0x4e, 0x44, 0x64, 0x51, 0x09, 0x5e, 0x8c, 0xbd, 0x98, 0x4b, 0xb3,
	0xfe, 0x01, 0xef, 0x96, 0x5e, 0x84, 0x1e, 0x4a, 0xf0, 0xe4, 0x25, 0x6c, 0x36, 0xd3, 0x74, 0x69,
	0x92, 0x0d, 0xbb, 0x9b, 0x96, 0x1c, 0xfc, 0xda, 0xe2, 0x51, 0x9a, 0xd4, 0xb6, 0x0a, 0xde, 0x96,
	0x37, 0x6f, 0xe7, 0xfd, 0x78, 0x03, 0x0f, 0x16, 0x4b, 0xa3, 0xf4, 0x2c, 0x57, 0x2b, 0x26, 0x94,
	0x46, 0x36, 0xd3, 0xbc, 0xc0, 0x95, 0xd2, 0x0b, 0x96, 0xe2, 0x52, 0x0a, 0x8c, 0xb9, 0xb5, 0x5a,
	0x26, 0xb5, 0x45, 0x13, 0x56, 0x5a, 0x59, 0x45, 0x61, 0xf7, 0x65, 0x70, 0x07, 0x67, 0xe3, 0xd6,
	0x36, 0x51, 0x82, 0xe7, 0xd2, 0x36, 0xf4, 0x12, 0x9c, 0xa4, 0x36, 0xb1, 0x4c, 0x3d, 0xe2, 0x93,
	0xa0, 0x1f, 0xf5, 0x93, 0xda, 0xbc, 0xa6, 0x83, 0x4f, 0x02, 0xe7, 0x9d, 0xf3, 0x65, 0xbb, 0x8f,
	0x52, 0xe8, 0x95, 0xbc, 0xc0, 0xd6, 0x79, 0x12, 0xb5, 0x6f, 0x7a, 0x03, 0xee, 0x26, 0xd8, 0x36,
	0x15, 0x7a, 0x07, 0xed, 0x08, 0x3a, 0xe9, 0xad, 0xa9, 0x90, 0xde, 0xc2, 0x69, 0x81, 0x85, 0xd2,
	0x4d, 0x9c, 0xcb, 0x42, 0x5a, 0xaf, 0xe7, 0x93, 0xe0, 0x30, 0x72, 0x3b, 0x6d, 0xb2, 0x96, 0xe8,
	0x33, 0x1c, 0xe7, 0x1b, 0x1e, 0xaf, 0xef, 0x93, 0xc0, 0x7d, 0xbc, 0x0e, 0x77, 0xd0, 0xe1, 0x6f,
	0xe2, 0x68, 0xeb, 0xa5, 0x3e, 0xb8, 0xb2, 0x14, 0x5c, 0x97, 0xdc, 0x4a, 0x55, 0x7a, 0x8e, 0x4f,
	0x02, 0x27, 0xda, 0x97, 0xe8, 0x3d, 0x5c, 0x54, 0xf3, 0xc6, 0x48, 0xc1, 0xf3, 0x78, 0x83, 0x99,
	0xa2, 0x11, 0xde, 0x51, 0x8b, 0x49, 0x7f, 0x66, 0x5d, 0xc2, 0x18, 0x8d, 0x18, 0x7d, 0x80, 0xa7,
	0x74, 0xb6, 0x1f, 0xbf, 0x6d, 0x78, 0x74, 0xf5, 0xb7, 0x91, 0xe9, 0xba, 0x60, 0x33, 0x25, 0xef,
	0xe3, 0x4c, 0xda, 0x79, 0x9d, 0x84, 0x42, 0x15, 0x2c, 0xd5, 0xd8, 0x2c, 0xd8, 0x6e, 0xc1, 0xd0,
	0xa0, 0x5e, 0xca, 0x32, 0x1b, 0x66, 0x8a, 0x55, 0x8b, 0x8c, 0xfd, 0x7b, 0xc1, 0x2f, 0x42, 0x12,
	0xa7, 0xbd, 0xd9, 0xd3, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xcd, 0x76, 0x03, 0xe3, 0xe8, 0x01,
	0x00, 0x00,
}