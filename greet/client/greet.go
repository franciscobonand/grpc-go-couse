package main

import (
	"context"
	"log"

	pb "github.com/franciscobonand/grpc-go-course/greet/proto"
)

func doGreet(ctx context.Context, c pb.GreetServiceClient) {
	log.Println("doGreet invoked")

	in := &pb.GreetRequest{
		FirstName: "Chico",
	}

	resp, err := c.Greet(ctx, in)
	if err != nil {
		log.Fatalf("Failed to greet server: %v\n", err)
	}

	log.Printf("Greeting: %s\n", resp.Result)
}
