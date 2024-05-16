package models

import "gorm.io/gorm"

type Poll struct {
	gorm.Model
	UserID      uint
	Title       string
	Description string
	PollType    string

	//This tells GORM that there is a one-to-many relationship between Poll and PollParty and between Poll and PollWedding
	PollParties  []PollParty   `gorm:"foreignKey:PollID"`
	PollWeddings []PollWedding `gorm:"foreignKey:PollID"`
}

type PollParty struct {
	gorm.Model
	PollID                uint
	Poll                  Poll   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	SongToBePlayed        string `gorm:"size:255;not null"`
	CurrentAlcoholLevel   int
	PreferredAlcoholLevel int
	FavoriteActivity      string `gorm:"size:255"`
	WishSnack             string `gorm:"size:255"`
}

type PollWedding struct {
	gorm.Model
	PollID              uint
	Poll                Poll   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	WeddingInvite       string `gorm:"size:255;not null"`
	KnowCoupleSince     int
	KnowCoupleFromWhere string `gorm:"size:255"`
	WeddingHighlight    string `gorm:"size:255"`
	CoupleWish          string `gorm:"size:255"`
}

const Dancing = "dancing"
const Drinking = "drinking"
const Eating = "eating"
const Singing = "singing"
const Beerpong = "beerpong"
