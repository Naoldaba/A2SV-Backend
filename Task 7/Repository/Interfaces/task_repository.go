package interfaces

import (
	domain "task_manager_api/Domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ITaskRepository interface {
	GetTasks() ([]*domain.Task, error)
	GetTaskById(id string) (*domain.Task, error)
	AddTask(task *domain.Task, id primitive.ObjectID) (*domain.Task, error)
	UpdateTask(id string, task *domain.Task) (*domain.Task, error)
	DeleteTask(id string) error
}