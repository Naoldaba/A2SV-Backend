package mocks

import (
	domain "task_manager_api/Domain"

	"github.com/stretchr/testify/mock"
)

type MockUserUseCase struct {
	mock.Mock
}

func (m *MockUserUseCase) Register(user *domain.User) error {
	_ = m.Called(user)
	return nil
}

func (m *MockUserUseCase) GetUser(email string) (*domain.User, error) {
	args := m.Called(email)
	return args.Get(0).(*domain.User), nil
}

func (m *MockUserUseCase) PromoteUserByID(id string, promoter *domain.User) (*domain.User, error){
	args := m.Called(id, promoter)
	return args.Get(0).(*domain.User), nil
}

func (m *MockUserUseCase) GetAllUsers() ([]*domain.User, error){
	args := m.Called()
	return args.Get(0).([]*domain.User), nil
}
