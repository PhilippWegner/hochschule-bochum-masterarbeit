package data

import (
	"context"
	"fmt"
	"log"

	"github.com/hasura/go-graphql-client"
)

type Plc struct {
	Time       string
	Machine    string
	Identifier []*Identifier
}

type Identifier struct {
	Name  string
	Value float64
}

func (identifier Identifier) String() string {
	return fmt.Sprintf("Name: %s, Value: %f", identifier.Name, identifier.Value)
}

func (plc Plc) String() string {
	return fmt.Sprintf("{Time: %s, Machine: %s, Identifier: %v}", plc.Time, plc.Machine, plc.Identifier)
}

type State struct {
	Time    string
	Machine string
	Name    string
	Color   string
	Value   int
}

type ApiRepository struct {
	graphql_api string
}

func NewApiRepository(graphql_api string) *ApiRepository {
	return &ApiRepository{graphql_api: graphql_api}
}

func (r *ApiRepository) GetPlcs(machine string, time string, limit int) ([]*Plc, error) {
	log.Println("GetPlcs")
	log.Println(r.graphql_api)
	client := graphql.NewClient(r.graphql_api, nil)
	var query struct {
		Plcs []*Plc `graphql:"plcs(machine: $machine, time: $time, limit: $limit, filter: {identifier: {in: $in}})"`
	}
	log.Println(query.Plcs)
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
	for _, plc := range query.Plcs {
		log.Println(plc)
	}
	return query.Plcs, nil
}

func (r *ApiRepository) GetStates(machine string, limit int) ([]*State, error) {
	log.Println("GetStates")
	client := graphql.NewClient(r.graphql_api, nil)
	var query struct {
		States []*State `graphql:"states(machine: $machine, limit: $limit)"`
	}
	log.Println(query.States)
	variables := map[string]interface{}{
		"machine": machine,
		"limit":   limit,
	}
	err := client.Query(context.Background(), &query, variables)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	for _, state := range query.States {
		log.Println(state)
	}
	return query.States, nil
}

func (r *ApiRepository) CreateState(state []*State) error {
	client := graphql.NewClient(r.graphql_api, nil)
	var mutation struct {
		CreateStates []*State `graphql:"createStates(input: $states)"`
	}
	variables := map[string]interface{}{
		"states": state,
	}
	err := client.Mutate(context.Background(), &mutation, variables)
	if err != nil {
		return err
	}
	log.Println(mutation.CreateStates)
	return nil
}
