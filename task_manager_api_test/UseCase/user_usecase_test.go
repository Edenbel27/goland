package UseCase

import (
	"errors"
	"task_manager_api_test/Domain"
	"task_manager_api_test/UseCase/usecaseMocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestUserUseCase(t *testing.T) {

	testUser := Domain.User{
		Email:    "email",
		Password: "password",
		Role:     "",
	}
t.Run("Registration", func(t *testing.T){

	t.Run("Register User Success" , func (t *testing.T){
		passwordMock := new(usecaseMocks.PasswordServiceMock)
		repoMock := new(usecaseMocks.RepoMock)
		jwtmock := new(usecaseMocks.JWTServiceMock)

		userUseCase := NewUserUseCase(repoMock,passwordMock, jwtmock)
		passwordMock.On("HashPassword", "password").Return("hashed_password", nil )
		repoMock.On("CheckEmailAvailablity", testUser.Email).Return (Domain.User{}, mongo.ErrNoDocuments)
		repoMock.On("StoreUser", mock.MatchedBy(func(u *Domain.User) bool{
			return u.Email == testUser.Email &&
					u.Password == "hashed_password" &&
					u.Role == "user"
		})).Return(nil)

		err := userUseCase.Register(&testUser)

		passwordMock.AssertExpectations(t)
		repoMock.AssertExpectations(t)

		assert.Equal(t, "hashed_password", testUser.Password)
		assert.Equal(t, "user", testUser.Role)
		assert.NoError(t, err)
		})


	t.Run("Register User with Email already existing" , func(t *testing.T){
		repoMock := new(usecaseMocks.RepoMock)


		userUseCase := NewUserUseCase(repoMock,nil, nil)
		repoMock.On("CheckEmailAvailablity", testUser.Email).Return (testUser, nil)
		err := userUseCase.Register(&testUser)


		assert.Error(t, err)
		assert.Contains(t, err.Error(), "email already exists")
		repoMock.AssertExpectations(t)

	})

	t.Run ("Hashing password failure when registering", func(t *testing.T){
		passwordMock := new(usecaseMocks.PasswordServiceMock)
		repoMock := new(usecaseMocks.RepoMock)

		userUseCase := NewUserUseCase(repoMock, passwordMock , nil)
		repoMock.On("CheckEmailAvailablity", testUser.Email).Return(Domain.User{}, errors.New("not found"))
		passwordMock.On("HashPassword", testUser.Password).Return("",errors.New("hash failed"))
		err := userUseCase.Register(&testUser)

		passwordMock.AssertExpectations(t)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "hash failed")
		repoMock.AssertNotCalled(t, "StoreUser")
	})

	t.Run("Register storing failure" , func(t *testing.T){
		passwordMock := new(usecaseMocks.PasswordServiceMock)
		repoMock := new(usecaseMocks.RepoMock)

		userUseCase := NewUserUseCase(repoMock, passwordMock , nil)
		repoMock.On("CheckEmailAvailablity", testUser.Email).Return(Domain.User{}, errors.New("not found"))
		passwordMock.On("HashPassword", testUser.Password).
            Return("hashed_password", nil)
		repoMock.On("StoreUser", &testUser).Return(errors.New("store failed"))
		err := userUseCase.Register(&testUser)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "store failed")
		repoMock.AssertExpectations(t)
		passwordMock.AssertExpectations(t)
	})

})

t.Run("Login", func(t *testing.T){
	t.Run("Login userSuccess", func(t *testing.T){
		passwordMock := new(usecaseMocks.PasswordServiceMock)
		repoMock := new(usecaseMocks.RepoMock)
		jwtMock := new(usecaseMocks.JWTServiceMock)

		userUseCase := NewUserUseCase(repoMock, passwordMock , jwtMock)
		passwordMock.On("CompareHashedPassword", "hashed_password", testUser.Password).Return(true)
		jwtMock.On("GenerateToken", testUser.Email, testUser.Role).Return("token", nil)
		repoMock.On("CheckEmailAvailablity", testUser.Email).Return(testUser, nil)

		userUseCase.Login(&testUser)

		passwordMock.AssertExpectations(t)
		repoMock.AssertExpectations(t)
		jwtMock.AssertExpectations(t)
	})

	t.Run("Login User with NO user found", func(t *testing.T){
		passwordMock := new(usecaseMocks.PasswordServiceMock)
		repoMock := new(usecaseMocks.RepoMock)
		jwtMock := new(usecaseMocks.JWTServiceMock)

		userUseCase := NewUserUseCase(repoMock, passwordMock , jwtMock)
		repoMock.On("CheckEmailAvailablity", "nonexistingemail").Return(Domain.User{}, errors.New("user not found"))

		userUseCase.Login(&Domain.User{Email:"nonexistingemail"})

		repoMock.AssertExpectations(t)
	})

	t.Run("Login User with hash comparison failure", func(t *testing.T){
		passwordMock := new(usecaseMocks.PasswordServiceMock)
		repoMock := new(usecaseMocks.RepoMock)
		jwtMock := new(usecaseMocks.JWTServiceMock)


		testUser.Password = "hashed_password"
		input_pass := "password"
		userUseCase := NewUserUseCase(repoMock, passwordMock , jwtMock)
		repoMock.On("CheckEmailAvailablity", testUser.Email).Return(testUser, nil)
		passwordMock.On("CompareHashedPassword", testUser.Password, input_pass).Return(false)

		loginUser := testUser
		loginUser.Password = input_pass
		userUseCase.Login(&loginUser)

		passwordMock.AssertExpectations(t)
	})

	t.Run("Login User with token generation failure", func(t *testing.T){
		passwordMock := new(usecaseMocks.PasswordServiceMock)
		repoMock := new(usecaseMocks.RepoMock)
		jwtMock := new(usecaseMocks.JWTServiceMock)

		userUseCase := NewUserUseCase(repoMock, passwordMock , jwtMock)

		passwordMock.On("CompareHashedPassword", "hashed_password", testUser.Password).Return(true)
		repoMock.On("CheckEmailAvailablity", testUser.Email).Return(testUser, nil)
		jwtMock.On("GenerateToken", testUser.Email, testUser.Role).Return("", errors.New("token generation failed"))


		_, _, err := userUseCase.Login(&testUser)
		assert.Contains(t, "token generation failed", err.Error())
		jwtMock.AssertExpectations(t)
	})

})

}