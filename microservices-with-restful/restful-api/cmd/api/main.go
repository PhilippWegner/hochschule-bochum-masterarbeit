package main

import (
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/PhilippWegner/hochschule-bochum-masterarbeit/microservices-with-restful/restful-api/docs"

	"github.com/PhilippWegner/hochschule-bochum-masterarbeit/microservices-with-restful/restful-api/database"
)

const defaultPort = "8080"

type Config struct {
	Repo database.Repository
}

// @title RESTful API
// @version 1.0
// @description This is a RESTful API for a plc and state data.
// @host localhost:8080
// @BasePath /api
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	app := Config{}
	app.Repo = database.ConnectInfluxdb()

	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      app.routes(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
