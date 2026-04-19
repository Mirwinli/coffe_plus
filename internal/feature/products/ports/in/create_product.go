package products_ports_in

import (
	"mime/multipart"

	"github.com/Mirwinli/coffe_plus/internal/core/domain"
	"github.com/google/uuid"
)

type CreateProductParams struct {
	Name        string
	Description *string
	Price       float64
	IsAvailable bool
	CategoryID  uuid.UUID
	ImageFile   multipart.File
	ImageName   string
}

func NewCreateProductParams(
	name string,
	description *string,
	price float64,
	isAvailable bool,
	categoryID uuid.UUID,
	imageFile multipart.File,
	imageName string,
) CreateProductParams {
	return CreateProductParams{
		Name:        name,
		Description: description,
		Price:       price,
		IsAvailable: isAvailable,
		CategoryID:  categoryID,
		ImageFile:   imageFile,
		ImageName:   imageName,
	}
}

type CreateProductResult struct {
	Product domain.Product
}

func NewCreateProductResult(product domain.Product) CreateProductResult {
	return CreateProductResult{Product: product}
}
