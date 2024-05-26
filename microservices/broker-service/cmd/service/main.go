package main

import (
	"log"
	"net/http"
	"time"
)

const webPort = "8080"

type Config struct{}

func main() {
	app := Config{}

	server := &http.Server{
		Addr:         ":" + webPort,
		Handler:      app.routes(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
