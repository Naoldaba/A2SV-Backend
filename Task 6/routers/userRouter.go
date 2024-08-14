package router

import (
	"task_manager_api/controllers"
	"task_manager_api/middleware"

	"github.com/gin-gonic/gin"
)


func UserRouter(router *gin.Engine) {
	router.POST("/user/register", controllers.Register)
	router.POST("/user/login", controllers.Login)
	router.GET("/users", middleware.JWTValidation(), middleware.RoleAuth("ADMIN"), controllers.GetAllUsers)
	router.POST("/user/promote_user/:id", middleware.JWTValidation(), middleware.RoleAuth("ADMIN"), controllers.PromoteUser)
}