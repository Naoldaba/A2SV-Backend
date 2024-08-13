package mocks

import (
	"task_manager_api/Domain"

	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockTaskRepository struct {
	mock.Mock
}

func (m *MockTaskRepository) AddTask(task *domain.Task, id primitive.ObjectID) (*domain.Task, error) {
	args := m.Called(task, id)
	return args.Get(0).(*domain.Task), args.Error(1)
}

func (m *MockTaskRepository) GetTasks() ([]*domain.Task, error) {
	args := m.Called()
	return args.Get(0).([]*domain.Task), args.Error(1)
}

func (m *MockTaskRepository) GetTaskById(id string) (*domain.Task, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*domain.Task), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockTaskRepository) UpdateTask(id string, task *domain.Task) (*domain.Task, error) {
	args := m.Called(id, task)
	if args.Get(0) != nil {
		return args.Get(0).(*domain.Task), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockTaskRepository) DeleteTask(id string) error {
	args := m.Called(id)
	return args.Error(0)
}