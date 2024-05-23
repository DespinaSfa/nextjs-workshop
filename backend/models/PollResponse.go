package models

import (
	"encoding/json"
)

type PollInfo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	PollType    string `json:"pollType"`
}

//PollResponses

type PollPartyResponse struct {
	SongToBePlayed        string `json:"songToBePlayed"`
	CurrentAlcoholLevel   int    `json:"currentAlcoholLevel"`
	PreferredAlcoholLevel int    `json:"preferredAlcoholLevel"`
	FavoriteActivity      string `json:"favoriteActivity"`
	WishSnack             string `json:"wishSnack"`
}

type PollWeddingResponse struct {
	WeddingInvite       string `json:"weddingInvite"`
	KnowCoupleSince     int    `json:"knowCoupleSince"`
	KnowCoupleFromWhere string `json:"knowCoupleFromWhere"`
	WeddingHighlight    string `json:"weddingHighlight"`
	CoupleWish          string `json:"coupleWish"`
}

type PollPlanningResponse struct {
	EssentialDrink  string `json:"essentialDrink"`
	EssentialFood   string `json:"essentialFood"`
	MusicToBePlayed string `json:"musicToBePlayed"`
	Activities      string `json:"activities"`
	EventWish       string `json:"eventWish"`
}

type GenericPollResponse struct {
	PollID   uint            `json:"poll_id"`
	PollType string          `json:"poll_type"`
	Data     json.RawMessage `json:"data"` // Use json.RawMessage for deferred unmarshalling
}
