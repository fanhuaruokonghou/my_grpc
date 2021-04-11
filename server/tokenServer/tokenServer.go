package tokenServer

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"io"
	tm "my_grpc/client/tokenClient/message"
)

type mathManager interface {
	Send(*tm.AddInfo) error
	Recv() (*tm.NumberRequest, error)
	grpc.ServerStream
}

type mathManagerImpl struct {
	grpc.ServerStream
}

func New() *mathManagerImpl {
	return new(mathManagerImpl)
}

func (i *mathManagerImpl) Send(addInfo *tm.AddInfo) error {
	return i.ServerStream.SendMsg(addInfo)
}

func (i *mathManagerImpl) Recv() (*tm.NumberRequest, error) {
	m := new(tm.NumberRequest)
	if err := i.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (i *mathManagerImpl) AddMethod(ctx context.Context, request *tm.NumberRequest) (response *tm.AddInfo, err error) {
	md, exist := metadata.FromIncomingContext(ctx)
	if !exist {
		return nil, status.Errorf(codes.Unauthenticated, "没有认证信息")
	}
	var appKey string
	var appSecret string
	if key, ok := md["appkey"]; ok {
		appKey = key[0]
	}
	if secret, ok := md["appsecret"]; ok {
		appSecret = secret[0]
	}

	if appSecret != "20210407" || appKey != "fanhua" {
		return nil, status.Errorf(codes.Unimplemented, "Token 不合法")
	}
	fmt.Println("服务端 Add 方法")
	result := request.GetArge2() + request.GetArge1()
	response = new(tm.AddInfo)
	response.Result = result
	return response, nil
}

func (i *mathManagerImpl) GetAddInfos(stream tm.AddService_GetAddInfosServer) error {
	for true {
		request, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("服务端读取完毕")
			return err
		}
		if err != nil {
			panic(err.Error())
			return nil
		}
		numberRequest := new(tm.NumberRequest)
		numberRequest.Arge1 = request.Arge1
		numberRequest.Arge2 = request.Arge2
		result, err := i.AddMethod(stream.Context(), numberRequest)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("计算结果：", result)
		err = stream.Send(result)
		if err == io.EOF {
			grpclog.Fatal("服务端发送完毕")
			return err
		}
		if err != nil {
			panic(err.Error())
			return err
		}
	}
	return nil
}

func TokenInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

	//通过metadata
	md, exist := metadata.FromIncomingContext(ctx)
	if !exist {
		return nil, status.Errorf(codes.Unauthenticated, "无Token认证信息")
	}
	fmt.Println(md)
	var appKey string
	var appSecret string

	if key, ok := md["appkey"]; ok {
		appKey = key[0]
	}

	if secret, ok := md["appsecret"]; ok {
		appSecret = secret[0]
	}

	if appKey != "fanhua" || appSecret != "20210407" {
		return nil, status.Errorf(codes.Unauthenticated, "Token 不合法")
	}

	//通过token验证，继续处理请求
	return handler(ctx, req)
}
