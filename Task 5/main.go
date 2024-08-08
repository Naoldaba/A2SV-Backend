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
	if err != nil{
		log.Fatal("err loading .env file")
	}

	dbName := os.Getenv("DB_NAME")
	colName := os.Getenv("COLLECTION_NAME")

	connString := os.Getenv("CONNECTION_STRING")

	taskService := data.CreateTaskService(dbName, colName, connString)

	if taskService == nil {
		log.Fatal("Failed to connect to the database")
	}

	r := gin.Default()
	router.CreateRouter(r, taskService)

	if err := r.Run("localhost:5050"); err!= nil{
		log.Fatal(err)
	}

	fmt.Println("server connnected")
}