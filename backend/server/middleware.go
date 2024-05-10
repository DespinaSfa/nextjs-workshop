package server

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"strings"
)

func setupMiddleware(r *chi.Mux) {
	r.Use(middleware.Logger)    //log every request
	r.Use(middleware.Recoverer) //recover from panics without crashing the server return 500
	//probably add something like
	//auth middleware
	r.Use(AuthenticationMiddleware)

	//CORSMiddleware
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
			http.Error(w, "Error parsing token", http.StatusUnauthorized)
			return
		}

		// Validate the token and get the claims
		if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Call the next handler
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
	})
}
