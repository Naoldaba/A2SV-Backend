package tests

import (
	"errors"
	"task_manager_api/Domain"
	"task_manager_api/Mocks"
	"task_manager_api/UseCases"
	"testing"

	"github.com/stretchr/testify/suite"
)

type UserUseCaseTestSuite struct {
	suite.Suite
	mockRepo   *mocks.MockUserRepository
	userUseCase *usecases.UserUseCase
}

func (suite *UserUseCaseTestSuite) SetupTest() {
	suite.mockRepo = new(mocks.MockUserRepository)
	suite.userUseCase = usecases.NewUserUseCase(suite.mockRepo)
}

func (suite *UserUseCaseTestSuite) TestRegister_UserAlreadyExists() {
	email := "test@example.com"
	existingUser := &domain.User{
		Email: email,
	}

	suite.mockRepo.On("GetUser", email).Return(existingUser, nil)

	err := suite.userUseCase.Register(existingUser)
	suite.Error(err)
	suite.EqualError(err, "user already exists")
	suite.mockRepo.AssertCalled(suite.T(), "GetUser", email)
}

func (suite *UserUseCaseTestSuite) TestRegister_Success() {
	email := "test@example.com"
	user := &domain.User{
		UserName: "subject_1",
		Email:    email,
		Password: "plain_password",
		Role:     "ADMIN",
	}

	suite.mockRepo.On("GetUser", email).Return(nil, nil)
	suite.mockRepo.On("Register", user).Return(nil)

	err := suite.userUseCase.Register(user)
	suite.NoError(err)
	suite.mockRepo.AssertCalled(suite.T(), "GetUser", email)
	suite.mockRepo.AssertCalled(suite.T(), "Register", user)
}

func (suite *UserUseCaseTestSuite) TestGetUserByEmail_Success() {
	email := "test@example.com"
	expectedUser := &domain.User{
		Email: email,
	}

	suite.mockRepo.On("GetUser", email).Return(expectedUser, nil)

	user, err := suite.userUseCase.GetUserByEmail(email)
	suite.NoError(err)
	suite.Equal(expectedUser, user)
	suite.mockRepo.AssertCalled(suite.T(), "GetUser", email)
}

func (suite *UserUseCaseTestSuite) TestGetUserByEmail_NotFound() {
	email := "test2@example.com"

	suite.mockRepo.On("GetUser", email).Return(nil, errors.New("error"))

	user, _ := suite.userUseCase.GetUserByEmail(email)
	suite.Nil(user)
}

func TestUserUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(UserUseCaseTestSuite))
}