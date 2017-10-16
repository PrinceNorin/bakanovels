package users

import (
	"github.com/PrinceNorin/bakanovels/models"
	"github.com/PrinceNorin/bakanovels/utils/messages"
	"github.com/PrinceNorin/bakanovels/utils/validator"
	"github.com/PrinceNorin/bakanovels/utils/validator/user"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *gin.Context) (*models.UserJSON, map[string][]string) {
	var form userValidator.UserRegisterForm
	msg := messages.GetMessages(c)

	err := validator.Validate(c, &form)
	if err != nil {
		return nil, err
	}

	password, er := bcrypt.GenerateFromPassword([]byte(form.Password), bcrypt.DefaultCost)
	if er != nil {
		msg.AddError("password", "can't generate password")
		return nil, msg.Errors
	}

	er = models.DB.Create(&models.User{
		Email:    form.Email,
		Username: form.Username,
		Password: string(password),
	}).Error
	if er != nil {
		msg.AddError("user", er.Error())
		return nil, msg.Errors
	}

	return &models.UserJSON{
		Email:    form.Email,
		Username: form.Username,
	}, nil
}
