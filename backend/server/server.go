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

	db.SetupDatabase(dbConfig)
	r := chi.NewRouter()

	setupMiddleware(r)
	setupRoutes(r)

	const port int = 3001

	fmt.Printf("Server running on http://localhost:%d\n", port)
	err := http.ListenAndServe(":3001", r)
	if err != nil {
		panic(err)
	}
}
