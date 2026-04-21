package cart_adapters_in

import "github.com/Mirwinli/coffe_plus/internal/core/domain"

type CartDTOResponse struct {
	Cart domain.Cart `json:"cart"`
}
