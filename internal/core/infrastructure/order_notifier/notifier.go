package core_infrastructure_ordernotifier

import "github.com/resend/resend-go/v2"

type OrderNotifier struct {
	EmailClient *resend.Client
	Config      Config
}

func NewOrderNotifier(config Config) *OrderNotifier {
	emailClient := resend.NewClient(config.ResendApiKey)

	return &OrderNotifier{
		EmailClient: emailClient,
		Config:      config,
	}
}
