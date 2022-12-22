package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/franciscobonand/grpc-go-course/calculator/proto"
)

const (
	addr = "0.0.0.0:50051"
)

func main() {
	ctx := context.Background()
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	defer conn.Close()

	c := pb.NewCalculatorServiceClient(conn)

	// doSum(ctx, c)
	// doPrimes(ctx, c)
	// doAvg(ctx, c)
	// doMax(ctx, c)
	doSqrt(ctx, 9, c)
	doSqrt(ctx, -25, c)
}
