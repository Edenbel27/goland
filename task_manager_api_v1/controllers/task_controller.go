package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"task_manager_api/models"
	"task_manager_api/data"
)

// func DisplayMenu(ctx *gin.Context){
// i := true
// for i{
// ctx.String(http.StatusOK , `{"message": "Welcome to the Task Manager API"}
// "1. Add Task"
// "2. View Tasks"
// "3. View Task by ID"
// "4. Update Task"
// "5. Delete Task"
// "6. Exit"
// Please select an option (1-6):
// `)
// var choice int
// ctx.BindJSON(&choice)
// switch choice {
// 	case 1:
// 		addTask(ctx)
// 	case 2:
// 		viewTasks(ctx)
// 	case 3:
// 		viewTaskByID(ctx)
// 	case 4:
// 		updateTask(ctx)
// 	case 5:
// 		deleteTask(ctx)
// 	case 6:
// 		ctx.String(http.StatusOK, `{"message": "Exiting the Task Manager API"}`)
// 		i = false	
// }

// }
// }

func AddTaskController(ctx *gin.Context) {
	var newTask models.Task
	if err := ctx.ShouldBind(&newTask); err != nil {
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
	task, err := data.ViewTaskByID(id)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusOK, task)
}

func UpdateTaskController(ctx *gin.Context) {
	id := ctx.Param("id")
	var updatedTask models.Task
	if err := ctx.ShouldBind(&updatedTask); err != nil {
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