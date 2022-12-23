package main

import (
	"context"
	"log"

	pb "github.com/franciscobonand/grpc-go-course/blog/proto"
)

func createBlog(ctx context.Context, c pb.BlogServiceClient) string {
	log.Println("Create Blog invoked")

	blog := &pb.Blog{
		AuthorId: "Chico",
		Title:    "My first blog",
		Content:  "Content of the first blog!",
	}

	resp, err := c.CreateBlog(ctx, blog)
	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}

	return resp.Id
}
