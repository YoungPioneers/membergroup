// Code generated by protoc-gen-go.
// source: message.proto
// DO NOT EDIT!

/*
Package membergroup is a generated protocol buffer package.

It is generated from these files:
	message.proto

It has these top-level messages:
	Node
	Message
*/
package membergroup

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

type MessageType int32

const (
	MessageType_PingMessage        MessageType = 0
	MessageType_PingReqMessage     MessageType = 1
	MessageType_JoinMessage        MessageType = 2
	MessageType_LeftMessage        MessageType = 3
	MessageType_FullMembersMessage MessageType = 4
)

var MessageType_name = map[int32]string{
	0: "PingMessage",
	1: "PingReqMessage",
	2: "JoinMessage",
	3: "LeftMessage",
	4: "FullMembersMessage",
}
var MessageType_value = map[string]int32{
	"PingMessage":        0,
	"PingReqMessage":     1,
	"JoinMessage":        2,
	"LeftMessage":        3,
	"FullMembersMessage": 4,
}

func (x MessageType) String() string {
	return proto.EnumName(MessageType_name, int32(x))
}
func (MessageType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type NodeStatus int32

const (
	NodeStatus_MemberAlive   NodeStatus = 0
	NodeStatus_MemberSuspect NodeStatus = 1
	NodeStatus_MemberDead    NodeStatus = 2
)

var NodeStatus_name = map[int32]string{
	0: "MemberAlive",
	1: "MemberSuspect",
	2: "MemberDead",
}
var NodeStatus_value = map[string]int32{
	"MemberAlive":   0,
	"MemberSuspect": 1,
	"MemberDead":    2,
}

func (x NodeStatus) String() string {
	return proto.EnumName(NodeStatus_name, int32(x))
}
func (NodeStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Node struct {
	Uid    uint64     `protobuf:"varint,1,opt,name=uid" json:"uid,omitempty"`
	Ip     uint64     `protobuf:"varint,2,opt,name=ip" json:"ip,omitempty"`
	Port   uint32     `protobuf:"varint,3,opt,name=port" json:"port,omitempty"`
	Status NodeStatus `protobuf:"varint,4,opt,name=status,enum=membergroup.NodeStatus" json:"status,omitempty"`
}

func (m *Node) Reset()                    { *m = Node{} }
func (m *Node) String() string            { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()               {}
func (*Node) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Message struct {
	MessageType MessageType `protobuf:"varint,1,opt,name=messageType,enum=membergroup.MessageType" json:"messageType,omitempty"`
	Source      []*Node     `protobuf:"bytes,2,rep,name=source" json:"source,omitempty"`
	Target      *Node       `protobuf:"bytes,3,opt,name=target" json:"target,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Message) GetSource() []*Node {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *Message) GetTarget() *Node {
	if m != nil {
		return m.Target
	}
	return nil
}

func init() {
	proto.RegisterType((*Node)(nil), "membergroup.Node")
	proto.RegisterType((*Message)(nil), "membergroup.Message")
	proto.RegisterEnum("membergroup.MessageType", MessageType_name, MessageType_value)
	proto.RegisterEnum("membergroup.NodeStatus", NodeStatus_name, NodeStatus_value)
}

func init() { proto.RegisterFile("message.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 291 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x91, 0xcb, 0x4a, 0xc3, 0x40,
	0x14, 0x86, 0xcd, 0x85, 0x0a, 0x27, 0x34, 0xb6, 0xb3, 0xd0, 0x2c, 0xa5, 0xab, 0xda, 0x45, 0x84,
	0xb8, 0x73, 0xa5, 0x20, 0x2e, 0xc4, 0x88, 0x4c, 0x7d, 0x81, 0xb4, 0x39, 0x86, 0x40, 0xd2, 0x19,
	0xe7, 0x22, 0xf8, 0x2c, 0xbe, 0xac, 0x73, 0x49, 0x68, 0x44, 0x77, 0x87, 0x6f, 0xbe, 0x9c, 0xff,
	0xe7, 0x04, 0xe6, 0x3d, 0x4a, 0x59, 0x35, 0x98, 0x73, 0xc1, 0x14, 0x23, 0x49, 0x8f, 0xfd, 0x0e,
	0x45, 0x23, 0x98, 0xe6, 0xab, 0x1e, 0xe2, 0x17, 0x56, 0x23, 0x59, 0x40, 0xa4, 0xdb, 0x3a, 0x0b,
	0x2e, 0x83, 0x75, 0x4c, 0xed, 0x48, 0x52, 0x08, 0x5b, 0x9e, 0x85, 0x0e, 0x98, 0x89, 0x10, 0x88,
	0x39, 0x13, 0x2a, 0x8b, 0x0c, 0x99, 0x53, 0x37, 0x93, 0x6b, 0x98, 0x49, 0x55, 0x29, 0x2d, 0xb3,
	0xd8, 0xd0, 0xb4, 0xb8, 0xc8, 0x27, 0xbb, 0x73, 0xbb, 0x78, 0xeb, 0x9e, 0xe9, 0xa0, 0xad, 0xbe,
	0x03, 0x38, 0x2d, 0x7d, 0x1b, 0x72, 0x0b, 0xc9, 0x50, 0xec, 0xed, 0x8b, 0xa3, 0x8b, 0x4e, 0x8b,
	0xec, 0xd7, 0x86, 0xf2, 0xf8, 0x4e, 0xa7, 0x32, 0xb9, 0x32, 0xc1, 0x4c, 0x8b, 0x3d, 0x9a, 0x82,
	0xd1, 0x3a, 0x29, 0x96, 0x7f, 0x82, 0xe9, 0x20, 0x58, 0x55, 0x55, 0xa2, 0x41, 0xdf, 0xfc, 0x7f,
	0xd5, 0x0b, 0x9b, 0x0e, 0x92, 0x49, 0x22, 0x39, 0x83, 0xe4, 0xb5, 0x3d, 0x34, 0x03, 0x5a, 0x9c,
	0x98, 0x13, 0xa4, 0x16, 0x50, 0xfc, 0x18, 0x59, 0x60, 0xa5, 0x27, 0xd6, 0x1e, 0x46, 0x10, 0x5a,
	0xf0, 0x8c, 0xef, 0x6a, 0x04, 0x11, 0x39, 0x07, 0xf2, 0xa8, 0xbb, 0xae, 0x74, 0xa9, 0x72, 0xe4,
	0xf1, 0xe6, 0x0e, 0xe0, 0x78, 0x21, 0xfb, 0x99, 0x37, 0xee, 0xbb, 0xf6, 0xd3, 0x86, 0x2d, 0x61,
	0xee, 0xc1, 0x56, 0x4b, 0x8e, 0x7b, 0x65, 0xb2, 0x52, 0x00, 0x8f, 0x1e, 0xb0, 0xaa, 0x17, 0xe1,
	0x6e, 0xe6, 0x7e, 0xe8, 0xcd, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xea, 0xb2, 0xf1, 0x4d, 0xe1,
	0x01, 0x00, 0x00,
}
