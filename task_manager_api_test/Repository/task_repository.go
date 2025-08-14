package Repository

import (
	"context"
	"fmt"
	"time"
	"task_manager_api_test/Domain"
	"task_manager_api_test/UseCase"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepo struct {
	collection *mongo.Collection
}

func NewTaskRepo(col *mongo.Collection) UseCase.ITaskRepo{
	return &TaskRepo{
		collection : col,
	}
}
 func (t *TaskRepo) StoreTask (task Domain.Task) error{
	ctx , cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_ , err := t.collection.InsertOne(ctx, task)
	return err
 }

 func (t *TaskRepo) RetriveAll () []Domain.Task {
	ctx , cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var results []Domain.Task

	filter := bson.D{}
	
	cursor, err := t.collection.Find(ctx, filter)
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
	ctx , cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var result Domain.Task
	id , err := primitive.ObjectIDFromHex(taskID)
	if err != nil{
		return Domain.Task{}
	}
	filter := bson.M{"_id":id}
	err = t.collection.FindOne(ctx, filter).Decode(&result)
	if err != nil{
		fmt.Println("Task not found or error decoding:", err)
		return Domain.Task{}
	}
	return result
  }

func (t *TaskRepo) UpdateOneTask(id  primitive.ObjectID, updatedTask Domain.Task) error{
	ctx , cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{"title": updatedTask.Title,
			"description": updatedTask.Description,
			"status":      updatedTask.Status,
			"due_date":    updatedTask.DueDate,
		}}
	result, err := t.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	fmt.Println("New record : ", result)
	return nil
}


func (t *TaskRepo) DeleteOneTask(id primitive.ObjectID) error{
	ctx , cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.M{"_id": id}
	result , err := t.collection.DeleteOne(ctx, filter)
	if err != nil{
		return err
	}
	fmt.Println("Deleted count: ", result.DeletedCount)
	return nil
}
