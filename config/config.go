package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

type Config struct {
	GOOGLE_KEY      string `env:"GOOGLE_KEY"`
	GOOGLE_SECRET   string `env:"GOOGLE_SECRET"`
	GOOGLE_CALLBACK string `env:"GOOGLE_CALLBACK"`
}

func (c *Config) ReadFromEnv() error {
	if err := godotenv.Load(); err != nil {
		return errors.WithStack(err)
	}
	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := env.Parse(c); err != nil {
		return errors.WithStack(err)
	} else {
		return validate.Struct(c)
	}
}
