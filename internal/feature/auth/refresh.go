package auth_service

import (
	"context"
	"fmt"
	"time"

	core_errors "github.com/Mirwinli/coffe_plus/internal/core/errors"
	core_http_tokens "github.com/Mirwinli/coffe_plus/internal/core/transport/http/tokens"
	core_http_jwt "github.com/Mirwinli/coffe_plus/internal/core/transport/http/tokens/jwt"
	auth_ports_in "github.com/Mirwinli/coffe_plus/internal/feature/auth/ports/in"
	auth_ports_out "github.com/Mirwinli/coffe_plus/internal/feature/auth/ports/out"
)

func (s *AuthService) Refresh(
	ctx context.Context,
	in auth_ports_in.RefreshAuthParams,
) (auth_ports_in.RefreshAuthResult, error) {

	hashedToken := core_http_tokens.HashToken(in.RefreshToken)

	result, err := s.authRepository.GetAndDeleteRefreshToken(
		ctx,
		auth_ports_out.NewGetRefreshParams(hashedToken, in.DeviceName),
	)
	if err != nil {
		return auth_ports_in.RefreshAuthResult{}, fmt.Errorf(
			"get refresh token: %w", core_errors.ErrUnauthorized,
		)
	}

	if result.ExpiresAt.Before(time.Now()) {
		return auth_ports_in.RefreshAuthResult{}, fmt.Errorf(
			"refresh token is expired: %w",
			core_errors.ErrUnauthorized,
		)
	}

	refreshToken, err := core_http_tokens.GenerateRefreshTokens()
	if err != nil {
		return auth_ports_in.RefreshAuthResult{}, fmt.Errorf(
			"generate refresh token: %w", err,
		)
	}

	expires := time.Now().Add(s.JWTConfig.RefreshTokenTTL)
	token := core_http_tokens.HashToken(refreshToken)

	params := auth_ports_out.NewSaveRefreshTokenAuthParams(
		token,
		result.UserID,
		in.DeviceName,
		in.IpAddress,
		expires,
	)

	if err = s.authRepository.SaveRefreshToken(ctx, params); err != nil {
		return auth_ports_in.RefreshAuthResult{}, fmt.Errorf(
			"save refresh token: %w", err,
		)
	}

	userResult, err := s.authRepository.GetUser(ctx, auth_ports_out.NewGetUserParams(result.UserID))
	if err != nil {
		return auth_ports_in.RefreshAuthResult{}, fmt.Errorf(
			"get user: %w", err,
		)
	}

	accessToken, err := core_http_jwt.CreateToken(
		result.UserID,
		userResult.Role,
		s.JWTConfig,
	)
	if err != nil {
		return auth_ports_in.RefreshAuthResult{}, fmt.Errorf(
			"create access token: %w", err,
		)
	}

	return auth_ports_in.NewRefreshAuthResult(refreshToken, accessToken), nil
}
