package products_adapters_in_products_transport_http

import (
	"github.com/Mirwinli/coffe_plus/internal/core/domain"
	"github.com/google/uuid"
)

type ProductDTOResponse struct {
	ID          uuid.UUID    `json:"id"`
	Version     int          `json:"version"`
	Name        string       `json:"name"`
	Description *string      `json:"description"`
	Price       domain.Money `json:"price"`
	IsAvailable bool         `json:"is_available"`
	CategoryID  uuid.UUID    `json:"category_id"`
	ImageURL    string       `json:"image_url"`
	PublicID    string       `json:"public_id"`
}

func productDTOsFromDomains(products []domain.Product) []ProductDTOResponse {
	dtos := make([]ProductDTOResponse, len(products))
	for i, product := range products {
		dtos[i] = productDTOFromDomain(product)
	}

	return dtos
}

func productDTOFromDomain(product domain.Product) ProductDTOResponse {
	return ProductDTOResponse{
		ID:          product.ID,
		Version:     product.Version,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		IsAvailable: product.IsAvailable,
		CategoryID:  product.CategoryID,
		ImageURL:    product.ImageURL,
		PublicID:    product.PublicID,
	}
}
