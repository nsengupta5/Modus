package routes

import (
	"net/http"

	"github.com/go-chi/chi"
)

func DefineRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/intro", introHandler)
	router.Get("/api/hello", helloHandler)
	return router
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Nikhil!"))
}

func introHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to FilmClub!"))
}
