package order_adapters_out_posgtres

import (
	"time"

	"github.com/Mirwinli/coffe_plus/internal/core/domain"
	"github.com/google/uuid"
)

type OrderModel struct {
	ID            uuid.UUID
	Version       int
	Status        string
	Price         domain.Money
	CreatedAt     time.Time
	Items         []OrderItemModel
	OrderReceiver OrderReceiver
}

func (o OrderModel) ToDomain() domain.Order {
	return domain.Order{
		ID:            o.ID,
		Version:       o.Version,
		Items:         itemModelsToDomains(o.Items),
		Status:        o.Status,
		Price:         o.Price,
		CreatedAt:     o.CreatedAt,
		OrderReceiver: o.OrderReceiver.ToDomain(),
	}

}

type OrderItemModel struct {
	ID             uuid.UUID    `json:"id"`
	OrderID        uuid.UUID    `json:"order_id"`
	ProductID      uuid.UUID    `json:"product_id"`
	Name           string       `json:"product_name"`
	ImageURL       string       `json:"image_url"`
	Quantity       int          `json:"quantity"`
	Price_Per_Unit domain.Money `json:"price_at_time"`
}

func (i OrderItemModel) ToDomain() domain.OrderItem {
	return domain.NewOrderItem(
		i.ID,
		i.ProductID,
		i.Name,
		i.ImageURL,
		i.Quantity,
		i.Price_Per_Unit,
	)
}

type OrderReceiver struct {
	CustomerID  uuid.UUID `json:"customer_id"`
	PhoneNumber string    `json:"phone_number"`
	Email       string    `json:"email"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
}

func (r OrderReceiver) ToDomain() domain.OrderReceiver {
	return domain.NewOrderReceiver(
		r.CustomerID,
		r.PhoneNumber,
		r.Email,
		r.FirstName,
		r.LastName,
	)
}

func orderModelsToDomains(orders []OrderModel) []domain.Order {
	domains := make([]domain.Order, len(orders))

	for i, order := range orders {
		domains[i] = order.ToDomain()
	}

	return domains
}

func itemModelsToDomains(items []OrderItemModel) []domain.OrderItem {
	domains := make([]domain.OrderItem, len(items))

	for i, item := range items {
		domains[i] = item.ToDomain()
	}

	return domains
}
