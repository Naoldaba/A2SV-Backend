package infrastructure

import (
	"fmt"
	"log"
	domain "task_manager_api/Domain"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

type IJWTService interface {
	ValidateToken(token string) (*jwt.Token, error)
	GenerateToken(user domain.User) (string, error)
}

type JWTService struct {
	secretKey []byte
}

func NewJWTService(secretKey []byte) IJWTService {
	return &JWTService{
		secretKey: secretKey,
	}
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

func (svc *JWTService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return svc.secretKey, nil
	})
}

func (svc *JWTService) GenerateToken(user domain.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"role":    user.Role,
	})

	signedToken, err := token.SignedString(svc.secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
