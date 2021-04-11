package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io"
	tc "my_grpc/client/tokenClient"
	tm "my_grpc/client/tokenClient/message"
)

func main() {
	cred, err := credentials.NewClientTLSFromFile("./Key/Server.pem", "*.fanhua.com")
	if err != nil {
		panic(err.Error())
	}
	auth := tc.TokenAuthentication{
		AppKey:    "fanhua",
		AppSecret: "20210407",
	}

	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(cred), grpc.WithPerRPCCredentials(&auth))
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()
	addService := map[int]tm.NumberRequest{
		0: tm.NumberRequest{Arge1: 1, Arge2: 3},
		1: tm.NumberRequest{Arge1: 5, Arge2: 7},
		2: tm.NumberRequest{Arge1: 9, Arge2: 2},
		3: tm.NumberRequest{Arge1: 4, Arge2: 6},
		4: tm.NumberRequest{Arge1: 8, Arge2: 0},
	}
	addServiceClient := tm.NewAddServiceClient(conn)

	add, err := addServiceClient.GetAddInfos(context.Background())
	if err != nil {
		panic(err.Error())
	}
	for _, info := range addService {
		err := add.Send(&info)
		if err != nil {
			panic(err.Error())
		}
	}
	add.CloseSend()
	for true {
		addInfo, err := add.Recv()
		if err == io.EOF {
			fmt.Println("数据读取结束...")
			return
		}
		if err != nil {
			//fmt.Println(err.Error())
			return
		}
		fmt.Println(addInfo)
	}

}
