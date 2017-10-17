package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Novel struct {
	gorm.Model

	UUID        string `gorm:"not null;unique"`
	Title       string `gorm:"not null"`
	Type        string `gorm:"not null";type:varchar(25)`
	Language    string `gorm:"not null;type:varchar(25)"`
	Description string `gorm:"type:text"`
	Translate   string `gorm:"type:varchar(25)"`
	Status      string `gorm:"type:varchar(25)"`
	Image       string
	PublishedAt *time.Time
	Approved    bool
}

type NovelJSON struct {
	UUID        string     `json:"id"`
	Title       string     `json:"title"`
	Type        string     `json:"type"`
	Language    string     `json:"language"`
	Description string     `json:"desc"`
	Translate   string     `json:"translate"`
	Status      string     `json:"status"`
	Image       string     `json:"image"`
	PublishedAt *time.Time `json:"published_at"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}
