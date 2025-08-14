package Infrustructure

import (
	"fmt"
	"os"
	"task_manager_api_test/UseCase"
	"github.com/dgrijalva/jwt-go"
)

type JWT struct {
}

func NewJWTService() UseCase.IJWTService {
	return &JWT{}
}

func (j *JWT) GenerateToken(email , role string ) (string , error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"role": role,
	})

	jwtToken, err := token.SignedString([]byte(os.Getenv("SECRETKEY")))
	if err != nil {
		fmt.Println("error : Token generation failed")
		return "", err
	}
	return jwtToken, nil
}