package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/franciscobonand/grpc-go-course/greet/proto"
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

	c := pb.NewGreetServiceClient(conn)

	// doGreet(ctx, c)
	// doGreetMany(ctx, c)
	// doLongGreet(ctx, c)
	// doGreetEveryone(ctx, c)
	doGreetDeadline(ctx, c, 5*time.Second)
	doGreetDeadline(ctx, c, time.Second)
}
