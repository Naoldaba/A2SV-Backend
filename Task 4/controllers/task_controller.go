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
    var updatedTask models.UpdateTask
    err := ctx.BindJSON(&updatedTask)
    if err != nil {
        ctx.IndentedJSON(http.StatusBadRequest, gin.H{
            "message": "Invalid format",
        })
        return
    }

    id, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        ctx.IndentedJSON(http.StatusBadRequest, gin.H{
            "message": "Invalid Id",
        })
        return
    }
    existingTask, err := taskService.GetTaskById(id)
    if err != nil {
        ctx.IndentedJSON(http.StatusBadRequest, gin.H{
            "message": "No such Task",
        })
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
        ctx.IndentedJSON(http.StatusBadRequest, gin.H{
            "message": "No such Task",
        })
        return 
    }
    
    ctx.IndentedJSON(http.StatusOK, gin.H{
        "updatedTask": updates,
    })

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