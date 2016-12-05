package handlers

import (

	"golang.org/x/net/websocket"
	"github.com/gin-gonic/gin"


	"io"
)

func WebSocket(c *gin.Context) {
	handler := websocket.Handler(EchoServer)
    	handler.ServeHTTP( c.Writer,c.Request)

}

func EchoServer(conn *websocket.Conn)  {
	io.Copy(conn, conn)
}