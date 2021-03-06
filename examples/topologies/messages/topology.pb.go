// Code generated by protoc-gen-go. DO NOT EDIT.
// source: messages/topology.proto

package messages

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

type BasicMessage struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BasicMessage) Reset()         { *m = BasicMessage{} }
func (m *BasicMessage) String() string { return proto.CompactTextString(m) }
func (*BasicMessage) ProtoMessage()    {}
func (*BasicMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_topology_d5992b0cbc0cf3b0, []int{0}
}
func (m *BasicMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BasicMessage.Unmarshal(m, b)
}
func (m *BasicMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BasicMessage.Marshal(b, m, deterministic)
}
func (dst *BasicMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BasicMessage.Merge(dst, src)
}
func (m *BasicMessage) XXX_Size() int {
	return xxx_messageInfo_BasicMessage.Size(m)
}
func (m *BasicMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_BasicMessage.DiscardUnknown(m)
}

var xxx_messageInfo_BasicMessage proto.InternalMessageInfo

func (m *BasicMessage) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*BasicMessage)(nil), "messages.BasicMessage")
}

func init() { proto.RegisterFile("messages/topology.proto", fileDescriptor_topology_d5992b0cbc0cf3b0) }

var fileDescriptor_topology_d5992b0cbc0cf3b0 = []byte{
	// 84 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcf, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0x4f, 0x2d, 0xd6, 0x2f, 0xc9, 0x2f, 0xc8, 0xcf, 0xc9, 0x4f, 0xaf, 0xd4, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0x49, 0x28, 0x69, 0x70, 0xf1, 0x38, 0x25, 0x16, 0x67, 0x26,
	0xfb, 0x42, 0x04, 0x84, 0x24, 0xb8, 0xd8, 0xa1, 0x72, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41,
	0x30, 0x6e, 0x12, 0x1b, 0x58, 0xab, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x87, 0xae, 0x48, 0x0a,
	0x55, 0x00, 0x00, 0x00,
}
