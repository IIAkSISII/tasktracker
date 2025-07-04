package config

type Database struct {
	Host     string `yaml:"host" env:"SERVER_HOST" env-default:"localhost"`
	Port     int    `yaml:"port" env:"SERVER_PORT" env-default:"5432"`
	User     string `yaml:"user" env:"SERVER_USER" env-default:"postgres"`
	Password string `yaml:"password" env:"SERVER_PASSWORD" env-required:"true"`
	Name     string `yaml:"name" env:"SERVER_NAME" env-required:"true"`
	SslMode  string `yaml:"ssl_mode" env:"SERVER_SSL_MODE" env-default:"disable"`
}
