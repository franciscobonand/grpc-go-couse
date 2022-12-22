package main

import (
	"context"
	"fmt"
	"log"
	"math"

	pb "github.com/franciscobonand/grpc-go-course/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) Sqrt(ctx context.Context, in *pb.SqrtRequest) (*pb.SqrtResponse, error) {
	log.Printf("Sqrt invoked with %v\n", in)

	if in.Number < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Received negative number '%d'", in.Number),
		)
	}

	return &pb.SqrtResponse{
		Result: math.Sqrt(float64(in.Number)),
	}, nil
}
