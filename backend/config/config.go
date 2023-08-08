package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config represents the application configuration.
type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
}

// ServerConfig represents the server configuration.
type ServerConfig struct {
	Port string `mapstructure:"port"`
}

// DatabaseConfig represents the database configuration.
type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
}

// LoadConfig loads the application configuration based on the environment.
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

	// Apply environment-specific overrides
	ApplyEnvironmentOverrides(&config, env)

	// Validate the configuration
	err = ValidateConfig(&config)
	if err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}

	return &config, nil
}

// ApplyEnvironmentOverrides applies environment-specific overrides to the configuration.
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

// ValidateConfig validates the configuration for any missing or invalid values.
func ValidateConfig(cfg *Config) error {
	if cfg.Server.Port == "" {
		return fmt.Errorf("server port is missing")
	}
	if cfg.Database.Host == "" {
		return fmt.Errorf("database host is missing")
	}
	// Add more validation checks for other configuration fields if needed
	return nil
}
