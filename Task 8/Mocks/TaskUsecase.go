package mocks

import (
	domain "task_manager_api/Domain"

	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockTaskUseCase struct {
	mock.Mock
}

func (m *MockTaskUseCase) AddTask(task *domain.Task, userID primitive.ObjectID) (*domain.Task, error) {
	args := m.Called(task, userID)
	if args.Get(0) != nil {
		return args.Get(0).(*domain.Task), nil
	}
	return nil, args.Error(1)
}

func (m *MockTaskUseCase) GetTasks() ([]*domain.Task, error) {
	args := m.Called()
	return args.Get(0).([]*domain.Task), nil
}

func (m *MockTaskUseCase) GetTaskById(id string) (*domain.Task, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*domain.Task), nil
	}
	return nil, args.Error(1)
}

func (m *MockTaskUseCase) UpdateTask(id string, task *domain.Task) (*domain.Task, error) {
	args := m.Called(id, task)
	if args.Get(0) != nil {
		return args.Get(0).(*domain.Task), nil
	}
	return nil, args.Error(1)
}

func (m *MockTaskUseCase) DeleteTask(id string) error {
	_ = m.Called(id)
	return nil
}
