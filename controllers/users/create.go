package userController

import (
	"net/http"

	"github.com/PrinceNorin/bakanovels/models/users"
	"github.com/gin-gonic/gin"
)

func UserRegisterHandler(c *gin.Context) {
	user, err := users.CreateUser(c)
	if err == nil {
		c.JSON(http.StatusOK, user)
	} else {
		c.JSON(http.StatusUnprocessableEntity, err)
	}
}
