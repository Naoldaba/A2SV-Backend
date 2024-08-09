package main

import (
	"fmt"
	"log"
	"task_manager_api/data"
	"task_manager_api/routers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil{
		log.Fatal("err loading .env file")
	}

	_, err = data.DbInstance()
	
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
		return 
	}

	r := gin.New()
	r.Use(gin.Logger())

	router.TaskRouter(r)
	router.UserRouter(r)

	if err := r.Run("localhost:5050"); err!= nil{
		log.Fatal(err)
	}

	fmt.Println("server connnected")
}