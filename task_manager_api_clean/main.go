package main

import (
	"task_manager_api_clean/Delivery/Controllers"
	"task_manager_api_clean/Infrustructure"
	"task_manager_api_clean/Repository"
	"task_manager_api_clean/Delivery/Routers"
	"task_manager_api_clean/UseCase"
)

func main() {
	
	Repository.ConnectDatabase() 
	taskRepo := Repository.NewTaskRepo()
	userRepo := Repository.NewUserRepo()
	passwordService := Infrustructure.NewPasswordService()
	jwtService := Infrustructure.NewJWTService() 

 	
	userUseCase := UseCase.NewUserUseCase(userRepo , passwordService , jwtService)
	taskUseCase := UseCase.NewTaskUseCase(taskRepo)


	userController := Controllers.NewUserController(userUseCase)
	taskController := Controllers.NewTaskController(taskUseCase)

    
	r := Routers.SetupRouter(userController , taskController)
	r.Run()
}

