package Infrustructure

import (
	"task_manager_api_test/UseCase"

	"golang.org/x/crypto/bcrypt"
)

type PS struct {
}

func NewPasswordService() UseCase.IPasswordService {
	return &PS{}
}

func (ps *PS) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	finalPassword := string(hashedPassword)
	return finalPassword, err
}

func (ps *PS) CompareHashedPassword(hashedPassword, plainPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword)) == nil
}
