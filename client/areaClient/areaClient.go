package areaClient

import (
	"google.golang.org/grpc"
	am "my_grpc/client/areaClient/areaMessage"
)

type roundnessAreaService_GetRoundnessAreaInfos interface {
	Send(*am.Roundness) error
	Recv() (*am.AreaInfo, error)
	grpc.ClientStream
}

type rectangleAreaService_GetRectangleAreaInfos interface {
	Send(*am.Rectangle) error
	Recv() (*am.AreaInfo, error)
	grpc.ClientStream
}

type roundnessAreaServiceGetRoundnessAreaInfos struct {
	grpc.ClientStream
}

type rectangleAreaServiceGetRectangleAreaInfos struct {
	grpc.ClientStream
}

func (receiver *roundnessAreaServiceGetRoundnessAreaInfos) Send(ro *am.Roundness) error {
	return receiver.ClientStream.SendMsg(ro)
}

func (receiver roundnessAreaServiceGetRoundnessAreaInfos) Recv() (*am.AreaInfo, error) {
	m := new(am.AreaInfo)
	if err := receiver.ClientStream.SendMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (receiver *rectangleAreaServiceGetRectangleAreaInfos) Send(re *am.Rectangle) error {
	return receiver.ClientStream.SendMsg(re)
}

func (receiver *rectangleAreaServiceGetRectangleAreaInfos) Recv() (*am.AreaInfo, error) {
	m := new(am.AreaInfo)
	if err := receiver.ClientStream.SendMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}
