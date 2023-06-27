package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"net"
	"strux_api/internal/config"
	"strux_api/services/protofiles/baseproto"
	"strux_api/services/protofiles/userproto"
	"strux_api/services/user_service/internal"
)

func main() {
	listener, err := net.Listen("tcp", config.UserServiceAddress)

	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	userproto.RegisterUserServer(grpcServer, &server{})
	err = grpcServer.Serve(listener)
	if err != nil {
		panic(err)
	}
}

type server struct {
	userproto.UnimplementedUserServer
}

func (s *server) CreateUser(c context.Context, request *userproto.RequestCreateUser) (*baseproto.BaseResponse, error) {
	resp := internal.CreateUser(request.Username, request.Password)
	return resp, nil
}

func (s *server) UserExist(c context.Context, request *userproto.RequestExistUser) (*baseproto.BaseResponse, error) {
	resp := internal.UserExist(request.Username)
	return resp, nil
}

func (s *server) UserDelete(c context.Context, request *userproto.RequestDeleteUser) (*baseproto.BaseResponse, error) {
	resp := internal.UserDelete(request.Username, request.Password)
	return resp, nil
}

func (s *server) UserUpdatePassword(c context.Context, request *userproto.RequestUpdatePassword) (*baseproto.BaseResponse, error) {
	resp := internal.PasswordUpdate(request.Username, request.Password, request.NewPassword)
	return resp, nil
}

func (s *server) UserLogIn(c context.Context, request *userproto.RequestUserLogIn) (*baseproto.BaseResponse, error) {
	resp := internal.UserLogIn(request.Username, request.Password)
	return resp, nil
}
