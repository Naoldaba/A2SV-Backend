package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct{
	ID			primitive.ObjectID		`json:"id" bson:"_id,omitempty"`
	UserName	string					`json:"username"`
	Email		string					`json:"email"`
	Role		string					`json:"role" validate:"required, oneof=ADMIN USER"`
	Password	string					`json:"-"`
}