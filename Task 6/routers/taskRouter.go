package router

import (
	"task_manager_api/controllers"
	"task_manager_api/middleware"
	"github.com/gin-gonic/gin"
)

func TaskRouter(router *gin.Engine){
	router.GET("/tasks", middleware.JWTValidation(), controllers.GetTasks)
	router.GET("/tasks/:id", controllers.GetTaskById)
	router.POST("/tasks",middleware.JWTValidation(), middleware.RoleAuth("ADMIN"), controllers.AddTask)
	router.PUT("/tasks/:id", controllers.UpdateTask)
	router.PATCH("/tasks/:id", controllers.UpdateSpecificField)
	router.DELETE("/tasks/:id", controllers.DeleteTask)
}