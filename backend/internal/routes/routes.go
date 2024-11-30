package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/nsengupta5/Modus/internal/database"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	jwt.RegisteredClaims
}

func DefineRoutes() http.Handler {
	router := http.NewServeMux()
	router.HandleFunc("/intro", introHandler)
	router.HandleFunc("/register", registerHandler)
	router.HandleFunc("/login", loginHandler)
	return router
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}

func introHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		w.Write([]byte("Welcome to Modus!"))
		return
	}

	tokenString := cookie.Value
	user, err := validateJWT(tokenString)
	if err != nil {
		w.Write([]byte("Welcome to Modus!"))
	} else {
		w.Write([]byte(fmt.Sprintf("Welcome back, %s!", user.Name)))
	}
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

	user, err := database.GetUser(email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	dbPass := user.Password

	err = comparePasswords(dbPass, password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	tokenString, expirationTime, err := generateJWT(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
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

func generateJWT(user *database.User) (string, time.Time, error) {
	jwtKey := viper.GetString("jwtKey")
	timeLimit := time.Duration(viper.GetInt("timeLimit"))
	expirationTime := time.Now().Add(timeLimit * time.Hour)

	claims := &Claims{
		Email: user.Email,
		Name:  user.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		return "", time.Time{}, fmt.Errorf("Error while generating JWT: %v", err)
	}

	return tokenString, expirationTime, nil
}

func validateJWT(tokenString string) (*database.User, error) {
	jwtKey := viper.GetString("jwtKey")
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, fmt.Errorf("Invalid JWT signature")
		}
		return nil, fmt.Errorf("Invalid JWT")
	}
	if !token.Valid {
		return nil, fmt.Errorf("Invalid JWT")
	}

	user := &database.User{
		Email: claims.Email,
		Name:  claims.Name,
	}

	return user, nil
}
