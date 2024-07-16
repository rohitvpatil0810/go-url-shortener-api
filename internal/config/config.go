package config

import (
	"log/slog"
	"strings"

	"github.com/joho/godotenv"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/v2"
)

type Config struct {
	Port string `koanf:"port"`
}

var k = koanf.New(".")

func LoadConfig() (*Config, error) {
	err := godotenv.Load()

	if err != nil {
		slog.Error("Error loading .env file")
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
