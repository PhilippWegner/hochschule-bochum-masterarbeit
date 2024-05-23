package main

import (
	"context"
	"log"

	"github.com/PhilippWegner/hochschule-bochum-masterarbeit/microservices-with-grpc/statemachine-service/data"
	"github.com/PhilippWegner/hochschule-bochum-masterarbeit/microservices-with-grpc/statemachine-service/model"
)

var DEFAULT_LAST_STATE = model.State{Time: "0"}

func (c *Config) calculate(machine string) error {
	// get last states
	lastState := &DEFAULT_LAST_STATE
	state, err := c.Client.GetStates(context.Background(), &model.GetStatesRequest{Machine: machine, Limit: 1})
	if err != nil {
		log.Println("calculate err:", err)
	}
	if err == nil && len(state.States) > 0 {
		lastState = state.States[0]
	}
	// log.Println("lastState:", lastState)
	// plcs, err := c.Client.GetPlcs(machine, lastState.Time, limit)
	plcs, err := c.Client.GetPlcs(context.Background(), &model.GetPlcsRequest{Machine: machine, Time: lastState.Time, Limit: int32(limit)})
	if err != nil {
		return err
	}
	var states []*model.State
	for _, plc := range plcs.Plcs {
		state := data.NewState(plc)
		states = append(states, state)
	}
	// err = c.ApiRepository.CreateState(states)
	_, err = c.Client.CreateStates(context.Background(), &model.CreateStatesRequest{States: states})
	if err != nil {
		return err
	}
	return nil
}
