package db

import (
	"backend/config"
	"backend/models"
	"fmt"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabase(config *config.Config) (*gorm.DB, error) {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)

	// Connect to database
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	// Test the connection
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("error getting underlying database connection: %w", err)
	}

	err = sqlDB.Ping()
	if err != nil {
		// Close the database connection if ping fails
		if closeErr := sqlDB.Close(); closeErr != nil {
			fmt.Println("error closing database connection:", closeErr)
		}
		return nil, fmt.Errorf("error pinging database: %w", err)
	}

	//Populate Database
	//TODO: Call only if database is empty
	populateDatabase(db)

	fmt.Printf("Connected to database %s running on %s:%s\n", config.DBName, config.DBHost, config.DBPort)

	return db, nil
}

func populateDatabase(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.PollWedding{})
	db.AutoMigrate(&models.PollParty{})

	addUser := models.User{Username: "Tom", Password: "12345", Token: "4321"}
	db.Create(&addUser)

	var getUser models.User
	db.First(&getUser, 1)
	pollWedding := models.Poll{User: getUser, Title: "Unsere Hochzeit", Description: "Hallo. Wir hoffen euch gefällt unsere Hochzeit. Für ein Spiel später füllt bitte diese kleine Umfrage aus. Vielen Dank! Euer Simon und eure Anna"}
	pollParty := models.Poll{User: getUser, Title: "Freds Fette Fete", Description: "Moin, moin! Diese Umfrage habe ich erstellt, damit ihr meine Party bewerten könnt. Die nächste wird dadurch noch geiler, versprochen!"}
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
