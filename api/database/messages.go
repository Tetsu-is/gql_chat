package database

import "gorm.io/gorm"

type Message struct {
	UserID   int `gorm:"not null unique"`
	UserName string
	Body     string
	gorm.Model
}

func MessageSeed(db *gorm.DB) error {
	var count int64
	db.Model(&Message{}).Count(&count)
	if count > 0 {
		return nil
	}
	messages := []Message{
		{
			UserID:   1,
			UserName: "Tetsu",
			Body:     "Hello, World!",
		},
		{
			UserID:   2,
			UserName: "JBoy-san",
			Body:     "Hello, World!",
		},
		{
			UserID:   3,
			UserName: "Icchy-san",
			Body:     "Hello, World!",
		},
	}
	for _, message := range messages {
		db.Create(&message)
	}
	return nil
}
