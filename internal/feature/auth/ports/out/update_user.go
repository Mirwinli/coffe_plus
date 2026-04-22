package auth_ports_out

import "github.com/Mirwinli/coffe_plus/internal/core/domain"

type UpdateUserParams struct {
	User domain.User
}

func NewUpdateUserParams(user domain.User) UpdateUserParams {
	return UpdateUserParams{
		User: user,
	}
}

type UpdateUserResult struct {
	User domain.User
}

func NewUpdateUserResult(user domain.User) UpdateUserResult {
	return UpdateUserResult{
		User: user,
	}
}
