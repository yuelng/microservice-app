package main

import (
	"log"
	"os"

	pb "base/protos/helloworld"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"fmt"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
)

const (
	address     = "localhost:50000"
	defaultName = "world"
)

func main() {
	context.Background()
	// Set up a connection to the server.
	conn, err := grpc.Dial(address,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_prometheus.UnaryClientInterceptor),
		grpc.WithStreamInterceptor(grpc_prometheus.StreamClientInterceptor))
	fmt.Println(conn)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name, Num: "2"})
	r1, err := c.SayHelloAgain(context.Background(), &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Println(r1.Message)
	log.Printf("Greeting: %s", r.Message)

}
