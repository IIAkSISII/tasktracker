package config

import "time"

type Auth struct {
	JWTSecret string        `yaml:"jwt_secret" env:"JWT_SECRET" env-required:"true"`
	TokenTtl  time.Duration `yaml:"token_ttl" env:"TOKEN_TTL" env-default:"10m"`
}
