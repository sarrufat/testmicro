// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/github.com/sarrufat/testmicro/hello.proto

package github_com_sarrufat_testmicro_hello

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Message struct {
	Say                  string   `protobuf:"bytes,1,opt,name=say,proto3" json:"say,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_59fc3fabf3c5b51c, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetSay() string {
	if m != nil {
		return m.Say
	}
	return ""
}

type Request struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_59fc3fabf3c5b51c, []int{1}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Response struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_59fc3fabf3c5b51c, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type StreamingRequest struct {
	Count                int64    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamingRequest) Reset()         { *m = StreamingRequest{} }
func (m *StreamingRequest) String() string { return proto.CompactTextString(m) }
func (*StreamingRequest) ProtoMessage()    {}
func (*StreamingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_59fc3fabf3c5b51c, []int{3}
}

func (m *StreamingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamingRequest.Unmarshal(m, b)
}
func (m *StreamingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamingRequest.Marshal(b, m, deterministic)
}
func (m *StreamingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamingRequest.Merge(m, src)
}
func (m *StreamingRequest) XXX_Size() int {
	return xxx_messageInfo_StreamingRequest.Size(m)
}
func (m *StreamingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamingRequest proto.InternalMessageInfo

func (m *StreamingRequest) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type StreamingResponse struct {
	Count                int64    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamingResponse) Reset()         { *m = StreamingResponse{} }
func (m *StreamingResponse) String() string { return proto.CompactTextString(m) }
func (*StreamingResponse) ProtoMessage()    {}
func (*StreamingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_59fc3fabf3c5b51c, []int{4}
}

func (m *StreamingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamingResponse.Unmarshal(m, b)
}
func (m *StreamingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamingResponse.Marshal(b, m, deterministic)
}
func (m *StreamingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamingResponse.Merge(m, src)
}
func (m *StreamingResponse) XXX_Size() int {
	return xxx_messageInfo_StreamingResponse.Size(m)
}
func (m *StreamingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamingResponse proto.InternalMessageInfo

func (m *StreamingResponse) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type Ping struct {
	Stroke               int64    `protobuf:"varint,1,opt,name=stroke,proto3" json:"stroke,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ping) Reset()         { *m = Ping{} }
func (m *Ping) String() string { return proto.CompactTextString(m) }
func (*Ping) ProtoMessage()    {}
func (*Ping) Descriptor() ([]byte, []int) {
	return fileDescriptor_59fc3fabf3c5b51c, []int{5}
}

func (m *Ping) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ping.Unmarshal(m, b)
}
func (m *Ping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ping.Marshal(b, m, deterministic)
}
func (m *Ping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ping.Merge(m, src)
}
func (m *Ping) XXX_Size() int {
	return xxx_messageInfo_Ping.Size(m)
}
func (m *Ping) XXX_DiscardUnknown() {
	xxx_messageInfo_Ping.DiscardUnknown(m)
}

var xxx_messageInfo_Ping proto.InternalMessageInfo

func (m *Ping) GetStroke() int64 {
	if m != nil {
		return m.Stroke
	}
	return 0
}

type Pong struct {
	Stroke               int64    `protobuf:"varint,1,opt,name=stroke,proto3" json:"stroke,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pong) Reset()         { *m = Pong{} }
func (m *Pong) String() string { return proto.CompactTextString(m) }
func (*Pong) ProtoMessage()    {}
func (*Pong) Descriptor() ([]byte, []int) {
	return fileDescriptor_59fc3fabf3c5b51c, []int{6}
}

func (m *Pong) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pong.Unmarshal(m, b)
}
func (m *Pong) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pong.Marshal(b, m, deterministic)
}
func (m *Pong) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pong.Merge(m, src)
}
func (m *Pong) XXX_Size() int {
	return xxx_messageInfo_Pong.Size(m)
}
func (m *Pong) XXX_DiscardUnknown() {
	xxx_messageInfo_Pong.DiscardUnknown(m)
}

var xxx_messageInfo_Pong proto.InternalMessageInfo

func (m *Pong) GetStroke() int64 {
	if m != nil {
		return m.Stroke
	}
	return 0
}

func init() {
	proto.RegisterType((*Message)(nil), "github.com.sarrufat.testmicro.hello.Message")
	proto.RegisterType((*Request)(nil), "github.com.sarrufat.testmicro.hello.Request")
	proto.RegisterType((*Response)(nil), "github.com.sarrufat.testmicro.hello.Response")
	proto.RegisterType((*StreamingRequest)(nil), "github.com.sarrufat.testmicro.hello.StreamingRequest")
	proto.RegisterType((*StreamingResponse)(nil), "github.com.sarrufat.testmicro.hello.StreamingResponse")
	proto.RegisterType((*Ping)(nil), "github.com.sarrufat.testmicro.hello.Ping")
	proto.RegisterType((*Pong)(nil), "github.com.sarrufat.testmicro.hello.Pong")
}

func init() {
	proto.RegisterFile("proto/github.com/sarrufat/testmicro/hello.proto", fileDescriptor_59fc3fabf3c5b51c)
}

var fileDescriptor_59fc3fabf3c5b51c = []byte{
	// 302 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x92, 0x4f, 0x4b, 0xc4, 0x30,
	0x10, 0xc5, 0xb7, 0xee, 0x6e, 0x77, 0x9d, 0xd3, 0x1a, 0x44, 0xa4, 0xfe, 0x41, 0xe2, 0xa5, 0x45,
	0x4d, 0x17, 0x45, 0x2f, 0xde, 0xf4, 0xe2, 0x45, 0x90, 0x7a, 0xf3, 0x96, 0x2d, 0xb3, 0xdd, 0x62,
	0xdb, 0xac, 0x49, 0x7a, 0x10, 0x3f, 0xa6, 0x5f, 0xc8, 0x36, 0x4d, 0x55, 0x44, 0x21, 0xde, 0x26,
	0xcc, 0xef, 0xbd, 0xf0, 0x1e, 0x03, 0xf1, 0x5a, 0x0a, 0x2d, 0xe2, 0x2c, 0xd7, 0xab, 0x7a, 0xc1,
	0x52, 0x51, 0xc6, 0x8a, 0x4b, 0x59, 0x2f, 0xb9, 0x8e, 0x35, 0x2a, 0x5d, 0xe6, 0xa9, 0x14, 0xf1,
	0x0a, 0x8b, 0x42, 0x30, 0x43, 0x92, 0xe3, 0x2f, 0x94, 0xf5, 0x28, 0xfb, 0x44, 0x99, 0x41, 0xe9,
	0x1e, 0x4c, 0xee, 0x51, 0x29, 0x9e, 0x21, 0x99, 0xc1, 0x50, 0xf1, 0xd7, 0x5d, 0xef, 0xc8, 0x0b,
	0x37, 0x93, 0x76, 0xa4, 0x07, 0x30, 0x49, 0xf0, 0xa5, 0x6e, 0x24, 0x84, 0xc0, 0xa8, 0xe2, 0x25,
	0xda, 0xad, 0x99, 0xe9, 0x3e, 0x4c, 0x13, 0x54, 0x6b, 0x51, 0x29, 0x23, 0x2e, 0x55, 0xd6, 0x8b,
	0x9b, 0x91, 0x86, 0x30, 0x7b, 0xd4, 0x12, 0x79, 0x99, 0x57, 0x59, 0xef, 0xb2, 0x0d, 0xe3, 0x54,
	0xd4, 0x95, 0x36, 0xdc, 0x30, 0xe9, 0x1e, 0x34, 0x82, 0xad, 0x6f, 0xa4, 0x35, 0xfc, 0x1d, 0x3d,
	0x84, 0xd1, 0x43, 0x43, 0x91, 0x1d, 0xf0, 0x95, 0x96, 0xe2, 0x19, 0xed, 0xda, 0xbe, 0xcc, 0x5e,
	0xfc, 0xbd, 0x3f, 0x7f, 0xdf, 0x80, 0xf1, 0x5d, 0x1b, 0x9c, 0x20, 0x8c, 0x6e, 0x79, 0x51, 0x90,
	0x53, 0xe6, 0x50, 0x13, 0xb3, 0x01, 0x82, 0x33, 0x47, 0xba, 0x0b, 0x41, 0x07, 0xe4, 0x0d, 0xfc,
	0x2e, 0x1b, 0xb9, 0x74, 0x92, 0xfe, 0xac, 0x2c, 0xb8, 0xfa, 0xaf, 0xac, 0xff, 0x7a, 0xee, 0x91,
	0x25, 0x4c, 0xdb, 0xb6, 0x4c, 0x23, 0x91, 0x93, 0x4f, 0x8b, 0x07, 0x8e, 0x68, 0xe3, 0x4a, 0x07,
	0xa1, 0x37, 0xf7, 0x6e, 0x4e, 0x9e, 0x22, 0x73, 0x72, 0xd7, 0x0e, 0xc7, 0xb9, 0xf0, 0x0d, 0x7a,
	0xf1, 0x11, 0x00, 0x00, 0xff, 0xff, 0x8e, 0x3c, 0x77, 0x2b, 0xd0, 0x02, 0x00, 0x00,
}