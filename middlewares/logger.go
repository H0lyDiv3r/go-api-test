package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("logginf before")
		c.Next()
		fmt.Println("logging after")
	}
}
