package database

import "github.com/PhilippWegner/hochschule-bochum-masterarbeit/microservices-with-restful/restful-api/model"

type Repository interface {
	CreateStates(states []*model.State) error
	GetStates(machine string, limit int) ([]*model.State, error)
	GetPlcs(machine string, time string, limit int) ([]*model.Plc, error)
}
