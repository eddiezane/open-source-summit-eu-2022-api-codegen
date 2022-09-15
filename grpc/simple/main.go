package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/eddiezane/open-source-summit-eu-2022-api-codegen/grpc/simple/hello"
)

type server struct {
	pb.UnimplementedGreeterServiceServer
}

func (s *server) Greet(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Greeting: fmt.Sprintf("Hello, %s", req.GetName())}, nil
}

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterGreeterServiceServer(s, &server{})

	log.Println("starting server...")
	log.Fatal(s.Serve(l))
}
