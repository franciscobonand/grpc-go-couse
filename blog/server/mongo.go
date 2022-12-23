package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func newMongoClient(ctx context.Context) (*mongo.Collection, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:root@localhost:27017/"))
	if err != nil {
		return nil, fmt.Errorf("error creating Mongo client: %v", err)
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, fmt.Errorf("error connecting to Mongo: %v", err)
	}

	coll := client.Database("blogdb").Collection("blog")
	return coll, nil
}
