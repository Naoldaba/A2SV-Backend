package infrastructure_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	infrastructure "task_manager_api/Infrastructure"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/suite"
)

type MiddlewareTestSuite struct {
	suite.Suite
	jwtService infrastructure.IJWTService
	secretKey  []byte
}

func (suite *MiddlewareTestSuite) SetupSuite() {
	err := godotenv.Load()
	suite.NoError(err)

	suite.secretKey = []byte(os.Getenv("SECRET_KEY")) 
	suite.jwtService = infrastructure.NewJWTService(suite.secretKey)
}

func (suite *MiddlewareTestSuite) TestJWTValidation_NoAuthorizationHeader() {
	router := gin.New()

	router.Use(infrastructure.JWTValidation())

	req, _ := http.NewRequest("GET", "/tasks", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	suite.Equal(http.StatusUnauthorized, resp.Code)
	suite.Contains(resp.Body.String(), "Authorization header is required")
}

func (suite *MiddlewareTestSuite) TestJWTValidation_InvalidAuthorizationHeader() {
	router := gin.New()
	router.Use(infrastructure.JWTValidation())

	req, _ := http.NewRequest("GET", "/tasks", nil)
	req.Header.Set("Authorization", "BearerInvalidToken") 
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	suite.Equal(http.StatusUnauthorized, resp.Code)
	suite.Contains(resp.Body.String(), "Invalid authorization header")
}

func (suite *MiddlewareTestSuite) TestJWTValidation_InvalidToken() {
	router := gin.New()
	router.Use(infrastructure.JWTValidation())

	req, _ := http.NewRequest("GET", "/tasks", nil)
	req.Header.Set("Authorization", "Bearer invalid_token") 
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	suite.Equal(http.StatusUnauthorized, resp.Code)
	suite.Contains(resp.Body.String(), "Invalid JWT")
}

func TestMiddlewareTestSuite(t *testing.T) {
	suite.Run(t, new(MiddlewareTestSuite))
}