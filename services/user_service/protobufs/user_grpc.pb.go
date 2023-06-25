// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: protobufs/user.proto

package protobufs

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserClient interface {
	CreateUser(ctx context.Context, in *RequestCreateUser, opts ...grpc.CallOption) (*BaseResponse, error)
	UserExist(ctx context.Context, in *RequestExistUser, opts ...grpc.CallOption) (*BaseResponse, error)
	UserDelete(ctx context.Context, in *RequestDeleteUser, opts ...grpc.CallOption) (*BaseResponse, error)
	UserUpdatePassword(ctx context.Context, in *RequestUpdatePassword, opts ...grpc.CallOption) (*BaseResponse, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) CreateUser(ctx context.Context, in *RequestCreateUser, opts ...grpc.CallOption) (*BaseResponse, error) {
	out := new(BaseResponse)
	err := c.cc.Invoke(ctx, "/protobufs.User/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserExist(ctx context.Context, in *RequestExistUser, opts ...grpc.CallOption) (*BaseResponse, error) {
	out := new(BaseResponse)
	err := c.cc.Invoke(ctx, "/protobufs.User/UserExist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserDelete(ctx context.Context, in *RequestDeleteUser, opts ...grpc.CallOption) (*BaseResponse, error) {
	out := new(BaseResponse)
	err := c.cc.Invoke(ctx, "/protobufs.User/UserDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserUpdatePassword(ctx context.Context, in *RequestUpdatePassword, opts ...grpc.CallOption) (*BaseResponse, error) {
	out := new(BaseResponse)
	err := c.cc.Invoke(ctx, "/protobufs.User/UserUpdatePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
// All implementations must embed UnimplementedUserServer
// for forward compatibility
type UserServer interface {
	CreateUser(context.Context, *RequestCreateUser) (*BaseResponse, error)
	UserExist(context.Context, *RequestExistUser) (*BaseResponse, error)
	UserDelete(context.Context, *RequestDeleteUser) (*BaseResponse, error)
	UserUpdatePassword(context.Context, *RequestUpdatePassword) (*BaseResponse, error)
	mustEmbedUnimplementedUserServer()
}

// UnimplementedUserServer must be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (UnimplementedUserServer) CreateUser(context.Context, *RequestCreateUser) (*BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUserServer) UserExist(context.Context, *RequestExistUser) (*BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserExist not implemented")
}
func (UnimplementedUserServer) UserDelete(context.Context, *RequestDeleteUser) (*BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserDelete not implemented")
}
func (UnimplementedUserServer) UserUpdatePassword(context.Context, *RequestUpdatePassword) (*BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserUpdatePassword not implemented")
}
func (UnimplementedUserServer) mustEmbedUnimplementedUserServer() {}

// UnsafeUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServer will
// result in compilation errors.
type UnsafeUserServer interface {
	mustEmbedUnimplementedUserServer()
}

func RegisterUserServer(s grpc.ServiceRegistrar, srv UserServer) {
	s.RegisterService(&User_ServiceDesc, srv)
}

func _User_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestCreateUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobufs.User/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).CreateUser(ctx, req.(*RequestCreateUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserExist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestExistUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserExist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobufs.User/UserExist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserExist(ctx, req.(*RequestExistUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestDeleteUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobufs.User/UserDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserDelete(ctx, req.(*RequestDeleteUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserUpdatePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestUpdatePassword)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserUpdatePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobufs.User/UserUpdatePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserUpdatePassword(ctx, req.(*RequestUpdatePassword))
	}
	return interceptor(ctx, in, info, handler)
}

// User_ServiceDesc is the grpc.ServiceDesc for User service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var User_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protobufs.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _User_CreateUser_Handler,
		},
		{
			MethodName: "UserExist",
			Handler:    _User_UserExist_Handler,
		},
		{
			MethodName: "UserDelete",
			Handler:    _User_UserDelete_Handler,
		},
		{
			MethodName: "UserUpdatePassword",
			Handler:    _User_UserUpdatePassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobufs/user.proto",
}
