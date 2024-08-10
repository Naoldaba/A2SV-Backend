package routers

import (
	"github.com/gin-gonic/gin"
	"task_manager_api/Infrastructure"
	"task_manager_api/Delivery/controllers"
)

func CreateTaskRouter(router *gin.Engine, taskController *controllers.TaskController){
	router.GET("/tasks", infrastructure.JWTValidation(), infrastructure.RoleAuth("ADMIN", "USER"), taskController.GetTasks)
	router.GET("/tasks/:id", infrastructure.JWTValidation(), infrastructure.RoleAuth("ADMIN", "USER"), taskController.GetTaskById)
	router.POST("/tasks",infrastructure.JWTValidation(), infrastructure.RoleAuth("ADMIN"), taskController.AddTask)
	router.PUT("/tasks/:id", infrastructure.JWTValidation(), infrastructure.RoleAuth("ADMIN"), taskController.UpdateTask)
	router.DELETE("/tasks/:id", infrastructure.JWTValidation(), infrastructure.RoleAuth("ADMIN"), taskController.DeleteTask)
}