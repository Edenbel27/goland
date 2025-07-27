package main

import (
	"log"
	"task_manager_api_v3/models"
	"task_manager_api_v3/router"
)

func main() {
	err := models.ConnectDatabase()
	if err != nil{
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	r := router.SetupRouter()
	r.Run(":8080")
	
}
