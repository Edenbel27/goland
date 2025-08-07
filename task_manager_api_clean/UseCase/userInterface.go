package UseCase

import "task_manager_api_clean/Domain"

type IUserRepo interface {
	StoreUser(user *Domain.User) error
	CheckEmailAvailablity(email string) (Domain.User, error)
}

type IPasswordService interface{
	HashPassword(password string) (string, error)
	CompareHashedPassword(dbuser , input string)bool

}

type IJWTService interface{
	GenerateToken(email , password string ) (string , error)
}

