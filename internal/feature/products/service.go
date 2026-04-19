package products_service

import (
	"github.com/Mirwinli/coffe_plus/internal/core/domain"
	repository "github.com/Mirwinli/coffe_plus/internal/feature/products/ports/out"
)

type ProductsService struct {
	ProductsRepository repository.ProductRepository
	ImageUploader      domain.ImageUploader
}

func NewProductService(
	productsRepository repository.ProductRepository,
	imageUploader domain.ImageUploader,
) *ProductsService {
	return &ProductsService{
		ProductsRepository: productsRepository,
		ImageUploader:      imageUploader,
	}
}
