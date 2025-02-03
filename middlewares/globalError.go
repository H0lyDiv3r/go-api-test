package middlewares

import (
	customErrors "api/errors"
	"errors"

	"github.com/gin-gonic/gin"
)

func GlobalError() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			var customErr *customErrors.CustomErorr
			err := c.Errors[0].Err
			if errors.As(err, &customErr) {
				c.JSON(customErr.StatusCode, customErr)
				c.Abort()
			}
		}
	}
}
