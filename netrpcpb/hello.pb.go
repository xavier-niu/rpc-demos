// Code generated by protoc-gen-go. DO NOT EDIT.
// source: netrpcpb/hello.proto

package netrpcpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import "net/rpc"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// define a struct-like struct, called String
type String struct {
	// define a value that is type of string, named value, and assign the id of it to 1
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *String) Reset()         { *m = String{} }
func (m *String) String() string { return proto.CompactTextString(m) }
func (*String) ProtoMessage()    {}
func (*String) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b6c848699c87aec, []int{0}
}

func (m *String) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_String.Unmarshal(m, b)
}
func (m *String) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_String.Marshal(b, m, deterministic)
}
func (m *String) XXX_Merge(src proto.Message) {
	xxx_messageInfo_String.Merge(m, src)
}
func (m *String) XXX_Size() int {
	return xxx_messageInfo_String.Size(m)
}
func (m *String) XXX_DiscardUnknown() {
	xxx_messageInfo_String.DiscardUnknown(m)
}

var xxx_messageInfo_String proto.InternalMessageInfo

func (m *String) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*String)(nil), "String")
}

func init() { proto.RegisterFile("netrpcpb/hello.proto", fileDescriptor_0b6c848699c87aec) }

var fileDescriptor_0b6c848699c87aec = []byte{
	// 114 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc9, 0x4b, 0x2d, 0x29,
	0x2a, 0x48, 0x2e, 0x48, 0xd2, 0xcf, 0x48, 0xcd, 0xc9, 0xc9, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x57, 0x92, 0xe3, 0x62, 0x0b, 0x2e, 0x29, 0xca, 0xcc, 0x4b, 0x17, 0x12, 0xe1, 0x62, 0x2d, 0x4b,
	0xcc, 0x29, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x70, 0x8c, 0x34, 0xb9, 0x78,
	0x3c, 0x40, 0xca, 0x83, 0x53, 0x8b, 0xca, 0x32, 0x93, 0x53, 0x85, 0x24, 0xb9, 0x58, 0xc1, 0x7c,
	0x21, 0x76, 0x3d, 0x88, 0x3e, 0x29, 0x18, 0xc3, 0x89, 0x3b, 0x8a, 0x53, 0x1f, 0x66, 0x47, 0x12,
	0x1b, 0xd8, 0x78, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x11, 0x5a, 0x47, 0x27, 0x76, 0x00,
	0x00, 0x00,
}

type HelloServiceInterface interface {
	Hello(*String, *String) error
}

func RegisterHelloService(
	srv *rpc.Server, x HelloServiceInterface,
) error {
	if err := srv.RegisterName("HelloService", x); err != nil {
		return err
	}
	return nil
}

type HelloServiceClient struct {
	*rpc.Client
}

var _ HelloServiceInterface = (*HelloServiceClient)(nil)

func DialHelloService(network, address string) (
	*HelloServiceClient, error,
) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{Client: c}, nil
}

func (p *HelloServiceClient) Hello(
	in *String, out *String,
) error {
	return p.Client.Call("HelloService.Hello", in, out)
}