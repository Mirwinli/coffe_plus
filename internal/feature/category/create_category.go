package category_service

import (
	"context"
	"fmt"

	category_ports_in "github.com/Mirwinli/coffe_plus/internal/feature/category/ports/in"
	categories_ports_out "github.com/Mirwinli/coffe_plus/internal/feature/category/ports/out"
)

func (s *CategoryService) CreateCategory(
	ctx context.Context,
	in category_ports_in.CreateCategoryParams,
) (category_ports_in.CreateCategoryResult, error) {

	params := categories_ports_out.NewSaveCategoryParams(in.Category)
	category, err := s.CategoryRepository.SaveCategory(ctx, params)
	if err != nil {
		return category_ports_in.CreateCategoryResult{}, fmt.Errorf(
			"save category by repository: %w",
			err,
		)
	}

	return category_ports_in.NewCreateCategoryResult(category.Category), nil
}
