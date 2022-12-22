package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/franciscobonand/grpc-go-course/greet/proto"
)

func doGreetEveryone(ctx context.Context, c pb.GreetServiceClient) {
	log.Println("doGreetEveryone invoked")

	names := []string{"Chico", "Roberto", "Matheus", "Amanda", "Maria"}

	stream, err := c.GreetEveryone(ctx)
	if err != nil {
		log.Fatalf("Failed creating stream Greet Everyone: %v\n", err)
	}

	go func() {
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
			log.Printf("Error receiving response from Greet Everyone: %v\n", err)
			break
		}

		log.Printf("Greet Everyone: %s\n", res.Result)
	}
}
