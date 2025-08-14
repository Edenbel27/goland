package Repository

import (
	"task_manager_api_test/Domain"
	"task_manager_api_test/Repository/repositoryMocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestTaskRepositoryInterface(t *testing.T) {

	t.Run("StoreTask Success", func(t *testing.T) {
		mockRepo := new(repositoryMocks.TaskRepositoryMock)
		testTask := Domain.Task{
			Title : "task1",
			Description : "description1",
			Status: "Inprogress",
			DueDate : time.Now().AddDate(0,0,23),
		}

		mockRepo.On("StoreTask", testTask).Return(nil)

		err := mockRepo.StoreTask(testTask)
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("StoreTask Fail", func(t *testing.T) {
	mockRepo := new(repositoryMocks.TaskRepositoryMock)
	expectedErr := assert.AnError
	testTask := Domain.Task{
		Title : "task1",
		Description : "description1",
		Status: "Inprogress",
		DueDate : time.Now().AddDate(0,0,23),
	}

	mockRepo.On("StoreTask", testTask).Return(expectedErr)

	err := mockRepo.StoreTask(testTask)
	assert.ErrorIs(t, err, expectedErr)
	mockRepo.AssertExpectations(t)
})

	t.Run("RetrieveAll Success" , func(t *testing.T){
		mockRepo := new(repositoryMocks.TaskRepositoryMock)
		testTask := Domain.Task{
			Title : "task1",
			Description : "description1",
			Status: "Inprogress",
			DueDate : time.Now().AddDate(0,0,23),
		}
		expectedTasks := []Domain.Task{testTask}

		mockRepo.On("RetriveAll").Return(expectedTasks)

		tasks := mockRepo.RetriveAll()
		assert.Equal(t,expectedTasks, tasks)
		mockRepo.AssertExpectations(t)

	})

	t.Run("ReteriveByID Success" , func(t *testing.T){
	mockRepo := new(repositoryMocks.TaskRepositoryMock)
	testTask := Domain.Task{
		ID: primitive.NewObjectID(),
		Title : "task1",
		Description : "description1",
		Status: "Inprogress",
		DueDate : time.Now().AddDate(0,0,23),
	}

	testID := testTask.ID.Hex()

	mockRepo.On("RetriveByID", testID).Return(testTask)

	task := mockRepo.RetriveByID(testID)

	assert.Equal(t,testTask, task)
	mockRepo.AssertExpectations(t)

})

	t.Run("UpdateOneTask success" , func(t *testing.T){
	mockRepo := new(repositoryMocks.TaskRepositoryMock)
	testTask := Domain.Task{
		ID: primitive.NewObjectID(),
		Title : "task1",
		Description : "description1",
		Status: "Inprogress",
		DueDate : time.Now().AddDate(0,0,23),
	}

	updatedTask := testTask
	updatedTask.Title = "task1 updated"
	mockRepo.On("UpdateOneTask", testTask.ID, updatedTask).Return (nil)
	err := mockRepo.UpdateOneTask(testTask.ID, updatedTask)

	assert.NoError(t,err)
	mockRepo.AssertExpectations(t)

})

	t.Run("UpdateOneTask - document not found" , func(t *testing.T){
	mockRepo := new(repositoryMocks.TaskRepositoryMock)
	nonExistingId := primitive.NewObjectID()
	expectedErr := assert.AnError

	mockRepo.On("UpdateOneTask", nonExistingId, mock.AnythingOfType("Domain.Task")).Return (expectedErr)
	err := mockRepo.UpdateOneTask(nonExistingId, Domain.Task{Title:"Fail"})

	assert.ErrorIs(t,err, expectedErr)
	mockRepo.AssertExpectations(t)

})

	t.Run("UpdateOneTask - invalid data" , func(t *testing.T){
	mockRepo := new(repositoryMocks.TaskRepositoryMock)
	testID := primitive.NewObjectID()
	expectedErr := mongo.WriteException{
		WriteErrors: []mongo.WriteError{{
				Code:121,
				Message: "Document failed validation",}},
	}
	invalidTask := Domain.Task{Title:"",}

	mockRepo.On("UpdateOneTask", testID, invalidTask).Return (expectedErr)
	err := mockRepo.UpdateOneTask(testID, invalidTask)

	var writeErr mongo.WriteException
	assert.ErrorAs(t,err, &writeErr)
	assert.Equal(t,121, writeErr.WriteErrors[0].Code)
	mockRepo.AssertExpectations(t)

})

	t.Run("DeleteOneTask success" , func(t *testing.T){
	mockRepo := new(repositoryMocks.TaskRepositoryMock)
	testTask := Domain.Task{
		ID: primitive.NewObjectID(),
		Title : "task1",
		Description : "description1",
		Status: "Inprogress",
		DueDate : time.Now().AddDate(0,0,23),
	}

	mockRepo.On("DeleteOneTask", testTask.ID).Return (nil)
	err := mockRepo.DeleteOneTask(testTask.ID)

	assert.NoError(t,err)
	mockRepo.AssertExpectations(t)

})

	t.Run("DeleteOneTask Fail" , func(t *testing.T){
	mockRepo := new(repositoryMocks.TaskRepositoryMock)
	nonExistingId := primitive.NewObjectID()
	expectedErr := assert.AnError

	mockRepo.On("DeleteOneTask", nonExistingId).Return (expectedErr)
	err := mockRepo.DeleteOneTask(nonExistingId)

	assert.ErrorIs(t,err, expectedErr)
	mockRepo.AssertExpectations(t)

})

	
}