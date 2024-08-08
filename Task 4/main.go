package main

import (
	"task_manager_api/router"
	"log"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.CreateRouter(r)

	if err := r.Run("localhost:5050"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
	fmt.Println("Server started")
}