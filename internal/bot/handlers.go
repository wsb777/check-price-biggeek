package bot

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/wsb777/check-price-biggeek/internal/parser"
	"github.com/wsb777/check-price-biggeek/internal/services"
)

type Handlers struct {
	userService services.UserService
	parser      *parser.Parser
}

func NewHandlers(
	userService services.UserService,
	parser *parser.Parser,
) *Handlers {
	return &Handlers{
		userService: userService,
		parser:      parser,
	}
}

func (h *Handlers) HandleAll(ctx context.Context, b *bot.Bot, update *models.Update) {
	handlers := []func(ctx context.Context, b *bot.Bot, update *models.Update) bool{
		h.HandleStart,
	}

	for _, handler := range handlers {
		if handler(ctx, b, update) {
			return
		}
	}
}

func (h *Handlers) HandleStart(ctx context.Context, b *bot.Bot, update *models.Update) bool {
	if update.Message == nil || !strings.HasPrefix(update.Message.Text, "/start") {
		return false
	}

	err := h.userService.RegisterUser(ctx, update.Message.Chat.ID, update.Message.Chat.ID, update.Message.From.Username)
	if err != nil {
		return false
	}

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   fmt.Sprintf("Привет. Это бот, который поможет тебе следить за обновлением цен товаров, которые ты сюда добавишь.\n\nПришли мне ссылку на товар и следи за ценой!"),
	})
	return true
}

func CreateParserHander(parser parser.Parser) bot.HandlerFunc {
	return func(ctx context.Context, b *bot.Bot, update *models.Update) {
		if update.Message == nil {
			return
		}

		productName, productPrice, err := parser.GetInfoByLink(update.Message.Text)

		if err != nil {
			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: update.Message.Chat.ID,
				Text:   fmt.Sprintf("Ошибка получения данных: %s", err),
			})
			return
		}

		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   fmt.Sprintf("Название продукта: %s\nЦена продукта:%v", productName, productPrice),
		})
	}
}
