package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Routes() http.Handler{
	router := chi.Router()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"htps://*", "https://*"},
		AllowedMethods: []string{"POST", "GET", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowedCredentials: true,
		MaxAge: 300,
	}))
}