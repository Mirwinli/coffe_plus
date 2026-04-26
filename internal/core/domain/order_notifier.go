package domain

type OrderNotifier interface {
	SendEmail(order Order, status string) error
}
