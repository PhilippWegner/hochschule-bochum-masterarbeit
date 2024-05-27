package main

import (
	"context"

	"github.com/PhilippWegner/hochschule-bochum-masterarbeit/microservices/grpc-logger-api/database"
	"github.com/PhilippWegner/hochschule-bochum-masterarbeit/microservices/grpc-logger-api/model"
)

type ModelServiceServer struct {
	model.UnimplementedModelServiceServer
}

var mongo = database.MongdbConnect()

func (s *ModelServiceServer) WriteLog(ctx context.Context, in *model.LogRequest) (*model.LogResponse, error) {
	log := database.LogEntry{
		Name: in.LogEntry.Name,
		Data: in.LogEntry.Data,
	}
	err := mongo.Insert(log)
	if err != nil {
		return nil, err
	}
	return &model.LogResponse{}, nil
}
