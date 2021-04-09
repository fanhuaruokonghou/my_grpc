package SSLServer

import (
"fmt"
"google.golang.org/grpc"
"io"
	me "my_grpc/client/streamClient/message"
)

type OrderService_GetOrderInfoServer interface {
	Send(*me.OrderInfo) error
	Recv() (*me.OrderRequest, error)
	grpc.ServerStream
}

type OrderServiceImpl struct {
}

type orderServerGetOrderInfosServer struct {
	grpc.ServerStream
}

func (o *orderServerGetOrderInfosServer) Send(m *me.OrderInfo) error {
	return o.ServerStream.SendMsg(m)
}

func (o orderServerGetOrderInfosServer) Recv() (*me.OrderRequest, error) {
	m := new(me.OrderRequest)
	if err := o.ServerStream.SendMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (i OrderServiceImpl) GetOrderInfos(stream me.OrderService_GetOrderInfosServer) error {
	for true {
		orderRequest, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("数据读取结束")
			return err
		}
		if err != nil {
			panic(err.Error())
			return err
		}

		fmt.Println(orderRequest.GetOrderId())
		orderMap := map[string]me.OrderInfo{
			"201907300001": me.OrderInfo{OrderId: "201907300001", OrderName: "衣服", OrderStatus: "已付款"},
			"201907310001": me.OrderInfo{OrderId: "201907310001", OrderName: "零食", OrderStatus: "已付款"},
			"201907310002": me.OrderInfo{OrderId: "201907310002", OrderName: "食品", OrderStatus: "未付款"},
		}

		result := orderMap[orderRequest.GetOrderId()]
		err = stream.Send(&result)
		if err == io.EOF {
			fmt.Println(err)
			return err
		}
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
	}
	return nil
}
