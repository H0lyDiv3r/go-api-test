package utils

import (
	"api/initializers"
	"api/models"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type Claims struct {
	Username string
	Email    string
	jwt.RegisteredClaims
}

func GenerateJwt(username, email string) (string, error) {
	claims := &Claims{
		Username:         username,
		Email:            email,
		RegisteredClaims: jwt.RegisteredClaims{},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseJwt(tknStr string) (Claims, error) {
	claims := &Claims{}

	_, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (any, error) {
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		return Claims{}, err
	}
	return *claims, nil
}

func GetUser(username string, c *gin.Context) (models.User, error) {
	var foundUser models.User
	err := initializers.DB.Where("username = ?", username).First(&foundUser).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(404, gin.H{
				"error": "user not found",
			})
			return models.User{}, err
		} else {
			c.JSON(404, gin.H{
				"error": "user not found",
			})
			return models.User{}, err
		}
	}
	return foundUser, nil
}
