package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/wsb777/check-price-biggeek/internal/bot"
	"github.com/wsb777/check-price-biggeek/internal/config"
	"github.com/wsb777/check-price-biggeek/internal/database"
	"github.com/wsb777/check-price-biggeek/internal/parser"
	"github.com/wsb777/check-price-biggeek/internal/services"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	log.Print("[INFO] Init config")
	cfg := config.NewConfig()

	log.Print("[INFO] Init parser")
	parser := parser.NewParser()

	log.Print("[INFO] Init database")
	db := database.DatabaseConnect()

	log.Print("[INFO] Init repository")
	repo := database.NewRepo(db)

	log.Print("[INFO] Init userservice")
	userService := services.NewUserService(repo)

	log.Print("[INFO] Init bot")
	b := bot.Init(cfg, &parser, userService)

	log.Print("[INFO] Starting bot")
	b.Run()
}
