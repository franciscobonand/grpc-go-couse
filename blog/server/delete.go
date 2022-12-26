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
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) DeleteBlog(ctx context.Context, in *pb.BlogId) (*emptypb.Empty, error) {
	log.Printf("Delete Blog invoked with: %v\n", in)

	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Invalid ID: %v\n", oid),
		)
	}

	filter := bson.M{"_id": oid}
	resp, err := s.coll.DeleteOne(ctx, filter)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Internal error: %v\n", err))
	}

	if resp.DeletedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			"Blog with provided OID not found",
		)
	}

	return &emptypb.Empty{}, nil
}
