package controllers

import (
	"api/errors"
	"api/initializers"
	"api/models"
	utils "api/utils/auth"

	"github.com/gin-gonic/gin"
)

func TodosCreate(c *gin.Context) {
	user, _ := c.Get("user")
	var body struct {
		Content string
		Status  bool
	}
	c.Bind(&body)
	foundUser, _ := utils.GetUser(user.(utils.Claims).Username, c)

	todo := models.Todo{Content: body.Content, Status: body.Status, UserID: foundUser.ID}
	result := initializers.DB.Create(&todo)
	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"todo": todo,
	})
	return
}

func TodosIndex(c *gin.Context) {
	var todos []models.Todo
	user, _ := c.Get("user")
	if user == nil {
		errors.InternalServerError(c)
		return
	}
	foundUser, e := utils.GetUser(user.(utils.Claims).Username, c)

	if e != nil {
		errors.InternalServerError(c)
	}
	initializers.DB.Where("user_id=?", foundUser.ID).Find(&todos)

	c.JSON(200, gin.H{
		"todos": todos,
	})
	return
}
func TodosShow(c *gin.Context) {
	// Get id from URL param
	id := c.Param("id")

	// Get a sing todo
	var todo models.Todo
	initializers.DB.First(&todo, id)

	// Return todo in response
	c.JSON(200, gin.H{
		"todo": todo,
	})
}

func TodosUpdate(c *gin.Context) {
	// Get id from URL param
	id := c.Param("id")

	// get the data of req body
	var body struct {
		Content string
		Status  bool
	}
	c.Bind(&body)

	// Get a single todo that we want to update
	var todo models.Todo
	initializers.DB.First(&todo, id)

	// Update it
	initializers.DB.Model(&todo).Updates(models.Todo{
		Content: body.Content,
		Status:  body.Status,
	})

	// Return response
	c.JSON(200, gin.H{
		"todo": todo,
	})
}

func TodosDelete(c *gin.Context) {
	// Get id from URL param
	id := c.Param("id")

	// Delete the Todo
	initializers.DB.Delete(&models.Todo{}, id)

	// Return response
	c.JSON(200, gin.H{
		"message": "Todo removed Successfully",
	})
}
