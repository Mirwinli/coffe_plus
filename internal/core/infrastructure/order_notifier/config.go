package core_infrastructure_ordernotifier

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	EmailAddress string `envconfig:"EMAIL_ADDRESS" required:"true"`
	ResendApiKey string `envconfig:"RESEND_API_KEY" required:"true"`
}

func NewConfig() (Config, error) {
	var config Config

	if err := envconfig.Process("", &config); err != nil {
		return Config{}, fmt.Errorf(
			"envconfig process: %w", err,
		)
	}

	return config, nil
}

func NewMustConfig() Config {
	config, err := NewConfig()
	if err != nil {
		err = fmt.Errorf(
			"get new email config: %w", err,
		)
		panic(err)
	}

	return config
}
