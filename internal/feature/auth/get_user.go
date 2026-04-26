package auth_service

import (
	"context"
	"fmt"

	auth_ports_in "github.com/Mirwinli/coffe_plus/internal/feature/auth/ports/in"
	auth_ports_out "github.com/Mirwinli/coffe_plus/internal/feature/auth/ports/out"
)

func (s *AuthService) GetUser(
	ctx context.Context,
	in auth_ports_in.GetUserParams,
) (auth_ports_in.GetUserResult, error) {
	params := auth_ports_out.NewGetUserParams(in.UserID)

	getUserResult, err := s.authRepository.GetUser(ctx, params)
	if err != nil {
		return auth_ports_in.GetUserResult{}, fmt.Errorf(
			"get user from repository: %w", err,
		)
	}

	return auth_ports_in.NewGetUserResult(getUserResult.User), nil
}
