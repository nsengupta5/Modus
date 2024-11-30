package main

import (
	"log"
	"net/http"

	"github.com/nsengupta5/Modus/internal/database"
	"github.com/nsengupta5/Modus/internal/routes"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../../internal")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal("Config file not found", err)
		} else {
			log.Fatal("Config file found but error reading it", err)
		}
	}
	if err := database.InitDB(); err != nil {
		log.Fatal("Failed to connect to database", err)
	}
	router := routes.DefineRoutes()
	logRouter := routes.LoggingMiddleware(router)

	http.ListenAndServe(":8080", logRouter)
}
