package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/PhilippWegner/hochschule-bochum-masterarbeit/microservices-with-restful/restful-api/database"
)

const defaultPort = "8080"

type Config struct {
	Repo database.Repository
}

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
