package novels

import (
	"database/sql"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/PrinceNorin/bakanovels/models"
	"github.com/PrinceNorin/bakanovels/utils/messages"
	"github.com/PrinceNorin/bakanovels/utils/validator"
	"github.com/PrinceNorin/bakanovels/utils/validator/novel"
	"github.com/gin-gonic/gin"
	uuid "github.com/nu7hatch/gouuid"
	"github.com/olahol/go-imageupload"
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

	var image string
	file, err := c.FormFile("image")
	if err == nil {
		img, err := imageupload.Process(c.Request, "image")
		if err != nil {
			msg.AddError("message", err.Error())
			return nil, msg.Errors
		}

		ext := filepath.Ext(file.Filename)
		path := fmt.Sprintf("uploads/novels/%s%s", u4.String(), ext)
		err = img.Save(path)
		if err != nil {
			msg.AddError("message", err.Error())
			return nil, msg.Errors
		}

		image = path
	}

	novel := models.Novel{
		UUID:        u4.String(),
		Title:       form.Title,
		Type:        strings.ToLower(form.Type),
		Language:    strings.ToLower(form.Language),
		Description: form.Description,
		Translate:   form.Translate,
		Status:      "ongoing",
		PublishedAt: form.PublishedAt,
	}

	if image != "" {
		novel.Image = models.NullString{
			NullString: sql.NullString{
				Valid:  true,
				String: image,
			},
		}
	}

	if err = models.DB.Create(&novel).Error; err != nil {
		msg.AddError("message", fmt.Sprintf("SQL: %s", err.Error()))
		return nil, msg.Errors
	}

	return novel.ToNovelJSON(), nil
}
