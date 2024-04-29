package db

import (
	"backend/config"
	"fmt"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sync"
)

var (
	db     *gorm.DB
	dbOnce sync.Once
)

func SetupDatabase(config *config.Config) (*gorm.DB, error) {
	if db == nil {
		dbOnce.Do(func() {
			connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
				config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)

			// Connect to database
			var err error
			db, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
			if err != nil {
				panic(fmt.Errorf("error connecting to database: %w", err))
			}

			// Test the connection
			sqlDB, err := db.DB()
			if err != nil {
				panic(fmt.Errorf("error getting underlying database connection: %w", err))
			}

			err = sqlDB.Ping()
			if err != nil {
				// Close the database connection if ping fails
				if closeErr := sqlDB.Close(); closeErr != nil {
					fmt.Println("error closing database connection:", closeErr)
				}
				panic(fmt.Errorf("error pinging database: %w", err))
			}

			// Populate Database
			// TODO: Call only if database is empty
			//populateDatabase(db)

			fmt.Printf("Connected to database %s running on %s:%s\n", config.DBName, config.DBHost, config.DBPort)
		})

		return db, nil
	} else {
		return nil, fmt.Errorf("Database already createddatabase")
	}
}
