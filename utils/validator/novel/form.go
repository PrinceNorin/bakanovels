package novelValidator

import "time"

type NovelCreateForm struct {
	Title       string     `json:"title" binding:"required"`
	Type        string     `json:"type" binding:"required"`
	Language    string     `json:"language" binding:"required"`
	Description string     `json:"desc"`
	Translate   string     `json:"translate"`
	PublishedAt *time.Time `json:"published_at"`
}
