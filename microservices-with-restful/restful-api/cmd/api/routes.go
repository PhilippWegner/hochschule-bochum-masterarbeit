package main

import "github.com/gin-gonic/gin"

func (app *Config) routes() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.GET("/plcs/:machine/:time/:limit", app.GetPlcs)
		api.GET("/states/:machine/:limit", app.GetStates)
		api.POST("/states", app.CreateStates)
	}
	return router
}
