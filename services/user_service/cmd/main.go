package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"net"
	"strux_api/internal/config"
	"strux_api/services/user_service/internal"
	protobufs2 "strux_api/services/user_service/protobufs"
)

func main() {
	listener, err := net.Listen("tcp", config.UserServiceAddress)

	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	protobufs2.RegisterUserServer(grpcServer, &server{})
	err = grpcServer.Serve(listener)
	if err != nil {
		panic(err)
	}
}

type server struct {
	protobufs2.UnimplementedUserServer
}

func (s *server) CreateUser(c context.Context, request *protobufs2.RequestCreateUser) (*protobufs2.BaseResponse, error) {
	resp := internal.CreateUser(request.Username, request.Password)
	return resp, nil
}

func (s *server) UserExist(c context.Context, request *protobufs2.RequestExistUser) (*protobufs2.BaseResponse, error) {
	resp := internal.UserExist(request.Username)
	return resp, nil
}
