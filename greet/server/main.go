package main

import (
	"log"
	"net"

	pb "github.com/franciscobonand/grpc-go-course/greet/proto"
	"google.golang.org/grpc"
)

const (
	addr = "0.0.0.0:50051"
)

type Server struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)
	srv := grpc.NewServer()

	if err = srv.Serve(lis); err != nil {
		log.Fatalf("Server error: %v\n", err)
	}
}
