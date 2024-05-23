package main

import (
	"log"

	"github.com/PhilippWegner/hochschule-bochum-masterarbeit/microservices-with-graphql/statemachine-service/data"
)

var DEFAULT_LAST_STATE = data.State{Time: "0"}

func (c *Config) calculate(machine string) error {
	// get last states
	lastState := DEFAULT_LAST_STATE
	state, err := c.ApiRepository.GetStates(machine, 1)
	// log.Println("calculate:", state)
	if err != nil {
		log.Println("calculate err:", err)
	}
	if len(state) > 0 {
		lastState = *state[0]
	}
	// log.Println("lastState:", lastState)
	plcs, err := c.ApiRepository.GetPlcs(machine, lastState.Time, limit)
	if err != nil {
		return err
	}
	var states []*data.State
	for _, plc := range plcs {
		state := data.NewState(plc)
		states = append(states, state)
	}
	var createStatesInput []*data.CreateStatesInput
	for _, state := range states {
		createStatesInput = append(createStatesInput, data.NewCreateStatesInput(state))
	}
	err = c.ApiRepository.CreateState(createStatesInput)
	if err != nil {
		return err
	}
	return nil
}
