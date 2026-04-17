package auth_service

import (
	"context"
	"errors"
	"fmt"

	core_errors "github.com/Mirwinli/coffe_plus/internal/core/errors"
	core_http_tokens "github.com/Mirwinli/coffe_plus/internal/core/transport/http/tokens"
	auth_ports_in "github.com/Mirwinli/coffe_plus/internal/feature/auth/ports/in"
	auth_ports_out "github.com/Mirwinli/coffe_plus/internal/feature/auth/ports/out"
)

func (s *AuthService) Logout(
	ctx context.Context,
	in auth_ports_in.LogoutAuthParams,
) error {
	refreshToken := core_http_tokens.HashToken(in.RefreshToken)

	params := auth_ports_out.NewLogoutUserAuthParams(in.DeviceName, refreshToken)
	if err := s.authRepository.LogoutUser(ctx, params); err != nil {
		if errors.Is(err, core_errors.ErrNotFound) {
			return fmt.Errorf("logout user: %w", core_errors.ErrUnauthorized)
		}

		return fmt.Errorf("logout user: %w", err)
	}

	return nil
}
