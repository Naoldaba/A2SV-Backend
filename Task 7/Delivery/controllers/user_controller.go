package controllers

import (
	"net/http"
	domain "task_manager_api/Domain"
	infrastructure "task_manager_api/Infrastructure"
	usecases "task_manager_api/UseCases"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userUsecase usecases.IUserUseCase
	jwtService    infrastructure.IJWTService
	passComparator infrastructure.ComparePassword
}

func NewUserController(userUsecase usecases.IUserUseCase, jwtService infrastructure.IJWTService, passCom infrastructure.ComparePassword) *UserController {
	return &UserController{
		userUsecase: userUsecase,
		jwtService:    jwtService,
		passComparator: passCom,
	}
}

func (uc *UserController) Register(c *gin.Context) {
	var user domain.User
	if err := c.BindJSON(&user); err != nil {
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

func (uc *UserController) Login(c *gin.Context) {

	var user domain.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	existingUser, err := uc.userUsecase.GetUser(user.Email)
	if err != nil {
		c.JSON(400, gin.H{"error": "user not found"})
		return
	}

	ok := uc.passComparator(&user, existingUser)
	if !ok {
		c.JSON(500, gin.H{"error": "Incorrect password or email"})
		return 
	}
	SignedToken, err := uc.jwtService.GenerateToken(*existingUser)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(200, gin.H{"message": "User logged in successfully", "token": SignedToken})
}

func (uc *UserController) PromoteUser(c *gin.Context) {
    id := c.Param("id")

    claims, exists := c.Get("claims")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

    claimsMap := claims.(jwt.MapClaims)
    promoter := &domain.User{
        Email: claimsMap["email"].(string),
        Role:  claimsMap["role"].(string),
    }

    updatedUser, err := uc.userUsecase.PromoteUserByID(id, promoter)
    if err != nil {
        c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"user": updatedUser})
}

func (uc *UserController) GetAllUsers(c *gin.Context) {
    users, err := uc.userUsecase.GetAllUsers()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"users": users})
}