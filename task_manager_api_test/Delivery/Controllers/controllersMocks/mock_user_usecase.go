package controllersMocks

import (
	"task_manager_api_test/Domain"

	"github.com/stretchr/testify/mock"
)

type MockUserUseCase struct {
	mock.Mock
}

func (m *MockUserUseCase) Register(user *Domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserUseCase) Login(user *Domain.User) (*Domain.User, string, error) {
	args := m.Called(user)
	return args.Get(0).(*Domain.User), args.String(1), args.Error(2)
}
