package app

import (

	//jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"api/handlers/web"
	"fmt"
	"golang.org/x/net/websocket"
	"strings"
)

func App() *gin.Engine {
	app := gin.Default()

	// Set 405 no method true
	// Reference: https://github.com/gin-gonic/gin/blob/develop/gin.go
	app.HandleMethodNotAllowed = true

	// Web API
	webAPI := app.Group("/api/web")
	webAPI.POST("/submission", web.SubmissionAdd)

	app.GET("/entry", func(c *gin.Context) {
		handler := websocket.Handler(countServer)
		handler.ServeHTTP(c.Writer, c.Request)
	})

	return app
}

type Count struct {
	Author string `json:"author"`
	Body   string `json:"body"`
}

func countServer(ws *websocket.Conn) {
	defer ws.Close()
	for {
		var count Count
		err := websocket.JSON.Receive(ws, &count)
		if err != nil {
			return
		}

		fmt.Println(count.Author)
		fmt.Println(count.Body)

		err = websocket.JSON.Send(ws, count)
		if err != nil {
			return
		}
	}
}
