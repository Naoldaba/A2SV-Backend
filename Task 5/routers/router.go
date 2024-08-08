package router

import (
	"task_manager_api/controllers"
	"task_manager_api/data"

	"github.com/gin-gonic/gin"
)

func CreateRouter(router *gin.Engine, taskService *data.TaskService){
	router.GET("/tasks", func(ctx *gin.Context) { controllers.GetTasks(ctx, taskService) })
	router.GET("/tasks/:id", func(ctx *gin.Context) { controllers.GetTaskById(ctx, taskService) })
	router.POST("/tasks", func(ctx *gin.Context) { controllers.AddTask(ctx, taskService) })
	router.PUT("/tasks/:id", func(ctx *gin.Context) { controllers.UpdateTask(ctx, taskService) })
	router.PATCH("/tasks/:id", func(ctx *gin.Context) { controllers.UpdateSpecificField(ctx, taskService) })
	router.DELETE("/tasks/:id", func(ctx *gin.Context) { controllers.DeleteTask(ctx, taskService) })
}