package server

import (
	"backend/db"
	_ "backend/docs"
	"backend/models"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"io"
	"net/http"
	"strconv"
)

type Server struct {
	DB *gorm.DB
}

func (s *Server) HealthHandler(w http.ResponseWriter, _ *http.Request) {
	_, err := w.Write([]byte("OK"))
	if err != nil {
		fmt.Println("Error writing response:", err)
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}
}

func (s *Server) HomeHandler(w http.ResponseWriter, _ *http.Request) {
	_, err := w.Write([]byte("Hello World :)"))
	if err != nil {
		fmt.Println("Error writing response:", err)
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}
}

func (s *Server) GetPollsHandler(w http.ResponseWriter, _ *http.Request) {
	// Read user polls
	polls, err := db.ReadUserPolls(2) // TODO User ID kommt von Maiks Auth :)
	if err != nil {
		fmt.Println("Error reading user polls:", err)
		http.Error(w, "Failed to read user polls", http.StatusInternalServerError)
		return
	}

	pollsJSON, err := json.Marshal(polls)
	if err != nil {
		fmt.Println("Error marshaling JSON response:", err)
		http.Error(w, "Failed to marshal JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(pollsJSON)
	if err != nil {
		fmt.Println("Error writing response:", err)
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}
}

func (s *Server) PostPollsHandler(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Print("Error reading request body:", err)
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	fmt.Print("Received POST request to /polls")
	fmt.Print("Request Body:", string(body))

	type PollPostBody struct {
		UserID      uint   `json:"userID"`
		TemplateNr  int    `json:"templateNr"`
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	var requestBody PollPostBody
	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		fmt.Print("Error parsing request body:", err)
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	userID := requestBody.UserID

	newPoll := &models.Poll{
		UserID:      userID,
		Title:       requestBody.Title,
		Description: requestBody.Description,
	}

	if err := db.CreatePoll(s.DB, newPoll); err != nil {
		fmt.Print("Error creating poll:", err)
		http.Error(w, "Failed to create poll", http.StatusInternalServerError)
		return
	}

	response := map[string]string{"message": "Poll created successfully", "status": "OK"}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		fmt.Print("Error marshaling JSON response:", err)
		http.Error(w, "Failed to marshal JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	// Write response
	if _, err := w.Write(jsonResponse); err != nil {
		fmt.Print("Error writing response:", err)
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s *Server) GetPollByIDHandler(w http.ResponseWriter, r *http.Request) {
	pollIDStr := chi.URLParam(r, "pollId")

	pollID, err := strconv.Atoi(pollIDStr)

	if err != nil {
		http.Error(w, "Invalid poll ID", http.StatusBadRequest)
		return
	}

	poll, err := db.ReadPollByID(s.DB, pollID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "Poll not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to retrieve poll", http.StatusInternalServerError) // <- err
		return
	}

	pollJSON, err := json.Marshal(poll)
	if err != nil {
		http.Error(w, "Failed to marshal poll to JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(pollJSON)
	if err != nil {
		fmt.Println("Error writing response:", err)
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}
}

func (s *Server) PostPollByIDHandler(w http.ResponseWriter, r *http.Request) {
	pollIDStr := chi.URLParam(r, "pollId")

	_, err := strconv.Atoi(pollIDStr)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid poll ID", http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading request body:", err)
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	var pollResponseData models.GenericPollResponse
	if err := json.Unmarshal(body, &pollResponseData); err != nil {
		fmt.Println("Error parsing request body:", err)
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	pollResponseJSON, err := json.Marshal(pollResponseData)
	if err != nil {
		fmt.Println("Error marshaling poll response data:", err)
		http.Error(w, "Failed to marshal poll response data", http.StatusInternalServerError)
		return
	}

	if err := db.CreatePollResponse(s.DB, pollResponseJSON); err != nil {
		fmt.Println("Error creating poll response:", err)
		http.Error(w, "Failed to create poll response", http.StatusInternalServerError)
		return
	}

	response := map[string]string{"message": "Poll response created successfully", "status": "OK"}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Error marshaling JSON response:", err)
		http.Error(w, "Failed to marshal JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonResponse)
	if err != nil {
		fmt.Println("Error writing response:", err)
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}
}

func (s *Server) DeletePollByIDHandler(w http.ResponseWriter, r *http.Request) {
	pollIDStr := chi.URLParam(r, "pollId")
	pollID, err := strconv.Atoi(pollIDStr)
	if err != nil {
		http.Error(w, "Invalid poll ID", http.StatusBadRequest)
		return
	}

	err = db.DeletePollByID(s.DB, pollID)
	if err != nil {
		http.Error(w, "Failed to delete poll", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) RefreshToken(w http.ResponseWriter, r *http.Request) {
	// Parse the refresh token from the request body
	var requestBody struct {
		RefreshToken string `json:"refreshToken"`
	}
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	userID := r.Context().Value("userID").(float64)
	// Verify the refresh token and get a new access token
	_, err = RefreshToken(requestBody.RefreshToken, userID)
	if err != nil {
		http.Error(w, "Failed to verify refresh token: "+err.Error(), http.StatusUnauthorized)
		return
	}

	newToken, err := CreateToken(userID)
	if err != nil {
		http.Error(w, "Failed to create token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	newRefreshToken, err := CreateRefreshToken(userID)
	if err != nil {
		http.Error(w, "Failed to create refresh token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the new access token
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": newToken, "refreshToken": newRefreshToken})
}

func (s *Server) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Look up the user in the database
	user, err := db.GetUserByUsername(credentials.Username)
	if err != nil {
		http.Error(w, "Failed to get user", http.StatusInternalServerError)
		return
	}

	// Check the password
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(credentials.Password))
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	token, err := CreateToken(float64(user.ID))
	if err != nil {
		http.Error(w, "Failed to create token", http.StatusInternalServerError)
		return
	}

	refreshToken, err := CreateRefreshToken(float64(user.ID))
	if err != nil {
		http.Error(w, "Failed to create refresh token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token, "refreshToken": refreshToken})
}

func setupRoutes(r *chi.Mux, dbInstance *gorm.DB) {
	server := &Server{DB: dbInstance}
	r.Mount("/swagger", httpSwagger.WrapHandler)
	r.Get("/", server.HomeHandler)
	r.Get("/health", server.HealthHandler)

	r.Post("/login", server.LoginHandler)
	r.Post("/refresh-token", server.RefreshToken)

	r.Get("/polls", server.GetPollsHandler)
	r.Post("/polls", server.PostPollsHandler)
	r.Get("/polls/{pollId}", server.GetPollByIDHandler)
	r.Post("/polls/{pollId}", server.PostPollByIDHandler)
	r.Delete("/polls/{pollId}", server.DeletePollByIDHandler)
}
