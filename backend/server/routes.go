package server

import (
	_ "backend/docs"
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

func setupRoutes(r *chi.Mux) {

	r.Get("/health", HealthHanlder)

	r.Mount("/swagger", httpSwagger.WrapHandler)

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
