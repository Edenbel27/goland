package Repository

import (
	"context"
	"fmt"

	"task_manager_api_clean/Domain"
	"task_manager_api_clean/UseCase"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepo struct {
	collection *mongo.Collection
	context context.Context
}

func NewTaskRepo() UseCase.ITaskRepo{
	col :=  GetCollection("tasks")
	ctx := context.Background()
	return &TaskRepo{
		collection : col,
		context : ctx,
	}
}
 func (t *TaskRepo) StoreTask (task Domain.Task) error{
	_ , err := t.collection.InsertOne(t.context, task)
	return err
 }

 func (t *TaskRepo) RetriveAll () []Domain.Task {
	var results []Domain.Task

	filter := bson.D{}
	
	cursor, err := t.collection.Find(t.context, filter)
	if err != nil{
		fmt.Println("Error finding tasks:", err)
		return []Domain.Task{} 
	}
	err = cursor.All(context.TODO(), &results)
	if err != nil{
		fmt.Println("Error decoding tasks:", err)
		return []Domain.Task{} 
	}
	return results
 }

func (t *TaskRepo) RetriveByID(taskID string) Domain.Task{
	var result Domain.Task
	id , err := primitive.ObjectIDFromHex(taskID)
	if err != nil{
		return Domain.Task{}
	}
	filter := bson.M{"_id":id}
	err = t.collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil{
		fmt.Println("Task not found or error decoding:", err)
		return Domain.Task{}
	}
	return result
  }

func (t *TaskRepo) UpdateOneTask(id  primitive.ObjectID, updatedTask Domain.Task) error{

	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{"title": updatedTask.Title,
			"description": updatedTask.Description,
			"status":      updatedTask.Status,
			"due_date":    updatedTask.DueDate,
		}}
	result, err := t.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	fmt.Println("New record : ", result)
	return nil
}


func (t *TaskRepo) DeleteOneTask(id primitive.ObjectID) error{

	filter := bson.M{"_id": id}
	result , err := t.collection.DeleteOne(context.TODO(), filter)
	if err != nil{
		return err
	}
	fmt.Println("Deleted count: ", result.DeletedCount)
	return nil
}
