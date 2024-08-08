package controllers

import (
	"net/http"
	"task_manager_api/data"
	"task_manager_api/models"
	"github.com/gin-gonic/gin"
)

func GetTasks(ctx *gin.Context, taskService *data.TaskService) {
	tasks, err := taskService.GetTasks()
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusOK, tasks)
}

func GetTaskById(ctx *gin.Context, taskService *data.TaskService) {
	id := ctx.Param("id")

	task, err := taskService.GetTaskById(id)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, task)
}

func AddTask(ctx *gin.Context, taskService *data.TaskService) {
	var newTask models.Task
	if err := ctx.BindJSON(&newTask); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := taskService.AddTask(newTask)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, task)
}

func UpdateSpecificField(ctx *gin.Context, taskService *data.TaskService) {
	var updatedTask models.UpdateTask
	err := ctx.BindJSON(&updatedTask)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid format"})
		return
	}

	id := ctx.Param("id")

	existingTask, err := taskService.GetTaskById(id)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "No such Task"})
		return
	}

	if updatedTask.Title != nil {
		existingTask.Title = *updatedTask.Title
	}
	if updatedTask.Description != nil {
		existingTask.Description = *updatedTask.Description
	}
	if updatedTask.DueDate != nil {
		existingTask.DueDate = *updatedTask.DueDate
	}
	if updatedTask.Status != nil {
		existingTask.Status = *updatedTask.Status
	}

	updates, err := taskService.UpdateTask(id, existingTask)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "No such Task"})
		return
	}

	ctx.IndentedJSON(http.StatusOK, updates)
}

func UpdateTask(ctx *gin.Context, taskService *data.TaskService) {
	var updatedTask models.Task
	if err := ctx.ShouldBindJSON(&updatedTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := ctx.Param("id")

	task, err := taskService.UpdateTask(id, updatedTask)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, task)
}

func DeleteTask(ctx *gin.Context, taskService *data.TaskService) {
	id := ctx.Param("id")

	err := taskService.DeleteTask(id)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
