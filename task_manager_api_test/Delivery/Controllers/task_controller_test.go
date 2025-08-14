package Controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"task_manager_api_test/Delivery/Controllers/controllersMocks"
	"task_manager_api_test/Domain"
	"testing"

	// "task_manager_api_test/Controllers/controllersMocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func createTestContext(method, path string, body interface{}) (*httptest.ResponseRecorder, *gin.Context) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	var buf bytes.Buffer
	if body != nil {
		json.NewEncoder(&buf).Encode(body)
	}

	c.Request, _ = http.NewRequest(method, path, &buf)
	c.Request.Header.Set("Content-Type", "application/json")
	return w, c
}

func TestAddTaskController_Success(t *testing.T) {
	mockUsecase := new(controllersMocks.MockTaskUseCase)
	controller := NewTaskController(mockUsecase)

	mockUsecase.On("AddTask", mock.AnythingOfType("Domain.Task")).Return(nil)

	w, c := createTestContext("POST", "/tasks", TaskDTO{
		Title:       "Test Task",
		Description: "Test Description",
		Status:      "pending",
	})

	controller.AddTaskController(c)

	assert.Equal(t, http.StatusOK, w.Code)
	mockUsecase.AssertExpectations(t)
}

func TestAddTaskController_InvalidPayload(t *testing.T) {
	mockUsecase := new(controllersMocks.MockTaskUseCase)
	controller := NewTaskController(mockUsecase)

	testCases := []struct {
		name        string
		payload     interface{}
		expectedErr string
	}{
		{
			name:        "empty body",
			payload:     nil,
			expectedErr: "EOF",
		},
		{
			name:        "invalid json",
			payload:     "{invalid}",
			expectedErr: "invalid character",
		},
		{
			name: "missing title",
			payload: map[string]interface{}{
				"description": "Test Description",
			},
			expectedErr: "Title",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			w, c := createTestContext("POST", "/tasks", tc.payload)

			controller.AddTaskController(c)

			assert.Equal(t, http.StatusBadRequest, w.Code)

			var response map[string]interface{}
			err := json.Unmarshal(w.Body.Bytes(), &response)
			if err != nil {
				t.Fatalf("Failed to unmarshal response: %v", err)
			}
			assert.Contains(t, response["error"], tc.expectedErr)
		})
	}
}

func TestViewTasksController_Success(t *testing.T) {
	mockUsecase := new(controllersMocks.MockTaskUseCase)
	controller := NewTaskController(mockUsecase)

	expectedTasks := []Domain.Task{
		{
			Title:       "Task 1",
			Description: "Description 1",
			Status:      "pending",
		},
	}
	mockUsecase.On("ViewTasks").Return(expectedTasks)

	w, c := createTestContext("GET", "/tasks", nil)
	controller.ViewTasksController(c)

	assert.Equal(t, http.StatusOK, w.Code)

	var tasks []Domain.Task
	json.Unmarshal(w.Body.Bytes(), &tasks)
	assert.Equal(t, expectedTasks, tasks)
	mockUsecase.AssertExpectations(t)
}

func TestViewTaskByIDController_Success(t *testing.T) {
	mockUsecase := new(controllersMocks.MockTaskUseCase)
	controller := NewTaskController(mockUsecase)

	taskID := primitive.NewObjectID().Hex()
	expectedTask := &Domain.Task{
		ID:          primitive.NewObjectID(),
		Title:       "Test Task",
		Description: "Test Description",
	}
	mockUsecase.On("ViewTaskByID", taskID).Return(expectedTask, nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("GET", "/tasks/"+taskID, nil)
	c.Params = []gin.Param{{Key: "id", Value: taskID}}

	controller.ViewTaskByIDController(c)

	assert.Equal(t, http.StatusOK, w.Code)

	var task Domain.Task
	json.Unmarshal(w.Body.Bytes(), &task)
	assert.Equal(t, *expectedTask, task)
	mockUsecase.AssertExpectations(t)
}

func TestUpdateTaskController_Success(t *testing.T) {
	mockUsecase := new(controllersMocks.MockTaskUseCase)
	controller := NewTaskController(mockUsecase)

	taskID := primitive.NewObjectID().Hex()
	mockUsecase.On("UpdateTask", taskID, mock.AnythingOfType("Domain.Task")).Return(nil)

	w, c := createTestContext("PUT", "/tasks/"+taskID, TaskDTO{
		Title:       "Updated Task",
		Description: "Updated Description",
	})
	c.Params = []gin.Param{{Key: "id", Value: taskID}}

	controller.UpdateTaskController(c)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "Task updated successfully", response["message"])
	mockUsecase.AssertExpectations(t)
}

func TestDeleteTaskController_Success(t *testing.T) {
	mockUsecase := new(controllersMocks.MockTaskUseCase)
	controller := NewTaskController(mockUsecase)

	taskID := primitive.NewObjectID().Hex()
	mockUsecase.On("DeleteTask", taskID).Return(nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("DELETE", "/tasks/"+taskID, nil)
	c.Params = []gin.Param{{Key: "id", Value: taskID}}

	controller.DeleteTaskController(c)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "Task deleted successfully", response["message"])
	mockUsecase.AssertExpectations(t)
}


