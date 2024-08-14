package infrastructure

import (
	"fmt"
	"os"
	"testing"

	domain "task_manager_api/Domain"
	infrastructure "task_manager_api/Infrastructure"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JWTServiceTestSuite struct {
	suite.Suite
	jwtService infrastructure.IJWTService
	secretKey  []byte
}

func (suite *JWTServiceTestSuite) SetupSuite() {
	err := godotenv.Load()
	suite.NoError(err)

	suite.secretKey = []byte(os.Getenv("SECRET_KEY"))
	suite.jwtService = infrastructure.NewJWTService(suite.secretKey)
}

func (suite *JWTServiceTestSuite) TestGenerateToken() {
	user := domain.User{
		ID:    primitive.NewObjectID(),
		Email: "test@example.com",
		Role:  "user",
	}

	token, err := suite.jwtService.GenerateToken(user)
	suite.NoError(err)
	suite.NotEmpty(token)

	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return suite.secretKey, nil
	})

	suite.NoError(err)
	suite.NotNil(parsedToken)

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	suite.True(ok)
	suite.Equal((user.ID).Hex(), claims["user_id"])
	suite.Equal("test@example.com", claims["email"])
	suite.Equal("user", claims["role"])
}

func (suite *JWTServiceTestSuite) TestValidateToken() {
	user := domain.User{
		ID:    primitive.NewObjectID(),
		Email: "test@example.com",
		Role:  "user",
	}

	token, err := suite.jwtService.GenerateToken(user)
	suite.NoError(err)

	parsedToken, err := suite.jwtService.ValidateToken(token)
	suite.NoError(err)
	suite.NotNil(parsedToken)

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	suite.True(ok)
	suite.Equal((user.ID).Hex(), claims["user_id"])
	suite.Equal("test@example.com", claims["email"])
	suite.Equal("user", claims["role"])
}

func TestJWTServiceTestSuite(t *testing.T) {
	suite.Run(t, new(JWTServiceTestSuite))
}
