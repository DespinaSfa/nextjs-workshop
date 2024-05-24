package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Poll struct {
	ID            string         `gorm:"type:uuid;primaryKey;"`
	UserID        uint           `gorm:"not null"`
	Title         string         `gorm:"size:255;not null"`
	Description   string         `gorm:"size:1024"`
	PollType      string         `gorm:"size:50"`
	PollParties   []PollParty    `gorm:"foreignKey:PollID"`
	PollWeddings  []PollWedding  `gorm:"foreignKey:PollID"`
	PollPlannings []PollPlanning `gorm:"foreignKey:PollID"`
}

func (poll *Poll) BeforeCreate(tx *gorm.DB) (err error) {
	poll.ID = uuid.New().String()
	return
}

type PollParty struct {
	gorm.Model
	PollID                string
	Poll                  Poll   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	SongToBePlayed        string `gorm:"size:255;not null"`
	CurrentAlcoholLevel   int
	PreferredAlcoholLevel int
	FavoriteActivity      string
	WishSnack             string `gorm:"size:255"`
}

type PollWedding struct {
	gorm.Model
	PollID              string
	Poll                Poll   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	WeddingInvite       string `gorm:"size:255;not null"`
	KnowCoupleSince     int
	KnowCoupleFromWhere string
	WeddingHighlight    string
	CoupleWish          string `gorm:"size:255"`
}

type PollPlanning struct {
	gorm.Model
	PollID          string
	Poll            Poll   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	EssentialDrink  string `gorm:"size:255"`
	EssentialFood   string `gorm:"size:255"`
	MusicToBePlayed string
	Activities      string
	EventWish       string `gorm:"size:255"`
}
