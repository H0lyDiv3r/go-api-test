package routes

import (
	"api/controllers"
	"api/middlewares"

	"github.com/gin-gonic/gin"
)

func TodoRoutes(r *gin.Engine) {
	todoGroup := r.Group("/todos")
	todoGroup.Use(middlewares.AuthGuard())
	todoGroup.Use(middlewares.LoggerMiddleWare())
	{
		todoGroup.POST("", controllers.TodosCreate)
		todoGroup.GET("", controllers.TodosIndex)
		todoGroup.GET("/:id", controllers.TodosShow)
		todoGroup.PUT("/:id", controllers.TodosUpdate)
		todoGroup.DELETE("/:id", controllers.TodosDelete)
	}
}
