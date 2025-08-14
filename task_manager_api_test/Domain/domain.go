package Domain

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID primitive.ObjectID
	Email    string
	Password string
	Role     string
}

type Task struct {
	ID          primitive.ObjectID
	Title       string
	Description string
	Status      string
	DueDate     time.Time
}