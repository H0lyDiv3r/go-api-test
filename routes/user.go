package routes

import (
	"api/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {

	userGroup := r.Group("/user")
	{
		userGroup.GET("/err", controllers.TryError)
		userGroup.POST("/signin", controllers.UserSignIn)
		userGroup.POST("", controllers.UserCreate)
		userGroup.PUT("/:id", controllers.UserUpdate)
	}
}
