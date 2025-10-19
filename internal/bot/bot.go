package bot

import (
	"context"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
	"github.com/wsb777/check-price-biggeek/internal/config"
	"github.com/wsb777/check-price-biggeek/internal/parser"
	"github.com/wsb777/check-price-biggeek/internal/services"
)

type Bot struct {
	Config      *config.Config
	Parser      *parser.Parser
	UserService services.UserService
}

func Init(cfg *config.Config, parser *parser.Parser, userService services.UserService) *Bot {
	return &Bot{
		Config:      cfg,
		Parser:      parser,
		UserService: userService,
	}
}

func (b *Bot) Run() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	botHandlers := NewHandlers(b.UserService, b.Parser)

	opts := []bot.Option{
		bot.WithDefaultHandler(botHandlers.HandleAll),
	}

	boot, err := bot.New(b.Config.BotToken, opts...)
	if err != nil {
		panic(err)
	}

	boot.Start(ctx)
}
