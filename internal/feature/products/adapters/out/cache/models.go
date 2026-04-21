package products_adapters_out_cache

import (
	"encoding/json"
	"fmt"
	"strconv"

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
	ImageURL    string
	PublicID    string
}

func domainToModel(product domain.Product) ProductModel {
	return ProductModel{
		ID:          product.ID,
		Version:     product.Version,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		IsAvaible:   product.IsAvailable,
		ImageURL:    product.ImageURL,
		PublicID:    product.PublicID,
	}
}

func modelToDomain(product ProductModel) domain.Product {
	return domain.Product{
		ID:          product.ID,
		Version:     product.Version,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		IsAvailable: product.IsAvaible,
		ImageURL:    product.ImageURL,
		PublicID:    product.PublicID,
	}
}

func productKey(id uuid.UUID) string {
	return fmt.Sprintf("product:%s", id)
}

func (m *ProductModel) Serialize() ([]byte, error) {
	bytes, err := json.Marshal(m)
	if err != nil {
		return nil, fmt.Errorf("serialize product: %w", err)
	}

	return bytes, nil
}

func (m *ProductModel) Deserialize(bytes []byte) error {
	err := json.Unmarshal(bytes, m)
	if err != nil {
		return fmt.Errorf("deserialize product: %w", err)
	}

	return nil
}

type ProductListModel []ProductModel

func (m *ProductListModel) Serialize() ([]byte, error) {
	bytes, err := json.Marshal(m)
	if err != nil {
		return nil, fmt.Errorf("serialize product list: %w", err)
	}

	return bytes, nil
}

func (m *ProductListModel) Deserialize(bytes []byte) error {
	err := json.Unmarshal(bytes, m)
	if err != nil {
		return fmt.Errorf("deserialize product list: %w", err)
	}

	return nil
}

func productsListKey(categoryID *uuid.UUID) string {
	if categoryID == nil {
		return "products:all"
	}

	return fmt.Sprintf("products:%s", *categoryID)
}

func productsListField(limit, offset *int, onlyAvailable bool) string {
	ptrStr := func(v *int) string {
		if v == nil {
			return "nil"
		}
		return strconv.Itoa(*v)
	}

	return fmt.Sprintf("%s:%s:%v", ptrStr(limit), ptrStr(offset), onlyAvailable)
}

func domainsToModel(productList []domain.Product) ProductListModel {
	models := make([]ProductModel, len(productList))

	for i, product := range productList {
		models[i] = domainToModel(product)
	}

	return models
}

func modelToDomains(model ProductListModel) []domain.Product {
	domains := make([]domain.Product, len(model))

	for i, product := range model {
		domains[i] = modelToDomain(product)
	}

	return domains
}
