package errors

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

type CustomErorr struct {
	StatusCode int      `json:"statusCode"`
	Status     string   `json:"status"`
	Messages   []string `json:"messages"`
}

func (c *CustomErorr) Error() string {
	jsonData, err := json.Marshal(c)
	if err != nil {
		fmt.Println("failed to marshal error")
		return ""
	}
	return string(jsonData)

}

func InternalServerError(c *gin.Context, messages ...string) {
	error := CustomErorr{
		StatusCode: 500,
		Status:     "Internal Server Error",
		Messages:   messages,
	}

	c.Error(&error)
	return
}

func BadRequest(c *gin.Context, messages ...string) {
	error := CustomErorr{
		StatusCode: 400,
		Status:     "Bad Request",
		Messages:   messages,
	}

	c.Error(&error)
	return
}
