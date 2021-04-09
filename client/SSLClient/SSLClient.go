package SSLClient

import (
	"google.golang.org/grpc"
	me "my_grpc/client/streamClient/message"
)

type orderService_GetOrderInfosService interface {
	Send(*me.OrderRequest) error
	Recv() (*me.OrderInfo, error)
	grpc.ClientStream
}

type orderServiceGetOrderInfosClient struct {
	grpc.ClientStream
}

func (o *orderServiceGetOrderInfosClient) Send(m *me.OrderRequest) error {
	return o.ClientStream.SendMsg(m)
}

func (o orderServiceGetOrderInfosClient) Recv() (*me.OrderInfo, error) {
	m := new(me.OrderInfo)
	if err := o.ClientStream.SendMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}
