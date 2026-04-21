package products_adapters_out_postgres

import (
	"github.com/Mirwinli/coffe_plus/internal/core/domain"
	"github.com/google/uuid"
)

type ProductModel struct {
	ID          uuid.UUID
	Version     int
	Name        string
	Description *string
	Price       domain.Money
	IsAvaible   bool
	CategoryID  uuid.UUID
	ImageURL    string
	PublicID    string
}
