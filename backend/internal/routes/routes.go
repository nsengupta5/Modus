package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/nsengupta5/Modus/internal/database"
)

func DefineRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/intro", introHandler)
	router.Post("/register", registerHandler)
	return router
}

func introHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to Modus!"))
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	println("Registering user")
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var data map[string]interface{}

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Println("Recieved data:", data)

	name := data["username"].(string)
	email := data["email"].(string)
	password := data["password"].(string)

	_, err := database.RegisterUser(name, email, password)
	if err != nil {
		log.Fatal(err)
	}
}
