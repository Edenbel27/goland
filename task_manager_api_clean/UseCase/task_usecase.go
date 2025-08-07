package UseCase

import (
	"fmt"
	"task_manager_api_clean/Domain"

	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskUseCase struct {
	Repo ITaskRepo
}

func NewTaskUseCase(repo ITaskRepo) *TaskUseCase {
	return &TaskUseCase{
		Repo: repo,
	}
}

func (t *TaskUseCase) AddTask(task Domain.Task) {
	err := t.Repo.StoreTask(task)
	if err != nil{
		fmt.Println("user insertion failed")
		return
	}
	fmt.Println("Inserted a task")
}

func (t *TaskUseCase) ViewTasks() []Domain.Task {
	return t.Repo.RetriveAll()
}

func (t *TaskUseCase) ViewTaskByID(taskID string) Domain.Task {
	return t.Repo.RetriveByID(taskID)
}

func (t *TaskUseCase) UpdateTask(taskID string, updatedTask Domain.Task) error {
	id, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		return err
	}
	result := t.Repo.UpdateOneTask(id , updatedTask)
	return result
}

func (t *TaskUseCase) DeleteTask(taskID string) error {

	id, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		return err
	}
	result := t.Repo.DeleteOneTask(id)
	return result
}