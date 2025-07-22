package main

import (
	"log"
	"task_manager_api_v2/models"
	"task_manager_api_v2/router"
)

func main() {
	err := models.ConnectDatabase()
	if err != nil{
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	router.SetupRouter()
}
