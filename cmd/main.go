package main

import (
	"api/initializers"
	"api/middlewares"
	"api/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDb()
}

func main() {
	r := gin.Default()
	r.Use(middlewares.GlobalError())
	routes.TodoRoutes(r)
	routes.UserRoutes(r)
	routes.TestRoute(r)
	r.Run()
}
