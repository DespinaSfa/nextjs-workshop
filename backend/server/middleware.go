package server

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"net/http"
	"strings"
)

func setupMiddleware(r *chi.Mux) {
	r.Use(middleware.Logger)    //log every request
	r.Use(middleware.Recoverer) //recover from panics without crashing the server return 500
	//probably add something like
	r.Use(corsMiddleware)
	//r.Use(AuthenticationMiddleware)
}

func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path == "/login" {
			next.ServeHTTP(w, r)
			return
		}
		// Get the token from the 'Authorization' header
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Authorization header is missing", http.StatusUnauthorized)
			return
		}

		// Remove the 'Bearer ' prefix
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		// Parse the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil {
			http.Error(w, "Error parsing token"+err.Error(), http.StatusUnauthorized)
			return
		}

		// Validate the token and get the claims
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Extract the user ID from the claims and add it to the request context
			userID := claims["user_id"].(float64)
			ctx := context.WithValue(r.Context(), "userID", userID)

			// Call the next handler with the new context
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
	})
}

func corsMiddleware(next http.Handler) http.Handler {
	corsOptions := cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}

	return cors.New(corsOptions).Handler(next)
}
