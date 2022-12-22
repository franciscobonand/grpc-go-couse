package main

import (
	"context"
	"log"

	pb "github.com/franciscobonand/grpc-go-course/calculator/proto"
)

func doAvg(ctx context.Context, c pb.CalculatorServiceClient) {
	log.Println("doAvg invoked")

	nums := []int32{1, 2, 3, 4, 5, 5, 8}

	stream, err := c.Average(ctx)
	if err != nil {
		log.Fatalf("Failed calling Average: %v\n", err)
	}

	for _, num := range nums {
		req := &pb.AvgRequest{
			Num: num,
		}
		err = stream.Send(req)
		if err != nil {
			log.Fatalf("Error while sending client stream: %v\n", err)
		}
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error receiving response from Average: %v\n", err)
	}

	log.Printf("Average result: %d\n", resp.Result)
}
