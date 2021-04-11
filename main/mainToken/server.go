package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
	tm "my_grpc/client/tokenClient/message"
	ts "my_grpc/server/tokenServer"
	"net"
)

func main() {
	cred, err := credentials.NewServerTLSFromFile("./Key/server.pem", "./Key/server.key")
	if err != nil {
		grpclog.Fatal("证书加载失败")
	}

	server := grpc.NewServer(grpc.Creds(cred), grpc.UnaryInterceptor(ts.TokenInterceptor))

	tm.RegisterAddServiceServer(server, ts.New())

	listen, err := net.Listen("tcp", ":8080")

	if err != nil {
		panic(err.Error())
	}
	server.Serve(listen)
}
