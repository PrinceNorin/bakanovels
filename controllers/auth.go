package controllers

import (
	"github.com/PrinceNorin/bakanovels/config"
	"github.com/PrinceNorin/bakanovels/models"
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"time"
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
	var user models.User
	models.DB.Where(&models.User{
		email:    email,
		password: password,
	}).First(&user)

	if user != nil {
		return user.ID, true
	}
	return user.ID, false
}

func authorizator(userId string, c *gin.Context) bool {
	if userId == "norin@example.com" {
		return true
	}
	return false
}

func unauthorizer(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}
