package main

import (
	"fmt"
	"log"
	"os"
	"task_manager_api/data"
	"task_manager_api/routers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	PORT := os.Getenv("PORT")

	if err != nil{
		log.Fatal("err loading .env file")
	}

	_, err = data.DbInstance()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
		return 
	}

	fmt.Println("Database Connected Successfully")

	r := gin.New()
	r.Use(gin.Logger())

	router.CreateRouter(r)

	if err := r.Run(":" + PORT); err!= nil{
		log.Fatal(err)
	}

	fmt.Println("server connnected")
}