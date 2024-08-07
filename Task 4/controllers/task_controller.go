package controllers

import (
	"net/http"
	"task_manager_api/data"
	"task_manager_api/models"
	"strconv"
	"github.com/gin-gonic/gin"
)


var taskService = data.CreateTaskSerive()


func GetTasks(ctx *gin.Context) {
    tasks := taskService.GetTasks()
    ctx.IndentedJSON(http.StatusOK, tasks)
}

func GetTaskById(ctx *gin.Context) {
    id := ctx.Param("id")
    taskId, err := strconv.Atoi(id)
    if err != nil {
        ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
        return
    }

    task, err := taskService.GetTaskById(taskId)
    if err != nil {
        ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    ctx.IndentedJSON(http.StatusOK, task)
}

func AddTask(ctx *gin.Context) {
    var newTask models.Task
    if err := ctx.BindJSON(&newTask); err != nil {
        ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    task := taskService.AddTask(newTask)
    ctx.IndentedJSON(http.StatusOK, task)
}

func UpdateTask(ctx *gin.Context) {
    var updatedTask models.Task
    if err := ctx.BindJSON(&updatedTask); err != nil {
        ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    id := ctx.Param("id")
    taskId, err := strconv.Atoi(id)
    if err != nil {
        ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
        return
    }

    task, err := taskService.UpdateTask(taskId, updatedTask)
    if err != nil {
        ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    ctx.IndentedJSON(http.StatusOK, task)
}

func DeleteTask(ctx *gin.Context) {
    id := ctx.Param("id")
    taskId, err := strconv.Atoi(id)
    if err != nil {
        ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
        return
    }

    err = taskService.DeleteTask(taskId)
    if err != nil {
        ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}