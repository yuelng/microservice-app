package handlers

import (
	"fmt"
	"golang.org/x/net/context"
	pb "base/protos/helloworld"
	//"google.golang.org/grpc/metadata"
	"golang.org/x/crypto/bcrypt"
	"encoding/json"
	"golang.org/x/tools/go/gcimporter15/testdata"
)

// server is used to implement helloworld.GreeterServer.
type Server struct{}

// SayHello implements helloworld.GreeterServer
func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	// md, _ := metadata.FromContext(ctx)
	fmt.Println("hello from server,hello")
	// bcrypt
	// passwordErr := bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(json.Password))
	// bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return &pb.HelloReply{Message: "Hello " + in.Name+in.Num}, nil
}

func (s *Server) SayHelloAgain(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	fmt.Println("hello from server,hello")
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}
