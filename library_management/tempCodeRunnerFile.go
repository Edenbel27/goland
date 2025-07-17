package main

import (
	"library_management/controllers"
	"library_management/services"
)

func main() {
	lib := services.NewLibrary()
	controllers.DisplayMenu(lib)
}
