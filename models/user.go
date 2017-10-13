package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model

	ID       int
	Email    string
	Username string
	Password string
}