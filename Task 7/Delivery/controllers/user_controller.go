package controllers

import (
	"net/http"
	"log"
	"task_manager_api/Domain"
	infrastructure "task_manager_api/Infrastructure"
	"task_manager_api/UseCases"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type UserController struct {
	userUsecase *usecases.UserUseCase
}

func NewUserController(userUsecase *usecases.UserUseCase) *UserController{
	return &UserController{
		userUsecase: userUsecase,
	}
}

func (uc *UserController) Register(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	err := uc.userUsecase.Register(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func (uc *UserController) GetUserByEmail(c *gin.Context) {
	email := c.Param("email")

	user, err := uc.userUsecase.GetUserByEmail(email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (uc *UserController) Login(c *gin.Context){
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
		return
	}
	
	var user domain.User
	if err = c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	existingUser, err := uc.userUsecase.GetUserByEmail(user.Email)
	if err != nil {
		log.Fatal(err)
		return
	}
	
	SignedToken, err := infrastructure.GenerateToken(*existingUser)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(200, gin.H{"message": "User logged in successfully", "token": SignedToken})
}