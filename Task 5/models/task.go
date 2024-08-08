package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID				primitive.ObjectID		`json:"id" bson:"_id,omitempty"`
	Title		    string  				`json:"title"`
	Description	    string					`json:"description"`
	DueDate		    time.Time				`json:"due_date"`
	Status		    string					`json:"status"`
}

type UpdateTask struct {
	Title		*string		`json:"title,omitempty"`
	Description	*string		`json:"description,omitempty"`
	DueDate		*time.Time	`json:"due_date,omitempty"`
	Status		*string		`json:"status,omitempty"`
}