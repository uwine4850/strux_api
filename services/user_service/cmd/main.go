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

	protobufs2.RegisterReverseServer(grpcServer, &server{})
	err = grpcServer.Serve(listener)
	if err != nil {
		panic(err)
	}
}

type server struct {
	protobufs2.UnimplementedReverseServer
}

func (s *server) CreateUser(c context.Context, request *protobufs2.RequestCreateUser) (*protobufs2.BaseResponse, error) {
	resp := internal.CreateUser(request.Username, request.Password)
	return resp, nil
}

func (s *server) GetUser(c context.Context, request *protobufs2.RequestGetUser) (*protobufs2.ResponseGetUser, error) {
	r := &protobufs2.ResponseGetUser{Username: "Tee", Success: true}
	return r, nil
}
