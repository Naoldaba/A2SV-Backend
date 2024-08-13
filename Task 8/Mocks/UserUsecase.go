package mocks

import (
	domain "task_manager_api/Domain"

	"github.com/stretchr/testify/mock"
)

type MockUserUseCase struct {
	mock.Mock
}


func (m *MockUserUseCase) Register(user *domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}


func (m *MockUserUseCase) GetUserByEmail(email string) (*domain.User, error) {
	args := m.Called(email)
	return args.Get(0).(*domain.User), args.Error(1)
}
