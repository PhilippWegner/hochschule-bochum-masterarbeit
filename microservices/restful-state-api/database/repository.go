package database

import "github.com/PhilippWegner/hochschule-bochum-masterarbeit/microservices/restful-state-api/model"

type Repository interface {
	CreateStates(states []*model.State) error
	GetStates(machine string, limit int) ([]*model.State, error)
}
