package router
import (
	"github.com/gin-gonic/gin"
	"task_manager_api_v2/controllers"
)

func SetupRouter() {
	router := gin.Default()
	router.GET("/tasks", controllers.ViewTasksController)
	router.GET("/tasks/:id", controllers.ViewTaskByIDController)
	router.POST("/tasks", controllers.AddTaskController)
	router.PUT("/tasks/:id", controllers.UpdateTaskController)
	router.DELETE("/tasks/:id", controllers.DeleteTaskController)
	router.Run()
}