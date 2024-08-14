package main

import (
	"fmt"
	"log"
	"os"
	"task_manager_api/Delivery/controllers"
	"task_manager_api/Delivery/routers"
	infrastructure "task_manager_api/Infrastructure"
	"task_manager_api/Repository/Implementation"
	usecases "task_manager_api/UseCases"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main(){
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
		return 
	}

	db_instance, err := infrastructure.DbInstance()
	if err  != nil {
		log.Fatal(err)
		return 
	}

	fmt.Println("Database connected successfully")

	router := gin.New()
	router.Use(gin.Logger())

	taskCollection := infrastructure.OpenCollection(db_instance, "Tasks")
	userCollection := infrastructure.OpenCollection(db_instance, "Users")

	tasKRepo := implemenation.NewTaskRepository(taskCollection)
	userRepo := implemenation.NewUserRepository(userCollection)

	taskUsecase := usecases.NewTaskUseCase(tasKRepo)
	userUsecase := usecases.NewUserUseCase(userRepo, infrastructure.PasswordHasher)

	jwtService := infrastructure.NewJWTService([]byte(os.Getenv("SECRET_KEY")))
	taskController := controllers.NewTaskController(taskUsecase)
	userController := controllers.NewUserController(userUsecase, jwtService, infrastructure.PasswordComparator)

	routers.CreateTaskRouter(router, taskController)
	routers.CreateUserRouter(router, userController)

	if err := router.Run(":" + os.Getenv("PORT")); err!= nil{
		log.Fatal(err)
	}

	fmt.Println("server connnected")
}	