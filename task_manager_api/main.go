package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Task struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	DueDate     time.Time `json:"duedate"`
}

var tasks = []Task{
	{ID: "1", Title: "Task 1", Description: "First task", DueDate: time.Now(), Status: "Pending"},
	{ID: "2", Title: "Task 2", Description: "Second task", DueDate: time.Now().AddDate(0, 0, 1), Status: "In Progress"},
	{ID: "3", Title: "Task 3", Description: "Third task", DueDate: time.Now().AddDate(0, 0, 2), Status: "Completed"},
}

// GET /tasks
// Description: Retrieve all tasks.
// Response: { "tasks": [ { "id": "...", "title": "...", "description": "...", "status": "...", "duedate": "..." }, ... ] }
func getTasks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"tasks": tasks})
}

// GET /tasks/:id
// Description: Retrieve a specific task by ID.
// Response: { "id": "...", "title": "...", "description": "...", "status": "...", "duedate": "..." } or { "error": "Task not found" }
func getTask(ctx *gin.Context) {
	id := ctx.Param("id")
	for _, task := range tasks {
		if task.ID == id {
			ctx.IndentedJSON(http.StatusOK, task)
			return
		}
	}

	ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": "Task not found"})

}

// PUT /tasks/:id
// Description: Update a specific task by ID.
// Request: { "title": "...", "description": "...", "status": "...", "duedate": "..." }
// Response: { "message": "Task updated" } or { "error": "...}
func updateTask(ctx *gin.Context) {
	id := ctx.Param("id")
	var updatedTask Task

	if err := ctx.ShouldBindJSON(&updatedTask); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i, task := range tasks {
		if task.ID == id {
			if updatedTask.Title != "" {
				tasks[i].Title = updatedTask.Title
			}

			if updatedTask.Description != "" {
				tasks[i].Description = updatedTask.Description
			}
			ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Task updated"})
			return
		}
	}
	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})

}

// DELETE /tasks/:id
// Description: Delete a specific task by ID.
// Response: { "message": "Task removed" } or { "message": "Task not found" }
func deleteTask(ctx *gin.Context) {
	id := ctx.Param("id")
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Task removed"})
			return
		}
	}
	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})

}

// POST /tasks
// Description: Create a new task.
// Request: { "id": "...", "title": "...", "description": "...", "status": "...", "duedate": "..." }
// Response: { "message": "Task created" } or { "error": "..." }
func addTask(ctx *gin.Context) {
	var newTask Task
	if err := ctx.ShouldBindJSON(&newTask); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tasks = append(tasks, newTask)
	ctx.IndentedJSON(http.StatusCreated, gin.H{"message": "Task created"})

}

func main() {
	//getting all the tasks
	router := gin.Default()
	router.GET("/tasks", getTasks)

	//getting a specific task based on ID
	router.GET("/tasks/:id", getTask)

	//Updating a specific task based on ID
	router.PUT("/tasks/:id", updateTask)

	//deleting a specific task by ID
	router.DELETE("/tasks/:id", deleteTask)

	// Addiing a new task
	router.POST("/tasks", addTask)

	router.Run()
}
