package Controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"task_manager_api_test/Delivery/Controllers/controllersMocks"
	"task_manager_api_test/Domain"

	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type UserControllerTestSuite struct {
	suite.Suite
	controller  *UserController
	mockUsecase *controllersMocks.MockUserUseCase
	router      *gin.Engine
}

func (suite *UserControllerTestSuite) SetupTest() {
	suite.mockUsecase = new(controllersMocks.MockUserUseCase)
	suite.controller = &UserController{
		UserUsecase: suite.mockUsecase,
	}

	gin.SetMode(gin.TestMode)
	suite.router = gin.Default()
	suite.router.POST("/register", suite.controller.RegisterHandler)
	suite.router.POST("/login", suite.controller.LoginHandler)
}

func (suite *UserControllerTestSuite) testInvalidPayload(endpoint string) {
	testCases := []struct {
		name  string
		input UserDTO
	}{
		{"Empty email", UserDTO{Password: "123", Role: "user"}},
		{"Empty password", UserDTO{Email: "email.com", Role: "user"}},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			body, _ := json.Marshal(tc.input)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", endpoint, bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			suite.router.ServeHTTP(w, req)
			suite.Equal(http.StatusBadRequest, w.Code)
			suite.Contains(w.Body.String(), "invalid request payload")
		})
	}
}

func (suite *UserControllerTestSuite) TestRegisterHandler_Success() {
	suite.Run("Success", func() {
		suite.mockUsecase.On("Register", mock.AnythingOfType("*Domain.User")).Return(nil)

		payload := UserDTO{Email: "test@example.com", Password: "password123", Role: "user"}
		body, _ := json.Marshal(payload)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		suite.router.ServeHTTP(w, req)

		suite.Equal(http.StatusOK, w.Code)
		suite.Contains(w.Body.String(), "test@example.com")
		suite.mockUsecase.AssertExpectations(suite.T())
	})
}

func (suite *UserControllerTestSuite) TestRegisterHandler_InvalidJSON() {
	suite.Run("InvalidJSON", func() {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer([]byte("{invalid_json}")))
		req.Header.Set("Content-Type", "application/json")

		suite.router.ServeHTTP(w, req)
		suite.Equal(http.StatusBadRequest, w.Code)
		suite.Contains(w.Body.String(), "invalid request payload")
	})
}

func (suite *UserControllerTestSuite) TestRegisterHandler_InvalidPayload() {
	suite.Run("InvalidPayload", func() {
		suite.testInvalidPayload("/register")
	})
}

// --- Tests for LoginHandler ---
func (suite *UserControllerTestSuite) TestLoginHandler_Success() {
	suite.Run("Success", func() {
		mockUser := &Domain.User{Email: "test@example.com", Password: "password123"}
		suite.mockUsecase.On("Login", mock.AnythingOfType("*Domain.User")).Return(mockUser, "mock-token", nil)

		payload := UserDTO{Email: "test@example.com", Password: "password123", Role: "user"}
		body, _ := json.Marshal(payload)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		suite.router.ServeHTTP(w, req)

		suite.Equal(http.StatusOK, w.Code)
		suite.Contains(w.Body.String(), "mock-token")
		suite.mockUsecase.AssertExpectations(suite.T())
	})
}

func (suite *UserControllerTestSuite) TestLoginHandler_InvalidCredentials() {
	suite.Run("InvalidCredentials", func() {
		suite.mockUsecase.On("Login", mock.AnythingOfType("*Domain.User")).Return(
			&Domain.User{}, "", errors.New("invalid credentials"),
		)

		payload := UserDTO{Email: "test@example.com", Password: "wrong-password", Role: "user"}
		body, _ := json.Marshal(payload)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		suite.router.ServeHTTP(w, req)

		suite.Equal(http.StatusUnauthorized, w.Code)
		suite.Contains(w.Body.String(), "invalid credentials")
	})
}

func (suite *UserControllerTestSuite) TestLoginHandler_InvalidPayload() {
	suite.Run("InvalidPayload", func() {
		suite.testInvalidPayload("/login")
	})
}

// --- Tests for ChangeToDomain ---
func (suite *UserControllerTestSuite) TestChangeToDomain() {
	suite.Run("ChangeToDomain", func() {
		dto := UserDTO{
			Email:    "test@example.com",
			Password: "password123",
			Role:     "admin",
		}

		domainUser := suite.controller.ChangeToDomain(dto)

		suite.Equal(dto.Email, domainUser.Email)
		suite.Equal(dto.Password, domainUser.Password)
		suite.Equal(dto.Role, domainUser.Role)
	})
}


func TestUserControllerTestSuite(t *testing.T) {
	suite.Run(t, new(UserControllerTestSuite))
}
