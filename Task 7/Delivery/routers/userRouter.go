package routers

import (
	"github.com/gin-gonic/gin"
	"task_manager_api/Delivery/controllers"
)

func CreateUserRouter(router *gin.Engine, userController *controllers.UserController){
	router.POST("/user/register", userController.Register)
	router.POST("/user/login", userController.Login)	
}