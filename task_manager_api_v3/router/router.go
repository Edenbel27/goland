package router
import (
	"github.com/gin-gonic/gin"
	"task_manager_api_v3/controllers"
	"task_manager_api_v3/middleware"
)

func SetupRouter() *gin.Engine{
	router := gin.Default()

	router.POST("/register", controllers.Register)
	router.POST("/login" , controllers.Login)

	auth := router.Group("/api" , middleware.JWTAuth())
	{
		auth.GET("/protected" , func (c *gin.Context){
			c.JSON(200 , gin.H{"message": "Protected route"})
		})
		admin := auth.Group ("/admin" , middleware.AuthorizeRole("admin"))
		{
			admin.GET("/dashboard" , func (c *gin.Context){
				c.JSON(200 , gin.H{"message": "Admin dashboard"})
			})
		}
	}




	router.GET("/tasks", controllers.ViewTasksController)
	router.GET("/tasks/:id", controllers.ViewTaskByIDController)
	router.POST("/tasks", controllers.AddTaskController)
	router.PUT("/tasks/:id", controllers.UpdateTaskController)
	router.DELETE("/tasks/:id", controllers.DeleteTaskController)
	
	return router
}