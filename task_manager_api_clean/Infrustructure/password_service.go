package Infrustructure

import (
	"task_manager_api_clean/UseCase"

	"golang.org/x/crypto/bcrypt"
)
type PS struct{
}
func NewPasswordService() UseCase.IPasswordService{
	return &PS{}
}


func (ps *PS) HashPassword(password string) (string , error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	password = string(hashedPassword)
	return password , err
}

func (ps *PS) CompareHashedPassword(dbuser , input string)bool{
		return bcrypt.CompareHashAndPassword([]byte(dbuser), []byte(input)) == nil		
}