package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type PollsRequestBody struct {
	UserID string `json:"userID"`
}

func setupRoutes(r *chi.Mux) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Hello World :)"))
		if err != nil {
			panic(err)
		}
	})

	// GET: /polls -> Get an overview of all polls for Dashboard
	r.Get("/polls", func(w http.ResponseWriter, r *http.Request) {

		// Read req body
		body, err := ioutil.ReadAll(r.Body)
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
		fmt.Print("Extracted UserID:", userID)

		// Respond with dummy JSON response
		response := map[string]string{"message": "Successfully processed request", "userID": userID}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			fmt.Print("Error marshaling JSON response:", err)
			http.Error(w, "Failed to marshal JSON response", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		w.Write(jsonResponse)
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
			UserID      string `json:"userID"`
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

		// Respond w dummy JSON response
		response := map[string]string{"message": "Request saved successfully", "status": "OK"}
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

	// GET: /polls/{pollId} -> Get a specific poll's results
	r.Get("/polls/{pollId}", func(w http.ResponseWriter, r *http.Request) {

		pollID := chi.URLParam(r, "pollId")

		// Respond w dummy JSON response
		response := map[string]string{"message": fmt.Sprintf("Getting results of poll with ID %s", pollID), "status": "OK"}
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

	// POST: /polls/{pollId} -> Post results of a poll
	r.Post("/polls/{pollId}", func(w http.ResponseWriter, r *http.Request) {
		pollID := chi.URLParam(r, "pollId")

		// Respond w dummy JSON response
		response := map[string]string{"message": fmt.Sprintf("Posting results of poll with ID %s", pollID), "status": "OK"}
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

	// DELETE: /polls/{pollId} -> Delete a specific poll
	r.Delete("/polls/{pollId}", func(w http.ResponseWriter, r *http.Request) {
		pollID := chi.URLParam(r, "pollId")

		// Respond w dummy JSON response
		response := map[string]string{"message": fmt.Sprintf("Deleting poll with ID %s", pollID), "status": "OK"}
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
}
