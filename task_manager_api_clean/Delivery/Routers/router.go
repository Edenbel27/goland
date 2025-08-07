package Routers

import (
	"task_manager_api_clean/Delivery/Controllers"
	"task_manager_api_clean/Infrustructure"

	"github.com/gin-gonic/gin"
)

func SetupRouter(uc *Controllers.UserController , tc *Controllers.TaskController) *gin.Engine{
	router := gin.Default()

	router.POST("/register", uc.RegisterHandler)
	router.POST("/login" , uc.LoginHandler)

	auth := router.Group("/api" , Infrustructure.JWTAuth())
	{
		auth.GET("/protected" , func (c *gin.Context){
			c.JSON(200 , gin.H{"message": "Protected route"})
		})
		admin := auth.Group ("/admin" , Infrustructure.AuthorizeRole("admin"))
		{
			admin.GET("/dashboard" , func (c *gin.Context){
				c.JSON(200 , gin.H{"message": "Admin dashboard"})
			})
		}
	}

	router.GET("/tasks", tc.ViewTasksController)
	router.GET("/tasks/:id", tc.ViewTaskByIDController)
	router.POST("/tasks", tc.AddTaskController)
	router.PUT("/tasks/:id", tc.UpdateTaskController)
	router.DELETE("/tasks/:id", tc.DeleteTaskController)
	
	return router
}