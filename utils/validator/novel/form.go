package novelValidator

import (
	"time"

	"github.com/PrinceNorin/bakanovels/models"
)

type NovelCreateForm struct {
	Title       string            `json:"title" form:"title" binding:"required"`
	Type        string            `json:"type" form:"type" binding:"required"`
	Language    string            `json:"language" form:"language" binding:"required"`
	Description models.NullString `json:"desc" form:"desc"`
	Translate   models.NullString `json:"translate" form:"translate"`
	PublishedAt *time.Time        `json:"published_at" form:"published_at"`
}
