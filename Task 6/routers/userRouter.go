package router

import (
	"task_manager_api/controllers"
	"github.com/gin-gonic/gin"
)


func UserRouter(router *gin.Engine) {
	router.POST("/user/register", controllers.Register)
	router.POST("/user/login", controllers.Login)
}