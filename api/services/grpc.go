package services

import (
	"base/protos/helloworld"
	"google.golang.org/grpc"
	"log"
	"sync"
)

const (
	//address     = "greeter:50000"
	greeterService = "localhost:50000"
	defaultName    = "world"
)

// 多个client同时用一个struct处理,集合的好处
type Client struct {
	helloworld.GreeterClient
}

// mustDial ensures a tcp connection to specified address.
func mustDial(addr *string) *grpc.ClientConn {
	conn, err := grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
		panic(err)
	}
	return conn
}

func InitGrpc() Client {
	var c Client
	// client with all grpc connections
	once.Do(func() {
		c = Client{
			GreeterClient: helloworld.NewGreeterClient((mustDial(greeterService))),
		}
	})

	return c
}
