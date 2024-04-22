package db

import (
	"backend/config"
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

	fmt.Printf("Connected to database %s running on %s:%s\n", config.DBName, config.DBHost, config.DBPort)

	return db, nil
}
