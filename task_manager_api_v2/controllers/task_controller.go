package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"task_manager_api_v2/models"
	"task_manager_api_v2/data"
)


func AddTaskController(ctx *gin.Context) {
	var newTask models.Task
	if err := ctx.ShouldBindJSON(&newTask); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	data.AddTask(newTask)
	ctx.IndentedJSON(http.StatusCreated, gin.H{"message": "Task created successfully"})
}

func ViewTasksController(ctx *gin.Context) {
	tasks := data.ViewTasks()
	if len(tasks) == 0 {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "No tasks found"})
		return
	}
	ctx.IndentedJSON(http.StatusOK, tasks)
}

func ViewTaskByIDController(ctx *gin.Context) {
	id := ctx.Param("id")
	task := data.ViewTaskByID(id)
		if task.ID.IsZero() {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	ctx.IndentedJSON(http.StatusOK, task)
}

func UpdateTaskController(ctx *gin.Context) {
	id := ctx.Param("id")
	var updatedTask models.Task
	if err := ctx.ShouldBindJSON(&updatedTask); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := data.UpdateTask(id, updatedTask)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Task updated successfully"})
}
func DeleteTaskController(ctx *gin.Context) {
	id := ctx.Param("id")
	err := data.DeleteTask(id)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}