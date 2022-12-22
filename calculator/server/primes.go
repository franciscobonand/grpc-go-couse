package main

import (
	"log"

	pb "github.com/franciscobonand/grpc-go-course/calculator/proto"
)

func (s *Server) Primes(in *pb.PrimeRequest, stream pb.CalculatorService_PrimesServer) error {
	log.Printf("Primes function invoked with %v\n", in)

	primes := getDecomposed(in.Num)

	for _, n := range primes {
		resp := &pb.PrimeResponse{
			Result: n,
		}
		if err := stream.Send(resp); err != nil {
			log.Fatalf("Error sending prime stream: %v\n", err)
		}
	}

	return nil
}

func getDecomposed(num int32) []int32 {
	var primes []int32
	var k int32 = 2
	n := num
	for n > 1 {
		if n%k == 0 {
			primes = append(primes, k)
			n = n / k
			continue
		}
		k++
	}
	return primes
}
