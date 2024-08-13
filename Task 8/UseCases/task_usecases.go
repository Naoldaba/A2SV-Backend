package usecases

import (
	"errors"
	"task_manager_api/Domain"
	"task_manager_api/Repository/Interfaces"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ITaskUseCase interface {
	AddTask(task *domain.Task, id primitive.ObjectID) (*domain.Task, error)
	GetTasks() ([]*domain.Task, error)
	GetTaskById(id string) (*domain.Task, error)
	UpdateTask(id string, task *domain.Task) (*domain.Task, error)
	DeleteTask(id string) error
}

type TaskUseCase struct{
	taskRepo interfaces.ITaskRepository
}

func NewTaskUseCase(taskRepo interfaces.ITaskRepository) ITaskUseCase{
	return &TaskUseCase{
		taskRepo: taskRepo,
	}
}

func (uc *TaskUseCase) AddTask(task *domain.Task, id primitive.ObjectID) (*domain.Task, error) {
	if task.Title == "" {
		return nil, errors.New("task title cannot be empty")
	}
	return uc.taskRepo.AddTask(task, id)
}

func (uc *TaskUseCase) GetTasks() ([]*domain.Task, error) {
	return uc.taskRepo.GetTasks()
}

func (uc *TaskUseCase) GetTaskById(id string) (*domain.Task, error) {
	return uc.taskRepo.GetTaskById(id)
}

func (uc *TaskUseCase) UpdateTask(id string, task *domain.Task) (*domain.Task, error) {
	existingTask, err := uc.taskRepo.GetTaskById(id)
	if err != nil {
		return nil, err
	}

	if task.Title != "" {
		existingTask.Title = task.Title
	}
	if task.Description != "" {
		existingTask.Description = task.Description
	}
	if !task.DueDate.IsZero() {
		existingTask.DueDate = task.DueDate
	}
	if task.Status != "" {
		existingTask.Status = task.Status
	}

	return uc.taskRepo.UpdateTask(id, existingTask)
}

func (uc *TaskUseCase) DeleteTask(id string) error {
	return uc.taskRepo.DeleteTask(id)
}