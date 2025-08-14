package controllersMocks

import (
	"task_manager_api_test/Domain"

	"github.com/stretchr/testify/mock"
)

type MockTaskUseCase struct {
	mock.Mock
}

func (m *MockTaskUseCase) AddTask(task Domain.Task) error {
	args := m.Called(task)
	return args.Error(0)
}

func (m *MockTaskUseCase) ViewTasks() []Domain.Task {
	args := m.Called()
	return args.Get(0).([]Domain.Task)
}

func (m *MockTaskUseCase) ViewTaskByID(id string) (*Domain.Task, error) {
	args := m.Called(id)
	return args.Get(0).(*Domain.Task), args.Error(1)
}

func (m *MockTaskUseCase) UpdateTask(id string, task Domain.Task) error {
	args := m.Called(id, task)
	return args.Error(0)
}

func (m *MockTaskUseCase) DeleteTask(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
