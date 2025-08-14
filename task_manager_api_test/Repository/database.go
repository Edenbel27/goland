package Repository

import (
	"context"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
var MongoClient *mongo.Client

func ConnectDatabase() error{
	ctx , cancel := context.WithTimeout(context.Background(), 10 *time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(ctx , clientOptions)
	if err != nil {
		panic(err)
	}
	MongoClient = client
	return nil
}

func GetCollection (name string) *mongo.Collection{
	return MongoClient.Database("task_manager_clean").Collection(name)
}