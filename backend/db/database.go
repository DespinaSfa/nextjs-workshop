package db

import (
	"backend/config"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func SetupDatabase(config *config.Config) *sql.DB {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)

	// Connect to database
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic("error connecting to database: " + err.Error())
	}

	// Check if connection is successful
	err = db.Ping()
	if err != nil {
		err := db.Close()
		if err != nil {
			fmt.Println("error pinging database: " + err.Error())
		} // Close the database connection if ping fails
		panic("error pinging database: " + err.Error())
	}

	fmt.Printf("Connected to database %s running on %s:%s\n", config.DBName, config.DBHost, config.DBPort)

	return db
}
