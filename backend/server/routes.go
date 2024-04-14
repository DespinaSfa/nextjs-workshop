package server

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func setupRoutes(r *chi.Mux) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Hello World :)"))
		if err != nil {
			panic(err)
		}
	})
}
