package main

import (
	"context"
	"log"

	pb "github.com/franciscobonand/grpc-go-course/calculator/proto"
)

func doSum(ctx context.Context, c pb.CalculatorServiceClient) {
	log.Println("doSum invoked")

	in := &pb.SumRequest{
		FirstNum:  4,
		SecondNum: 5,
	}

	resp, err := c.Sum(ctx, in)
	if err != nil {
		log.Fatalf("Failed to sum: %v\n", err)
	}

	log.Printf("Sum result: %d\n", resp.Result)
}
