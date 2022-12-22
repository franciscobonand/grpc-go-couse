package main

import (
	"io"
	"log"

	pb "github.com/franciscobonand/grpc-go-course/calculator/proto"
)

func (s *Server) Average(stream pb.CalculatorService_AverageServer) error {
	log.Print("Average function invoked")
	nums := []int32{}

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			avg := getAvg(nums)
			return stream.SendAndClose(&pb.AvgResponse{Result: avg})
		}
		if err != nil {
			log.Fatalf("Error reading client stream: %v\n", err)
		}
		nums = append(nums, req.Num)
	}
}

func getAvg(nums []int32) float32 {
	var sum int32
	length := float32(len(nums))
	for _, num := range nums {
		sum += num
	}
	return float32(sum) / length
}
