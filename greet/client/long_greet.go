package main

import (
	"context"
	"log"
	"time"

	pb "github.com/franciscobonand/grpc-go-course/greet/proto"
)

func doLongGreet(ctx context.Context, c pb.GreetServiceClient) {
	log.Println("doLongGreet invoked")

	names := []string{"Chico", "Roberto", "Matheus", "Amanda", "Maria"}

	stream, err := c.LongGreet(ctx)
	if err != nil {
		log.Fatalf("Failed creating stream Long Greet: %v\n", err)
	}

	for _, name := range names {
		in := &pb.GreetRequest{
			FirstName: name,
		}
		log.Printf("Sending req: %v\n", in)
		err := stream.Send(in)
		if err != nil {
			log.Fatalf("Error while sending client stream: %v\n", err)
		}
		time.Sleep(time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error receiving response from Long Greet: %v\n", err)
	}

	log.Printf("Long Greet: %s\n", res.Result)
}
