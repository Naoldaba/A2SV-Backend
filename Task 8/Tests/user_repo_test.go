package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"task_manager_api/Domain"
	"task_manager_api/Repository/Implementation"
	"task_manager_api/Repository/Interfaces"
)

type UserRepositoryTestSuite struct {
	suite.Suite
	client *mongo.Client
	db     *mongo.Database
	repo   interfaces.IUserRepository
}

func (suite *UserRepositoryTestSuite) SetupSuite() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	suite.NoError(err, "no error connecting to the MongoDB server")

	err = client.Ping(context.TODO(), nil)
	suite.NoError(err, "no error pinging the MongoDB server")

	suite.client = client
	suite.db = client.Database("test_db")
	suite.repo = implemenation.NewUserRepository(suite.db.Collection("Users"))
}

func (suite *UserRepositoryTestSuite) TearDownSuite() {
	err := suite.client.Disconnect(context.TODO())
	suite.NoError(err, "no error disconnecting from the MongoDB server")
}

func (suite *UserRepositoryTestSuite) SetupTest() {
	err := suite.db.Collection("users").Drop(context.TODO())
	suite.NoError(err, "no error dropping the users collection")
}

func (suite *UserRepositoryTestSuite) TestRegisterUser() {
	user := &domain.User{
		Email:    "test@example.com",
		UserName: "Test User",
		Role:     "ADMIN",
		Password: "asafsd1123", 
	}

	err := suite.repo.Register(user)
	suite.NoError(err, "no error when registering the user")

	var insertedUser *domain.User
	insertedUser, err = suite.repo.GetUser(user.Email)
	suite.NoError(err, "no error finding the registered user")
	suite.Equal(user.Email, insertedUser.Email, "the registered user's email should match")
	suite.Equal(user.UserName, insertedUser.UserName, "the registered user's username should match")
	suite.Equal(user.Role, insertedUser.Role, "the registered user's role should match")
	suite.Equal(user.Password, insertedUser.Password, "the registered user's password should match")
}

func (suite *UserRepositoryTestSuite) TestGetUser() {
	user := &domain.User{
		Email:    "test@example.com",
		UserName: "Test User",
		Role:     "ADMIN",
		Password: "asafsd1123",
	}
	err := suite.repo.Register(user)
	suite.NoError(err, "no error inserting the user for testing retrieval")

	retrievedUser, err := suite.repo.GetUser(user.Email)
	suite.NoError(err, "no error when getting the user")
	suite.NotNil(retrievedUser, "retrieved user should not be nil")
	suite.Equal(user.Email, retrievedUser.Email, "the retrieved user's email should match")
	suite.Equal(user.UserName, retrievedUser.UserName, "the retrieved user's username should match")
	suite.Equal(user.Role, retrievedUser.Role, "the retrieved user's role should match")
	suite.Equal(user.Password, retrievedUser.Password, "the retrieved user's password should match")
}

func (suite *UserRepositoryTestSuite) TestGetUserNotFound() {
	_, err := suite.repo.GetUser("nonexistent@example.com")
	suite.Error(err, "an error should be returned when the user is not found")
	suite.EqualError(err, "user not found", "the error message should be 'user not found'")
}

func TestUserRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}