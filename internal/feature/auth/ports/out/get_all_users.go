package auth_ports_out

import "github.com/Mirwinli/coffe_plus/internal/core/domain"

type GetAllUsersParams struct {
	Limit  *int
	Offset *int
}

func NewGetAllUsersParams(limit, offset *int) GetAllUsersParams {
	return GetAllUsersParams{
		Limit:  limit,
		Offset: offset,
	}
}

type GetAllUsersResult struct {
	Users []domain.User
}

func NewGetAllUsersResult(users []domain.User) GetAllUsersResult {
	return GetAllUsersResult{
		Users: users,
	}
}
