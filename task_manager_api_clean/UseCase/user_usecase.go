package UseCase

import (
	"fmt"
	"task_manager_api_clean/Domain"

	"go.mongodb.org/mongo-driver/mongo"
)
type UserUseCase struct{
	Repo IUserRepo
	PS IPasswordService
	JWTS IJWTService
}

func NewUserUseCase(repo IUserRepo , ps IPasswordService , jwts IJWTService ) *UserUseCase{
	return &UserUseCase{
		Repo : repo,
		PS:ps,
		JWTS:jwts,
	}
}

func (u *UserUseCase) Register(user *Domain.User){
	// logic
	// check availablity
	_, err := u.Repo.CheckEmailAvailablity(user.Email)
	if err == nil {
		fmt.Println("email already exist")
		return 
}
    
	// hash password
	hashed, err:= u.PS.HashPassword(user.Password)
	if err != nil{
		fmt.Println("error in hashing")
		return
	}
	user.Password = hashed
	if user.Role ==""{
		user.Role = "user"
	}
	
	// store the user
	err = u.Repo.StoreUser(user)
	if err != nil{
		fmt.Println("user registration failed")
		return 
	}
}

func (u *UserUseCase) Login (user *Domain.User)(*Domain.User, string , error){
	existing , err := u.Repo.CheckEmailAvailablity(user.Email)
	if err != nil{
		
		if err == mongo.ErrNoDocuments {
			// User with that email doesn't exist
			return &Domain.User{},"",  fmt.Errorf("user does not exist")

		} else {
			return &Domain.User{},"", fmt.Errorf("database error: %v", err)

		} 
	} 
	if !u.PS.CompareHashedPassword(existing.Password , user.Password){
		return &Domain.User{},"", fmt.Errorf("invalid password")

	}

	jwtToken, err := u.JWTS.GenerateToken(existing.Email , existing.Role)
	if err != nil{

		return &Domain.User{},"", fmt.Errorf("token generation failed")

	}

	return &existing, jwtToken, nil

}
