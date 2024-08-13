package tests

import (
	"context"
	"testing"
	"time"

	domain "task_manager_api/Domain"
	implemenation "task_manager_api/Repository/Implementation"
	interfaces "task_manager_api/Repository/Interfaces"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TaskRepositoryTestSuite struct {
	suite.Suite
	db     *mongo.Database
	client *mongo.Client
	repo   interfaces.ITaskRepository
}

func (suite *TaskRepositoryTestSuite) SetupSuite() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
    suite.NoError(err, "cannot make db connection")

	err = client.Ping(context.TODO(), nil)
	suite.NoError(err, "cannot make db connection")

	suite.db = client.Database("test_db")
	suite.client = client
	suite.repo = implemenation.NewTaskRepository(suite.db.Collection("Task"))
}

func (suite *TaskRepositoryTestSuite) TearDownSuite() {
	err := suite.db.Drop(context.TODO())
	suite.NoError(err, "unable to drop db")

	err = suite.client.Disconnect(context.TODO())
	suite.NoError(err, "unable to disconnect")
}

func (suite *TaskRepositoryTestSuite) TestAddTask() {
	task := &domain.Task{
		Title:       "Test Task",
		Description: "Test Description",
		DueDate:     time.Now(),
		Status:      "pending",
	}

	userID := primitive.NewObjectID()
	result, err := suite.repo.AddTask(task, userID)

	suite.NoError(err, "error while adding task")
	assert.NotNil(suite.T(), result.ID)
	assert.Equal(suite.T(), task.Title, result.Title)
}

func (suite *TaskRepositoryTestSuite) TestGetTasks() {
	task := &domain.Task{
		Title:       "Test Task",
		Description: "Test Description",
		DueDate:     time.Now(),
		Status:      "pending",
	}
	userID := primitive.NewObjectID()
	_, err := suite.repo.AddTask(task, userID)
	suite.NoError(err, "error while adding task")

	tasks, err := suite.repo.GetTasks()
	suite.NoError(err, "error while retriving tasks")
	assert.GreaterOrEqual(suite.T(), len(tasks), 1)
}

func (suite *TaskRepositoryTestSuite) TestGetTaskById() {
	task := &domain.Task{
		Title:       "Test Task",
		Description: "Test Description",
		DueDate:     time.Now(),
		Status:      "pending",
	}

	userID := primitive.NewObjectID()
	insertedTask, err := suite.repo.AddTask(task, userID)
	suite.NoError(err, "error while adding task")

	result, err := suite.repo.GetTaskById(insertedTask.ID.Hex())
	suite.NoError(err, "error while retriving task")
	assert.Equal(suite.T(), insertedTask.ID, result.ID)
}

func (suite *TaskRepositoryTestSuite) TestUpdateTask() {
	task := &domain.Task{
		Title:       "Initial Task",
		Description: "Initial Description",
		DueDate:     time.Now(),
		Status:      "pending",
	}
	userID := primitive.NewObjectID()
	insertedTask, err := suite.repo.AddTask(task, userID)
	suite.NoError(err, "error while adding task")

	updatedTask := &domain.Task{
		Title:       "Updated Task",
		Description: "Updated Description",
		DueDate:     time.Now(),
		Status:      "done",
	}

	result, err := suite.repo.UpdateTask(insertedTask.ID.Hex(), updatedTask)
	suite.NoError(err, "error while updating task")
	assert.Equal(suite.T(), "Updated Task", result.Title)
}

func (suite *TaskRepositoryTestSuite) TestDeleteTask() {
	task := &domain.Task{
		Title:       "Task to Delete",
		Description: "Description",
		DueDate:     time.Now(),
		Status:      "pending",
	}
	userID := primitive.NewObjectID()
	insertedTask, err := suite.repo.AddTask(task, userID)
	suite.NoError(err, "error while adding task")

	err = suite.repo.DeleteTask(insertedTask.ID.Hex())
	suite.NoError(err, "error while deleting task")

	_, err = suite.repo.GetTaskById(insertedTask.ID.Hex())
	assert.Error(suite.T(), err)
}

func TestTaskRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(TaskRepositoryTestSuite))
}