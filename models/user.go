package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model

	UUID     string `gorm:"not null;unique"`
	Email    string `gorm:"not null;unique"`
	Username string `gorm:"not null"`
	Password string `gorm:"not null"`
	Role     string `sql:"DEFAULT:user"`
}

type UserJSON struct {
	UUID     string `json:"uuid"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type AdminJSON struct {
	UserJSON

	Role string `json:"role"`
}

func (u *User) IsAdmin() bool {
	return u.Role == "admin"
}
