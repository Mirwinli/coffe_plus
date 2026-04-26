package core_infrastructure_ordernotifier

import (
	"fmt"
	"strings"

	"github.com/Mirwinli/coffe_plus/internal/core/domain"
)

func buildEmail(order domain.Order, status string) (subject, html string) {
	switch status {
	case domain.StatusCreated:
		subject = "📋 Замовлення отримано — Coffee Plus"
		html = renderEmail(order, emailData{
			Icon:    "📋",
			Title:   "Замовлення отримано!",
			Message: fmt.Sprintf("Привіт, <strong>%s</strong>! Ваше замовлення отримано. Очікуємо підтвердження від кухні.", order.OrderReceiver.FirstName),
			Color:   "#3B82F6",
		})
	case domain.StatusCooking:
		subject = "👨‍🍳 Готуємо ваше замовлення — Coffee Plus"
		html = renderEmail(order, emailData{
			Icon:    "👨‍🍳",
			Title:   "Замовлення підтверджено!",
			Message: fmt.Sprintf("Привіт, <strong>%s</strong>! Замовлення підтверджено — наш кухар вже готує його для вас.", order.OrderReceiver.FirstName),
			Color:   "#F59E0B",
		})
	case domain.StatusWaitingForCustomer:
		subject = "🔔 Ваше замовлення готове — Coffee Plus"
		html = renderEmail(order, emailData{
			Icon:    "🛎️",
			Title:   "Готово! Можна забирати.",
			Message: fmt.Sprintf("Привіт, <strong>%s</strong>! Ваше замовлення готове і чекає на вас.", order.OrderReceiver.FirstName),
			Color:   "#10B981",
		})
	case domain.StatusCanceled:
		subject = "❌ Замовлення скасовано — Coffee Plus"
		html = renderEmail(order, emailData{
			Icon:    "😔",
			Title:   "Замовлення скасовано",
			Message: fmt.Sprintf("Привіт, <strong>%s</strong>! На жаль, ваше замовлення скасовано. Зв'яжіться з нами якщо маєте питання.", order.OrderReceiver.FirstName),
			Color:   "#EF4444",
		})
	case domain.StatusCompleted:
		subject = "⭐ Дякуємо за замовлення — Coffee Plus"
		html = renderEmail(order, emailData{
			Icon:    "🙏",
			Title:   "Дякуємо!",
			Message: fmt.Sprintf("Привіт, <strong>%s</strong>! Дякуємо що обрали Coffee Plus. Будемо раді бачити вас знову!", order.OrderReceiver.FirstName),
			Color:   "#6366F1",
		})
	default:
		subject = "📦 Оновлення замовлення — Coffee Plus"
		html = renderEmail(order, emailData{
			Icon:    "📦",
			Title:   "Оновлення замовлення",
			Message: fmt.Sprintf("Привіт, <strong>%s</strong>! Статус вашого замовлення оновлено.", order.OrderReceiver.FirstName),
			Color:   "#6B7280",
		})
	}
	return
}

type emailData struct {
	Icon    string
	Title   string
	Message string
	Color   string
}

func renderEmail(order domain.Order, data emailData) string {
	return fmt.Sprintf(`<!DOCTYPE html>
<html lang="uk">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Coffee Plus</title>
</head>
<body style="margin:0;padding:0;background:#F9FAFB;font-family:'Helvetica Neue',Helvetica,Arial,sans-serif;">
  <table width="100%%" cellpadding="0" cellspacing="0" style="background:#F9FAFB;padding:40px 0;">
    <tr>
      <td align="center">
        <table width="560" cellpadding="0" cellspacing="0" style="background:#ffffff;border-radius:12px;overflow:hidden;box-shadow:0 1px 3px rgba(0,0,0,0.08);">
 
          <!-- Header -->
          <tr>
            <td style="background:%s;padding:32px;text-align:center;">
              <p style="margin:0;font-size:40px;">%s</p>
              <h1 style="margin:12px 0 0;color:#ffffff;font-size:22px;font-weight:600;letter-spacing:-0.3px;">%s</h1>
            </td>
          </tr>
 
          <!-- Body -->
          <tr>
            <td style="padding:32px;">
              <p style="margin:0 0 24px;color:#374151;font-size:15px;line-height:1.6;">%s</p>
 
              <!-- Order info -->
              <table width="100%%" cellpadding="0" cellspacing="0" style="background:#F9FAFB;border-radius:8px;padding:16px;margin-bottom:24px;">
                <tr>
                  <td style="color:#6B7280;font-size:12px;text-transform:uppercase;letter-spacing:0.5px;padding-bottom:12px;">
                    Замовлення #%s
                  </td>
                </tr>
                %s
                <tr>
                  <td style="border-top:1px solid #E5E7EB;padding-top:12px;margin-top:12px;">
                    <table width="100%%">
                      <tr>
                        <td style="color:#374151;font-size:14px;font-weight:600;">Сума</td>
                        <td align="right" style="color:#374151;font-size:14px;font-weight:600;">%s грн</td>
                      </tr>
                    </table>
                  </td>
                </tr>
              </table>
 
              <p style="margin:0;color:#9CA3AF;font-size:13px;line-height:1.5;">
                Якщо у вас є питання — зв'яжіться з нами за номером <strong>%s</strong>
              </p>
            </td>
          </tr>
 
          <!-- Footer -->
          <tr>
            <td style="background:#F9FAFB;padding:20px 32px;border-top:1px solid #F3F4F6;text-align:center;">
              <p style="margin:0;color:#9CA3AF;font-size:12px;">© Coffee Plus. Всі права захищені.</p>
            </td>
          </tr>
 
        </table>
      </td>
    </tr>
  </table>
</body>
</html>`,
		data.Color,
		data.Icon,
		data.Title,
		data.Message,
		order.ID.String()[:8],
		buildItemsRows(order.Items),
		order.Price,
		order.OrderReceiver.PhoneNumber,
	)
}

func buildItemsRows(items []domain.OrderItem) string {
	var sb strings.Builder
	for _, item := range items {
		sb.WriteString(fmt.Sprintf(`
      <tr>
        <td style="padding:6px 0;">
          <table width="100%%">
            <tr>
              <td style="color:#374151;font-size:14px;">%s <span style="color:#9CA3AF;">x%d</span></td>
              <td align="right" style="color:#374151;font-size:14px;">%s грн</td>
            </tr>
          </table>
        </td>
      </tr>`, item.Name, item.Quantity, item.Price_Per_Unit))
	}
	return sb.String()
}
