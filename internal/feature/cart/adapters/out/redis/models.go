package cart_adapters_out_redis

import (
	"fmt"

	"github.com/google/uuid"
)

func cartKey(cartID uuid.UUID) string {
	return fmt.Sprintf("cart:%s", cartID)
}
