package main

import (
	"log"
	"net"

	pb "github.com/franciscobonand/grpc-go-course/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	addr = "0.0.0.0:50051"
)

type Server struct {
	pb.CalculatorServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)
	srv := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(srv, &Server{})
	reflection.Register(srv)

	if err = srv.Serve(lis); err != nil {
		log.Fatalf("Server error: %v\n", err)
	}
}
