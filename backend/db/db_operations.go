package db

import (
	"backend/models"
	"fmt"
	"gorm.io/gorm"
)

func CreatePoll(addPoll *models.Poll) error {
	poll := db.Create(addPoll)
	if poll.Error != nil {
		return fmt.Errorf("error creating poll: %w", poll.Error)
	}
	return nil
}

func CreateResult(poll *models.Poll, addPollResult *models.PollParty) error {
	// insert a new Poll result
	return nil
}

func ReadUserPolls(userID int) ([]*models.PollResponse, error) {
	var user *models.User
	// Check if the user exists
	result := db.First(&user, userID)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to find user: %w", result.Error)
	}

	var pollResponse []*models.PollResponse
	// Select only the "title" and "description" columns from the database
	result = db.Model(&models.Poll{}).Select("title, description, poll_type").Where("user_id = ?", userID).Scan(&pollResponse)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to get user's polls: %w", result.Error)
	}

	return pollResponse, nil
}

func ReadPollByID(pollID int) error {
	// return one poll with all results
	return nil
}

func DeletePollByID(pollID int) error {
	// delete a poll by its id
	return nil
}

func populateDatabase(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.PollWedding{})
	db.AutoMigrate(&models.PollParty{})

	addUser := models.User{Username: "Tom", Password: "12345", Token: "4321"}
	db.Create(&addUser)

	var getUser models.User
	db.First(&getUser, 1)
	pollWedding := models.Poll{User: getUser, Title: "Unsere Hochzeit", Description: "Hallo. Wir hoffen euch gefällt unsere Hochzeit. Für ein Spiel später füllt bitte diese kleine Umfrage aus. Vielen Dank! Euer Simon und eure Anna", PollType: "wedding"}
	pollParty := models.Poll{User: getUser, Title: "Freds Fette Fete", Description: "Moin, moin! Diese Umfrage habe ich erstellt, damit ihr meine Party bewerten könnt. Die nächste wird dadurch noch geiler, versprochen!", PollType: "party"}
	db.Create(&pollWedding)
	db.Create(&pollParty)

	var getPollWedding models.Poll
	db.First(&getPollWedding, 1)
	var getPollParty models.Poll
	db.First(&getPollParty, 2)
	pollWeddingResults1 := models.PollWedding{Poll: getPollWedding, WeddingInvite: "bride", KnowCoupleSince: 10, KnowCoupleFromWhere: "Universität", WeddingHighlight: "afterParty", CoupleWish: "Glück und Gesundheit"}
	pollPartyResults1 := models.PollParty{Poll: getPollParty, SongToBePlayed: "tempo - cro", CurrentAlcoholLevel: 1, PreferredAlcoholLevel: 3, FavoriteActivity: "dance", WishSnack: "Pizza"}
	db.Create(&pollWeddingResults1)
	db.Create(&pollPartyResults1)
	pollWeddingResults2 := models.PollWedding{Poll: getPollWedding, WeddingInvite: "groom", KnowCoupleSince: 20, KnowCoupleFromWhere: "In einem Café", WeddingHighlight: "food", CoupleWish: "Super Flitterwochen "}
	pollPartyResults2 := models.PollParty{Poll: getPollParty, SongToBePlayed: "Friesenjung - Ski Aggu", CurrentAlcoholLevel: 5, PreferredAlcoholLevel: 1, FavoriteActivity: "karaoke", WishSnack: "Brownies"}
	db.Create(&pollWeddingResults2)
	db.Create(&pollPartyResults2)
}
