package tests

import (
	"testing"
	"time"

	domain "task_manager_api/Domain"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DomainTestSuite struct {
	suite.Suite
	validate *validator.Validate
}

func (suite *DomainTestSuite) SetupSuite() {
	suite.validate = validator.New()
}

func (suite *DomainTestSuite) TestTaskInitialization() {
	id := primitive.NewObjectID()
	userID := primitive.NewObjectID()
	dueDate := time.Now().Add(24 * time.Hour)

	task := domain.Task{
		ID:          id,
		Title:       "Test Task",
		Description: "This is a test task.",
		DueDate:     dueDate,
		Status:      "Pending",
		UserID:      userID,
	}

	suite.Equal(id, task.ID, "The ID should match")
	suite.Equal("Test Task", task.Title, "The title should match")
	suite.Equal("Pending", task.Status, "The status should match")
	suite.True(task.DueDate.Equal(dueDate), "The due date should match")
	suite.Equal(userID, task.UserID, "The user ID should match")
}

func (suite *DomainTestSuite) TestUserInitialization() {
	id := primitive.NewObjectID()

	user := domain.User{
		ID:       id,
		UserName: "testuser",
		Email:    "testuser@example.com",
		Role:     "USER",
		Password: "password",
	}

	suite.Equal(id, user.ID, "The ID should match")
	suite.Equal("testuser", user.UserName, "The username should match")
	suite.Equal("testuser@example.com", user.Email, "The email should match")
	suite.Equal("USER", user.Role, "The role should match")
}

func (suite *DomainTestSuite) TestUserValidation() {
	validUser := domain.User{
		UserName: "testuser",
		Email:    "testuser@example.com",
		Role:     "USER",
		Password: "password",
	}

	err := suite.validate.Struct(validUser)
	suite.NoError(err, "The valid user should not produce a validation error")

	invalidUser := domain.User{
		UserName: "testuser",
		Email:    "testuser@example.com",
		Role:     "INVALID_ROLE",
		Password: "password",
	}

	err = suite.validate.Struct(invalidUser)
	suite.Error(err, "An invalid role should produce a validation error")
}

func TestDomainTestSuite(t *testing.T) {
	suite.Run(t, new(DomainTestSuite))
}
