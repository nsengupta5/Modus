package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

var db *sql.DB

func InitDB() error {
	cfg := mysql.Config{
		User:   viper.GetString("dbuser"),
		Passwd: viper.GetString("dbpass"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "modus",
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return err
	}

	if pingErr := db.Ping(); pingErr != nil {
		return pingErr
	}

	log.Println("Connected to database")
	return nil
}

func RegisterUser(name, email, password string) (int64, error) {
	result, err := db.Exec("INSERT INTO users (name, email, password) VALUES (?, ?, ?)", name, email, password)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func GetUserPassword(email string) (string, error) {
	var password string
	row := db.QueryRow("SELECT password FROM users WHERE email = ?", email)
	if err := row.Scan(&password); err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("Email not found; please try again")
		}
		return "", fmt.Errorf("Internal error: [%s]", err)
	}
	return password, nil
}
