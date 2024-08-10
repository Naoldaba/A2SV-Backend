package controllers

import (
	"net/http"
	"task_manager_api/Domain"
	"task_manager_api/UseCases"

	"github.com/gin-gonic/gin"
)

type TaskController struct{
	taskUsecase *usecases.TaskUseCase
}

func NewTaskController(taskUsecase *usecases.TaskUseCase) *TaskController{
	return &TaskController{
		taskUsecase: taskUsecase,
	}
}

func (tc *TaskController) AddTask(c *gin.Context) {
	var task domain.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	userID, exists := c.Get("userID")
	user_id := userID.(string)
	
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "NO user ID"})
		return
	}
	createdTask, err := tc.taskUsecase.AddTask(&task, user_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, createdTask)
}

func (tc *TaskController) GetTasks(c *gin.Context) {
	tasks, err := tc.taskUsecase.GetTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func (tc *TaskController) GetTaskById(c *gin.Context) {
	id := c.Param("id")

	task, err := tc.taskUsecase.GetTaskById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (tc *TaskController) UpdateTask(c *gin.Context) {
	id := c.Param("id")

	var task domain.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	updatedTask, err := tc.taskUsecase.UpdateTask(id, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedTask)
}

func (tc *TaskController) DeleteTask(c *gin.Context) {
	id := c.Param("id")

	err := tc.taskUsecase.DeleteTask(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
