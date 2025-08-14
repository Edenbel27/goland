package Routers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Define interfaces matching the real controllers for type compatibility
type userController interface {
	RegisterHandler(*gin.Context)
	LoginHandler(*gin.Context)
}
type taskController interface {
	ViewTasksController(*gin.Context)
	ViewTaskByIDController(*gin.Context)
	AddTaskController(*gin.Context)
	UpdateTaskController(*gin.Context)
	DeleteTaskController(*gin.Context)
}

// Mock controllers with minimal handlers for testing
type mockUserController struct{}
func (m *mockUserController) RegisterHandler(c *gin.Context) { c.JSON(200, gin.H{"message": "register"}) }
func (m *mockUserController) LoginHandler(c *gin.Context)    { c.JSON(200, gin.H{"message": "login"}) }

type mockTaskController struct{}
func (m *mockTaskController) ViewTasksController(c *gin.Context)      { c.JSON(200, gin.H{"tasks": []string{}}) }
func (m *mockTaskController) ViewTaskByIDController(c *gin.Context)   { c.JSON(200, gin.H{"task": "task"}) }
func (m *mockTaskController) AddTaskController(c *gin.Context)        { c.JSON(201, gin.H{"message": "added"}) }
func (m *mockTaskController) UpdateTaskController(c *gin.Context)     { c.JSON(200, gin.H{"message": "updated"}) }
func (m *mockTaskController) DeleteTaskController(c *gin.Context)     { c.JSON(200, gin.H{"message": "deleted"}) }

// Provide a version of SetupRouter for testing that accepts interfaces
func setupTestRouter(uc userController, tc taskController) *gin.Engine {
	router := gin.Default()
	router.POST("/register", uc.RegisterHandler)
	router.POST("/login", uc.LoginHandler)
	router.GET("/tasks", tc.ViewTasksController)
	router.GET("/tasks/:id", tc.ViewTaskByIDController)
	router.POST("/tasks", tc.AddTaskController)
	router.PUT("/tasks/:id", tc.UpdateTaskController)
	router.DELETE("/tasks/:id", tc.DeleteTaskController)
	return router
}

func TestRouterRoutes(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := setupTestRouter(&mockUserController{}, &mockTaskController{})

	tests := []struct {
		method string
		path   string
		status int
	}{
		{"POST", "/register", 200},
		{"POST", "/login", 200},
		{"GET", "/tasks", 200},
		{"GET", "/tasks/123", 200},
		{"POST", "/tasks", 201},
		{"PUT", "/tasks/123", 200},
		{"DELETE", "/tasks/123", 200},
	}

	for _, tt := range tests {
		req, _ := http.NewRequest(tt.method, tt.path, nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, tt.status, w.Code, tt.method+" "+tt.path)
	}
}
