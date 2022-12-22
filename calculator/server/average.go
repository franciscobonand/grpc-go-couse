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
		nums = append(nums, req.Num)
	}
}

func getAvg(nums []int32) int32 {
	var sum int32
	length := int32(len(nums))
	for _, num := range nums {
		sum += num
	}
	return sum / length
}
