package repositoryMocks

import (
	"github.com/stretchr/testify/mock"
	"task_manager_api_test/Domain"
)

type UserRepoSitoryMock struct{
	mock.Mock
}

func (m * UserRepoSitoryMock) StoreUser (user *Domain.User) error{
	args := m.Called(user)
	return args.Error(0)
}

func (m *UserRepoSitoryMock) CheckEmailAvailablity(email string)(Domain.User, error){
	args := m.Called(email)
	return args.Get(0).(Domain.User), args.Error(1)
} 

