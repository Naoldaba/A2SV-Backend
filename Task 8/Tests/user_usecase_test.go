package tests

import (
	"errors"
	"task_manager_api/Domain"
	infrastructure "task_manager_api/Infrastructure"
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
	suite.userUseCase = usecases.NewUserUseCase(suite.mockRepo, infrastructure.PasswordHasher)
}

func (suite *UserUseCaseTestSuite) TestRegister_UserAlreadyExists() {
	email := "test@example.com"
	existingUser := &domain.User{
		Email: email,
	}
	
	users := []*domain.User{
		{
			UserName: "subject_1",
			Email:    email,
			Password: "plain_password",
			Role:     "ADMIN",
		},
		{
			UserName: "subject_2",
			Email:    email,
			Password: "plain_password",
			Role:     "ADMIN",
		},

	}

	suite.mockRepo.On("GetUser", email).Return(existingUser, nil)
	suite.mockRepo.On("GetAllUsers").Return(users, nil)

	err := suite.userUseCase.Register(existingUser)
	suite.Error(err)
	suite.EqualError(err, "user already exists")
	suite.mockRepo.AssertCalled(suite.T(), "GetUser", email)
}

func (suite *UserUseCaseTestSuite) TestRegister_Success() {
	email := "test@example.com"
	users := []*domain.User{
		{
			UserName: "subject_1",
			Email:    email,
			Password: "plain_password",
			Role:     "ADMIN",
		},
		{
			UserName: "subject_2",
			Email:    email,
			Password: "plain_password",
			Role:     "ADMIN",
		},

	}

	suite.mockRepo.On("GetUser", email).Return(nil, nil)
	suite.mockRepo.On("Register", users[0]).Return(nil)
	suite.mockRepo.On("GetAllUsers").Return(users, nil)

	err := suite.userUseCase.Register(users[0])
	suite.NoError(err)
	suite.mockRepo.AssertCalled(suite.T(), "GetUser", email)
	suite.mockRepo.AssertCalled(suite.T(), "Register", users[0])
}

func (suite *UserUseCaseTestSuite) TestGetUserByEmail_Success() {
	email := "test@example.com"
	expectedUser := &domain.User{
		Email: email,
	}

	suite.mockRepo.On("GetUser", email).Return(expectedUser, nil)

	user, err := suite.userUseCase.GetUser(email)
	suite.NoError(err)
	suite.Equal(expectedUser, user)
	suite.mockRepo.AssertCalled(suite.T(), "GetUser", email)
}

func (suite *UserUseCaseTestSuite) TestGetUserByEmail_NotFound() {
	email := "test2@example.com"

	suite.mockRepo.On("GetUser", email).Return(nil, errors.New("error"))

	user, _ := suite.userUseCase.GetUser(email)
	suite.Nil(user)
}

func TestUserUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(UserUseCaseTestSuite))
}