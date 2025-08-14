package UseCase

import (
	"task_manager_api_test/Domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ITaskRepo interface {
	StoreTask(task Domain.Task) error
	RetriveAll () []Domain.Task
	RetriveByID(id string) Domain.Task 
	UpdateOneTask(id primitive.ObjectID, updatedTask Domain.Task) error
	DeleteOneTask(id primitive.ObjectID) error
}