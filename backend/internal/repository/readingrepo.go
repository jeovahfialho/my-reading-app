package repository

import (
	"context"
	"my-reading-app/internal/domain"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ReadingRepository interface {
	GetReadingByDay(day int) (*domain.Reading, error)
}

type mongoRepository struct {
	client *mongo.Client
}

func NewMongoRepository(client *mongo.Client) ReadingRepository {
	return &mongoRepository{client: client}
}

func (m *mongoRepository) GetReadingByDay(day int) (*domain.Reading, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var reading domain.Reading
	collection := m.client.Database("bibleLectures").Collection("lecturesbyperiod")
	filter := bson.M{"day": day}
	err := collection.FindOne(ctx, filter).Decode(&reading)

	if err != nil {
		return nil, err
	}
	return &reading, nil
}
