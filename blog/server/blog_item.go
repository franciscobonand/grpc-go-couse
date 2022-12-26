package main

import (
	pb "github.com/franciscobonand/grpc-go-course/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogItem struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	AuthorID string             `bson:"author_id"`
	Title    string             `bson:"title"`
	Content  string             `bson:"content"`
}

func documentToBlog(data *BlogItem) *pb.Blog {
	return &pb.Blog{
		Id:       data.ID.Hex(),
		AuthorId: data.AuthorID,
		Title:    data.Title,
		Content:  data.Content,
	}
}

func blogToDocument(data *pb.Blog) (*BlogItem, error) {
	bi := &BlogItem{
		AuthorID: data.AuthorId,
		Title:    data.Title,
		Content:  data.Content,
	}

	if data.Id != "" {
		oid, err := primitive.ObjectIDFromHex(data.Id)
		if err != nil {
			return nil, err
		}
		bi.ID = oid
	}

	return bi, nil
}
