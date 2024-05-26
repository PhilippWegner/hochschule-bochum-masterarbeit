package main

import "github.com/gin-gonic/gin"

func (app *Config) routes() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/handle", app.Handle)
	}

	return router
}
