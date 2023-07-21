package main

import (
	"context"
	"fmt"
	"github.com/uwine4850/strux_api/internal/config"
	"github.com/uwine4850/strux_api/services/package_service/internal"
	"github.com/uwine4850/strux_api/services/protofiles/baseproto"
	"github.com/uwine4850/strux_api/services/protofiles/pkgproto"
	"google.golang.org/grpc"
	"net"
)

func main() {
	fmt.Println("Start package service.")
	listener, err := net.Listen("tcp", config.GetPkgServiceAddress())
	if err != nil {
		panic(err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pkgproto.RegisterPackageServer(grpcServer, &server{})
	err = grpcServer.Serve(listener)
	if err != nil {
		panic(err)
	}
}

type server struct {
	pkgproto.UnimplementedPackageServer
}

func (s *server) UploadPackage(c context.Context, request *pkgproto.RequestUploadPackage) (*baseproto.BaseResponse, error) {
	resp := internal.UploadPkg(request)
	return resp, nil
}

func (s *server) ExistsPackage(c context.Context, request *pkgproto.RequestPackageExists) (*baseproto.BaseResponse, error) {
	resp := internal.ExistsPackage(request)
	return resp, nil
}

func (s *server) DownloadPackage(c context.Context, request *pkgproto.RequestDownloadPackage) (*pkgproto.MutateDownloadBaseResponse, error) {
	resp := internal.DownloadPackage(request)
	return resp, nil
}

func (s *server) ShowVersions(c context.Context, request *pkgproto.RequestShowVersions) (*pkgproto.MutateShowVersionBaseResponse, error) {
	resp := internal.ShowVersions(request)
	return resp, nil
}
