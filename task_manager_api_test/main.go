package main

import (
	"task_manager_api_test/Delivery/Controllers"
	"task_manager_api_test/Delivery/Routers"
	"task_manager_api_test/Infrustructure"
	"task_manager_api_test/Repository"
	"task_manager_api_test/UseCase"
)

func main() {

	Repository.ConnectDatabase()
	userCol := Repository.GetCollection("users")
	taskCol := Repository.GetCollection("tasks")
	taskRepo := Repository.NewTaskRepo(taskCol)
	userRepo := Repository.NewUserRepo(userCol)

	passwordService := Infrustructure.NewPasswordService()
	jwtService := Infrustructure.NewJWTService()

	userUseCase := UseCase.NewUserUseCase(userRepo, passwordService, jwtService)
	taskUseCase := UseCase.NewTaskUseCase(taskRepo)

	userController := Controllers.NewUserController(userUseCase)
	taskController := Controllers.NewTaskController(taskUseCase)

	r := Routers.SetupRouter(userController, taskController)
	r.Run()
}
