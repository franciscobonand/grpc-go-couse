package main

import (
	"fmt"
	"log"

	pb "github.com/franciscobonand/grpc-go-course/greet/proto"
)

func (s *Server) GreetManyTimes(in *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	log.Printf("GreetManyTimes function invoked with: %v\n", in)

	for i := 0; i < 10; i++ {
		resp := &pb.GreetResponse{
			Result: fmt.Sprintf("Hello %s, for the %d time!", in.FirstName, i),
		}

		if err := stream.Send(resp); err != nil {
			return err
		}
	}

	return nil
}
