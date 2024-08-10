package routers

import (
	"github.com/gin-gonic/gin"
	"task_manager_api/Delivery/controllers"
)

func CreateUserRouter(router *gin.Engine, userController *controllers.UserController){
	router.POST("/register", userController.Register)
	router.POST("/login", userController.Login)	
}