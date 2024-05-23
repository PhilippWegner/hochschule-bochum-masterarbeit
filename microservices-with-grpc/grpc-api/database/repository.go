package database

import (
	"github.com/PhilippWegner/hochschule-bochum-masterarbeit/microservices-with-grpc/grpc-api/model"
)

type Repository interface {
	GetPlcs(machine string, time string, limit int) ([]*model.Plc, error)
	GetStates(machine string, limit int) ([]*model.State, error)
	CreateStates(states []*model.State) error
}
