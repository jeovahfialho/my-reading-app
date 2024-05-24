package domain

import "time"

type User struct {
	ID          string    `bson:"_id,omitempty"`
	Name        string    `bson:"name"`
	Email       string    `bson:"email"`
	Password    string    `bson:"password"`
	DateOfBirth string    `bson:"dateOfBirth"`
	CreatedAt   time.Time `bson:"createdAt"`
	Role        string    `bson:"role"`
	Status      string    `bson:"status"`
}
