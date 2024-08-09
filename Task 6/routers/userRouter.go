package router

import (
	"task_manager_api/controllers"
	"github.com/gin-gonic/gin"
)


func UserRouter(router *gin.Engine) {
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)
}