package category_service

import (
	"context"
	"fmt"

	category_ports_in "github.com/Mirwinli/coffe_plus/internal/feature/category/ports/in"
	categories_ports_out "github.com/Mirwinli/coffe_plus/internal/feature/category/ports/out"
)

func (s *CategoryService) GetAllCategories(
	ctx context.Context,
	in category_ports_in.GetAllCategoriesParams,
) (category_ports_in.GetAllCategoriesResult, error) {
	params := categories_ports_out.NewGetAllCategoriesParams(in.Limit, in.Offset)

	categories, err := s.CategoryRepository.GetAllCategories(ctx, params)
	if err != nil {
		return category_ports_in.GetAllCategoriesResult{}, fmt.Errorf(
			"get all categories from repository: %w", err,
		)
	}

	return category_ports_in.NewGetAllCategoriesResult(categories.Categories), nil
}
