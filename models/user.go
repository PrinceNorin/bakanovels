package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model

	Email    string
	Username string
	Password string
}

type UserJSON struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}
