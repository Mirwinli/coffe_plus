package category_service

import (
	"context"
	"fmt"

	categories_ports_in "github.com/Mirwinli/coffe_plus/internal/feature/category/ports/in"
	categories_ports_out "github.com/Mirwinli/coffe_plus/internal/feature/category/ports/out"
)

func (s *CategoryService) PatchCategory(
	ctx context.Context,
	in categories_ports_in.PatchCategoryParams,
) (categories_ports_in.PatchCategoryResult, error) {
	getParams := categories_ports_out.NewGetCategoryParams(in.ID)
	result, err := s.CategoryRepository.GetCategory(ctx, getParams)
	if err != nil {
		return categories_ports_in.PatchCategoryResult{}, fmt.Errorf(
			"get category from repository: %w",
			err,
		)
	}

	category := result.Category
	if err = category.ApplyPatch(in.Patch); err != nil {
		return categories_ports_in.PatchCategoryResult{}, fmt.Errorf(
			"apply patch category: %w",
			err,
		)
	}

	params := categories_ports_out.NewUpdateCategoryParams(category)
	patchedCategory, err := s.CategoryRepository.UpdateCategory(ctx, params)
	if err != nil {
		return categories_ports_in.PatchCategoryResult{}, fmt.Errorf(
			"update category from repository: %w",
			err,
		)
	}

	return categories_ports_in.NewPatchCategoryResult(patchedCategory.Category), nil
}
