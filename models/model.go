package models

import (
	"github.com/PrinceNorin/bakanovels/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
)

var DB *gorm.DB

func init() {
	c := config.Get()
	db, err := gorm.Open(c.DB_ADAPTER, c.DB_CON_STR)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Register model here
	db.AutoMigrate(new(User))

	DB = db
}
