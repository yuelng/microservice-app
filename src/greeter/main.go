// 认证 jwt logging
// handler service model util
package main

import (
	"log"
	"net"

	"greeter/handlers"

	pb "base/protos/helloworld"
	"google.golang.org/grpc"
)

const (
	grpcPort = ":50000"
)

func main() {
	// start grpc server
	server := &handlers.Server{}

	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterGreeterServer(s, server)
	log.Println("start grpc server listen " + grpcPort)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
