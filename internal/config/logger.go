package config

type Logger struct {
	Level  string `yaml:"level" env:"LOG_LEVEL" env-default:"debug"`
	Format string `yaml:"format" env:"LOG_FORMAT" env-default:"text"`
}
