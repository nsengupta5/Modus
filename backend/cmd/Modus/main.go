package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/nsengupta5/Modus/internal/database"
	"github.com/nsengupta5/Modus/internal/routes"
)

func main() {
	if err := database.InitDB(); err != nil {
		log.Fatal("Failed to connect to database", err)
	}
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Mount("/", routes.DefineRoutes())
	http.ListenAndServe(":8080", r)
}
