package categories_ports_in

import (
	"github.com/Mirwinli/coffe_plus/internal/core/domain"
	"github.com/google/uuid"
)

type PatchCategoryParams struct {
	ID    uuid.UUID
	Patch domain.CategoryPatch
}

func NewPatchCategoryParams(id uuid.UUID, patch domain.CategoryPatch) PatchCategoryParams {
	return PatchCategoryParams{
		ID:    id,
		Patch: patch,
	}
}

type PatchCategoryResult struct {
	Category domain.Category
}

func NewPatchCategoryResult(category domain.Category) PatchCategoryResult {
	return PatchCategoryResult{
		Category: category,
	}
}
