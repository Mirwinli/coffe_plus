package domain

import (
	"fmt"

	core_errors "github.com/Mirwinli/coffe_plus/internal/core/errors"
	"github.com/google/uuid"
)

type Category struct {
	ID      uuid.UUID
	Version int
	Name    string
}

func NewCategoryUninitialized(name string) Category {
	return Category{
		ID:   uuid.New(),
		Name: name,
	}
}

func NewCategory(name string, id uuid.UUID, version int) Category {
	return Category{
		ID:      id,
		Name:    name,
		Version: version,
	}
}

func (c *Category) Validate() error {
	if c.Name == "" {
		return fmt.Errorf("Category name cannot be empty")
	}

	return nil
}

type CategoryPatch struct {
	Name Nullable[string]
}

func NewCategoryPatch(name Nullable[string]) CategoryPatch {
	return CategoryPatch{
		Name: name,
	}
}

func (p *CategoryPatch) Validate() error {
	if p.Name.Set && p.Name.Value == nil {
		return fmt.Errorf(
			"name can't be NULL: %w",
			core_errors.ErrInvalidArgument,
		)
	}
	return nil
}

func (c *Category) ApplyPatch(patch CategoryPatch) error {
	if err := patch.Validate(); err != nil {
		return fmt.Errorf(
			"validate patch: %w",
			err,
		)
	}

	tmp := *c

	if patch.Name.Set {
		tmp.Name = *patch.Name.Value
	}

	if err := tmp.Validate(); err != nil {
		return fmt.Errorf(
			"validate patched category: %w",
			err,
		)
	}

	*c = tmp

	return nil
}
