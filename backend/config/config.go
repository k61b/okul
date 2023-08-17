package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	Utils    UtilsConfig    `mapstructure:"utils"`
}

type ServerConfig struct {
	Port string `mapstructure:"port"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
}

type UtilsConfig struct {
	JWT_Secret        string `mapstructure:"jwt_secret"`
	JWT_TokenDuration int    `mapstructure:"jwt_token_duration"`
}

func LoadConfig(env string) (*Config, error) {
	cfg := viper.New()

	cfg.SetConfigName("config." + env)
	cfg.SetConfigType("yaml")

	cfg.AddConfigPath(".")
	cfg.AddConfigPath("./config")

	err := cfg.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config
	err = cfg.Unmarshal(&config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	ApplyEnvironmentOverrides(&config, env)

	err = ValidateConfig(&config)
	if err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}

	return &config, nil
}

func ApplyEnvironmentOverrides(cfg *Config, env string) {
	switch env {
	case "dev":
		cfg.Server.Port = "8080"
	case "prod":
		cfg.Server.Port = "80"
	default:
		cfg.Server.Port = "8080"
	}
}

func ValidateConfig(cfg *Config) error {
	if cfg.Server.Port == "" {
		return fmt.Errorf("server port is missing")
	}
	if cfg.Database.Host == "" {
		return fmt.Errorf("database host is missing")
	}

	return nil
}
