package core_infrastructure_telegram_bot

import (
	"context"
	"fmt"
	"strings"

	"github.com/Mirwinli/coffe_plus/internal/core/domain"
	order_ports_in "github.com/Mirwinli/coffe_plus/internal/feature/order/ports/in"
	shop_ports_in "github.com/Mirwinli/coffe_plus/internal/feature/shop/ports/in"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/uuid"
)

const (
	getCreatedOrdersCallbackData = "get_created_orders" // impl
	getCookingOrdersCallbackData = "get_cooking_orders" // impl
	closeShopCallbackData        = "close_shop"         // impl
	openShopCallbackData         = "open_shop"          // impl
	acceptOrderCallbackData      = "accept:"            // impl
	rejectOrderCallbackData      = "reject:"            // impl
	readyOrderCallbackData       = "ready:"             // impl
	menuCallbackData             = "menu"

	orderStatusCreated = "created"
	shopStatusClosed   = "closed"
	shopStatusOpen     = "open"
)

func (b *Bot) handleCallback(ctx context.Context, callback tgbotapi.CallbackQuery) error {
	data := callback.Data
	b.bot.Request(tgbotapi.NewCallback(callback.ID, ""))

	switch {
	case data == getCookingOrdersCallbackData:
		if err := b.sendOrdersByStatus(ctx, domain.StatusCooking); err != nil {
			return fmt.Errorf(
				"get orders by status=cooking: %w", err,
			)
		}
	case data == getCreatedOrdersCallbackData:
		if err := b.sendOrdersByStatus(ctx, domain.StatusCreated); err != nil {
			return fmt.Errorf(
				"get orders by status=created: %w", err,
			)
		}
	case data == closeShopCallbackData:
		params := shop_ports_in.NewCangeShopStatusParams(shopStatusClosed)

		if err := b.shopService.CangeShopStatus(ctx, params); err != nil {
			return fmt.Errorf(
				"closed shop: %w", err,
			)
		}

		keyboard := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("🏠 Головне меню", menuCallbackData),
			),
		)

		msg := tgbotapi.NewMessage(b.chatID, "Магазин успішно закритий!")
		msg.ReplyMarkup = keyboard

		if _, err := b.bot.Send(msg); err != nil {
			return fmt.Errorf(
				"send message: %w", err,
			)
		}
	case data == openShopCallbackData:
		params := shop_ports_in.NewCangeShopStatusParams(shopStatusOpen)

		if err := b.shopService.CangeShopStatus(ctx, params); err != nil {
			return fmt.Errorf(
				"closed shop: %w", err,
			)
		}

		keyboard := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("🏠 Головне меню", menuCallbackData),
			),
		)

		msg := tgbotapi.NewMessage(b.chatID, "Магазин успішно відкритий!")
		msg.ReplyMarkup = keyboard

		if _, err := b.bot.Send(msg); err != nil {
			return fmt.Errorf(
				"send message: %w", err,
			)
		}
	case strings.HasPrefix(data, acceptOrderCallbackData):
		orderID, err := b.updateStatusOrder(ctx, callback, domain.StatusCooking)
		if err != nil {
			return err
		}

		keyboard := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("🏠 Головне меню", menuCallbackData),
			),
		)

		text := fmt.Sprintf("✅ Замовлення #%s прийнято! Готуємо 👨‍🍳", orderID)
		msg := tgbotapi.NewMessage(b.chatID, text)
		msg.ReplyMarkup = keyboard

		if _, err := b.bot.Send(msg); err != nil {
			return fmt.Errorf(
				"send message: %w", err,
			)
		}
	case strings.HasPrefix(data, rejectOrderCallbackData):
		orderID, err := b.updateStatusOrder(ctx, callback, domain.StatusCanceled)
		if err != nil {
			return err
		}

		keyboard := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("🏠 Головне меню", menuCallbackData),
			),
		)

		text := fmt.Sprintf("❌ Замовлення #%s скасовано!", orderID)
		msg := tgbotapi.NewMessage(b.chatID, text)
		msg.ReplyMarkup = keyboard

		if _, err := b.bot.Send(msg); err != nil {
			return fmt.Errorf(
				"send message: %w", err,
			)
		}
	case strings.HasPrefix(data, readyOrderCallbackData):
		orderID, err := b.updateStatusOrder(ctx, callback, domain.StatusWaitingForCustomer)
		if err != nil {
			return err
		}

		keyboard := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("🏠 Головне меню", menuCallbackData),
			),
		)

		text := fmt.Sprintf("✅ Замовлення #%s приготовано!", orderID)
		msg := tgbotapi.NewMessage(b.chatID, text)
		msg.ReplyMarkup = keyboard

		if _, err := b.bot.Send(msg); err != nil {
			return fmt.Errorf(
				"send message: %w", err,
			)
		}
	case data == menuCallbackData:
		b.sendMenu()
	}
	return nil
}

func (b *Bot) sendOrdersByStatus(ctx context.Context, status string) error {
	params := order_ports_in.NewAdminListOrdersParams(&status)

	getOrdersResult, err := b.orderService.AdminListOrders(ctx, params)
	if err != nil {
		return fmt.Errorf(
			"get orders with status=created: %w", err,
		)
	}

	orders := getOrdersResult.Orders

	for _, order := range orders {
		text := fmt.Sprintf(
			"📦 Замовлення #%s\n👤 %s %s\n📞 %s\n\n",
			order.ID.String(),
			order.OrderReceiver.FirstName,
			order.OrderReceiver.LastName,
			order.OrderReceiver.PhoneNumber,
		)

		for _, item := range order.Items {
			text += fmt.Sprintf("• %s x%d — %s грн\n", item.Name, item.Quantity, item.Price_Per_Unit)
		}

		text += fmt.Sprintf("\n💰 Всього: %s грн", order.Price)

		var keyboard tgbotapi.InlineKeyboardMarkup

		switch status {
		case domain.StatusCreated:
			keyboard = tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("✅ Прийняти", acceptOrderCallbackData+order.ID.String()),
					tgbotapi.NewInlineKeyboardButtonData("❌ Відхилити", rejectOrderCallbackData+order.ID.String()),
				),
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("🏠 Головне меню", menuCallbackData),
				),
			)
		case domain.StatusCooking:
			keyboard = tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("✅ Готово", readyOrderCallbackData+order.ID.String()),
					tgbotapi.NewInlineKeyboardButtonData("❌ Скасувати", rejectOrderCallbackData+order.ID.String()),
				),
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("🏠 Головне меню", menuCallbackData),
				),
			)
		}

		msg := tgbotapi.NewMessage(b.chatID, text)
		msg.ReplyMarkup = keyboard

		if _, err := b.bot.Send(msg); err != nil {
			return fmt.Errorf(
				"send message: %w", err,
			)
		}
	}
	return nil
}

func (b *Bot) updateStatusOrder(ctx context.Context, callback tgbotapi.CallbackQuery, status string) (uuid.UUID, error) {
	orderID, err := b.getOrderIDFromCallback(callback)
	if err != nil {
		return uuid.Nil, fmt.Errorf(
			"get orderID from callback data: %w", err,
		)
	}

	params := order_ports_in.NewUpdateOrderParams(status, orderID)
	if _, err := b.orderService.UpdateOrder(ctx, params); err != nil {
		if _, err := b.bot.Send(tgbotapi.NewMessage(b.chatID, "Помилка!,не змогли оновити статус,покупець не проінформований!!!")); err != nil {
			return uuid.Nil, fmt.Errorf(
				"send message: %w", err,
			)
		}
		return uuid.Nil, fmt.Errorf(
			"update order status: %w", err,
		)
	}

	return orderID, nil
}

func (b *Bot) getOrderIDFromCallback(callback tgbotapi.CallbackQuery) (uuid.UUID, error) {
	parts := strings.Split(callback.Data, ":")

	orderID, err := uuid.Parse(parts[1])
	if err != nil {
		b.bot.Send(tgbotapi.NewMessage(b.chatID, "Помилка!,не змогли получити ID замовлення"))
		return uuid.Nil, fmt.Errorf(
			"parse UUID: %w", err,
		)
	}

	return orderID, nil
}
