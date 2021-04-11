package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io"
	me "my_grpc/client/streamClient/message"
)

func main() {

	creds, err := credentials.NewClientTLSFromFile("./Key/server.pem", "*.fanhua.com")
	if err != nil {
		panic(err.Error())
	}

	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(creds))
	if err != nil {
		panic(err.Error())
	}

	defer conn.Close()

	orderServiceClient := me.NewOrderServiceClient(conn)
	fmt.Println("客户端请求RPC调用：双向流模式")
	orderIds := []string{"201907300001", "201907310001", "201907310002"}

	orderInfoClient, err := orderServiceClient.GetOrderInfos(context.Background())

	for _, orderId := range orderIds {
		orderRequest := me.OrderRequest{OrderId: orderId}
		err := orderInfoClient.Send(&orderRequest)
		if err != nil {
			panic(err.Error())
		}
	}

	orderInfoClient.CloseSend()

	for true {
		orderInfo, err := orderInfoClient.Recv()
		if err == io.EOF {
			fmt.Println("读取结束")
			return
		}
		if err != nil {
			return
		}
		fmt.Println("读取到的信息：", orderInfo)
	}

}
