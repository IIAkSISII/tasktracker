package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

type Config struct {
	Server   Server   `yaml:"server" env:"SERVER" env-required:"true"`
	Database Database `yaml:"database" env:"DATABASE" env-required:"true"`
	Auth     Auth     `yaml:"auth" env:"AUTH" env-required:"true"`
	Logger   Logger   `yaml:"logger" env:"LOGGER" env-required:"true"`
}

func NewConfig(path string) (*Config, error) {
	var cfg Config

	err := cleanenv.ReadConfig(path, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

func GetConfigPath() string {
	if path := os.Getenv("CONFIG_PATH"); path != "" {
		return path
	}
	return "remote/config.yaml"
}
