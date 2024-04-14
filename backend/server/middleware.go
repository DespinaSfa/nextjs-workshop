package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func setupMiddleware(r *chi.Mux) {
	r.Use(middleware.Logger)    //log every request
	r.Use(middleware.Recoverer) //recover from panics without crashing the server return 500
	//probably add something like
	//auth middleware
	//CORSMiddleware
}
