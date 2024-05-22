package main

import (
	"log"

	"github.com/PhilippWegner/hochschule-bochum-masterarbeit/microservices-with-graphql/statemachine-service/data"
)

var (
	graphql_api = "http://localhost:8080/query"
	machines    = []string{"presse_11"}
	limit       = 10
)

type Config struct {
	ApiRepository data.Repository
	Machines      []string
}

func main() {
	log.Println("Starting application")
	app := Config{}
	app.ApiRepository = data.NewApiRepository(graphql_api)
	app.Machines = machines

	for _, machine := range app.Machines {
		err := app.calculate(machine)
		if err != nil {
			log.Println(err)
		}
	}
}
