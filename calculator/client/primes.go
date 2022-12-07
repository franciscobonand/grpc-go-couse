package main

import (
	"context"
	"fmt"
	"io"
	"log"

	pb "github.com/franciscobonand/grpc-go-course/calculator/proto"
)

func doPrimes(ctx context.Context, c pb.CalculatorServiceClient) {
	log.Println("doPrimes invoked")

	in := &pb.PrimeRequest{
		Num: 120,
	}

	stream, err := c.Primes(ctx, in)
	if err != nil {
		log.Fatalf("Failed to get primes decomposition: %v\n", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error receiving stream response: %v\n", err)
		}
		fmt.Printf("Primes: %d\n", msg.Result)
	}
}
