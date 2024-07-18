package config

import (
	"strings"

	"github.com/joho/godotenv"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/v2"
	"github.com/rohitvpatil0810/go-url-shortener-api/pkg/logger"
)

type Config struct {
	Port           string         `koanf:"port"`
	PostgresConfig PostgresConfig `koanf:"postgres"`
}

type PostgresConfig struct {
	User     string `koanf:"user"`
	Password string `koanf:"password"`
	DB       string `koanf:"db"`
	Url      string `koanf:"url"`
}

var k = koanf.New(".")

func LoadConfig(envPath ...string) (*Config, error) {
	var err error
	if len(envPath) > 0 {
		err = godotenv.Load(envPath[0])
	} else {
		err = godotenv.Load()
	}

	if err != nil {
		logger.Logger.Error("Error loading .env file")
		return nil, err
	}

	k.Load(env.Provider("", ".", func(s string) string {
		return strings.ReplaceAll(strings.ToLower(s), "_", ".")
	}), nil)

	var cfg Config
	if err := k.Unmarshal("", &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
