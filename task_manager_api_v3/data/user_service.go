package data

import (
	"context"
	"task_manager_api_v3/models"

	"go.mongodb.org/mongo-driver/bson"
)
func FindOneEmail(user models.User) (models.User, error){
	var dbuser models.User
	userCol := models.GetCollection("users")
	filter := bson.M{"email": user.Email}
	ok := userCol.FindOne(context.TODO(), filter).Decode(&dbuser)
	if ok == nil{
		return dbuser ,nil
	}
	return models.User{}, ok

}
func InsertOneUser(user models.User) error{
	userCol := models.GetCollection("users")
	_, err := userCol.InsertOne(context.TODO(), user)
	return err
}

	