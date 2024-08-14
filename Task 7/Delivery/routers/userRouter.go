package routers

import (
	"task_manager_api/Delivery/controllers"
	infrastructure "task_manager_api/Infrastructure"

	"github.com/gin-gonic/gin"
)

func CreateUserRouter(router *gin.Engine, userController *controllers.UserController){
	router.POST("/user/register", userController.Register)
	router.POST("/user/login", userController.Login)	
	router.GET("/users", infrastructure.JWTValidation(), infrastructure.RoleAuth("ADMIN"), userController.GetAllUsers)
	router.POST("/user/promote_user/:id", infrastructure.JWTValidation(), infrastructure.RoleAuth("ADMIN"), userController.PromoteUser)	
}