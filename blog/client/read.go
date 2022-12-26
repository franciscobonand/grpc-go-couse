package main

import (
	"context"
	"log"

	pb "github.com/franciscobonand/grpc-go-course/blog/proto"
)

func readBlog(ctx context.Context, c pb.BlogServiceClient, id string) *pb.Blog {
	log.Println("Read Blog invoked")

	req := &pb.BlogId{Id: id}
	resp, err := c.ReadBlog(ctx, req)
	if err != nil {
		log.Printf("Blog with provided ID not found: %v\n", err)
		return nil
	}

	log.Printf("Blog found: %v\n", resp)
	return resp
}
