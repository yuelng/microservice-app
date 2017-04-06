package app

import (

	//jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"api/handlers"

	"api/handlers/web"
)

func App() *gin.Engine {
	app := gin.Default()

	// Set 405 no method true
	// Reference: https://github.com/gin-gonic/gin/blob/develop/gin.go
	app.HandleMethodNotAllowed = true

	// Web API
	webAPI := app.Group("/api/web")
	webAPI.POST("/submission", web.SubmissionAdd)

	app.GET("/entry", handlers.WebSocket2)

	return app
}
