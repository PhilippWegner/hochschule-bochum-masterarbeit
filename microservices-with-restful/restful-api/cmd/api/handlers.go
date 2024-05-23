package main

import (
	"net/http"
	"strconv"

	"github.com/PhilippWegner/hochschule-bochum-masterarbeit/microservices-with-restful/restful-api/model"
	"github.com/gin-gonic/gin"
)

func (app *Config) CreateStates(ctx *gin.Context) {
	var states []*model.State
	err := ctx.BindJSON(&states)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = app.Repo.CreateStates(states)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (app *Config) GetStates(ctx *gin.Context) {
	machine := ctx.Param("machine")
	limit, _ := strconv.Atoi(ctx.Param("limit"))
	states, err := app.Repo.GetStates(machine, limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, states)
}

func (app *Config) GetPlcs(ctx *gin.Context) {
	machine := ctx.Param("machine")
	time := ctx.Param("time")
	limit, _ := strconv.Atoi(ctx.Param("limit"))
	plcs, err := app.Repo.GetPlcs(machine, time, limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, plcs)
}
