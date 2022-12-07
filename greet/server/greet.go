package main

import (
	"context"
	"log"

	pb "github.com/franciscobonand/grpc-go-course/greet/proto"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet function invoked with %v\n", in)

	return &pb.GreetResponse{
		Result: "Heyo " + in.FirstName,
	}, nil
}
