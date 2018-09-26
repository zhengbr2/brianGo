// Code generated by protoc-gen-go. DO NOT EDIT.
// source: data.proto

package example

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Data struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Data) Reset()         { *m = Data{} }
func (m *Data) String() string { return proto.CompactTextString(m) }
func (*Data) ProtoMessage()    {}
func (*Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_data_452b729f4e02bf2e, []int{0}
}
func (m *Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Data.Unmarshal(m, b)
}
func (m *Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Data.Marshal(b, m, deterministic)
}
func (dst *Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Data.Merge(dst, src)
}
func (m *Data) XXX_Size() int {
	return xxx_messageInfo_Data.Size(m)
}
func (m *Data) XXX_DiscardUnknown() {
	xxx_messageInfo_Data.DiscardUnknown(m)
}

var xxx_messageInfo_Data proto.InternalMessageInfo

func (m *Data) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func init() {
	proto.RegisterType((*Data)(nil), "example.Data")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FormatDataClient is the client API for FormatData service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FormatDataClient interface {
	DoFormat(ctx context.Context, in *Data, opts ...grpc.CallOption) (*Data, error)
}

type formatDataClient struct {
	cc *grpc.ClientConn
}

func NewFormatDataClient(cc *grpc.ClientConn) FormatDataClient {
	return &formatDataClient{cc}
}

func (c *formatDataClient) DoFormat(ctx context.Context, in *Data, opts ...grpc.CallOption) (*Data, error) {
	out := new(Data)
	err := c.cc.Invoke(ctx, "/example.FormatData/DoFormat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FormatDataServer is the server API for FormatData service.
type FormatDataServer interface {
	DoFormat(context.Context, *Data) (*Data, error)
}

func RegisterFormatDataServer(s *grpc.Server, srv FormatDataServer) {
	s.RegisterService(&_FormatData_serviceDesc, srv)
}

func _FormatData_DoFormat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Data)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FormatDataServer).DoFormat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.FormatData/DoFormat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FormatDataServer).DoFormat(ctx, req.(*Data))
	}
	return interceptor(ctx, in, info, handler)
}

var _FormatData_serviceDesc = grpc.ServiceDesc{
	ServiceName: "example.FormatData",
	HandlerType: (*FormatDataServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DoFormat",
			Handler:    _FormatData_DoFormat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "data.proto",
}

func init() { proto.RegisterFile("data.proto", fileDescriptor_data_452b729f4e02bf2e) }

var fileDescriptor_data_452b729f4e02bf2e = []byte{
	// 106 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x49, 0x2c, 0x49,
	0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4f, 0xad, 0x48, 0xcc, 0x2d, 0xc8, 0x49, 0x55,
	0x92, 0xe2, 0x62, 0x71, 0x49, 0x2c, 0x49, 0x14, 0x12, 0xe2, 0x62, 0x29, 0x49, 0xad, 0x28, 0x91,
	0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x8d, 0x2c, 0xb8, 0xb8, 0xdc, 0xf2, 0x8b, 0x72,
	0x13, 0x4b, 0xc0, 0x2a, 0xb4, 0xb8, 0x38, 0x5c, 0xf2, 0x21, 0x7c, 0x21, 0x5e, 0x3d, 0xa8, 0x7e,
	0x3d, 0x90, 0x94, 0x14, 0x2a, 0x57, 0x89, 0x21, 0x89, 0x0d, 0x6c, 0x8b, 0x31, 0x20, 0x00, 0x00,
	0xff, 0xff, 0x33, 0xcb, 0xa7, 0xd5, 0x73, 0x00, 0x00, 0x00,
}
