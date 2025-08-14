package UseCase

import (
	"fmt"
	"task_manager_api_test/Domain"
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

type ITaskUseCase interface {
	AddTask(task Domain.Task) error
	ViewTasks() []Domain.Task
	ViewTaskByID(id string) (*Domain.Task, error)
	UpdateTask(id string, task Domain.Task) error
	DeleteTask(id string) error
}

func (t *TaskUseCase) AddTask(task Domain.Task) error {
	err := t.Repo.StoreTask(task)
	if err != nil {
		fmt.Println("user insertion failed")
		return err
	}
	fmt.Println("Inserted a task")
	return nil
}

func (t *TaskUseCase) ViewTasks() []Domain.Task {
	return t.Repo.RetriveAll()
}

func (t *TaskUseCase) ViewTaskByID(taskID string) (*Domain.Task, error) {
	task := t.Repo.RetriveByID(taskID)
	return &task, nil
}

func (t *TaskUseCase) UpdateTask(taskID string, updatedTask Domain.Task) error {
	id, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		return err
	}
	return t.Repo.UpdateOneTask(id, updatedTask)
}

func (t *TaskUseCase) DeleteTask(taskID string) error {
	id, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		return err
	}
	return t.Repo.DeleteOneTask(id)
}