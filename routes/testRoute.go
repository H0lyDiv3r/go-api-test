package routes

import (
	"api/controllers"
	"api/initializers"

	"github.com/gin-gonic/gin"
)

func TestRoute(r *gin.Engine) {
	service := controllers.NewContentService(initializers.DB)
	controller := controllers.NewContentController(service)
	testRoute := r.Group("/api/:test")
	{
		testRoute.GET("", controller.GetAll)
	}
}
