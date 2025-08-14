package Repository

import (
	"context"
	"fmt"
	"task_manager_api_test/Domain"
	"task_manager_api_test/UseCase"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)
type UserRepo struct{
	collection *mongo.Collection
}

func NewUserRepo(col *mongo.Collection) UseCase.IUserRepo{
	return &UserRepo{
		collection : col,
	}
}

func (u *UserRepo) StoreUser(user *Domain.User) error {
	
	ctx , cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_ , err := u.collection.InsertOne(ctx, user)
	fmt.Println("in store: ",user)
	return err
}

func (u *UserRepo) CheckEmailAvailablity(email string) (Domain.User,error){
	ctx , cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	var user Domain.User
	filter := bson.M{"email":email}
	result := u.collection.FindOne(ctx, filter)
	err := result.Decode(&user)
	
	return user,err
}