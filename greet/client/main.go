package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/franciscobonand/grpc-go-course/greet/proto"
)

const (
	addr = "0.0.0.0:50051"
)

func main() {
	ctx := context.Background()
	tls := false
	opts := []grpc.DialOption{}

	if tls {
		certFile := "ssl/ca.crt"
		creds, err := credentials.NewClientTLSFromFile(certFile, "")
		if err != nil {
			log.Fatalf("Failed to load CA trust certificate: %v\n", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	}

	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)

	doGreet(ctx, c)
	// doGreetMany(ctx, c)
	// doLongGreet(ctx, c)
	// doGreetEveryone(ctx, c)
	// doGreetDeadline(ctx, c, 5*time.Second)
	// doGreetDeadline(ctx, c, time.Second)
}
