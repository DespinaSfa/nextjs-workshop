package db

import (
	"backend/models"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func CreatePoll(db *gorm.DB, addPoll *models.Poll) error {
	result := db.Create(addPoll)
	if result.Error != nil {
		return fmt.Errorf("error creating poll: %w", result.Error)
	}
	return nil
}

func CreateResult(db *gorm.DB, poll *models.Poll, addPollResult *models.PollParty) error {
	// Validate that the poll exists
	if err := db.First(&models.Poll{}, poll.ID).Error; err != nil {
		return fmt.Errorf("poll not found: %w", err)
	}

	// Set the PollID of the PollParty to the given poll's ID
	addPollResult.PollID = poll.ID

	// Insert the new PollParty result
	if err := db.Create(addPollResult).Error; err != nil {
		return fmt.Errorf("failed to create poll result: %w", err)
	}

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

func ReadPollByID(db *gorm.DB, pollID int) (*models.Poll, error) {
	var poll models.Poll
	result := db.Preload("PollParties").Preload("PollWeddings").First(&poll, pollID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("poll not found: %w", result.Error)
		}
		return nil, fmt.Errorf("error retrieving poll: %w", result.Error)
	}
	return &poll, nil
}

func DeletePollByID(db *gorm.DB, pollID int) error {
	// Delete the poll; associated PollParty and PollWedding records will be deleted automatically
	result := db.Delete(&models.Poll{}, pollID)
	if result.Error != nil {
		return fmt.Errorf("failed to delete poll: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no poll found with ID %d", pollID)
	}

	return nil
}

func populateDatabase(db *gorm.DB) {
	// Drop existing schema
	if err := db.Migrator().DropTable(&models.PollWedding{}, &models.PollParty{}, &models.Poll{}, &models.User{}); err != nil {
		fmt.Println("Failed to drop tables:", err)
		return
	}

	// Automatically migrate your schema
	migrate := []interface{}{&models.User{}, &models.Poll{}, &models.PollParty{}, &models.PollWedding{}}
	for _, model := range migrate {
		if err := db.AutoMigrate(model); err != nil {
			fmt.Printf("Failed to migrate %T: %v\n", model, err)
			return
		}
	}

	// Populate users
	users := []models.User{
		{Username: "CrazyCatLady", Password: "meowmix", Token: "catnip4life"},
		{Username: "TheRealElvis", Password: "thankyouverymuch", Token: "blueSuedeShoes"},
		{Username: "WannabeWizard", Password: "alohomora", Token: "muggleStruggles"},
	}

	for i := range users {
		if err := db.Create(&users[i]).Error; err != nil {
			fmt.Printf("Failed to create user %s: %v\n", users[i].Username, err)
			return
		}
	}

	// Create polls
	polls := []models.Poll{
		{UserID: users[0].ID, Title: "The Ultimate Cat Quiz", Description: "Decide which cat is the best: Garfield, Tom, or Sylvester?", PollType: "fun"},
		{UserID: users[1].ID, Title: "Elvis Song Poll", Description: "Which song rocks your blue suede shoes?", PollType: "music"},
		{UserID: users[2].ID, Title: "Wizardry Wonders", Description: "Which spell would you cast in a duel?", PollType: "fantasy"},
	}
	for i := range polls {
		if err := db.Create(&polls[i]).Error; err != nil {
			fmt.Printf("Failed to create poll %s: %v\n", polls[i].Title, err)
			return
		}
	}

	// Create related party and wedding details
	partyDetails := []models.PollParty{
		{PollID: polls[0].ID, SongToBePlayed: "Meow Mix Theme", CurrentAlcoholLevel: 0, PreferredAlcoholLevel: 0, FavoriteActivity: "Cat Napping", WishSnack: "Tuna"},
		{PollID: polls[1].ID, SongToBePlayed: "Hound Dog", CurrentAlcoholLevel: 5, PreferredAlcoholLevel: 10, FavoriteActivity: "Dancing", WishSnack: "Peanut Butter Banana Sandwich"},
	}
	for _, detail := range partyDetails {
		if err := db.Create(&detail).Error; err != nil {
			fmt.Printf("Failed to create poll party details for poll ID %d: %v\n", detail.PollID, err)
			return
		}
	}

	weddingDetails := []models.PollWedding{
		{PollID: polls[2].ID, WeddingInvite: "Fellow Wizards and Witches", KnowCoupleSince: 394, KnowCoupleFromWhere: "Hogwarts", WeddingHighlight: "Dueling Bridesmaids", CoupleWish: "May our wands never falter"},
	}
	for _, detail := range weddingDetails {
		if err := db.Create(&detail).Error; err != nil {
			fmt.Printf("Failed to create poll wedding details for poll ID %d: %v\n", detail.PollID, err)
			return
		}
	}

	fmt.Println("\nDatabase populated successfully ;)")

}
