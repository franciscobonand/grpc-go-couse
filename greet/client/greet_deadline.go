package main

import (
	"context"
	"log"
	"time"

	pb "github.com/franciscobonand/grpc-go-course/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doGreetDeadline(ctx context.Context, client pb.GreetServiceClient, timeout time.Duration) {
	log.Println("doGreetDeadline invoked")
	c, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	in := &pb.GreetRequest{
		FirstName: "Chico",
	}

	resp, err := client.GreetWithDeadline(c, in)
	if err != nil {
		e, ok := status.FromError(err)
		if !ok {
			log.Fatalf("Non-gRPC error greeting server: %v\n", err)
		}

		if e.Code() == codes.DeadlineExceeded {
			log.Println("Deadline exceeded")
			return
		}
		log.Fatalf("Unexpected gRPC error: %v\n", e)
	}

	log.Printf("Greeting with Deadline: %s\n", resp.Result)
}
