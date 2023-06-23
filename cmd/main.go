package main

import (
	"net/http"
	"strux_api/internal/config"
	"strux_api/internal/rest_api/routes"
)

func main() {
	err := http.ListenAndServe(config.Host+":"+config.Port, routes.UsersInit())
	if err != nil {
		panic(err)
	}
}

//func main() {
//	opts := []grpc.DialOption{
//		grpc.WithTransportCredentials(insecure.NewCredentials()),
//	}
//	conn, err := grpc.Dial("127.0.0.1:5300", opts...)
//
//	if err != nil {
//		grpclog.Fatalf("fail to dial: %v", err)
//	}
//
//	defer conn.Close()
//
//	client := protobufs.NewReverseClient(conn)
//	//request := &protobufs.RequestSetUser{
//	//	Username: "tee",
//	//	Password: "pass",
//	//}
//	//response, err := cmd.SetUser(context.Background(), request)
//	response, err := client.GetUser(context.Background(), &protobufs.RequestGetUser{Username: "tee"})
//
//	if err != nil {
//		grpclog.Fatalf("fail to dial: %v", err)
//	}
//
//	fmt.Println(response)
//}
