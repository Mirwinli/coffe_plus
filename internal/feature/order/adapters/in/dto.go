package order_adapters_in

import "github.com/Mirwinli/coffe_plus/internal/core/domain"

type OrderDTOResponse struct {
	Order domain.Order `json:"order"`
}

func orderDTOsFromDomains(orders []domain.Order) []OrderDTOResponse {
	dtos := make([]OrderDTOResponse, len(orders))

	for I, order := range orders {
		dtos[I] = OrderDTOResponse{
			Order: order,
		}
	}

	return dtos
}
