package bot

import (
	"context"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
	"github.com/wsb777/check-price-biggeek/internal/config"
	"github.com/wsb777/check-price-biggeek/internal/parser"
)

type Bot struct {
	Config *config.Config
	Parser *parser.Parser
}

func Init(cfg *config.Config, parser *parser.Parser) *Bot {
	return &Bot{
		Config: cfg,
		Parser: parser,
	}
}

func (b *Bot) Run() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	parserHandler := CreateParserHander(*b.Parser)
	startedHandler :=
	mainHandler :=
	opts := []bot.Option{
		bot.WithDefaultHandler(handler),
	}

	boot, err := bot.New(b.Config.BotToken, opts...)
	if err != nil {
		panic(err)
	}

	boot.Start(ctx)
}
