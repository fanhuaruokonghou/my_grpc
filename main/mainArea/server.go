package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
	am "my_grpc/client/areaClient/areaMessage"
	as "my_grpc/server/areaServer"
	"net"
)

func main() {
	cred, err := credentials.NewServerTLSFromFile("/Users/wenpanpan/Desktop/go_source/my_grpc/Key/server.pem", "/Users/wenpanpan/Desktop/go_source/my_grpc/Key/server.key")
	if err != nil {
		grpclog.Fatal("服务端证书加载失败：", err.Error())
	}

	server := grpc.NewServer(grpc.Creds(cred))

	am.RegisterRoundnessAreaServiceServer(server, as.NewRoundness())
	am.RegisterRectangleAreaServiceServer(server, as.NewRectangle())
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err.Error())
	}
	server.Serve(listen)
}
