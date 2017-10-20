package novelValidator

import (
	"time"

	"github.com/PrinceNorin/bakanovels/models"
)

type NovelCreateForm struct {
	Title              string            `json:"title" binding:"required"`
	Type               string            `json:"type" binding:"required"`
	OriginalLanguage   string            `json:"original_language" binding:"required"`
	Description        models.NullString `json:"desc"`
	TranslatedLanguage models.NullString `json:"translate_language"`
	PublishedAt        *time.Time        `json:"published_at"`
	Image              string            `json:"image"`
}
