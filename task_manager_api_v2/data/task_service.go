package data

import (
	"context"
	"fmt"
	"task_manager_api_v2/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// var Tasks []models.Task

func AddTask(newTask models.Task) error {
 
	// Tasks = append(Tasks, newTask)
	collection := models.MongoClient.Database(models.Db).Collection(models.CollName)
	inserted, err := collection.InsertOne(context.TODO(), newTask)
	if err != nil {
		return err
	}
	fmt .Println("Inserted a task with ID:", inserted.InsertedID)
	return err
}

func AddTasks(tasks []models.Task) error{
	newTasks := make([]interface{}, len(tasks))
	for i, task := range tasks {
		newTasks[i] = task
	}
	collection := models.MongoClient.Database(models.Db).Collection(models.CollName)
	inserted, err := collection.InsertMany(context.TODO(), newTasks)
	if err != nil {
		return err
	}
	fmt.Println("Inserted tasks with IDs:", inserted.InsertedIDs)
	return err
}

func ViewTasks() []models.Task {
	var results []models.Task

	filter := bson.D{}
	
	collection := models.MongoClient.Database(models.Db).Collection(models.CollName)
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil{
		fmt.Println("Error finding tasks:", err)
		return []models.Task{} 
	}
	err = cursor.All(context.TODO(), &results)
	if err != nil{
		fmt.Println("Error decoding tasks:", err)
		return []models.Task{} 
	}
	return results
}

func ViewTaskByID(taskID string) models.Task {
	var result models.Task
	id , err := primitive.ObjectIDFromHex(taskID)
	if err != nil{
		return models.Task{}
	}
	filter := bson.M{"_id":id}
	collection :=  models.MongoClient.Database(models.Db).Collection(models.CollName)
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil{
		fmt.Println("Task not found or error decoding:", err)
		return models.Task{}
	}
	return result
	// for _, task := range Tasks {
	// 	if task.ID == id {
	// 		return task, nil
	// 	}
	// }
	// return models.Task{}, errors.New("task not found")
}

func UpdateTask(taskID string, updatedTask models.Task) error {
	id, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{"title": updatedTask.Title,
			"description": updatedTask.Description,
			"status":      updatedTask.Status,
			"due_date":    updatedTask.DueDate,
		}}
	collection := models.MongoClient.Database(models.Db).Collection(models.CollName)
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	fmt.Println("New record : ", result)
	return nil
	// for i, task := range Tasks {
	// 	if task.ID == id {
	// 		if updatedTask.Title != "" {
	// 			Tasks[i].Title = updatedTask.Title
	// 		}
	// 		if updatedTask.Description != "" {
	// 			Tasks[i].Description = updatedTask.Description
	// 		}
	// 		if updatedTask.Status != "" {
	// 			Tasks[i].Status = updatedTask.Status
	// 		}
	// 		if !updatedTask.DueDate.IsZero() {
	// 			Tasks[i].DueDate = updatedTask.DueDate
	// 		}
	// 		return nil
	// 	}
	// }
	// return errors.New("task not found")

}

func DeleteTask(taskID string) error {

	id, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": id}
	collection := models.MongoClient.Database(models.Db).Collection(models.CollName)
	result , err := collection.DeleteOne(context.TODO(), filter)
	if err != nil{
		return err
	}
	fmt.Println("Deleted count: ", result.DeletedCount)
	return nil
	// for i, task := range Tasks {
	// 	if task.ID == id {
	// 		Tasks = append(Tasks[:i], Tasks[i+1:]...)
	// 		return nil
	// 	}
	// }
	// return errors.New("task not found")
}