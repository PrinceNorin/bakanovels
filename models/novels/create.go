package novels

import (
	"fmt"

	"github.com/PrinceNorin/bakanovels/models"
	"github.com/PrinceNorin/bakanovels/utils/messages"
	"github.com/PrinceNorin/bakanovels/utils/validator"
	"github.com/PrinceNorin/bakanovels/utils/validator/novel"
	"github.com/gin-gonic/gin"
	uuid "github.com/nu7hatch/gouuid"
)

func CreateNovel(c *gin.Context) (*models.NovelJSON, map[string][]string) {
	var form novelValidator.NovelCreateForm
	msg := messages.GetMessages(c)

	if err := validator.Validate(c, &form); err != nil {
		return nil, err
	}

	u4, err := uuid.NewV4()
	if err != nil {
		msg.AddErrorT("message", "internal.error")
		return nil, msg.Errors
	}

	err = models.DB.Create(&models.Novel{
		UUID:        u4.String(),
		Title:       form.Title,
		Type:        form.Type,
		Language:    form.Language,
		Description: form.Description,
		Translate:   form.Translate,
	}).Error

	if err != nil {
		msg.AddError("message", fmt.Sprintf("SQL: %s", err.Error()))
		return nil, msg.Errors
	}

	return &models.NovelJSON{
		UUID:        u4.String(),
		Title:       form.Title,
		Type:        form.Type,
		Language:    form.Language,
		Description: form.Description,
		Translate:   form.Translate,
	}, nil
}
