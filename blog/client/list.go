package main

import (
	"context"
	"io"
	"log"

	pb "github.com/franciscobonand/grpc-go-course/blog/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func listBlog(ctx context.Context, c pb.BlogServiceClient) {
	log.Println("List Blogs invoked")

	stream, err := c.ListBlogs(ctx, &emptypb.Empty{})
	if err != nil {
		log.Fatalf("Unexpected error calling list blogs stream: %v\n", err)
	}

	for {
		data, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error receiving data: %v\n", err)
		}
		log.Println(data)
	}

}
