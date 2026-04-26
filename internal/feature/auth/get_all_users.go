package auth_service

import (
	"context"
	"fmt"

	auth_ports_in "github.com/Mirwinli/coffe_plus/internal/feature/auth/ports/in"
	auth_ports_out "github.com/Mirwinli/coffe_plus/internal/feature/auth/ports/out"
)

func (s *AuthService) GetAllUsers(
	ctx context.Context,
	in auth_ports_in.GetAllUsersParams,
) (auth_ports_in.GetAllUsersResult, error) {
	params := auth_ports_out.NewGetAllUsersParams(in.Limit, in.Offset)

	getAllUsersResult, err := s.authRepository.GetAllUsers(ctx, params)
	if err != nil {
		return auth_ports_in.GetAllUsersResult{}, fmt.Errorf(
			"get all users from repository: %w", err,
		)
	}

	return auth_ports_in.NewGetAllUsersResult(getAllUsersResult.Users), nil
}
