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

	Client := data.Client

	r := gin.New()
	r.Use(gin.Logger())

	router.CreateRouter(r, Client)

	if err := r.Run("localhost:5050"); err!= nil{
		log.Fatal(err)
	}

	fmt.Println("server connnected")
}