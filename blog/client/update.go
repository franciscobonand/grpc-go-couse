package main

import (
	"context"
	"log"

	pb "github.com/franciscobonand/grpc-go-course/blog/proto"
)

func updateBlog(ctx context.Context, c pb.BlogServiceClient, id string) {
	log.Println("Update Blog invoked")

	blog := &pb.Blog{
		Id:       id,
		AuthorId: "Roberto",
		Title:    "Brand new title",
		Content:  "Updated content of the first blog!",
	}

	_, err := c.UpdateBlog(ctx, blog)
	if err != nil {
		log.Fatalf("Failed to update Blog: %v\n", err)
	}

	log.Println("Blog updated!")
}
