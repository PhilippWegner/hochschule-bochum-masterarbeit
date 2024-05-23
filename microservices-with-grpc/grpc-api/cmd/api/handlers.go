package main

import (
	"context"

	"github.com/PhilippWegner/hochschule-bochum-masterarbeit/microservices-with-grpc/grpc-api/database"
	"github.com/PhilippWegner/hochschule-bochum-masterarbeit/microservices-with-grpc/grpc-api/model"
)

type ModelServiceServer struct {
	model.UnimplementedModelServiceServer
}

var influx = database.ConnectInfluxdb()

func (api *ModelServiceServer) GetPlcs(ctx context.Context, in *model.GetPlcsRequest) (*model.GetPlcsResponse, error) {
	machine := in.GetMachine()
	time := in.GetTime()
	limit := int(in.GetLimit())

	plcs, err := influx.GetPlcs(machine, time, limit)
	if err != nil {
		return nil, err
	}
	return &model.GetPlcsResponse{Plcs: plcs}, nil
}
