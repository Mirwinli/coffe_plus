package auth_service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Mirwinli/coffe_plus/internal/core/domain"
	core_errors "github.com/Mirwinli/coffe_plus/internal/core/errors"
	core_http_tokens "github.com/Mirwinli/coffe_plus/internal/core/transport/http/tokens"
	core_http_jwt "github.com/Mirwinli/coffe_plus/internal/core/transport/http/tokens/jwt"
	auth_ports_in "github.com/Mirwinli/coffe_plus/internal/feature/auth/ports/in"
	auth_ports_out "github.com/Mirwinli/coffe_plus/internal/feature/auth/ports/out"
	"golang.org/x/crypto/bcrypt"
)

func (s *AuthService) Login(
	ctx context.Context,
	in auth_ports_in.LoginAuthParams,
) (auth_ports_in.LoginAuthResult, error) {
	if err := domain.ValidateEmail(in.Email); err != nil {
		return auth_ports_in.LoginAuthResult{}, fmt.Errorf(
			"user validate email: %w", err,
		)
	}

	params := auth_ports_out.NewLoginUserAuthParams(in.Email, in.Password)
	result, err := s.authRepository.LoginUser(ctx, params)
	if err != nil {
		if errors.Is(err, core_errors.ErrNotFound) {
			return auth_ports_in.LoginAuthResult{}, fmt.Errorf(
				"invalidd email or password: %w",
				core_errors.ErrInvalidCredentials,
			)
		}
		return auth_ports_in.LoginAuthResult{}, fmt.Errorf("login user: %w", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(result.PasswordHash), []byte(params.Password))
	if err != nil {
		return auth_ports_in.LoginAuthResult{}, fmt.Errorf(
			"invalid email or password: %w",
			core_errors.ErrInvalidCredentials,
		)
	}

	refreshToken, err := core_http_tokens.GenerateRefreshTokens()
	if err != nil {
		return auth_ports_in.LoginAuthResult{}, fmt.Errorf(
			"generate refresh token: %w", err,
		)
	}

	hashedToken := core_http_tokens.HashToken(refreshToken)
	ttl := time.Now().Add(s.JWTConfig.RefreshTokenTTL)

	refParams := auth_ports_out.NewSaveRefreshTokenAuthParams(
		hashedToken,
		result.ID,
		in.DeviceName,
		in.IpAddress,
		ttl,
	)

	if err = s.authRepository.SaveRefreshToken(ctx, refParams); err != nil {
		return auth_ports_in.LoginAuthResult{}, fmt.Errorf(
			"generate refresh token: %w", err,
		)
	}

	accessToken, err := core_http_jwt.CreateToken(
		result.ID,
		result.Role,
		s.JWTConfig,
	)
	if err != nil {
		return auth_ports_in.LoginAuthResult{}, fmt.Errorf("create token: %w", err)
	}

	return auth_ports_in.NewLoginAuthResult(
		result.ID,
		result.Version,
		result.FirstName,
		result.LastName,
		result.Email,
		result.PhoneNumber,
		result.Role,
		result.CreatedAt,
		accessToken,
		refreshToken,
	), nil
}
