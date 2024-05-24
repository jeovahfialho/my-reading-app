package db

import (
	"context"
	"my-reading-app/pkg/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectMongo creates a new client and connects to MongoDB
func ConnectMongo(ctx context.Context) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(config.GetMongoURI())
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}
	return client, nil
}
