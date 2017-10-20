package novels

import (
	"fmt"
	"strings"

	"github.com/PrinceNorin/bakanovels/models"
	"github.com/PrinceNorin/bakanovels/utils/messages"
	"github.com/PrinceNorin/bakanovels/utils/uploader"
	"github.com/PrinceNorin/bakanovels/utils/validator"
	"github.com/PrinceNorin/bakanovels/utils/validator/novel"
	"github.com/gin-gonic/gin"
	uuid "github.com/nu7hatch/gouuid"
)

func CreateNovel(c *gin.Context) (*models.Novel, map[string][]string) {
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

	novel := models.Novel{
		UUID:             u4.String(),
		Title:            form.Title,
		Type:             strings.ToLower(form.Type),
		OriginalLanguage: strings.ToLower(form.OriginalLanguage),
		Description:      form.Description,
		Status:           "ongoing",
		PublishedAt:      form.PublishedAt,
	}

	if form.Image != "" {
		if err := uploader.NewImageUploader().Upload(form.Image, &novel); err != nil {
			msg.AddError("message", err.Error())
			return nil, msg.Errors
		}
	}

	if tl := form.TranslatedLanguage; tl.Valid {
		tl.String = strings.ToLower(tl.String)
		novel.TranslatedLanguage = tl
	}

	if err = models.DB.Create(&novel).Error; err != nil {
		msg.AddError("message", fmt.Sprintf("SQL: %s", err.Error()))
		return nil, msg.Errors
	}

	return &novel, nil
}
