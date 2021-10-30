package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `json:"_id"`
	Username string             `json:"username"`
	Password string             `json:"password"`
}
