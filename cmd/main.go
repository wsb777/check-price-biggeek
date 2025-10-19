package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/wsb777/check-price-biggeek/internal/bot"
	"github.com/wsb777/check-price-biggeek/internal/config"
	"github.com/wsb777/check-price-biggeek/internal/parser"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	log.Print("[INFO] Init config")
	cfg := config.NewConfig()
	log.Print("[INFO] Init parser")
	parser := parser.NewParser()
	log.Print("[INFO] Init bot")
	b := bot.Init(cfg, &parser)
	log.Print("[INFO] Starting bot")
	b.Run()
}
