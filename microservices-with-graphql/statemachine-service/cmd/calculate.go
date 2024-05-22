package main

import (
	"log"

	"github.com/PhilippWegner/hochschule-bochum-masterarbeit/microservices-with-graphql/statemachine-service/data"
)

func (c *Config) calculate(machine string) error {
	plcs, err := c.ApiRepository.GetPlcs(machine, "0", limit)
	if err != nil {
		return err
	}
	var states []*data.State
	for _, plc := range plcs {
		state := data.NewState(plc)
		states = append(states, state)
	}
	for _, state := range states {
		log.Println(state)
	}
	return nil
}
