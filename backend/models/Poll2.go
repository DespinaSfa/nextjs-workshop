package models

import "gorm.io/gorm"

type Poll2 struct {
	gorm.Model
	User        User
	UserID      uint
	Title       string
	Description string
	text2       string
}
