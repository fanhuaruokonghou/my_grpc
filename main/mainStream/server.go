package main

import (
	"google.golang.org/grpc"
	me "my_grpc/client/streamClient/message"
	"my_grpc/server/streamServer"
	"net"
)

func main() {
	server := grpc.NewServer()
	me.RegisterOrderServiceServer(server, new(streamServer.OrderServiceImpl))
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err.Error())
	}
	server.Serve(listen)

}
