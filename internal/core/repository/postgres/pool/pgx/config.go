package core_postgres_pgx

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Host     string        `envconfig:"HOST" required:"true"`
	Port     string        `envconfig:"PORT" default:"5432"`
	User     string        `envconfig:"USER" required:"true"`
	Password string        `envconfig:"PASSWORD" required:"true"`
	Database string        `envconfig:"DB" required:"true"`
	Timeout  time.Duration `envconfig:"TIMEOUT" required:"true"`
}

func NewConfig() (Config, error) {
	var config Config

	err := envconfig.Process("POSTGRES", &config)
	if err != nil {
		return config, fmt.Errorf("process envconfig: %w", err)
	}

	return config, nil
}

func NewMustConfig() Config {
	config, err := NewConfig()
	if err != nil {
		err = fmt.Errorf("get new Postgres pool config: %w", err)
		panic(err)
	}

	return config
}
