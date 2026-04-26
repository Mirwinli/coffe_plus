package core_infrastructure_telegram_bot

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) sendMenu() error {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("📋 Замовлення що очікують", getCreatedOrdersCallbackData),
			tgbotapi.NewInlineKeyboardButtonData("📋 Замовлення що готуются", getCookingOrdersCallbackData),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🔴 Закрити магазин", closeShopCallbackData),
			tgbotapi.NewInlineKeyboardButtonData("🟢 Відкрити магазин", openShopCallbackData),
		),
	)

	msg := tgbotapi.NewMessage(b.chatID, "☕ Coffee Plus — панель кухаря")
	msg.ReplyMarkup = keyboard

	if _, err := b.bot.Send(msg); err != nil {
		return fmt.Errorf(
			"send message: %w", err,
		)
	}
	return nil
}
