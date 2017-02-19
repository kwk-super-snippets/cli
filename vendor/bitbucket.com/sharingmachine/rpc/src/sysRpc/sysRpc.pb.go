// Code generated by protoc-gen-go.
// source: sysRpc.proto
// DO NOT EDIT!

/*
Package sysRpc is a generated protocol buffer package.

It is generated from these files:
	sysRpc.proto

It has these top-level messages:
	InfoRequest
	InfoResponse
	ErrorRequest
	ErrorResponse
*/
package sysRpc

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

type InfoRequest struct {
	ClientVersion string `protobuf:"bytes,1,opt,name=clientVersion" json:"clientVersion,omitempty"`
}

func (m *InfoRequest) Reset()                    { *m = InfoRequest{} }
func (m *InfoRequest) String() string            { return proto.CompactTextString(m) }
func (*InfoRequest) ProtoMessage()               {}
func (*InfoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *InfoRequest) GetClientVersion() string {
	if m != nil {
		return m.ClientVersion
	}
	return ""
}

type InfoResponse struct {
	Version         string `protobuf:"bytes,1,opt,name=version" json:"version,omitempty"`
	Build           string `protobuf:"bytes,2,opt,name=build" json:"build,omitempty"`
	ClientSupported string `protobuf:"bytes,3,opt,name=clientSupported" json:"clientSupported,omitempty"`
	Dc              string `protobuf:"bytes,4,opt,name=dc" json:"dc,omitempty"`
	Revision        string `protobuf:"bytes,5,opt,name=revision" json:"revision,omitempty"`
}

func (m *InfoResponse) Reset()                    { *m = InfoResponse{} }
func (m *InfoResponse) String() string            { return proto.CompactTextString(m) }
func (*InfoResponse) ProtoMessage()               {}
func (*InfoResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *InfoResponse) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *InfoResponse) GetBuild() string {
	if m != nil {
		return m.Build
	}
	return ""
}

func (m *InfoResponse) GetClientSupported() string {
	if m != nil {
		return m.ClientSupported
	}
	return ""
}

func (m *InfoResponse) GetDc() string {
	if m != nil {
		return m.Dc
	}
	return ""
}

func (m *InfoResponse) GetRevision() string {
	if m != nil {
		return m.Revision
	}
	return ""
}

type ErrorRequest struct {
	Multi bool `protobuf:"varint,1,opt,name=multi" json:"multi,omitempty"`
}

func (m *ErrorRequest) Reset()                    { *m = ErrorRequest{} }
func (m *ErrorRequest) String() string            { return proto.CompactTextString(m) }
func (*ErrorRequest) ProtoMessage()               {}
func (*ErrorRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ErrorRequest) GetMulti() bool {
	if m != nil {
		return m.Multi
	}
	return false
}

type ErrorResponse struct {
}

func (m *ErrorResponse) Reset()                    { *m = ErrorResponse{} }
func (m *ErrorResponse) String() string            { return proto.CompactTextString(m) }
func (*ErrorResponse) ProtoMessage()               {}
func (*ErrorResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func init() {
	proto.RegisterType((*InfoRequest)(nil), "sysRpc.InfoRequest")
	proto.RegisterType((*InfoResponse)(nil), "sysRpc.InfoResponse")
	proto.RegisterType((*ErrorRequest)(nil), "sysRpc.ErrorRequest")
	proto.RegisterType((*ErrorResponse)(nil), "sysRpc.ErrorResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for SysRpc service

type SysRpcClient interface {
	GetApiInfo(ctx context.Context, in *InfoRequest, opts ...grpc.CallOption) (*InfoResponse, error)
	TestTransportError(ctx context.Context, in *ErrorRequest, opts ...grpc.CallOption) (*ErrorResponse, error)
	TestAppError(ctx context.Context, in *ErrorRequest, opts ...grpc.CallOption) (*ErrorResponse, error)
}

type sysRpcClient struct {
	cc *grpc.ClientConn
}

func NewSysRpcClient(cc *grpc.ClientConn) SysRpcClient {
	return &sysRpcClient{cc}
}

func (c *sysRpcClient) GetApiInfo(ctx context.Context, in *InfoRequest, opts ...grpc.CallOption) (*InfoResponse, error) {
	out := new(InfoResponse)
	err := grpc.Invoke(ctx, "/sysRpc.SysRpc/GetApiInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysRpcClient) TestTransportError(ctx context.Context, in *ErrorRequest, opts ...grpc.CallOption) (*ErrorResponse, error) {
	out := new(ErrorResponse)
	err := grpc.Invoke(ctx, "/sysRpc.SysRpc/TestTransportError", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysRpcClient) TestAppError(ctx context.Context, in *ErrorRequest, opts ...grpc.CallOption) (*ErrorResponse, error) {
	out := new(ErrorResponse)
	err := grpc.Invoke(ctx, "/sysRpc.SysRpc/TestAppError", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SysRpc service

type SysRpcServer interface {
	GetApiInfo(context.Context, *InfoRequest) (*InfoResponse, error)
	TestTransportError(context.Context, *ErrorRequest) (*ErrorResponse, error)
	TestAppError(context.Context, *ErrorRequest) (*ErrorResponse, error)
}

func RegisterSysRpcServer(s *grpc.Server, srv SysRpcServer) {
	s.RegisterService(&_SysRpc_serviceDesc, srv)
}

func _SysRpc_GetApiInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysRpcServer).GetApiInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sysRpc.SysRpc/GetApiInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysRpcServer).GetApiInfo(ctx, req.(*InfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysRpc_TestTransportError_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ErrorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysRpcServer).TestTransportError(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sysRpc.SysRpc/TestTransportError",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysRpcServer).TestTransportError(ctx, req.(*ErrorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysRpc_TestAppError_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ErrorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysRpcServer).TestAppError(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sysRpc.SysRpc/TestAppError",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysRpcServer).TestAppError(ctx, req.(*ErrorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SysRpc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sysRpc.SysRpc",
	HandlerType: (*SysRpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetApiInfo",
			Handler:    _SysRpc_GetApiInfo_Handler,
		},
		{
			MethodName: "TestTransportError",
			Handler:    _SysRpc_TestTransportError_Handler,
		},
		{
			MethodName: "TestAppError",
			Handler:    _SysRpc_TestAppError_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sysRpc.proto",
}

func init() { proto.RegisterFile("sysRpc.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 304 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x52, 0x31, 0x4f, 0xf3, 0x30,
	0x10, 0xfd, 0xd2, 0x8f, 0x96, 0xf6, 0x48, 0xa9, 0x74, 0xb4, 0x52, 0x94, 0x09, 0x45, 0x19, 0x3a,
	0x65, 0xa0, 0x13, 0x03, 0x43, 0x8b, 0x10, 0x62, 0x43, 0x49, 0xc5, 0xde, 0x3a, 0x87, 0xb0, 0x1a,
	0x62, 0x63, 0x3b, 0x41, 0xfc, 0x0f, 0xfe, 0x13, 0x7f, 0x0b, 0xc5, 0x6e, 0xa3, 0xa6, 0x23, 0x93,
	0xf5, 0xde, 0xdd, 0xbd, 0x7b, 0xf7, 0x64, 0xf0, 0xf5, 0x97, 0x4e, 0x25, 0x4b, 0xa4, 0x12, 0x46,
	0xe0, 0xc0, 0xa1, 0x68, 0x01, 0x17, 0x4f, 0xe5, 0xab, 0x48, 0xe9, 0xa3, 0x22, 0x6d, 0x30, 0x86,
	0x31, 0x2b, 0x38, 0x95, 0xe6, 0x85, 0x94, 0xe6, 0xa2, 0x0c, 0xbc, 0x6b, 0x6f, 0x3e, 0x4a, 0xbb,
	0x64, 0xf4, 0xed, 0x81, 0xef, 0xa6, 0xb4, 0x14, 0xa5, 0x26, 0x0c, 0xe0, 0xbc, 0xee, 0x0c, 0x1c,
	0x20, 0x4e, 0xa1, 0xbf, 0xad, 0x78, 0x91, 0x07, 0x3d, 0xcb, 0x3b, 0x80, 0x73, 0x98, 0x38, 0xc5,
	0xac, 0x92, 0x52, 0x28, 0x43, 0x79, 0xf0, 0xdf, 0xd6, 0x4f, 0x69, 0xbc, 0x84, 0x5e, 0xce, 0x82,
	0x33, 0x5b, 0xec, 0xe5, 0x0c, 0x43, 0x18, 0x2a, 0xaa, 0xb9, 0x5d, 0xd5, 0xb7, 0x6c, 0x8b, 0xa3,
	0x18, 0xfc, 0x07, 0xa5, 0x84, 0x3a, 0x1c, 0x33, 0x85, 0xfe, 0x7b, 0x55, 0x18, 0x6e, 0x3d, 0x0d,
	0x53, 0x07, 0xa2, 0x09, 0x8c, 0xf7, 0x5d, 0xce, 0xfc, 0xcd, 0x8f, 0x07, 0x83, 0xcc, 0xa6, 0x81,
	0xb7, 0x00, 0x8f, 0x64, 0x96, 0x92, 0x37, 0xd7, 0xe1, 0x55, 0xb2, 0x8f, 0xec, 0x28, 0xa1, 0x70,
	0xda, 0x25, 0x9d, 0x46, 0xf4, 0x0f, 0xef, 0x01, 0xd7, 0xa4, 0xcd, 0x5a, 0x6d, 0x4a, 0xdd, 0x78,
	0xb7, 0x3b, 0xb0, 0xed, 0x3e, 0x36, 0x16, 0xce, 0x4e, 0xd8, 0x56, 0xe4, 0x0e, 0xfc, 0x46, 0x64,
	0x29, 0xe5, 0x5f, 0xc6, 0x57, 0x31, 0xcc, 0xb8, 0x48, 0x76, 0x9f, 0xbb, 0x44, 0x93, 0xaa, 0x39,
	0xa3, 0x44, 0xd3, 0x46, 0xb1, 0xb7, 0xd5, 0x28, 0xb3, 0x6f, 0x2a, 0xd9, 0xb3, 0xb7, 0x1d, 0xd8,
	0x1f, 0xb0, 0xf8, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x96, 0xde, 0x38, 0x78, 0x11, 0x02, 0x00, 0x00,
}