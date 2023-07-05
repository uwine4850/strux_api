// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: package.proto

package pkgproto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	baseproto "strux_api/services/protofiles/baseproto"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PackageClient is the client API for Package service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PackageClient interface {
	UploadPackage(ctx context.Context, in *RequestUploadPackage, opts ...grpc.CallOption) (*baseproto.BaseResponse, error)
	ExistsPackage(ctx context.Context, in *RequestPackageExists, opts ...grpc.CallOption) (*baseproto.BaseResponse, error)
	DownloadPackage(ctx context.Context, in *RequestDownloadPackage, opts ...grpc.CallOption) (*MutateDownloadBaseResponse, error)
}

type packageClient struct {
	cc grpc.ClientConnInterface
}

func NewPackageClient(cc grpc.ClientConnInterface) PackageClient {
	return &packageClient{cc}
}

func (c *packageClient) UploadPackage(ctx context.Context, in *RequestUploadPackage, opts ...grpc.CallOption) (*baseproto.BaseResponse, error) {
	out := new(baseproto.BaseResponse)
	err := c.cc.Invoke(ctx, "/pkgproto.Package/UploadPackage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *packageClient) ExistsPackage(ctx context.Context, in *RequestPackageExists, opts ...grpc.CallOption) (*baseproto.BaseResponse, error) {
	out := new(baseproto.BaseResponse)
	err := c.cc.Invoke(ctx, "/pkgproto.Package/ExistsPackage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *packageClient) DownloadPackage(ctx context.Context, in *RequestDownloadPackage, opts ...grpc.CallOption) (*MutateDownloadBaseResponse, error) {
	out := new(MutateDownloadBaseResponse)
	err := c.cc.Invoke(ctx, "/pkgproto.Package/DownloadPackage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PackageServer is the server API for Package service.
// All implementations must embed UnimplementedPackageServer
// for forward compatibility
type PackageServer interface {
	UploadPackage(context.Context, *RequestUploadPackage) (*baseproto.BaseResponse, error)
	ExistsPackage(context.Context, *RequestPackageExists) (*baseproto.BaseResponse, error)
	DownloadPackage(context.Context, *RequestDownloadPackage) (*MutateDownloadBaseResponse, error)
	mustEmbedUnimplementedPackageServer()
}

// UnimplementedPackageServer must be embedded to have forward compatible implementations.
type UnimplementedPackageServer struct {
}

func (UnimplementedPackageServer) UploadPackage(context.Context, *RequestUploadPackage) (*baseproto.BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadPackage not implemented")
}
func (UnimplementedPackageServer) ExistsPackage(context.Context, *RequestPackageExists) (*baseproto.BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistsPackage not implemented")
}
func (UnimplementedPackageServer) DownloadPackage(context.Context, *RequestDownloadPackage) (*MutateDownloadBaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownloadPackage not implemented")
}
func (UnimplementedPackageServer) mustEmbedUnimplementedPackageServer() {}

// UnsafePackageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PackageServer will
// result in compilation errors.
type UnsafePackageServer interface {
	mustEmbedUnimplementedPackageServer()
}

func RegisterPackageServer(s grpc.ServiceRegistrar, srv PackageServer) {
	s.RegisterService(&Package_ServiceDesc, srv)
}

func _Package_UploadPackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestUploadPackage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PackageServer).UploadPackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pkgproto.Package/UploadPackage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PackageServer).UploadPackage(ctx, req.(*RequestUploadPackage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Package_ExistsPackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestPackageExists)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PackageServer).ExistsPackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pkgproto.Package/ExistsPackage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PackageServer).ExistsPackage(ctx, req.(*RequestPackageExists))
	}
	return interceptor(ctx, in, info, handler)
}

func _Package_DownloadPackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestDownloadPackage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PackageServer).DownloadPackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pkgproto.Package/DownloadPackage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PackageServer).DownloadPackage(ctx, req.(*RequestDownloadPackage))
	}
	return interceptor(ctx, in, info, handler)
}

// Package_ServiceDesc is the grpc.ServiceDesc for Package service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Package_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pkgproto.Package",
	HandlerType: (*PackageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UploadPackage",
			Handler:    _Package_UploadPackage_Handler,
		},
		{
			MethodName: "ExistsPackage",
			Handler:    _Package_ExistsPackage_Handler,
		},
		{
			MethodName: "DownloadPackage",
			Handler:    _Package_DownloadPackage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "package.proto",
}
