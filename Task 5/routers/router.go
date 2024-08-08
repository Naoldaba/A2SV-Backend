package router

import (
	"task_manager_api/controllers"

	"github.com/gin-gonic/gin"
)

func CreateRouter(router *gin.Engine){
	router.GET("/tasks", controllers.GetTasks)
	router.GET("/tasks/:id", controllers.GetTaskById)
	router.POST("/tasks", controllers.AddTask)
	router.PUT("/tasks/:id", controllers.UpdateTask)
	router.PATCH("/tasks/:id", controllers.UpdateSpecificField)
	router.DELETE("/tasks/:id", controllers.DeleteTask)
}