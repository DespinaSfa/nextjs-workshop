package models

type PollInfoResponse struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	PollType    string `json:"pollType"`
}

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
