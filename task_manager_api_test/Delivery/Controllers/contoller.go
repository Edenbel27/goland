package Controllers

import (
	"task_manager_api_test/Domain"
	"task_manager_api_test/UseCase"
	"net/http"
	"github.com/gin-gonic/gin"
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserController struct {
	UserUsecase UseCase.IUserUseCase
}

type TaskController struct{
	TaskUsecase UseCase.ITaskUseCase
}

func NewUserController(user UseCase.IUserUseCase) *UserController{
	return &UserController{
		UserUsecase : user,
	}
}

func NewTaskController(task UseCase.ITaskUseCase) *TaskController{
	return &TaskController{
		TaskUsecase : task,
	}
}

type UserDTO struct{
	Email string 	`json:"email"`
	Password string `json:"password"`
	Role string 	`json:"role"`
}

type TaskDTO struct{
	ID          primitive.ObjectID `json:"id,omitempty"`
	Title       string    			`json:"title"`
	Description string    			`json:"description"`
	Status      string   			`json:"status"`
	DueDate     time.Time 			`json:"due_date"`
}


func(uc *UserController) RegisterHandler(ctx *gin.Context){
	var user UserDTO
	err := ctx.ShouldBindJSON(&user)
	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error":"invalid request payload"})
		return 
	}

	if user.Email == "" || user.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error":"invalid request payload"})
		return 
	}
	uc.UserUsecase.Register(uc.ChangeToDomain(user))
	ctx.JSON(http.StatusOK, gin.H{
		"user":  user,
	})
}

func (uc *UserController) ChangeToDomain(userDTO UserDTO) *Domain.User{
	var user Domain.User
	user.Email = userDTO.Email
	user.Password = userDTO.Password
	user.Role = userDTO.Role
	return &user
} 

func (uc *UserController)LoginHandler(ctx *gin.Context){
	var user UserDTO

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid json format"})
		return
	}

	if user.Email == "" || user.Password == ""{
		ctx.JSON(http.StatusBadRequest, gin.H{"error":"invalid request payload"})
		return 
	}
	domUser := uc.ChangeToDomain(user)
	domUser, token , err := uc.UserUsecase.Login(domUser)
		if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	
	
	ctx.JSON(http.StatusOK, gin.H{"token": token})


}

func (tc *TaskController) ChangeToTask(taskDTO TaskDTO) *Domain.Task{
	var task Domain.Task
	task.Description = taskDTO.Description
	task.DueDate = taskDTO.DueDate
	task.Title = taskDTO.Title
	task.Status = taskDTO.Status
	return &task
} 


func (tc *TaskController) AddTaskController (ctx *gin.Context){
	var newTask TaskDTO
	if err := ctx.ShouldBindJSON(&newTask); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	task := tc.ChangeToTask(newTask)
	tc.TaskUsecase.AddTask(*task)
}

func (tc *TaskController) ViewTasksController(ctx *gin.Context) {
	tasks := tc.TaskUsecase.ViewTasks()
	if len(tasks) == 0 {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "No tasks found"})
		return
	}
	ctx.IndentedJSON(http.StatusOK, tasks)
}

func (tc *TaskController) ViewTaskByIDController(ctx *gin.Context) {
	id := ctx.Param("id")
	task, _ := tc.TaskUsecase.ViewTaskByID(id)
	ctx.IndentedJSON(http.StatusOK, task)
}

func (tc *TaskController) UpdateTaskController(ctx *gin.Context) {
	id := ctx.Param("id")
	var updatedTask TaskDTO
	if err := ctx.ShouldBindJSON(&updatedTask); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	domainUpdatedTask := tc.ChangeToTask(updatedTask)
	err := tc.TaskUsecase.UpdateTask(id,*domainUpdatedTask)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Task updated successfully"})
}
func (tc *TaskController) DeleteTaskController(ctx *gin.Context) {
	id := ctx.Param("id")
	err := tc.TaskUsecase.DeleteTask(id)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}