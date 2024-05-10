package db

import (
	"backend/models"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CreatePoll(db *gorm.DB, addPoll *models.Poll) error {
	result := db.Create(addPoll)
	if result.Error != nil {
		return fmt.Errorf("error creating poll: %w", result.Error)
	}
	return nil
}

func CreatePollResponse(db *gorm.DB, jsonResponse []byte) error {

	//example usage:
	//partyJson := `{
	//    "poll_id": 1,
	//    "poll_type": "party",
	//    "data": {
	//        "songToBePlayed": "Dancing Queen",
	//        "currentAlcoholLevel": 5,
	//        "preferredAlcoholLevel": 7,
	//        "favoriteActivity": "Karaoke",
	//        "wishSnack": "Chips"
	//    }
	//}`
	//if err := db.AddPollResponse(dataBase, []byte(partyJson)); err != nil {
	//	fmt.Println("Error adding party poll response:", err)
	//} else {
	//	fmt.Println("Party poll response added successfully")
	//}

	var response models.GenericPollResponse
	if err := json.Unmarshal(jsonResponse, &response); err != nil {
		return err
	}

	switch response.PollType {
	case "party":
		var partyResponse models.PollPartyResponse
		if err := json.Unmarshal(response.Data, &partyResponse); err != nil {
			return err
		}
		// Convert PollPartyResponse to PollParty (your DB model)
		dbModel := models.PollParty{
			PollID:                response.PollID,
			SongToBePlayed:        partyResponse.SongToBePlayed,
			CurrentAlcoholLevel:   partyResponse.CurrentAlcoholLevel,
			PreferredAlcoholLevel: partyResponse.PreferredAlcoholLevel,
			FavoriteActivity:      partyResponse.FavoriteActivity,
			WishSnack:             partyResponse.WishSnack,
		}

		if dbModel.CurrentAlcoholLevel < 0 || dbModel.CurrentAlcoholLevel > 5 {
			return errors.New("current alcohol level must be between 0 and 5")
		}
		if dbModel.PreferredAlcoholLevel < 0 || dbModel.PreferredAlcoholLevel > 5 {
			return errors.New("preferred alcohol level must be between 0 and 5")
		}

		if dbModel.FavoriteActivity != models.Dancing &&
			dbModel.FavoriteActivity != models.Drinking &&
			dbModel.FavoriteActivity != models.Eating &&
			dbModel.FavoriteActivity != models.Singing &&
			dbModel.FavoriteActivity != models.Beerpong {
			return errors.New("favorite activity must be one of: dancing, drinking, eating, singing, beerpong")
		}

		return db.Create(&dbModel).Error

	case "wedding":
		var weddingResponse models.PollWeddingResponse
		if err := json.Unmarshal(response.Data, &weddingResponse); err != nil {
			return err
		}

		// Convert PollWeddingResponse to PollWedding (your DB model)
		dbModel := models.PollWedding{
			PollID:              response.PollID,
			WeddingInvite:       weddingResponse.WeddingInvite,
			KnowCoupleSince:     weddingResponse.KnowCoupleSince,
			KnowCoupleFromWhere: weddingResponse.KnowCoupleFromWhere,
			WeddingHighlight:    weddingResponse.WeddingHighlight,
			CoupleWish:          weddingResponse.CoupleWish,
		}

		// Check wedding invite
		if dbModel.WeddingInvite != "bride" && dbModel.WeddingInvite != "groom" && dbModel.WeddingInvite != "both" {
			return errors.New("wedding invite must be one of: bride, groom, both")
		}

		// Check acquaintance duration
		if dbModel.KnowCoupleSince < 1 || dbModel.KnowCoupleSince > 30 {
			return errors.New("know couple since must be between 1 and 30 years")
		}

		// Wedding highlight check
		validHighlights := map[string]bool{"wedding": true, "food": true, "dance": true, "program": true, "afterParty": true}
		if _, ok := validHighlights[dbModel.WeddingHighlight]; !ok {
			return errors.New("wedding highlight must be one of: wedding, food, dance, program, afterParty")
		}

		// Create record in the database
		return db.Create(&dbModel).Error

	default:
		return errors.New("unknown poll type")
	}
}

func GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	result := db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return &models.User{}, nil
		}
		return nil, fmt.Errorf("error querying database: %w", result.Error)
	}
	return &user, nil
}

func ReadUserPolls(userID int) ([]*models.PollInfo, error) {
	var user *models.User
	// Check if the user exists
	result := db.First(&user, userID)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to find user: %w", result.Error)
	}

	var pollResponse []*models.PollInfo
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
		{Username: "CrazyCatLady", Password: hashPassword("meowmix"), Token: "catnip4life"},
		{Username: "TheRealElvis", Password: hashPassword("thankyouverymuch"), Token: "blueSuedeShoes"},
		{Username: "WannabeWizard", Password: hashPassword("alohomora"), Token: "muggleStruggles"},
	}

	for i := range users {
		if err := db.Create(&users[i]).Error; err != nil {
			fmt.Printf("Failed to create user %s: %v\n", users[i].Username, err)
			return
		}
	}

	// Create polls
	polls := []models.Poll{
		{UserID: users[1].ID, Title: "Unsere Hochzeit", Description: "Hallo. Wir hoffen euch gefällt unsere Hochzeit. Für ein Spiel später füllt bitte diese kleine Umfrage aus. Vielen Dank! Euer Simon und eure Anna", PollType: "wedding"},
		{UserID: users[1].ID, Title: "Freds Fette Fete", Description: "Moin, moin! Diese Umfrage habe ich erstellt, damit ihr meine Party bewerten könnt. Die nächste wird dadurch noch geiler, versprochen!", PollType: "party"}}
	for i := range polls {
		if err := db.Create(&polls[i]).Error; err != nil {
			fmt.Printf("Failed to create poll %s: %v\n", polls[i].Title, err)
			return
		}
	}

	// Create related party and wedding details
	partyDetails := []models.PollParty{
		{PollID: polls[1].ID, SongToBePlayed: "tempo - cro", CurrentAlcoholLevel: 1, PreferredAlcoholLevel: 3, FavoriteActivity: "dance", WishSnack: "Pizza"},
		{PollID: polls[1].ID, SongToBePlayed: "Friesenjung - Ski Aggu", CurrentAlcoholLevel: 5, PreferredAlcoholLevel: 1, FavoriteActivity: "karaoke", WishSnack: "Brownies"},
	}
	for _, detail := range partyDetails {
		if err := db.Create(&detail).Error; err != nil {
			fmt.Printf("Failed to create poll party details for poll ID %d: %v\n", detail.PollID, err)
			return
		}
	}

	weddingDetails := []models.PollWedding{
		{PollID: polls[0].ID, WeddingInvite: "groom", KnowCoupleSince: 20, KnowCoupleFromWhere: "In einem Café", WeddingHighlight: "food", CoupleWish: "Super Flitterwochen "},
		{PollID: polls[0].ID, WeddingInvite: "bride", KnowCoupleSince: 10, KnowCoupleFromWhere: "Universität", WeddingHighlight: "afterParty", CoupleWish: "Glück und Gesundheit"},
	}
	for _, detail := range weddingDetails {
		if err := db.Create(&detail).Error; err != nil {
			fmt.Printf("Failed to create poll wedding details for poll ID %d: %v\n", detail.PollID, err)
			return
		}
	}

	fmt.Println("\nDatabase populated successfully ;)")

}

func hashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic("Failed to hash password: " + err.Error())
	}
	return string(hashedPassword)
}
