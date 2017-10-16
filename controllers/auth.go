package controllers

import (
	"strconv"
	"time"

	"bitbucket.org/bakauta/server/model"

	"github.com/PrinceNorin/bakanovels/config"
	"github.com/PrinceNorin/bakanovels/models"
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

func buildAuthMiddleware() *jwt.GinJWTMiddleware {
	return &jwt.GinJWTMiddleware{
		Realm:         "bakanovels",
		Key:           []byte(config.Get().SecretKey),
		Timeout:       30 * time.Second,
		MaxRefresh:    time.Hour,
		Authenticator: authenticator,
		Authorizator:  authorizator,
		Unauthorized:  unauthorizer,
		TokenLookup:   "header:Authorization",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	}
}

func authenticator(email string, password string, c *gin.Context) (string, bool) {
	var user model.User
	err := models.DB.Where(&models.User{
		Email:    email,
		Password: password,
	}).First(&user).Error

	if err != nil {
		return "0", false
	}
	return strconv.Itoa(int(user.ID)), true
}

func authorizator(userId string, c *gin.Context) bool {
	var user model.User
	id, err := strconv.Atoi(userId)
	if err != nil {
		return false
	}

	err = models.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		return false
	}
	return true
}

func unauthorizer(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}
