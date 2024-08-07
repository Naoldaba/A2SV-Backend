package models

import (
	"time"
)

type Task struct {
	ID			int			`json:"id"`
	Title		string  	`json:"title"`
	Description	string		`json:"description"`
	DueDate		time.Time	`json:"due_date"`
	Status		string		`json:"status"`
}

type UpdateTask struct {
	Title		*string		`json:"title,omitempty"`
	Description	*string		`json:"description,omitempty"`
	DueDate		*time.Time	`json:"due_date,omitempty"`
	Status		*string		`json:"status,omitempty"`
}