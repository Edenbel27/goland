package repositoryMocks

import (
	"task_manager_api_test/Domain"

	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
type TaskRepositoryMock struct {
	mock.Mock
}

 func (t *TaskRepositoryMock) StoreTask (task Domain.Task) error{
	args := t.Called(task)
	return args.Error(0)
 }

 func (t *TaskRepositoryMock) RetriveAll () []Domain.Task {
	args := t.Called()
	return args.Get(0).([]Domain.Task)
 }

func (t *TaskRepositoryMock) RetriveByID(taskID string) Domain.Task{
	args := t.Called(taskID)
	return args.Get(0).(Domain.Task)
  }

func (t *TaskRepositoryMock) UpdateOneTask(id  primitive.ObjectID, updatedTask Domain.Task) error{
	args := t.Called(id , updatedTask)
	return args.Error(0)
}


func (t *TaskRepositoryMock) DeleteOneTask(id primitive.ObjectID) error{

	args := t.Called(id)
	return args.Error(0)
}