package controllers

import (
	customErrors "api/errors"
	"api/initializers"
	"api/models"
	utils "api/utils/auth"
	"fmt"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func UserCreate(c *gin.Context) {
	var body struct {
		Username string `json:"username"`
		Email    string
		Password string
	}

	c.ShouldBindJSON(&body)

	password, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err,
		})
	}
	user := models.User{Username: body.Username, Email: body.Email, Password: string(password)}
	err = initializers.DB.Where("username = ?", user.Username).Order("id").First(&models.User{}).Error
	fmt.Println(user)

	fmt.Println(err)
	if err == nil {
		c.JSON(409, gin.H{
			"error": "user already exists",
		})
		return
	}
	result := initializers.DB.Create(&user)
	if result.Error != nil {
		c.JSON(400, gin.H{
			"error": result.Error,
		})
		return
	}
	accessToken, err := utils.GenerateJwt(user.Username, user.Email)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "internal server error",
		})
	}
	c.JSON(201, gin.H{
		"user":         user,
		"access_token": accessToken,
	})
	return
}

func UserSignIn(c *gin.Context) {
	fmt.Println("here")
	var body struct {
		Username string
		Password string
	}

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(403, gin.H{
			"error": "bad Request",
		})
		return
	}
	var user models.User

	err = initializers.DB.Where("username = ?", body.Username).Order("id").First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(404, gin.H{
				"error": "record not found",
			})
			return
		} else {
			c.JSON(500, gin.H{
				"error": "internal server error",
			})
			return
		}
	}

	fmt.Println(user)
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(403, gin.H{
			"error": "wrong credentials",
		})
		return
	}

	token, err := utils.GenerateJwt(user.Username, user.Email)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "internal server error",
		})
		return
	}
	c.JSON(200, gin.H{
		"accessToken": token,
	})
	return
}

func UserUpdate(c *gin.Context) {

	id := c.Param("id")
	var body struct {
		gorm.Model
		models.User
	}

	err := c.ShouldBindJSON(&body)

	if err != nil {
		return
	}

	var found models.User
	initializers.DB.First(&found, id)

	result := initializers.DB.Model(&found).Updates(models.User{
		Username: body.Username,
		Email:    body.Email,
	})
	if result.Error != nil {

		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(404, gin.H{
				"error": "notfound",
			})
			return
		}
		c.JSON(500, gin.H{
			"error": result.Error,
		})
		return
	}
	c.JSON(200, gin.H{
		"user": found,
	})
}

func TryJwt(c *gin.Context) {
	token, err := utils.GenerateJwt("buttFace", "faceButt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(token)
	tkn, err := utils.ParseJwt(token)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(tkn)
}
func TryError(c *gin.Context) {
	// err := errors.TestError()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	customErrors.InternalServerError(c, "a", "b", "d")
}
