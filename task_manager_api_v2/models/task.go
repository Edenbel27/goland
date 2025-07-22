package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
type Task struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title       string    `json:"title" bson:"title"`
	Description string    `json:"description" bson:"description"`
	Status      string    `json:"status" bson:"status"`
	DueDate     time.Time `json:"due_date" bson:"due_date"`
}