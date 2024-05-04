package models

import "gorm.io/gorm"

type Poll struct {
	gorm.Model
	UserID      uint `gorm:"index;not null"` // Ensures faster queries on UserID and that it must be provided
	User        User
	Title       string `gorm:"size:255;not null"`
	Description string `gorm:"type:text"` // Allows longer text for descriptions
	PollType    string `gorm:"size:100;index;not null"`
}

type PollParty struct {
	gorm.Model
	PollID                uint   `gorm:"uniqueIndex"`
	Poll                  Poll   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	SongToBePlayed        string `gorm:"size:255;not null"`
	CurrentAlcoholLevel   int
	PreferredAlcoholLevel int
	FavoriteActivity      string `gorm:"size:255"`
	WishSnack             string `gorm:"size:255"`
}

type PollWedding struct {
	gorm.Model
	PollID              uint   `gorm:"uniqueIndex"`
	Poll                Poll   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	WeddingInvite       string `gorm:"size:255;not null"`
	KnowCoupleSince     int
	KnowCoupleFromWhere string `gorm:"size:255"`
	WeddingHighlight    string `gorm:"size:255"`
	CoupleWish          string `gorm:"size:255"`
}
