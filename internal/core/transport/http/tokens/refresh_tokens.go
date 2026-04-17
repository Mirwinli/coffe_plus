package core_http_tokens

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"

	core_errors "github.com/Mirwinli/coffe_plus/internal/core/errors"
)

func GenerateRefreshTokens() (string, error) {
	b := make([]byte, 32)

	if _, err := rand.Read(b); err != nil {
		return "", fmt.Errorf("generate refresh token: %w", core_errors.ErrInternalServerError)
	}

	return base64.RawURLEncoding.EncodeToString(b), nil
}

func HashToken(token string) string {
	h := sha256.New()
	h.Write([]byte(token))
	return fmt.Sprintf("%x", h.Sum(nil))
}
