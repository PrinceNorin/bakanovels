package users

import (
	"fmt"

	"github.com/PrinceNorin/bakanovels/models"
	"github.com/PrinceNorin/bakanovels/utils/messages"
	"github.com/PrinceNorin/bakanovels/utils/validator"
	"github.com/PrinceNorin/bakanovels/utils/validator/user"
	"github.com/gin-gonic/gin"
	"github.com/nu7hatch/gouuid"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *gin.Context) (*models.UserJSON, map[string][]string) {
	var form userValidator.UserRegisterForm
	msg := messages.GetMessages(c)

	err := validator.Validate(c, &form)
	if err != nil {
		return nil, err
	}

	if userValidator.CheckEmail(form.Email) {
		msg.AddErrorT("email", "user.email.exists")
		return nil, msg.Errors
	}

	u4, er := uuid.NewV4()
	if er != nil {
		msg.AddErrorT("message", "internal.error")
		return nil, msg.Errors
	}

	password, er := bcrypt.GenerateFromPassword([]byte(form.Password), bcrypt.DefaultCost)
	if er != nil {
		msg.AddErrorT("message", "internal.error")
		return nil, msg.Errors
	}

	er = models.DB.Create(&models.User{
		UUID:     u4.String(),
		Email:    form.Email,
		Username: form.Username,
		Password: string(password),
	}).Error
	if er != nil {
		msg.AddError("message", fmt.Sprintf("SQL: %s", er.Error()))
		return nil, msg.Errors
	}

	return &models.UserJSON{
		UUID:     u4.String(),
		Email:    form.Email,
		Username: form.Username,
	}, nil
}
