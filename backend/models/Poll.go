package models

import "gorm.io/gorm"

type Poll struct {
	gorm.Model
	UserID      uint
	Title       string
	Description string
	PollType    string

	//This tells GORM that there is a one-to-many relationship between Poll and PollParty and between Poll and PollWedding
	PollParties   []PollParty    `gorm:"foreignKey:PollID"`
	PollWeddings  []PollWedding  `gorm:"foreignKey:PollID"`
	PollPlannings []PollPlanning `gorm:"foreignKey:PollID"`
}

type PollParty struct {
	gorm.Model
	PollID                uint
	Poll                  Poll   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	SongToBePlayed        string `gorm:"size:255;not null"`
	CurrentAlcoholLevel   int
	PreferredAlcoholLevel int
	FavoriteActivity      string
	WishSnack             string `gorm:"size:255"`
}

type PollWedding struct {
	gorm.Model
	PollID              uint
	Poll                Poll   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	WeddingInvite       string `gorm:"size:255;not null"`
	KnowCoupleSince     int
	KnowCoupleFromWhere string
	WeddingHighlight    string
	CoupleWish          string `gorm:"size:255"`
}

type PollPlanning struct {
	gorm.Model
	PollID          uint
	Poll            Poll   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	EssentialDrink  string `gorm:"size:255"`
	EssentialFood   string `gorm:"size:255"`
	MusicToBePlayed string
	Activities      string
	EventWish       string `gorm:"size:255"`
}
