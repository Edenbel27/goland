package UseCase

import (
	"task_manager_api_test/Domain"
	"task_manager_api_test/UseCase/usecaseMocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestTaskUseCase(t *testing.T) {

	t.Run("Add Task success", func (t *testing.T){
		mockRepo := new(usecaseMocks.TaskRepoMock)
		mockRepo.On("StoreTask", mock.Anything).Return(nil)

		taskUseCase := NewTaskUseCase(mockRepo)
		taskUseCase.AddTask(Domain.Task{
				Title:       "Test Task",
				Description: "This is a test task",
				Status:      "Pending",
				DueDate:     time.Now(),
		})

		mockRepo.AssertExpectations(t)
		
	})

	t.Run("Retrieve All Tasks success", func (t *testing.T){
		mockRepo := new(usecaseMocks.TaskRepoMock)
		testTasks := []Domain.Task{
			{
				Title:       "Test Task",
				Description: "This is a test task",
				Status:      "Pending",
				DueDate:     time.Now(),
			},
		}
		mockRepo.On("RetriveAll").Return(testTasks)

		taskUseCase := NewTaskUseCase(mockRepo)
		tasks := taskUseCase.ViewTasks()


		assert.Equal(t, tasks, testTasks)
		mockRepo.AssertExpectations(t)
		
	})

	t.Run("View By ID success", func (t *testing.T){
		mockRepo := new(usecaseMocks.TaskRepoMock)
		mockRepo.On("RetriveByID", mock.Anything).Return(Domain.Task{Title:"Title"}, nil)

		taskUseCase := NewTaskUseCase(mockRepo)
		task, _ := taskUseCase.ViewTaskByID("1")

		assert.Equal(t, task.Title, "Title")
		mockRepo.AssertExpectations(t)
		
	})

	t.Run("Update task success", func (t *testing.T){
		validID := primitive.NewObjectID()
        hexID := validID.Hex()
		mockRepo := new(usecaseMocks.TaskRepoMock)
		mockRepo.On("UpdateOneTask", validID,mock.AnythingOfType("Domain.Task")).Return(nil)

		taskUseCase := NewTaskUseCase(mockRepo)
		err := taskUseCase.UpdateTask(hexID, Domain.Task{Title: "Updated Title"})

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
		
	})

	t.Run("Delete task success", func (t *testing.T){
		validID := primitive.NewObjectID()
        hexID := validID.Hex()
		mockRepo := new(usecaseMocks.TaskRepoMock)
		mockRepo.On("DeleteOneTask", validID).Return(nil)

		taskUseCase := NewTaskUseCase(mockRepo)
		err := taskUseCase.DeleteTask(hexID)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)

	})
}