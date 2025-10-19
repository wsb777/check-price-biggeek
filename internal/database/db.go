package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func DatabaseConnect() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("[ERROR] DB: %v", err)
	}
	db.AutoMigrate(&User{}, &Link{})

	return db
}
