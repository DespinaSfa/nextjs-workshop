package server

import (
	"backend/config"
	"backend/db"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func InitServer() {
	dbConfig := config.LoadConfig()

	_, err := db.SetupDatabase(dbConfig)
	if err != nil {
		panic("error setting up database: " + err.Error())
	}

	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	setupMiddleware(r)
	setupRoutes(r)
	r.Post("/login", loginHandler)

	const port int = 3001

	fmt.Printf("\nServer running on http://localhost:%d", port)
	err = http.ListenAndServe(":3001", r)
	if err != nil {
		panic(err)
	}
}
