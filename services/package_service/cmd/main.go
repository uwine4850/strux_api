package main

import (
	"context"
	"google.golang.org/grpc"
	"net"
	"strux_api/internal/config"
	"strux_api/services/package_service/internal"
	"strux_api/services/protofiles/baseproto"
	"strux_api/services/protofiles/pkgproto"
)

func main() {
	listener, err := net.Listen("tcp", config.PkgServiceAddress)
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
	internal.DownloadPackage(request)
	baseResp := &baseproto.BaseResponse{
		Message: "OK",
		Success: true,
		Status:  0,
	}
	resp := &pkgproto.MutateDownloadBaseResponse{
		BaseResponse: baseResp,
		UplFiles:     nil,
		UplDirInfo:   nil,
	}
	return resp, nil
}
