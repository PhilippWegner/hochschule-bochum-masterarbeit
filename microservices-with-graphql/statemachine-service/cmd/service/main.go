package main

import (
	"log"
	"sync"
	"time"

	"github.com/PhilippWegner/hochschule-bochum-masterarbeit/microservices-with-graphql/statemachine-service/data"
)

var (
	graphql_api = "http://localhost:8080/query"
	machines    = []string{"presse_11"}
	limit       = 10000
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

	log.Println("Starting looprunner")
	for {
		start := time.Now()
		app.looprunner()
		duration := time.Since(start)
		log.Printf("looprunner took %v\n", duration)
	}
}

func (c *Config) looprunner() {
	var wg sync.WaitGroup
	wg.Add(len(c.Machines))
	for _, machine := range c.Machines {
		go func(machine string) {
			defer wg.Done()
			err := c.calculate(machine)
			if err != nil {
				log.Printf("calculate(%v) failed: %v\n", machine, err)
			}
		}(machine)
	}
	// Wait for all goroutines to finish
	wg.Wait()
}
