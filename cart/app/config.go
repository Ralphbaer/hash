package app

import "github.com/Ralphbaer/hash/common"

// Config is the top level configuration struct for entire application
type Config struct {
	EnvName       string `env:"ENV_NAME"`
	ServerAddress string `env:"SERVER_ADDRESS"`
	SpecURL       string `env:"SPEC_URL"`
}

// NewConfig creates a instance of Config
func NewConfig() *Config {
	cfg := &Config{}
	common.SetConfigFromEnvVars(cfg)
	return cfg
}
