package core_infrastructure_ordernotifier

import (
	"fmt"

	"github.com/Mirwinli/coffe_plus/internal/core/domain"
	"github.com/resend/resend-go/v2"
)

func (n *OrderNotifier) SendEmail(order domain.Order, status string) error {
	subject, html := buildEmail(order, status)

	params := &resend.SendEmailRequest{
		From: "Coffee Plus <onboarding@resend.dev>",
		//From:    n.Config.EmailAddress,
		To:      []string{order.OrderReceiver.Email},
		Subject: subject,
		Html:    html,
	}

	_, err := n.EmailClient.Emails.Send(params)
	if err != nil {
		return fmt.Errorf(
			"send email: %w", err,
		)
	}

	return nil
}
