package models

import (
	"fmt"
	"time"

	"database/sql"
)

type Novel struct {
	ID                 uint       `gorm:"primary_key" json:"-"`
	UUID               string     `gorm:"not null;unique" json:"uuid"`
	Title              string     `gorm:"not null" json:"title"`
	Type               string     `gorm:"not null";type:varchar(25) json:"type"`
	OriginalLanguage   string     `gorm:"not null;type:varchar(25)" json:"original_language"`
	Description        NullString `gorm:"type:text" json:"description"`
	TranslatedLanguage NullString `gorm:"type:varchar(25)" json:"translated_language"`
	Status             string     `gorm:"type:varchar(25)" json:"st atus"`
	Image              NullString `json:"image"`
	PublishedAt        *time.Time `json:"published_at"`
	Approved           bool       `json:"-"`
	CreatedAt          time.Time  `json:"created_at"`
	UpdatedAt          time.Time  `json:"updated_at"`
	DeletedAt          *time.Time `json:"-"`
}

func (n *Novel) SavePath() string {
	return fmt.Sprintf("novels/%s", n.UUID)
}

func (n *Novel) SaveColumnName() string {
	return "image"
}

func (n *Novel) SetValue(val string) {
	n.Image = NullString{
		NullString: sql.NullString{
			Valid:  true,
			String: val,
		},
	}
}
