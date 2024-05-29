package repository

import (
	"context"
	"my-reading-app/internal/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ReadingStatusRepository define as operações que podem ser realizadas no repositório de status de leitura.
type ReadingStatusRepository interface {
	GetStatus(userId string) ([]domain.ReadingStatus, error)
	UpdateStatus(userId string, day int, status string) error
}

type mongoReadingStatusRepository struct {
	client *mongo.Client
}

func NewMongoReadingStatusRepository(client *mongo.Client) ReadingStatusRepository {
	return &mongoReadingStatusRepository{client: client}
}

func (r *mongoReadingStatusRepository) GetStatus(userId string) ([]domain.ReadingStatus, error) {
	collection := r.client.Database("bibleLectures").Collection("readingStatus")
	filter := bson.M{"userId": userId}
	var statuses []domain.ReadingStatus

	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var status domain.ReadingStatus
		if err := cursor.Decode(&status); err != nil {
			return nil, err
		}
		statuses = append(statuses, status)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return statuses, nil
}

func (r *mongoReadingStatusRepository) UpdateStatus(userId string, day int, status string) error {
	collection := r.client.Database("bibleLectures").Collection("readingStatus")
	filter := bson.M{"userId": userId, "day": day}
	update := bson.M{"$set": bson.M{"status": status}}
	_, err := collection.UpdateOne(context.Background(), filter, update, options.Update().SetUpsert(true))
	return err
}
