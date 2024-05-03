package database

import "gorm.io/gorm"

type User struct {
	UserName string `gorm:"not null"`
	Email    string `gorm:"not null unique"`
	gorm.Model
}

func UserSeed(db *gorm.DB) error {
	var count int64
	db.Model(&User{}).Count(&count)
	if count > 0 {
		return nil
	}
	users := []User{
		{
			UserName: "Tetsu",
			Email:    "tetsu@tetsu.com",
		},
		{
			UserName: "JBoy-san",
			Email:    "jboy@jboy.com",
		},
		{
			UserName: "Icchy-san",
			Email:    "icchy@icchy.com",
		},
	}
	for _, user := range users {
		db.Create(&user)
	}
	return nil
}
