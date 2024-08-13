package mocks

import (
    "task_manager_api/Domain"
    "github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
    mock.Mock
}

func (m *MockUserRepository) GetUser(email string) (*domain.User, error){
	args := m.Called(email)
	if args.Get(0) != nil {
		return args.Get(0).(*domain.User), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockUserRepository) Register(user *domain.User) error {
    args := m.Called(user)
    return args.Error(0)
}