package db

import "gorm.io/gorm"

type Message struct {
	UserID   int `gorm:"not null unique"`
	UserName string
	Body     string
	gorm.Model
}
