package main

import (
	"api/initializers"
	"api/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDb()
}

func main() {
	initializers.DB.AutoMigrate(&models.Todo{})
	initializers.DB.AutoMigrate(&models.User{})
}
