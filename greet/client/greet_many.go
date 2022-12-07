package main

import (
	"context"
	"fmt"
	"io"
	"log"

	pb "github.com/franciscobonand/grpc-go-course/greet/proto"
)

func doGreetMany(ctx context.Context, c pb.GreetServiceClient) {
	log.Println("doGreetMany invoked")

	in := &pb.GreetRequest{
		FirstName: "Chico",
	}

	stream, err := c.GreetManyTimes(ctx, in)
	if err != nil {
		log.Fatalf("Failed to greet many server: %v\n", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error receiving stream response: %v\n", err)
		}
		fmt.Printf("Greet Many: %s\n", msg.Result)
	}
}
