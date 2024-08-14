package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"task_manager_api/Delivery/controllers"
	"task_manager_api/Domain"
	"task_manager_api/Mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type UserControllerTestSuite struct {
    suite.Suite
    controller *controllers.UserController
    router     *gin.Engine
    mockUsecase *mocks.MockUserUseCase
	mockInfra   *mocks.MockInfrastructure
}

func (suite *UserControllerTestSuite) SetupSuite() {
    suite.mockUsecase = new(mocks.MockUserUseCase)
	suite.mockInfra = new(mocks.MockInfrastructure)

	mockPassComp := func(user, existingUser *domain.User) bool {
		return true
	}

    suite.controller = controllers.NewUserController(suite.mockUsecase, suite.mockInfra, mockPassComp)

    suite.router = gin.Default()
    suite.router.POST("/user/register", suite.controller.Register)
    suite.router.POST("/user/login", suite.controller.Login)
}

func (suite *UserControllerTestSuite) TestRegister() {

    newUser := &domain.User{
        UserName: "John Doe",
        Email:    "john@example.com",
        Password: "password123",
		Role: "ADMIN",
    }

    suite.mockUsecase.On("Register", newUser).Return(nil)

    w := httptest.NewRecorder()
    req, _ := http.NewRequest("POST", "/user/register", strings.NewReader(`{"username":"John Doe","email":"john@example.com","password":"password123", "role": "ADMIN"}`))
    req.Header.Set("Content-Type", "application/json")
    suite.router.ServeHTTP(w, req)

    suite.Equal(http.StatusOK, w.Code)
    suite.mockUsecase.AssertExpectations(suite.T())
}

func (suite *UserControllerTestSuite) TestLogin() {
    existingUser := &domain.User{
        UserName: "John Doe",
        Email:    "john@example.com",
        Password: "password123",
		Role: "ADMIN",
    }

	token := "adsfdfweredsafasdfadfasfsdfasdfas"
    suite.mockUsecase.On("GetUser", "john@example.com").Return(existingUser, nil)
	suite.mockInfra.On("GenerateToken", *existingUser).Return(token, nil)

    w := httptest.NewRecorder()
    req, _ := http.NewRequest("POST", "/user/login", strings.NewReader(`{"email":"john@example.com","password":"password123"}`))
    req.Header.Set("Content-Type", "application/json")
    suite.router.ServeHTTP(w, req)

    suite.Equal(http.StatusOK, w.Code)
    suite.mockUsecase.AssertExpectations(suite.T())
}

func TestUserControllerTestSuite(t *testing.T) {
    suite.Run(t, new(UserControllerTestSuite))
}