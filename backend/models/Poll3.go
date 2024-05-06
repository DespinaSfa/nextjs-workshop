package models

import "gorm.io/gorm"

type Poll3 struct {
	gorm.Model
	Poll   Poll
	PollID uint
	Text   string
}
