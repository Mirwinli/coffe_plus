package auth_service

import (
	core_http_jwt "github.com/Mirwinli/coffe_plus/internal/core/transport/http/tokens/jwt"
	repository "github.com/Mirwinli/coffe_plus/internal/feature/auth/ports/out"
)

type AuthService struct {
	authRepository repository.AuthRepository
	JWTConfig      core_http_jwt.Config
}

func NewAuthService(
	authRepository repository.AuthRepository,
	config core_http_jwt.Config,
) *AuthService {
	return &AuthService{
		authRepository: authRepository,
		JWTConfig:      config,
	}
}
