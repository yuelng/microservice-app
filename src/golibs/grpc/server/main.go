package main

import (
	"log"
	"net"

	"fmt"
	"golang.org/x/net/context"
	pb "golibs/grpc/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	md, _ := metadata.FromContext(ctx)
	fmt.Println(md["key1"])
	fmt.Println("hello from server")
	return &pb.HelloReply{Message: "Hello " + in.Name + in.Num}, nil
}

func (s *server) SayHelloAgain(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	fmt.Println("hello from server")
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
