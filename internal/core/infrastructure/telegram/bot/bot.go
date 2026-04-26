package core_infrastructure_telegram_bot

import (
	"context"
	"fmt"

	core_logger "github.com/Mirwinli/coffe_plus/internal/core/logger"
	order_ports_in "github.com/Mirwinli/coffe_plus/internal/feature/order/ports/in"
	shop_ports_in "github.com/Mirwinli/coffe_plus/internal/feature/shop/ports/in"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
)

type Bot struct {
	bot          *tgbotapi.BotAPI
	chatID       int64
	orderService order_ports_in.OrderService
	shopService  shop_ports_in.ShopService
	log          core_logger.Logger
}

func NewBot(
	tokenString string,
	chatID int64,
	orderService order_ports_in.OrderService,
	shopService shop_ports_in.ShopService,
	log core_logger.Logger,
) (Bot, error) {
	bot, err := tgbotapi.NewBotAPI(tokenString)
	if err != nil {
		return Bot{}, fmt.Errorf("create telegram bot: %w", err)
	}

	return Bot{
		bot:          bot,
		chatID:       chatID,
		orderService: orderService,
		shopService:  shopService,
		log:          log,
	}, nil
}

func (b *Bot) Start(ctx context.Context) {
	b.bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.bot.GetUpdatesChan(u)

	for {
		select {
		case <-ctx.Done():
			return
		case update := <-updates:
			if update.Message != nil && update.Message.IsCommand() {
				switch update.Message.Command() {
				case "start":
					b.log.Info("chat id", zap.Int("ChatID", int(update.Message.Chat.ID)))
					if err := b.sendMenu(); err != nil {
						b.log.Error("failed send menu", zap.Error(err))
					}
				}
			}
			if update.CallbackQuery != nil {
				if err := b.handleCallback(ctx, *update.CallbackQuery); err != nil {
					b.log.Error("failed hanle calback error", zap.Error(err))
				}
			}
		}
	}
}
