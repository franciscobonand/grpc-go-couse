package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/franciscobonand/grpc-go-course/greet/proto"
)

func (s *Server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {
	log.Println("GreetEveryone function invoked")

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error reading client stream: %v\n", err)
		}

		log.Printf("Receiving: %s\n", req.FirstName)

		err = stream.Send(&pb.GreetResponse{
			Result: fmt.Sprintf("Hello %s!\n", req.FirstName),
		})
		if err != nil {
			log.Fatalf("Error sending response to client: %v\n", err)
		}
	}

	return nil
}
