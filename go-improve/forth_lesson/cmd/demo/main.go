package main

import (
	"google.golang.org/grpc/reflection"
	"net"
)

import (
	"golang.org/x/net/context"
	//"google.golang.org/grpc/reflection"
	"log"
	//"net"
	//"net/http"

	pb "github.com/gin-gonic/examples/grpc/pb"
	//"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type server struct {
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
