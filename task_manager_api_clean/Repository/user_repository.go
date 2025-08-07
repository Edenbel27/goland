package Repository

import (
	"fmt"
	"context"
	"task_manager_api_clean/Domain"
	"task_manager_api_clean/UseCase"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"

)
type UserRepo struct{
	collection *mongo.Collection
	context context.Context
}

func NewUserRepo() UseCase.IUserRepo{
	col :=  GetCollection("users")
	ctx := context.Background()
	return &UserRepo{
		collection : col,
		context : ctx,
	}
}

func (u *UserRepo) StoreUser(user *Domain.User) error {
	
	_ , err := u.collection.InsertOne(u.context, user)
	fmt.Println("in store: ",user)
	return err
}

func (u *UserRepo) CheckEmailAvailablity(email string) (Domain.User,error){
	var user Domain.User
	filter := bson.M{"email":email}
	result := u.collection.FindOne(u.context, filter)
	err := result.Decode(&user)
	
	return user,err
}