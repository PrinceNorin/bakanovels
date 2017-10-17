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
