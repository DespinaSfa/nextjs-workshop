package server

import (
	"backend/db"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"os"
)

func InitServer() {
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println("Current directory:", currentDir)

	db.SetupDatabase()
	r := chi.NewRouter()

	setupMiddleware(r)
	setupRoutes(r)

	const port int = 3001
	fmt.Printf("Server running on http://localhost:%d\n", port)
	err = http.ListenAndServe(":3001", r)
	if err != nil {
		panic(err)
	}
}
