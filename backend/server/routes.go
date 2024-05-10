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
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"strconv"
)

type PollsRequestBody struct {
	UserID string `json:"userID"`
}

// HealthHanlder @Summary Health check endpoint
// @Description This is an example of a GET endpoint
// @Success 200 {string} string "OK"
// @Router /health [get]
func HealthHanlder(w http.ResponseWriter, _ *http.Request) {
	_, err := w.Write([]byte("OK"))
	if err != nil {
		panic(err)
	}
}

func setupRoutes(r *chi.Mux, dbInstance *gorm.DB) {

	r.Get("/health", HealthHanlder)

	r.Mount("/swagger", httpSwagger.WrapHandler)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Hello World :)"))
		if err != nil {
			panic(err)
		}
	})

	// GET: /polls -> Get an overview of all polls for Dashboard
	r.Get("/polls", func(w http.ResponseWriter, r *http.Request) {

		// Read req body
		/*body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Print("Error reading request body:", err)
			http.Error(w, "Failed to read request body", http.StatusInternalServerError)
			return
		}

		// Log req body
		fmt.Print("Request Body:", string(body))

		// Parse req body
		var requestBody PollsRequestBody
		err = json.Unmarshal(body, &requestBody)
		if err != nil {
			fmt.Print("Error parsing request body:", err)
			http.Error(w, "Failed to parse request body", http.StatusBadRequest)
			return
		}

		// Extract userID parsed req body
		userID := requestBody.UserID
		fmt.Print("Extracted UserID:", userID) */

		// Read user polls
		polls, err := db.ReadUserPolls(2) // TODO woher kommt die userID dann?
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
		w.Write(pollsJSON)
	})

	// POST: /polls -> Create a poll
	r.Post("/polls", func(w http.ResponseWriter, r *http.Request) {
		// Read the request body
		body, err := ioutil.ReadAll(r.Body)
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

		if err := db.CreatePoll(dbInstance, newPoll); err != nil {
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
	})

	// TODO fix
	// GET: /polls/{pollId} -> Get a specific poll's results
	r.Get("/polls/{pollId}", func(w http.ResponseWriter, r *http.Request) {

		pollIDStr := chi.URLParam(r, "pollId")
		fmt.Println("Poll ID:", pollIDStr) // Debugging line

		pollID, err := strconv.Atoi(pollIDStr)
		fmt.Print("pollIDStr!!!!!!!!>")
		fmt.Println(pollIDStr)

		if err != nil {
			http.Error(w, "Invalid poll ID", http.StatusBadRequest)
			return
		}

		poll, err := db.ReadPollByID(dbInstance, pollID)
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
		w.Write(pollJSON)
	})

	// POST: /polls/{pollId} -> Post results of a poll
	r.Post("/polls/{pollId}", func(w http.ResponseWriter, r *http.Request) {

		pollIDStr := chi.URLParam(r, "pollId")

		_, err := strconv.Atoi(pollIDStr)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Invalid poll ID", http.StatusBadRequest)
			return
		}

		body, err := ioutil.ReadAll(r.Body)
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

		if err := db.CreatePollResponse(dbInstance, pollResponseJSON); err != nil {
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
		w.Write(jsonResponse)
	})

	// DELETE: /polls/{pollId} -> Delete a specific poll
	r.Delete("/polls/{pollId}", func(w http.ResponseWriter, r *http.Request) {

		pollIDStr := chi.URLParam(r, "pollId")
		pollID, err := strconv.Atoi(pollIDStr)
		if err != nil {
			http.Error(w, "Invalid poll ID", http.StatusBadRequest)
			return
		}

		err = db.DeletePollByID(dbInstance, pollID)
		if err != nil {
			http.Error(w, "Failed to delete poll", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	})
}
