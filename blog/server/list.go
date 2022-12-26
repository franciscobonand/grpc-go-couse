package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/franciscobonand/grpc-go-course/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) ListBlogs(in *emptypb.Empty, stream pb.BlogService_ListBlogsServer) error {
	log.Println("List Blogs invoked")
	ctx := context.Background()

	reader, err := s.coll.Find(ctx, bson.D{})
	if err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Internal error: %v\n", err))
	}
	defer reader.Close(ctx)

	for reader.Next(ctx) {
		data := &BlogItem{}
		err := reader.Decode(data)
		if err != nil {
			return status.Errorf(
				codes.Internal,
				"Failed to decode",
			)
		}

		err = stream.Send(documentToBlog(data))
		if err != nil {
			return status.Errorf(
				codes.Internal,
				"Failed to send data stream",
			)
		}
	}

	if err = reader.Err(); err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Internal error: %v\n", err))
	}

	return nil
}
