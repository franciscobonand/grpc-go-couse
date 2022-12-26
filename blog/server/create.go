package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/franciscobonand/grpc-go-course/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateBlog(ctx context.Context, in *pb.Blog) (*pb.BlogId, error) {
	log.Printf("Create Blog invoked with: %v\n", in)

	data, _ := blogToDocument(in) // Structure doesn't contain an ID, so no error can occur

	resp, err := s.coll.InsertOne(ctx, data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Internal error: %v\n", err))
	}

	id, ok := resp.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Cannot convert to OID: %v\n", err))
	}

	return &pb.BlogId{
		Id: id.Hex(),
	}, nil
}
