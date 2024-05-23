package main

import (
	"log"
	"sync"
	"time"

	"github.com/PhilippWegner/hochschule-bochum-masterarbeit/microservices-with-grpc/statemachine-service/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	grpc_api = "localhost:8080"
	machines = []string{"presse_11"}
	limit    = 10000
)

type Config struct {
	Client   model.ModelServiceClient
	Machines []string
}

func main() {
	app := Config{}
	conn, err := grpc.NewClient(grpc_api, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := model.NewModelServiceClient(conn)
	app.Client = client
	app.Machines = machines

	// Start the service
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
