package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	tests := []struct {
		name        string
		env         string
		expectedErr bool
	}{
		{
			name:        "Load Dev Config",
			env:         "dev",
			expectedErr: false,
		},
		{
			name:        "Load Prod Config",
			env:         "prod",
			expectedErr: false,
		},
		{
			name:        "Invalid Config",
			env:         "invalid_env",
			expectedErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Load the configuration
			cfg, err := LoadConfig(tt.env)
			if tt.expectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, cfg)
			}
		})
	}
}

func TestValidateConfig(t *testing.T) {
	tests := []struct {
		name        string
		config      Config
		expectedErr bool
	}{
		{
			name: "Valid Configuration",
			config: Config{
				Server:   ServerConfig{Port: "8080"},
				Database: DatabaseConfig{Host: "localhost", Port: "5432", User: "user", Password: "pass", Name: "db"},
			},
			expectedErr: false,
		},
		{
			name:        "Invalid Server Port",
			config:      Config{Server: ServerConfig{Port: ""}, Database: DatabaseConfig{Host: "localhost"}},
			expectedErr: true,
		},
		{
			name:        "Invalid Database Host",
			config:      Config{Server: ServerConfig{Port: "8080"}, Database: DatabaseConfig{Host: ""}},
			expectedErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateConfig(&tt.config)
			if tt.expectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
