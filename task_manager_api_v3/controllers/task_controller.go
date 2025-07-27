package controllers

import (
	"fmt"
	"net/http"
	"os"
	"task_manager_api_v3/data"
	"task_manager_api_v3/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
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

// Register user
func Register(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil || (input.Email == "" || input.Password == "" || input.Role == "") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user type"})
		return
	}
	_, err := data.FindOneEmail(input)

	if err == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": "User Already Exists"})
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Error on password hashing"})
	}
	input.Password = string(hashedPassword)
	err = data.InsertOneUser(input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "User registration failed"})
	}

	c.IndentedJSON(200, gin.H{"message": "user registered successfully"})
}

func Login(c *gin.Context) {
	var input models.User
	var dbuser models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid json format"})
		return
	}

	dbuser, err := data.FindOneEmail(input)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			// User with that email doesn't exist
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		} else {
			// Some DB or decoding error
			fmt.Println("Database error:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}
	}

	if bcrypt.CompareHashAndPassword([]byte(dbuser.Password), []byte(input.Password)) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": dbuser.Email,
		"role":  dbuser.Role,
	})

	jwtToken, err := token.SignedString([]byte(os.Getenv("SECRETKEY")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token generation failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user logged in successfully", "token": jwtToken})
}
