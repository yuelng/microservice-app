package handlers

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/net/websocket"

	"fmt"
	"io"
)

func WebSocket(c *gin.Context) {
	handler := websocket.Handler(EchoServer)
	handler.ServeHTTP(c.Writer, c.Request)

}

func EchoServer(conn *websocket.Conn) {
	io.Copy(conn, conn)
}

func WebSocket2(c *gin.Context) {
	handler := websocket.Handler(countServer)
	handler.ServeHTTP(c.Writer, c.Request)
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
