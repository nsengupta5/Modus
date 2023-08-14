package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/nsengupta5/FilmClub/internal/routes"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Mount("/", routes.DefineRoutes())
	http.ListenAndServe(":8080", r)
}
