package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// LoadConfig loads the application configuration using viper
func LoadConfig() (*viper.Viper, error) {
	cfg := viper.New()

	cfg.SetConfigName("config") // Specify the name of the config file without extension
	cfg.SetConfigType("yaml")   // Set the config file type

	// Add the paths to search for the config file (current directory and ./config)
	cfg.AddConfigPath(".")
	cfg.AddConfigPath("./config")

	// Read the configuration file
	if err := cfg.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	return cfg, nil
}
