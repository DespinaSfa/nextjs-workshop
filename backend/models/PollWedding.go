package models

import "gorm.io/gorm"

type PollWedding struct {
	gorm.Model
	Poll                Poll
	PollID              uint
	WeddingInvite       string
	KnowCoupleSince     int
	KnowCoupleFromWhere string
	WeddingHighlight    string
	CoupleWish          string
}
