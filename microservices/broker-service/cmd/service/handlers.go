package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"

	"encoding/json"

	"github.com/PhilippWegner/hochschule-bochum-masterarbeit/microservices/broker-service/model"
	"github.com/gin-gonic/gin"
	"github.com/hasura/go-graphql-client"
)

const DEFAULT_LIMIT = 1000

const DEFAULT_GET_PLCS_URL = "http://localhost:8082/query"
const DEFAULT_GET_LAST_STATE_URL = "http://localhost:8081/api/states/%s/%d"
const DEFAULT_INSERT_STATES_URL = "http://localhost:8081/api/states"

func (app *Config) Handle(ctx *gin.Context) {
	log.Println("Handle request")
	var requestPayload model.RequestPayload
	if err := ctx.ShouldBindJSON(&requestPayload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	switch requestPayload.Action {
	case "last-state":
		log.Println("last-state")
		app.LastState(ctx, requestPayload.Machine)
	case "next-plcs":
		log.Println("next-plcs")
		app.NextPlcs(ctx, requestPayload.Machine, requestPayload.LastState)
	case "insert-states":
		log.Println("insert-states")
		app.InsertStates(ctx, requestPayload.States)
	default:
		log.Println("Invalid action")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid action"})
		return
	}
}

func (app *Config) LastState(ctx *gin.Context, machine string) {
	// rest client
	request, err := http.NewRequest("GET", fmt.Sprintf(DEFAULT_GET_LAST_STATE_URL, machine, 1), nil)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)
	log.Println("response:", response.StatusCode)
	log.Println("err:", err)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer response.Body.Close()
	var states []*model.StatePayload
	err = json.NewDecoder(response.Body).Decode(&states)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, states)
}

func (app *Config) NextPlcs(ctx *gin.Context, machine string, state model.StatePayload) {
	// graphql client
	client := graphql.NewClient(DEFAULT_GET_PLCS_URL, nil)
	var query struct {
		Plcs []*model.PlcPayload `graphql:"plcs(machine: $machine, time: $time, limit: $limit, filter: {identifier: {in: $in}})"`
	}
	variables := map[string]interface{}{
		"machine": machine,
		"time":    state.Time,
		"limit":   DEFAULT_LIMIT,
		"in":      []string{"heizzeit_ist", "heizzeit_soll", "einspritzzeit_ist", "einspritzzeit_soll", "position_presse_geoeffnet"},
	}
	log.Println("variables:", variables)
	err := client.Query(context.Background(), &query, variables)
	if err != nil {
		log.Println("error:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, query.Plcs)

}

func (app *Config) InsertStates(ctx *gin.Context, states []*model.StatePayload) {
	// rest client
	statesJson, err := json.Marshal(states)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	request, err := http.NewRequest("POST", DEFAULT_INSERT_STATES_URL, bytes.NewBuffer(statesJson))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)
	log.Println("response:", response)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer response.Body.Close()
	ctx.JSON(http.StatusOK, gin.H{"message": "States inserted"})
}
