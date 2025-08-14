package usecaseMocks

import (
	"task_manager_api_test/Domain"

	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskRepoMock struct {
	mock.Mock
}


func (m *TaskRepoMock) StoreTask(task Domain.Task) error {
	args := m.Called(task)
	return args.Error(0)
}

func (m *TaskRepoMock) RetriveAll () []Domain.Task {
	args := m.Called()
	return args.Get(0).([]Domain.Task)
}
func (m *TaskRepoMock) RetriveByID(id string) Domain.Task{
	args := m.Called(id)
	return args.Get(0).(Domain.Task)
}
func (m *TaskRepoMock) UpdateOneTask(id primitive.ObjectID, updatedTask Domain.Task) error {
	args := m.Called(id, updatedTask)
	return args.Error(0)
}

func (m *TaskRepoMock) DeleteOneTask(id primitive.ObjectID) error {
	args := m.Called(id)
	return args.Error(0)
}
