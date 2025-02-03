package controllers

import (
	"api/models"
	"errors"
	"fmt"
	"reflect"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ContentService struct {
	db *gorm.DB
}

func NewContentService(db *gorm.DB) *ContentService {
	return &ContentService{db: db}
}

type ContentController struct {
	Service *ContentService
}

func NewContentController(service *ContentService) *ContentController {
	return &ContentController{Service: service}
}

func (c *ContentService) FindAll(contentType string) (interface{}, error) {
	model, ok := getModel(contentType)

	modelType := reflect.TypeOf(model).Elem()
	sliceType := reflect.SliceOf(modelType)
	target := reflect.New(sliceType).Interface()
	if !ok {
		return nil, errors.New("this is not ok man")
	}
	err := c.db.Find(target).Error
	if err != nil {
		fmt.Println("we have a problem bruh", err)
	}
	return target, nil
}

func (c *ContentController) GetAll(ctx *gin.Context) {

	content := ctx.Param("test")
	fmt.Println(ctx.Param("test"))
	result, err := c.Service.FindAll(content)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": "error is an error",
		})
	}
	ctx.JSON(200, gin.H{
		"data": result,
	})
}

func getModel(contentType string) (interface{}, bool) {
	switch contentType {
	case "todo":
		return &models.Todo{}, true
	case "user":
		return &models.User{}, true
	default:
		return nil, false
	}
}
