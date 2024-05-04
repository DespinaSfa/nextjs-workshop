package server

import (
	"backend/config"
	"backend/db"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func InitServer() {
	dbConfig := config.LoadConfig()

	_, err := db.SetupDatabase(dbConfig)
	if err != nil {
		panic("error setting up database: " + err.Error())
	}
	r := chi.NewRouter()

	setupMiddleware(r)
	setupRoutes(r)

	const port int = 3001

	fmt.Printf("\nServer running on http://localhost:%d", port)
	err = http.ListenAndServe(":3001", r)
	if err != nil {
		panic(err)
	}
}
