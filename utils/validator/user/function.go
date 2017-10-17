package userValidator

import (
	"github.com/PrinceNorin/bakanovels/models"
)

func CheckEmail(email string) bool {
	var count int
	models.DB.Model(&models.User{}).Where("email = ?", email).Count(&count)
	return count != 0
}
