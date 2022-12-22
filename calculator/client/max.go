package main

import (
	"context"
	"io"
	"log"

	pb "github.com/franciscobonand/grpc-go-course/calculator/proto"
)

func doMax(ctx context.Context, c pb.CalculatorServiceClient) {
	log.Println("doMax invoked")

	nums := []int32{1, 5, 3, 6, 2, 20}

	stream, err := c.Max(ctx)
	if err != nil {
		log.Fatalf("Failed calling Max: %v\n", err)
	}

	go func() {
		for _, num := range nums {
			req := &pb.MaxRequest{
				Num: num,
			}
			err = stream.Send(req)
			if err != nil {
				log.Fatalf("Error while sending client stream: %v\n", err)
			}
		}
		err := stream.CloseSend()
		if err != nil {
			log.Fatalf("Error closing send stream: %v\n", err)
		}
	}()

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Error reading server response stream: %v\n", err)
			break
		}

		log.Printf("Max: %d\n", res.Max)
	}
}
