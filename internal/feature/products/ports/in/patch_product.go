package products_ports_in

import (
	"github.com/Mirwinli/coffe_plus/internal/core/domain"
	"github.com/google/uuid"
)

type PatchProductParams struct {
	ID    uuid.UUID
	Patch domain.ProductPatch
}

func NewPatchProductParams(
	id uuid.UUID,
	patch domain.ProductPatch,
) PatchProductParams {
	return PatchProductParams{
		ID:    id,
		Patch: patch,
	}
}

type PatchProductResult struct {
	Product domain.Product
}

func NewPatchProductResult(product domain.Product) PatchProductResult {
	return PatchProductResult{
		Product: product,
	}
}
