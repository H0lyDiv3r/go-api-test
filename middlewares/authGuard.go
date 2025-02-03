package middlewares

import (
	"api/errors"
	utils "api/utils/auth"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthGuard() gin.HandlerFunc {
	return func(c *gin.Context) {
		tkn := c.GetHeader("Authorization")
		token := strings.Split(tkn, " ")
		if len(token) < 2 {
			errors.InternalServerError(c)
			c.Abort()
			return
		}
		user, err := utils.ParseJwt(token[1])
		if err != nil {
			errors.InternalServerError(c)
			c.JSON(403, gin.H{
				"error": "unauthorized",
			})
			c.Abort()
		}
		c.Set("user", user)
		c.Next()
	}
}
