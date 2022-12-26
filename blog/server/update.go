package main

import (
	"context"
	"log"

	pb "github.com/franciscobonand/grpc-go-course/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) UpdateBlog(ctx context.Context, in *pb.Blog) (*emptypb.Empty, error) {
	log.Printf("Update Blog invoked with: %v\n", in)

	data, err := blogToDocument(in)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Cannot parse provided ID",
		)
	}

	resp, err := s.coll.UpdateOne(
		ctx,
		bson.M{"_id": data.ID},
		bson.M{"$set": data},
	)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"Could not update",
		)
	}
	if resp.MatchedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			"Blog with provided OID not found",
		)
	}

	return &emptypb.Empty{}, nil
}
