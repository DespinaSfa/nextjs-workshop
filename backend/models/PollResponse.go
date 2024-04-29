package models

type PollResponse struct {
	Title       string
	Description string
	PollType    string
}

type PollPartyResponse struct {
	SongToBePlayed        string
	CurrentAlcoholLevel   int
	PreferredAlcoholLevel int
	FavoriteActivity      string
	WishSnack             string
}

type PollWeddingResponse struct {
	WeddingInvite       string
	KnowCoupleSince     int
	KnowCoupleFromWhere string
	WeddingHighlight    string
	CoupleWish          string
}
