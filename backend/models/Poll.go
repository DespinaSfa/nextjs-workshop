package models

import "gorm.io/gorm"

type Poll struct {
	gorm.Model
	User        User
	UserID      uint
	Title       string
	Description string
}
