package main

import (
	"io"
	"log"
	"math"

	pb "github.com/franciscobonand/grpc-go-course/calculator/proto"
)

func (s *Server) Max(stream pb.CalculatorService_MaxServer) error {
	log.Print("Max function invoked")
	max := int32(math.MinInt32)

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Error reading client stream: %v\n", err)
			return err
		}
		if req.Num > max {
			max = req.Num
			err = stream.Send(&pb.MaxResponse{Max: max})
			if err != nil {
				log.Printf("Error sending response: %v\n", err)
				return err
			}
		}
	}
	return nil
}
