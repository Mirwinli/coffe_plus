package domain

import (
	"time"

	"github.com/google/uuid"
)

type Cart struct {
	ID        uuid.UUID `json:"id"`
	Items     []Item    `json:"items"`
	Price     Money     `json:"total_price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewCart(
	id uuid.UUID,
	items []Item,
	price Money,
	createdAt time.Time,
	updatedAt time.Time,
) Cart {
	return Cart{
		ID:        id,
		Items:     items,
		Price:     price,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

type Item struct {
	ProductID      uuid.UUID `json:"product_id"`
	Name           string    `json:"name"`
	ImageURL       string    `json:"image_url"`
	Quantity       int       `json:"quantity"`
	Price_Per_Unit Money     `json:"price_per_unit"`
}

func NewItem(
	productID uuid.UUID,
	quantity int,
	price_per_unit Money,
	name string,
	imageURL string,
) Item {
	return Item{
		ProductID:      productID,
		Quantity:       quantity,
		Price_Per_Unit: price_per_unit,
		Name:           name,
		ImageURL:       imageURL,
	}
}
