package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/franciscobonand/grpc-go-course/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ReadBlog(ctx context.Context, in *pb.BlogId) (*pb.Blog, error) {
	log.Printf("Read Blog invoked with: %v\n", in)

	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Invalid ID: %v\n", oid),
		)
	}

	data := &BlogItem{}
	filter := bson.M{"_id": oid}
	resp := s.coll.FindOne(ctx, filter)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Internal error: %v\n", err))
	}

	if err := resp.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			"Blog not found",
		)
	}

	return documentToBlog(data), nil
}
