package domain

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct{
	ID				primitive.ObjectID		`json:"id" bson:"_id,omitempty"`
	Title		    string  				`json:"title"`
	Description	    string					`json:"description"`
	DueDate		    time.Time				`json:"due_date"`
	Status		    string					`json:"status"`
	UserID			primitive.ObjectID		`json:"user_id"`
}




type User struct{
	ID			primitive.ObjectID		`json:"id" bson:"_id,omitempty"`
	UserName	string					`json:"username"`
	Email		string					`json:"email"`
	Role		string					`json:"role" validate:"required,oneof=ADMIN USER"`
	Password	string					`json:"password"`
}	