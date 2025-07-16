package config

import "time"

type Server struct {
	Host         string        `yaml:"host" env:"SERVER_HOST" env-default:"localhost"`
	Port         string        `yaml:"port" env:"SERVER_PORT" env-default:"8080"`
	ReadTimeout  time.Duration `yaml:"read_timeout" env:"SERVER_READ_TIMEOUT" env-default:"10s"`
	WriteTimeout time.Duration `yaml:"write_timeout" env:"SERVER_WRITE_TIMEOUT" env-default:"10s"`
}
