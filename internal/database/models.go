package database

import "gorm.io/gorm"

type Link struct {
	gorm.Model
	Link   string
	Price  uint64
	UserID int64
}

type User struct {
	gorm.Model
	ID       uint  `gorm:"primaryKey"`
	UserID   int64 `gorm:"unique"`
	ChatID   int64 `gorm:"unique"`
	Username string
	Links    []Link `gorm:"foreignKey:UserID"`
}
