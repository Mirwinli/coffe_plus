package auth_ports_out

import (
	"github.com/Mirwinli/coffe_plus/internal/core/domain"
	"github.com/google/uuid"
)

type GetUserParams struct {
	UserID uuid.UUID
}

func NewGetUserParams(userID uuid.UUID) GetUserParams {
	return GetUserParams{
		UserID: userID,
	}
}

type GetUserResult struct {
	User domain.User
}

func NewGetUserResult(user domain.User) GetUserResult {
	return GetUserResult{
		User: user,
	}
}
