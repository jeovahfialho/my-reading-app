package repository

import (
	"context"
	"my-reading-app/internal/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	CreateUser(user domain.User) (string, error)
	GetUserByEmail(email string) (domain.User, error)
}

type mongoUserRepository struct {
	client *mongo.Client
}

func NewMongoUserRepository(client *mongo.Client) UserRepository {
	return &mongoUserRepository{client: client}
}

func (m *mongoUserRepository) CreateUser(user domain.User) (string, error) {
	collection := m.client.Database("bibleLectures").Collection("users")
	result, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		return "", err
	}
	id := result.InsertedID.(primitive.ObjectID).Hex()
	return id, nil
}

func (m *mongoUserRepository) GetUserByEmail(email string) (domain.User, error) {
	var user domain.User
	collection := m.client.Database("bibleLectures").Collection("users")
	filter := bson.M{"email": email}
	err := collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}
