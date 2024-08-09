package controllers

import (
	"log"
	"os"
	"task_manager_api/data"
	"task_manager_api/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection
var userService data.UserService

func init() {
	var err error
	userService = data.UserService{}

    Client, err := data.DbInstance() 
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	userCollection = data.OpenCollection(Client, "Users")
}

func Register(c *gin.Context){
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
	c.JSON(500, gin.H{"error": "Internal server error"})
	return
	}
	user.Password = string(hashedPassword)

	err = userService.Register(userCollection, user)
	if err !=  nil {
		c.JSON(500, gin.H{"message": err})
		return 
	}

	c.JSON(200, gin.H{"message": "User registered successfully"})
}

func Login(c *gin.Context){
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
		return
	}
	
	var user models.User
	if err = c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	var jwtSecret = []byte(os.Getenv("SECRET_KEY"))

	existingUser, err := userService.GetUser(userCollection, user.Email)
	if err != nil || bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password)) != nil {
		c.JSON(401, gin.H{"error": "Invalid email or password"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": existingUser.ID,
		"email": existingUser.Email,
		"role": existingUser.Role,
	})

	SignedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(200, gin.H{"message": "User logged in successfully", "token": SignedToken})
}