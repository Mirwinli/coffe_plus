package core_http_jwt

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Secret          string        `envconfig:"SECRET"`
	RefreshTokenTTL time.Duration `envconfig:"REFRESH_TOKEN_TTL"`
	AccessTokenTTL  time.Duration `envconfig:"ACCESS_TOKEN_TTL"`
}

func NewConfig() (Config, error) {
	var config Config

	if err := envconfig.Process("JWT", &config); err != nil {
		return Config{}, fmt.Errorf("process envconfig: %w", err)
	}

	return config, nil
}

func NewMustConfig() Config {
	config, err := NewConfig()
	if err != nil {
		err = fmt.Errorf("get new JWT config: %w", err)
		panic(err)
	}

	return config
}
