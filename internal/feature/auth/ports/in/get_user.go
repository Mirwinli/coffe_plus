package auth_ports_in

import (
	"github.com/Mirwinli/coffe_plus/internal/core/domain"
	"github.com/google/uuid"
)

type GetUserParams struct {
	UserID uuid.UUID
}

func NewGetUserParams(id uuid.UUID) GetUserParams {
	return GetUserParams{UserID: id}
}

type GetUserResult struct {
	User domain.User
}

func NewGetUserResult(user domain.User) GetUserResult {
	return GetUserResult{
		User: user,
	}
}
