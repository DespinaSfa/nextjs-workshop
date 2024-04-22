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
	//TODO: Check if Database is already populated
	//populateDatabase(db)

	fmt.Printf("Connected to database %s running on %s:%s\n", config.DBName, config.DBHost, config.DBPort)

	return db, nil
}

// Call only when Database is empty
func populateDatabase(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Poll1{})
	db.AutoMigrate(&models.Poll2{})
	db.AutoMigrate(&models.Poll3{})

	addUser := models.User{Username: "Tom", Password: "12345", Token: "4321"}
	db.Create(&addUser)

	var getUser models.User
	db.First(&getUser, 1)
	poll1 := models.Poll1{User: getUser, Title: "Title1", Description: "Description", Text: "Text"}
	poll2 := models.Poll2{User: getUser, Title: "Title2", Description: "Description", Text: "Text"}
	poll3 := models.Poll3{User: getUser, Title: "Title3", Description: "Description", Text: "Text"}

	db.Create(&poll1)
	db.Create(&poll2)
	db.Create(&poll3)
}
