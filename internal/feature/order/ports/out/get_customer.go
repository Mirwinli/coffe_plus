package order_ports_out

import (
	"github.com/Mirwinli/coffe_plus/internal/core/domain"
	"github.com/google/uuid"
)

type GetCustomerParams struct {
	UserID uuid.UUID
}

func NewGetCustomerParams(userID uuid.UUID) GetCustomerParams {
	return GetCustomerParams{
		UserID: userID,
	}
}

type GetCustomerResult struct {
	User domain.User
}

func NewGetCustomerResult(user domain.User) GetCustomerResult {
	return GetCustomerResult{
		User: user,
	}
}
