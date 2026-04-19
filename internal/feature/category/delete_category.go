package category_service

import (
	"context"
	"fmt"

	category_ports_in "github.com/Mirwinli/coffe_plus/internal/feature/category/ports/in"
	categories_ports_out "github.com/Mirwinli/coffe_plus/internal/feature/category/ports/out"
)

func (s *CategoryService) DeleteCategory(
	ctx context.Context,
	in category_ports_in.DeleteCategoryParams,
) error {
	params := categories_ports_out.NewDeleteCategoryParams(in.ID)

	if err := s.CategoryRepository.DeleteCategory(ctx, params); err != nil {
		return fmt.Errorf("delete category from repository: %w", err)
	}

	return nil
}
