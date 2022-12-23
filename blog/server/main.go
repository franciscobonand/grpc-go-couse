package main

import (
	"context"
	"log"
	"net"

	pb "github.com/franciscobonand/grpc-go-course/blog/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

const (
	addr = "0.0.0.0:50051"
)

type Server struct {
	pb.BlogServiceServer
	coll *mongo.Collection
}

func main() {
	ctx := context.Background()
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	coll, err := newMongoClient(ctx)
	if err != nil {
		log.Fatalf("Mongo err: %v\n", err)
	}

	srv := Server{
		coll: coll,
	}
	grpcSrv := grpc.NewServer()
	pb.RegisterBlogServiceServer(grpcSrv, &srv)

	log.Printf("Listening on %s\n", addr)
	if err = grpcSrv.Serve(lis); err != nil {
		log.Fatalf("Server error: %v\n", err)
	}
}
