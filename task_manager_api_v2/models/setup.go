package models

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
const ConnectionString = "mongodb://localhost:27017/tasks"
var Db = "tasks"
var CollName = "tasks"
var MongoClient *mongo.Client

func ConnectDatabase() error{
	clientOptions := options.Client().ApplyURI(ConnectionString)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}
	MongoClient = client
	return nil
}