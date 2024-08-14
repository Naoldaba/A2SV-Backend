package mocks

import (
	"task_manager_api/Domain"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/mock"
	//"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockInfrastructure struct {
	mock.Mock
}

func (m *MockInfrastructure) ValidateToken(token string) (*jwt.Token, error){
	args := m.Called(token)
	return args.Get(0).(*jwt.Token), nil
}

func (m *MockInfrastructure) GenerateToken(user domain.User) (string, error){
	args := m.Called(user)
	return args.Get(0).(string), nil
}

