package infrastructure

import (
	"fmt"
	"log"
	"os"
	"task_manager_api/Domain"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

var jwtSecret = []byte(os.Getenv("SECRET_KEY"))

type TokenGenerator func(user domain.User) (string, error)

var GenerateToken TokenGenerator = func(existingUser domain.User) (string, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": existingUser.ID,
		"email":   existingUser.Email,
		"role":    existingUser.Role,
	})

	SignedToken, err := token.SignedString(jwtSecret)

	return SignedToken, err
}

func ValidateToken(req_token string) (*jwt.Token, error)   {
	token, err := jwt.Parse(req_token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		  return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	  })

	return token, err
}