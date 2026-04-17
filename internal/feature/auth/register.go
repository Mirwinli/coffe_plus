package auth_service

import (
	"context"
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

func (s *AuthService) Register(
	ctx context.Context,
	params auth_ports_in.RegisterAuthParams,
) (auth_ports_in.RegisterAuthResult, error) {

	passHash, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		return auth_ports_in.RegisterAuthResult{}, fmt.Errorf(
			"hashing password: %w",
			core_errors.ErrInternalServerError,
		)
	}

	user := domain.NewUserUninitialized(
		params.FirstName,
		params.LastName,
		string(passHash),
		params.Email,
		params.PhoneNumber,
	)

	if err = user.Validate(); err != nil {
		return auth_ports_in.RegisterAuthResult{}, fmt.Errorf(
			"user validate: %w",
			err,
		)
	}

	in := auth_ports_out.NewSaveUserAuthParams(user)

	result, err := s.authRepository.SaveUser(ctx, in)
	if err != nil {
		return auth_ports_in.RegisterAuthResult{}, fmt.Errorf(
			"save user in repository: %w",
			err,
		)
	}

	accessToken, err := core_http_jwt.CreateToken(result.ID, result.Role, s.JWTConfig)
	if err != nil {
		return auth_ports_in.RegisterAuthResult{}, fmt.Errorf(
			"create access token: %w",
			err,
		)
	}

	refreshToken, err := core_http_tokens.GenerateRefreshTokens()
	if err != nil {
		return auth_ports_in.RegisterAuthResult{}, fmt.Errorf(
			"generate refresh token: %w",
			err,
		)
	}

	hashedToken := core_http_tokens.HashToken(refreshToken)
	RefreshExpiresAT := time.Now().Add(s.JWTConfig.RefreshTokenTTL)

	saveRefreshIn := auth_ports_out.NewSaveRefreshTokenAuthParams(
		hashedToken,
		result.ID,
		params.DeviceName,
		params.IpAddress,
		RefreshExpiresAT,
	)

	if err = s.authRepository.SaveRefreshToken(ctx, saveRefreshIn); err != nil {
		return auth_ports_in.RegisterAuthResult{}, fmt.Errorf(
			"save refresh token in repository: %w",
			err,
		)
	}

	return auth_ports_in.NewRegisterAuthResult(
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
