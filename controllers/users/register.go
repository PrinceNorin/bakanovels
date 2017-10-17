package userController

import (
	"net/http"

	"github.com/PrinceNorin/bakanovels/models/users"
	"github.com/gin-gonic/gin"
)

func UserRegisterHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	user, err := users.CreateUser(c)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"user": user})
	} else {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"errors": err,
		})
	}
}
