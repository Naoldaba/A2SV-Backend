package interfaces

import domain "task_manager_api/Domain"

type ITaskRepository interface {
	GetTasks() ([]*domain.Task, error)
	GetTaskById(id string) (*domain.Task, error)
	AddTask(task *domain.Task, id string) (*domain.Task, error)
	UpdateTask(id string, task *domain.Task) (*domain.Task, error)
	DeleteTask(id string) error
}