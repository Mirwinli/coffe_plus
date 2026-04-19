package category_service

import (
	categories_ports_out "github.com/Mirwinli/coffe_plus/internal/feature/category/ports/out"
)

type CategoryService struct {
	CategoryRepository categories_ports_out.CategoriesRepository
}

func NewCategoryService(categoryRepository categories_ports_out.CategoriesRepository) *CategoryService {
	return &CategoryService{
		CategoryRepository: categoryRepository,
	}
}
