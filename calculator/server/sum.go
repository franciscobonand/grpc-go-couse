package main

import (
	"context"
	"log"

	pb "github.com/franciscobonand/grpc-go-course/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Sum function invoked with %v\n", in)

	return &pb.SumResponse{
		Result: in.FirstNum + in.SecondNum,
	}, nil
}
