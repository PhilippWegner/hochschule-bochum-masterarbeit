package main

import (
	"net/http"
	"strconv"

	"github.com/PhilippWegner/hochschule-bochum-masterarbeit/microservices-with-restful/restful-api/model"
	"github.com/gin-gonic/gin"
)

// CreateStates godoc
// @Summary Create states
// @Description Create states
// @Tags state
// @Accept  application/json
// @Produce  application/json
// @Param states body []model.State true "States"
// @Success 200 {object} string "status"
// @Failure 500 {object} string "error"
// @Router /states [post]
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

// GetStates godoc
// @Summary Get state data
// @Description Get state data
// @Tags state
// @Produce  application/json
// @Param machine path string true "Machine"
// @Param limit path int true "Limit"
// @Success 200 {array} model.State
// @Failure 500 {object} string "error"
// @Router /states/{machine}/{limit} [get]
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

// GetPlcs godoc
// @Summary Get plc data
// @Description Get plc data
// @Tags plc
// @Produce  application/json
// @Param machine path string true "Machine"
// @Param time path string true "Time"
// @Param limit path int true "Limit"
// @Success 200 {array} model.Plc
// @Failure 500 {object} string "error"
// @Router /plcs/{machine}/{time}/{limit} [get]
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
