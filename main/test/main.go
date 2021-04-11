package main

import (
	"fmt"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
)

func main() {
	cred, err := credentials.NewServerTLSFromFile("./Key/server.pem", "./Key/server.key")
	if err != nil {
		grpclog.Fatal("证书加载失败")
	}
	fmt.Println(cred)
}
