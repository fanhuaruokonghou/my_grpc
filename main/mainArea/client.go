package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io"
	am "my_grpc/client/areaClient/areaMessage"
)

func main() {
	cred, err := credentials.NewClientTLSFromFile("/Users/wenpanpan/Desktop/go_source/my_grpc/Key/server.pem", "*.fanhua.com")
	if err != nil {
		panic(err.Error())
	}
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(cred))
	if err != nil {
		panic(err.Error())
	}

	roundnessServiceClient := am.NewRoundnessAreaServiceClient(conn)
	rectangleServiceClient := am.NewRectangleAreaServiceClient(conn)
	fmt.Println("客户端请求：双向流模式")

	roundnesses := []float32{float32(1), float32(3), float32(5), float32(7)}
	//rectangle := [][]float32{}
	rectangles := map[float32]float32{float32(1): float32(3), float32(5): float32(7), float32(9): float32(11), float32(13): float32(15)}

	roundnessAreaClient, err := roundnessServiceClient.GetRoundnessAreaInfos(context.Background())
	for _, roundness := range roundnesses {
		roundnessRequest := am.Roundness{R: roundness}
		err := roundnessAreaClient.Send(&roundnessRequest)
		if err != nil {
			panic(err.Error())
		}
	}
	roundnessAreaClient.CloseSend()

	rectangleAreaClient, err := rectangleServiceClient.GetRectangleAreaInfos(context.Background())
	for width, high := range rectangles {
		rectangleRequest := am.Rectangle{Width: width, High: high}
		err := rectangleAreaClient.Send(&rectangleRequest)
		if err != nil {
			panic(err.Error())
		}
	}

	rectangleAreaClient.CloseSend()

	for true {
		area, err := roundnessAreaClient.Recv()
		if err == io.EOF {
			fmt.Println("客户端读取圆面积计算结果结束...")
			break
		}
		if err != nil {
			break
		}
		fmt.Println("读取到圆面积计算结果：", area)
	}
	for true {
		area, err := rectangleAreaClient.Recv()
		if err == io.EOF {
			fmt.Println("客户端读取矩形面积计算结果结束...")
			return
		}
		if err != nil {
			return
		}
		fmt.Println("读取到矩形面积计算结果：", area)
	}
}
