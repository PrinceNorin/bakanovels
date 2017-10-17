package messages

import (
	"errors"

	"github.com/PrinceNorin/bakanovels/utils/locale"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/i18n"
)

type Messages struct {
	Errors map[string][]string
	c      *gin.Context
	T      i18n.TranslateFunc
}

func GetMessages(c *gin.Context) *Messages {
	T, _ := locale.GetTfuncWithMissing(c)
	return &Messages{
		c: c,
		T: T,
	}
}

func (msg *Messages) AddError(key string, text string) error {
	if msg.Errors == nil {
		msg.Errors = make(map[string][]string)
	}

	msg.Errors[key] = append(msg.Errors[key], text)
	return errors.New(text)
}

func (msg *Messages) AddErrorT(key string, tranID string, args ...interface{}) error {
	return msg.AddError(key, msg.T(tranID, args...))
}
