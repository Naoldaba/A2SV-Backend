package tests

import (
	"testing"
	"time"

	"task_manager_api/Domain"
	"task_manager_api/Mocks"
	"task_manager_api/UseCases"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskUseCaseTestSuite struct {
	suite.Suite
	taskRepo *mocks.MockTaskRepository
	taskUC   usecases.ITaskUseCase
}

func (suite *TaskUseCaseTestSuite) SetupTest() {
	suite.taskRepo = new(mocks.MockTaskRepository)
	suite.taskUC = usecases.NewTaskUseCase(suite.taskRepo)
}

func (suite *TaskUseCaseTestSuite) TestAddTask() {
	task := &domain.Task{
		Title: "Test Task",
		Description: "Test description",
		DueDate: time.Now(),
		Status: "done",
		UserID: primitive.NewObjectID(),
	}
	id := primitive.NewObjectID()
	suite.taskRepo.On("AddTask", task, id).Return(task, nil)

	result, err := suite.taskUC.AddTask(task, id)

	suite.NoError(err)
	suite.Equal(task, result)
	suite.taskRepo.AssertExpectations(suite.T())
}

func (suite *TaskUseCaseTestSuite) TestAddTask_EmptyTitle() {
	task := &domain.Task{
		Title: "",
		Description: "Test description",
		DueDate: time.Now(),
		Status: "done",
		UserID: primitive.NewObjectID(),
	}

	id := primitive.NewObjectID()
	_, err := suite.taskUC.AddTask(task, id)

	suite.Error(err)
	suite.EqualError(err, "task title cannot be empty")
}

func (suite *TaskUseCaseTestSuite) TestGetTasks() {
	tasks := []*domain.Task{
		{
				Title: "Test Task 1",
				Description: "Test description",
				DueDate: time.Now(),
				Status: "done",
				UserID: primitive.NewObjectID(),
		}, 
		{
				Title: "Test Task 2",
				Description: "Test description",
				DueDate: time.Now(),
				Status: "done",
				UserID: primitive.NewObjectID(),
		},
	}

	suite.taskRepo.On("GetTasks").Return(tasks, nil)

	result, err := suite.taskUC.GetTasks()

	suite.NoError(err)
	suite.Equal(tasks, result)
	suite.taskRepo.AssertExpectations(suite.T())
}

func (suite *TaskUseCaseTestSuite) TestGetTaskById() {
	task := &domain.Task{
		Title: "Test Task",
		Description: "Test description",
		DueDate: time.Now(),
		Status: "done",
		UserID: primitive.NewObjectID(),
	}

	suite.taskRepo.On("GetTaskById", "123").Return(task, nil)

	result, err := suite.taskUC.GetTaskById("123")

	suite.NoError(err)
	suite.Equal(task, result)
	suite.taskRepo.AssertExpectations(suite.T())
}

func (suite *TaskUseCaseTestSuite) TestUpdateTask() {
	task := &domain.Task{
		ID : primitive.NewObjectID(),
		Title:       "Updated Task",
		Description: "Updated Description",
		DueDate:     time.Now(),
		Status:      "In Progress",
		UserID: primitive.NewObjectID(),
	}

	updatedTask := &domain.Task{
		ID : task.ID,
		Title:       "Updated Task",
		Description: "Updated Description",
		DueDate:     task.DueDate,
		Status:      "In Progress",
		UserID: task.UserID,
	}

	suite.taskRepo.On("GetTaskById", "123").Return(task, nil)
	suite.taskRepo.On("UpdateTask", "123", updatedTask).Return(updatedTask, nil)

	result, err := suite.taskUC.UpdateTask("123", updatedTask)

	suite.NoError(err)
	suite.Equal(updatedTask, result)
	suite.taskRepo.AssertExpectations(suite.T())
}

func (suite *TaskUseCaseTestSuite) TestDeleteTask() {
	suite.taskRepo.On("DeleteTask", "123").Return(nil)

	err := suite.taskUC.DeleteTask("123")

	suite.NoError(err)
	suite.taskRepo.AssertExpectations(suite.T())
}

func TestTaskUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(TaskUseCaseTestSuite))
}