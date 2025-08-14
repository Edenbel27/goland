package usecaseMocks

import (
	"task_manager_api_test/Domain"

	"github.com/stretchr/testify/mock"
)
type PasswordServiceMock struct {
	mock.Mock
}

type RepoMock struct{
	mock.Mock
}

type JWTServiceMock struct{
	mock.Mock
}

// password mock

func (m *PasswordServiceMock) HashPassword (password string)(string, error){
	args := m.Called(password)
	return args.String(0), args.Error(1)
}

func (m *PasswordServiceMock) CompareHashedPassword(hashedPassword string , password string)bool{
	args := m.Called(hashedPassword, password)
	return args.Bool(0)
}

//repo mock

func (m *RepoMock) StoreUser(user *Domain.User)error{
	args := m.Called(user)
	return args.Error(0)
}

func (m *RepoMock) CheckEmailAvailablity(email string) (Domain.User, error){
	args := m.Called(email)
	return args.Get(0).(Domain.User), args.Error(1)
}


//JWT mock

func (m *JWTServiceMock) GenerateToken (email , role string)(string, error){
	args := m.Called(email , role)
	return args.String(0), args.Error(1)
}
