package areaServer

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"math"
)
import am "my_grpc/client/areaClient/areaMessage"

type Roundness interface {
	Area(context.Context, *am.Roundness) (*am.AreaInfo, error)
	Send(roundness *am.AreaInfo) error
	Recv() (*am.Roundness, error)
	grpc.ServerStream
}

type Rectangle interface {
	Area(context.Context, *am.Rectangle) (*am.AreaInfo, error)
	Send(*am.AreaInfo) error
	Recv() (*am.Rectangle, error)
	grpc.ServerStream
}

type roundnessImpl struct {
	grpc.ServerStream
}

type rectangleImpl struct {
	grpc.ServerStream
}

func NewRoundness() *roundnessImpl {
	return new(roundnessImpl)
}

func NewRectangle() *rectangleImpl {
	return new(rectangleImpl)
}

func (ro *roundnessImpl) Area(ri *am.Roundness) (*am.AreaInfo, error) {
	fmt.Println("服务端计算圆面积 Area 方法")
	area := ri.R * ri.R * math.Pi
	fmt.Println("圆面积计算结果为：", area)
	areaResponse := new(am.AreaInfo)
	areaResponse.Area = area
	return areaResponse, nil
}

func (ro *roundnessImpl) Send(ri *am.AreaInfo) error {
	return ro.ServerStream.SendMsg(ri)
}

func (ro roundnessImpl) Recv() (*am.Roundness, error) {
	m := new(am.Roundness)
	if err := ro.ServerStream.SendMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (r *rectangleImpl) Area(re *am.Rectangle) (*am.AreaInfo, error) {
	fmt.Println("服务端计算矩形面积 Area 方法")
	area := re.Width * re.High
	fmt.Println("矩形面积计算结果为：", area)
	areaResponse := new(am.AreaInfo)
	areaResponse.Area = area
	return areaResponse, nil
}

func (r *rectangleImpl) Send(re *am.AreaInfo) error {
	return r.ServerStream.SendMsg(re)
}

func (r rectangleImpl) Recv() (*am.Rectangle, error) {
	m := new(am.Rectangle)
	if err := r.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (ro *roundnessImpl) GetRoundnessAreaInfos(stream am.RoundnessAreaService_GetRoundnessAreaInfosServer) error {
	for true {
		areaRequest, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("圆半径数据读取结束")
			return err
		}
		if err != nil {
			panic(err.Error())
			return nil
		}
		fmt.Printf("读取到的圆的半径%.2f\n", areaRequest.GetR())
		m := new(am.Roundness)
		m.R = areaRequest.GetR()
		result, err := ro.Area(m)
		if err != nil {
			panic(err.Error())
		}
		err = stream.Send(result)
		if err == io.EOF {
			fmt.Println("服务端发送完毕")
			return err
		}
		if err != nil {
			fmt.Println(err.Error())
			return err
		}

	}
	return nil
}

func (r *rectangleImpl) GetRectangleAreaInfos(stream am.RectangleAreaService_GetRectangleAreaInfosServer) error {
	for true {
		areaRequest, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("矩形半径数据读取结束")
			return err
		}
		if err != nil {
			panic(err.Error())
			return nil
		}
		fmt.Printf("矩形的长：%.2f，矩形的宽：%.2f\n", areaRequest.GetHigh(), areaRequest.GetWidth())
		m := new(am.Rectangle)
		m.Width = areaRequest.GetWidth()
		m.High = areaRequest.GetHigh()
		result, err := r.Area(m)
		if err != nil {
			panic(err.Error())
		}
		err = stream.Send(result)
		if err == io.EOF {
			fmt.Println("服务端发送完毕")
			return err
		}
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
	}
	return nil
}
