package validator

import (
	"fmt"
	"strings"

	"github.com/PrinceNorin/bakanovels/utils/locale"
	"github.com/PrinceNorin/bakanovels/utils/messages"
	"github.com/PrinceNorin/bakanovels/utils/strings"
	"github.com/gin-gonic/gin"
	v "gopkg.in/go-playground/validator.v8"
)

func Validate(c *gin.Context, form interface{}) map[string][]string {
	err := c.BindJSON(form)
	if err != nil {
		var e interface{} = err

		switch e.(type) {
		case v.ValidationErrors:
			return buildValidationErrors(c, err)
		default:
			msg := messages.GetMessages(c)
			msg.AddError("error", err.Error())
			return msg.Errors
		}
	}

	return nil
}

func buildValidationErrors(c *gin.Context, err error) map[string][]string {
	var msg string

	T, _ := locale.GetTfunc(c)
	ve := make(map[string][]string)

	for _, fieldErr := range err.(v.ValidationErrors) {
		key := translationKey(fieldErr)

		if msg = T(key); msg == key {
			newKey := fmt.Sprintf("validation.%s", fieldErr.Tag)

			if msg = T(newKey); msg == newKey {
				msg = T("translation.missing", map[string]interface{}{"Key": key})
			}
		}

		f := stringsutils.ToSnake(fieldErr.Field)
		ve[f] = append(ve[f], msg)
	}

	return ve
}

func translationKey(fe *v.FieldError) string {
	name := fe.FieldNamespace
	field := stringsutils.ToSnake(fe.Field)
	model := strings.Split(stringsutils.ToSnake(name), "_")[0]
	return fmt.Sprintf("%s.%s.%s", model, field, fe.Tag)
}
