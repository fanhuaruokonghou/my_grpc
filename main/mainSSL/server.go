package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
	me "my_grpc/client/streamClient/message"
	"my_grpc/server/streamServer"
	"net"
)

func main() {

	cred, err := credentials.NewServerTLSFromFile("./Key/server.pem", "./Key/server.key")
	if err != nil {
		grpclog.Fatal("加载证书失败", err)
	}

	server := grpc.NewServer(grpc.Creds(cred))
	me.RegisterOrderServiceServer(server, new(streamServer.OrderServiceImpl))
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err.Error())
	}
	server.Serve(listen)

}
