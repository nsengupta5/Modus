package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/nsengupta5/Modus/internal/database"
	"golang.org/x/crypto/bcrypt"
)

func DefineRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/intro", introHandler)
	router.Post("/register", registerHandler)
	router.Post("/login", loginHandler)
	return router
}

func introHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to Modus!"))
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var data map[string]interface{}

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	name := data["username"].(string)
	email := data["email"].(string)
	password := data["password"].(string)

	password, err := hashAndSalt(password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = database.RegisterUser(name, email, password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var data map[string]interface{}

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	email := data["email"].(string)
	password := data["password"].(string)

	dbPass, err := database.GetUserPassword(email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	err = comparePasswords(dbPass, password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
}

func hashAndSalt(password string) (string, error) {
	bytePassword := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.MinCost)
	if err != nil {
		return "", fmt.Errorf("Error while hashing password: %v", err)
	}
	return string(hash), nil
}

func comparePasswords(hashedPassword string, password string) error {
	byteHash := []byte(hashedPassword)
	bytePassword := []byte(password)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePassword)
	if err != nil {
		return fmt.Errorf("Invalid password: %v", err)
	}
	return nil
}
