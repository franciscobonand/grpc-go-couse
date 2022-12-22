package main

import (
	"context"
	"log"

	pb "github.com/franciscobonand/grpc-go-course/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doSqrt(ctx context.Context, n int32, c pb.CalculatorServiceClient) {
	log.Println("doSqrt invoked")
	in := &pb.SqrtRequest{
		Number: n,
	}

	resp, err := c.Sqrt(ctx, in)
	if err != nil {
		e, ok := status.FromError(err)
		if !ok {
			log.Fatalf("Non-gRPC error: %v\n", err)
		}
		log.Printf("Error msg from server: %s\n", e.Message())
		log.Printf("Error code from server: %s\n", e.Code())
		if e.Code() == codes.InvalidArgument {
			log.Println("Probably sent negative number :P")
			return
		}
	}

	log.Printf("Sqrt result: %.2f\n", resp.Result)
}
