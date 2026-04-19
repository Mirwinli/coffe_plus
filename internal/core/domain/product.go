package domain

import (
	"fmt"

	core_errors "github.com/Mirwinli/coffe_plus/internal/core/errors"
	"github.com/google/uuid"
)

type Product struct {
	ID          uuid.UUID
	Version     int
	Name        string
	Description *string
	Price       float64
	IsAvaible   bool
	CategoryID  uuid.UUID
	ImageURL    string
	PublicID    string
}

func NewProductUninitialized(
	name string,
	description *string,
	price float64,
	isAvaible bool,
	categoryID uuid.UUID,
	imageURL string,
	publicID string,
) Product {
	return Product{
		ID:          uuid.New(),
		Version:     versionUnitialized,
		Name:        name,
		Description: description,
		Price:       price,
		CategoryID:  categoryID,
		IsAvaible:   isAvaible,
		ImageURL:    imageURL,
		PublicID:    publicID,
	}
}

func NewProduct(
	id uuid.UUID,
	version int,
	name string,
	description *string,
	price float64,
	isAvaible bool,
	categoryID uuid.UUID,
	imageURL string,
	publicID string,
) Product {
	return Product{
		ID:          id,
		Version:     version,
		Name:        name,
		Description: description,
		Price:       price,
		CategoryID:  categoryID,
		IsAvaible:   isAvaible,
		ImageURL:    imageURL,
		PublicID:    publicID,
	}
}

func (p *Product) Validate() error {
	nameLen := len([]rune(p.Name))
	if nameLen < 3 || nameLen > 100 {
		return fmt.Errorf(
			"name length must be between 3 and 100: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	if p.Description != nil {
		DescriptionLen := len([]rune(*p.Description))

		if DescriptionLen < 3 || DescriptionLen > 1000 {
			return fmt.Errorf(
				"description length must be between 3 and 1000: %w",
				core_errors.ErrInvalidArgument,
			)
		}
	}

	if p.CategoryID == uuid.Nil {
		return fmt.Errorf(
			"category id cannot be NULL: %w",
			core_errors.ErrInvalidArgument,
		)
	}
	return nil
}

type ProductPatch struct {
	Name        Nullable[string]
	Description Nullable[string]
	Price       Nullable[float64]
	CategoryID  Nullable[uuid.UUID]
	IsAvailable Nullable[bool]
}

func NewProductPatch(
	name Nullable[string],
	description Nullable[string],
	price Nullable[float64],
	categoryID Nullable[uuid.UUID],
	isAvailable Nullable[bool],
) ProductPatch {
	return ProductPatch{
		Name:        name,
		Description: description,
		Price:       price,
		CategoryID:  categoryID,
		IsAvailable: isAvailable,
	}
}

func (p *ProductPatch) Validate() error {
	if p.Name.Set && p.Name.Value == nil {
		return fmt.Errorf(
			"name cannot be NULL: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	if p.CategoryID.Set && p.CategoryID.Value == nil {
		return fmt.Errorf(
			"category cannot be NULL: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	return nil
}

func (p *Product) ApplyPatch(patch ProductPatch) error {
	if err := patch.Validate(); err != nil {
		return fmt.Errorf("validate product patch: %w", err)
	}

	tmp := *p

	if patch.Name.Set {
		tmp.Name = *patch.Name.Value
	}

	if patch.Description.Set {
		tmp.Description = patch.Description.Value
	}

	if patch.Price.Set {
		tmp.Price = *patch.Price.Value
	}

	if patch.IsAvailable.Set {
		tmp.IsAvaible = *patch.IsAvailable.Value
	}

	if patch.CategoryID.Set {
		tmp.CategoryID = *patch.CategoryID.Value
	}

	if err := tmp.Validate(); err != nil {
		return fmt.Errorf("validate patched product: %w", err)
	}

	*p = tmp

	return nil
}
