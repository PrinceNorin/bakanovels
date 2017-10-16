package userController

import (
	"github.com/PrinceNorin/bakanovels/controllers/router"
)

func init() {
	router.Get().POST("/register", UserRegisterHandler)
}
