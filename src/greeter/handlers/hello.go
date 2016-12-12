package handlers

import (
	"fmt"
	"golang.org/x/net/context"
	pb "base/protos/helloworld"
	"google.golang.org/grpc/metadata"
)

// server is used to implement helloworld.GreeterServer.
type Server struct{}

// SayHello implements helloworld.GreeterServer
func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	md, _ := metadata.FromContext(ctx)
	fmt.Println(md["key1"])
	fmt.Println("hello from server,hello")
	return &pb.HelloReply{Message: "Hello " + in.Name+in.Num}, nil
}

func (s *Server) SayHelloAgain(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	fmt.Println("hello from server,hello")
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}
