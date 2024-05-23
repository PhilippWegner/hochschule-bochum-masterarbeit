package data

import (
	"context"
	"fmt"
	"log"

	"github.com/hasura/go-graphql-client"
)

type Plc struct {
	Time       string `json:"time"`
	Machine    string `json:"machine"`
	Identifier []*Identifier
}

type Identifier struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

func (identifier Identifier) String() string {
	return fmt.Sprintf("Name: %s, Value: %f", identifier.Name, identifier.Value)
}

func (plc Plc) String() string {
	return fmt.Sprintf("{Time: %s, Machine: %s, Identifier: %v}", plc.Time, plc.Machine, plc.Identifier)
}

type State struct {
	Time    string `json:"time"`
	Machine string `json:"machine"`
	Name    string `json:"name"`
	Color   string `json:"color"`
	Value   int    `json:"value"`
}

type CreateStatesInput struct {
	Time    string `json:"time"`
	Machine string `json:"machine"`
	Name    string `json:"name"`
	Color   string `json:"color"`
	Value   int    `json:"value"`
}

func NewCreateStatesInput(state *State) *CreateStatesInput {
	return &CreateStatesInput{
		Time:    state.Time,
		Machine: state.Machine,
		Name:    state.Name,
		Color:   state.Color,
		Value:   state.Value,
	}
}

type ApiRepository struct {
	graphql_api string
}

func NewApiRepository(graphql_api string) *ApiRepository {
	return &ApiRepository{graphql_api: graphql_api}
}

func (r *ApiRepository) GetPlcs(machine string, time string, limit int) ([]*Plc, error) {
	client := graphql.NewClient(r.graphql_api, nil)
	var query struct {
		Plcs []*Plc `graphql:"plcs(machine: $machine, time: $time, limit: $limit, filter: {identifier: {in: $in}})"`
	}
	variables := map[string]interface{}{
		"machine": machine,
		"time":    time,
		"limit":   limit,
		"in":      []string{"heizzeit_ist", "heizzeit_soll", "einspritzzeit_ist", "einspritzzeit_soll", "position_presse_geoeffnet"},
	}
	err := client.Query(context.Background(), &query, variables)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return query.Plcs, nil
}

func (r *ApiRepository) GetStates(machine string, limit int) ([]*State, error) {
	client := graphql.NewClient(r.graphql_api, nil)
	var query struct {
		States []*State `graphql:"states(machine: $machine, limit: $limit)"`
	}
	variables := map[string]interface{}{
		"machine": machine,
		"limit":   limit,
	}
	err := client.Query(context.Background(), &query, variables)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return query.States, nil
}

func (r *ApiRepository) CreateState(statesInput []*CreateStatesInput) error {
	// log count of statesInput
	client := graphql.NewClient(r.graphql_api, nil)
	var mutation struct {
		CreateStates []*State `graphql:"createStates(input: $input)"`
	}
	variables := map[string]interface{}{
		"input": statesInput,
	}
	err := client.Mutate(context.Background(), &mutation, variables)
	if err != nil {
		return err
	}
	return nil
}
