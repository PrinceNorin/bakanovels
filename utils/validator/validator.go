package validator

import (
	"fmt"
	"strings"

	"github.com/PrinceNorin/bakanovels/utils/locale"
	"github.com/PrinceNorin/bakanovels/utils/strings"
	"github.com/gin-gonic/gin"
	v "gopkg.in/go-playground/validator.v8"
)

func Validate(c *gin.Context, form interface{}) map[string][]string {
	err := c.Bind(form)
	if err != nil {
		T, _ := locale.GetTfunc(c)
		ve := make(map[string][]string)

		for _, fieldErr := range err.(v.ValidationErrors) {
			tag := fieldErr.Tag
			field := fieldErr.FieldNamespace
			key := fmt.Sprintf("%s.%s", field, tag)
			msg := T(key)
			if msg == key {
				msg = T("validation.required")
			}

			parts := strings.Split(field, ".")
			f := stringsutils.ToSnake(parts[len(parts)-1])
			ve[f] = append(ve[f], msg)
		}

		return ve
	}

	return nil
}
