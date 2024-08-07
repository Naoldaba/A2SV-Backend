package router

import (
	"github.com/gin-gonic/gin"
	"task_manager_api/controllers"
)

func Router(){
	router := gin.Default()

	router.GET("/tasks", controllers.GetTasks)
	router.GET("/tasks/:id", controllers.GetTaskById)
	router.POST("/tasks/", controllers.AddTask)
	router.PATCH("/tasks/:id", controllers.UpdateTask)
	router.DELETE("/tasks/:id", controllers.DeleteTask)

	router.Run("localhost:5050")
}

