package tests

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"task_manager_api/Delivery/controllers"
	domain "task_manager_api/Domain"
	mocks "task_manager_api/Mocks"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskControllerTestSuite struct {
	suite.Suite
	mockUsecase    *mocks.MockTaskUseCase
	taskController *controllers.TaskController
	router         *gin.Engine
}

func (s *TaskControllerTestSuite) SetupSuite() {
	s.mockUsecase = new(mocks.MockTaskUseCase)
	s.taskController = controllers.NewTaskController(s.mockUsecase)

	s.router = gin.Default()

	s.router.POST("/tasks", s.taskController.AddTask)
	s.router.GET("/tasks", s.taskController.GetTasks)
	s.router.GET("/tasks/:id", s.taskController.GetTaskById)
	s.router.PUT("/tasks/:id", s.taskController.UpdateTask)
	s.router.DELETE("/tasks/:id", s.taskController.DeleteTask)
}

// func (s *TaskControllerTestSuite) TestAddTask() {
// 	task := &domain.Task{
// 		Title:       "Test Task",
// 		Description: "Test description",
// 		DueDate:     time.Now(),
// 		Status:      "done",
// 	}
// 	user_id := primitive.NewObjectID()
// 	s.mockUsecase.On("AddTask", mock.AnythingOfType("*domain.Task"), user_id).Return(&task, nil)

// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("POST", "/tasks", strings.NewReader(`{"title":"Test Task","description":"Test description","due_date":"`+task.DueDate.Format(time.RFC3339)+`","status":"done"}`))
// 	req.Header.Set("Content-Type", "application/json")

// 	ctx, _ := gin.CreateTestContext(w)
// 	ctx.Request = req
// 	ctx.Set("userID", user_id.Hex())

// 	s.router.ServeHTTP(w, req)

// 	fmt.Println("Response Body:", w.Body.String())
// 	s.Equal(http.StatusOK, w.Code)
// 	s.mockUsecase.AssertExpectations(s.T())
// }

func (s *TaskControllerTestSuite) TestGetTasks() {
	tasks := []*domain.Task{
		{
			Title:       "Task 1",
			Description: "Description 1",
			DueDate:     time.Now(),
			Status:      "done",
			UserID:      primitive.NewObjectID(),
		},
		{
			Title:       "Task 2",
			Description: "Description 2",
			DueDate:     time.Now(),
			Status:      "pending",
			UserID:      primitive.NewObjectID(),
		},
	}

	s.mockUsecase.On("GetTasks").Return(tasks, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/tasks", nil)

	s.router.ServeHTTP(w, req)

	s.Equal(http.StatusOK, w.Code)
	s.mockUsecase.AssertExpectations(s.T())
}

func (s *TaskControllerTestSuite) TestGetTaskById() {
	task := &domain.Task{
		Title:       "Test Task",
		Description: "Test description",
		DueDate:     time.Now(),
		Status:      "done",
		UserID:      primitive.NewObjectID(),
	}

	taskID := primitive.NewObjectID().Hex()

	s.mockUsecase.On("GetTaskById", taskID).Return(task, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/tasks/"+taskID, nil)

	s.router.ServeHTTP(w, req)

	s.Equal(http.StatusOK, w.Code)
	s.mockUsecase.AssertExpectations(s.T())
}

func (s *TaskControllerTestSuite) TestUpdateTask() {
	taskID := primitive.NewObjectID().Hex()

	updatedTask := &domain.Task{
		Title:       "Updated Task",
		Description: "Updated description",
		DueDate:     time.Now(),
		Status:      "in-progress",
		UserID:      primitive.NewObjectID(),
	}

	s.mockUsecase.On("UpdateTask", taskID, mock.AnythingOfType("*domain.Task")).Return(updatedTask, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/tasks/"+taskID, strings.NewReader(`{"title":"Updated Task","description":"Updated description","due_date":"`+updatedTask.DueDate.Format(time.RFC3339)+`","status":"in-progress","user_id":"`+updatedTask.UserID.Hex()+`"}`))
	req.Header.Set("Content-Type", "application/json")

	s.router.ServeHTTP(w, req)
	fmt.Println("Response Body:", w.Body.String())
	s.Equal(http.StatusOK, w.Code)
	s.mockUsecase.AssertExpectations(s.T())
}

func (s *TaskControllerTestSuite) TestDeleteTask() {
	taskID := primitive.NewObjectID().Hex()

	s.mockUsecase.On("DeleteTask", taskID).Return(nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/tasks/"+taskID, nil)

	s.router.ServeHTTP(w, req)

	s.Equal(http.StatusOK, w.Code)
	s.mockUsecase.AssertExpectations(s.T())
}

func TestTaskControllerTestSuite(t *testing.T) {
	suite.Run(t, new(TaskControllerTestSuite))
}
