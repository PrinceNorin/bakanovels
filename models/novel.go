package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Novel struct {
	gorm.Model

	UUID        string     `gorm:"not null;unique" json:"uuid"`
	Title       string     `gorm:"not null" json:"title"`
	Type        string     `gorm:"not null";type:varchar(25) json:"type"`
	Language    string     `gorm:"not null;type:varchar(25)" json:"language"`
	Description NullString `gorm:"type:text" json:"description"`
	Translate   NullString `gorm:"type:varchar(25)" json:"translate"`
	Status      string     `gorm:"type:varchar(25)" json:"status"`
	Image       NullString `json:"image"`
	PublishedAt *time.Time `json:"published_at"`
	Approved    bool       `json:"approved"`
}

type NovelJSON struct {
	Novel

	ID       uint `json:"-"`
	Approved bool `json:"-"`
}

func (n *Novel) ToNovelJSON() *NovelJSON {
	return &NovelJSON{
		Novel:    *n,
		ID:       n.ID,
		Approved: n.Approved,
	}
}
