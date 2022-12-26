package main

import (
	"context"
	"log"

	pb "github.com/franciscobonand/grpc-go-course/blog/proto"
)

func deleteBlog(ctx context.Context, c pb.BlogServiceClient, id string) {
	log.Println("Delete Blog invoked")

	req := &pb.BlogId{Id: id}
	_, err := c.DeleteBlog(ctx, req)
	if err != nil {
		log.Fatalf("Error while deleting: %v\n", err)
	}

	log.Printf("Blog '%s' deleted", id)
}
