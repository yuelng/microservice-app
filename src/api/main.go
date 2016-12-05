package main

import (
	"api/handlers"
	"api/migration"
	"api/models"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"base/rpc"
	pb "base/protos/helloworld"
)

const (
	address     = "localhost:50051"
)
func init() {
	migration.CreateDatabase()
	db := models.InitDB()
	defer db.Close()

	db.InitSchema()
	db.Seed()
}

func main() {
	p := fmt.Println

	// flag 解析命令行参数
	// flag.String("e", "default", "help message")
	// 使用 flag.String()  Bool(), Int()
	// 最后调用flag.Parse() 对命令行参数进行解析
	enviroment := flag.String("e", "development", "")
	flag.Parse()
	p(*enviroment)

	r := gin.Default()
	// Ping test
	r.GET("/ping", func(c *gin.Context) {

		c.String(200, "pong")
	})

	// Ping test
	r.GET("/hello", handlers.Hello)

	// Ping test
	r.GET("/location", handlers.Location)

	r.GET("/ws", handlers.WebSocket)

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")

	rpc.StartServiceConns("","")
	defer conn.Close()
	c := pb.NewGreeterClient(conn)
}
