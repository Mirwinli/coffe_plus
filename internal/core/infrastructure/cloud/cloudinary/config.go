package core_infrastructure_cloudinary

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	CloudName string `envconfig:"CLOUD_NAME" required:"true"`
	ApiKey    string `envconfig:"API_KEY"    required:"true"`
	ApiSecret string `envconfig:"API_SECRET" required:"true"`
}

func NewConfig() (*Config, error) {
	var config Config
	err := envconfig.Process("CLOUDINARY", &config)
	if err != nil {
		return nil, fmt.Errorf(
			"env process error: %w",
			err,
		)
	}

	return &config, nil
}

func NewMustConfig() *Config {
	config, err := NewConfig()
	if err != nil {
		err = fmt.Errorf(
			"get new cloudinary config: %w",
			err,
		)
		panic(err)
	}

	return config
}
