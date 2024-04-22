package models

import "gorm.io/gorm"

type Poll1 struct {
	gorm.Model
	User        User
	UserID      uint
	Title       string
	Description string
	text1       string
}
