package category_adapters_out

import (
	"github.com/Mirwinli/coffe_plus/internal/core/domain"
	"github.com/google/uuid"
)

type CategoryModel struct {
	ID      uuid.UUID
	Version int
	Name    string
}

func domainFromModel(m CategoryModel) domain.Category {
	return domain.Category{
		ID:      m.ID,
		Name:    m.Name,
		Version: m.Version,
	}
}

func domainsFromModels(m []CategoryModel) []domain.Category {
	domains := make([]domain.Category, len(m))

	for i, model := range m {
		domains[i] = domainFromModel(model)
	}

	return domains
}
