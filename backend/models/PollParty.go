package models

import "gorm.io/gorm"

type PollParty struct {
	gorm.Model
	Poll                  Poll
	PollID                uint
	SongToBePlayed        string
	CurrentAlcoholLevel   int
	PreferredAlcoholLevel int
	FavoriteActivity      string
	WishSnack             string
}
