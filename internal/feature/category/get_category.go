package category_service

import (
	"context"
	"fmt"

	categories_ports_in "github.com/Mirwinli/coffe_plus/internal/feature/category/ports/in"
	categories_ports_out "github.com/Mirwinli/coffe_plus/internal/feature/category/ports/out"
)

func (s *CategoryService) GetCategory(
	ctx context.Context,
	in categories_ports_in.GetCategoryParams,
) (categories_ports_in.GetCategoryResult, error) {

	params := categories_ports_out.NewGetCategoryParams(in.ID)

	category, err := s.CategoryRepository.GetCategory(ctx, params)
	if err != nil {
		return categories_ports_in.GetCategoryResult{}, fmt.Errorf(
			"get category by repository: %w", err,
		)
	}

	return categories_ports_in.NewGetCategoryResult(category.Category), nil
}
