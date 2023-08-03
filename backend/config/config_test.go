package config

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	// Set environment variable for the config file path
	os.Setenv("CONFIG_FILE", "config.yaml")

	cfg, err := LoadConfig()
	assert.NoError(t, err)

	// Test values from the example config.yaml
	assert.Equal(t, "your-db-username", cfg.GetString("database.username"))
	assert.Equal(t, "your-db-password", cfg.GetString("database.password"))
	assert.Equal(t, "localhost", cfg.GetString("database.host"))
	assert.Equal(t, 5432, cfg.GetInt("database.port"))
	assert.Equal(t, "your-db-name", cfg.GetString("database.name"))

	assert.Equal(t, 8080, cfg.GetInt("server.port"))

	assert.Equal(t, "your-jwt-secret-key", cfg.GetString("jwt.secret_key"))
	assert.Equal(t, time.Hour*24, cfg.GetDuration("jwt.expiration"))
}

func TestInvalidConfigFile(t *testing.T) {
	// Set environment variable for a non-existent config file
	os.Setenv("CONFIG_FILE", "nonexistent.yaml")

	cfg, err := LoadConfig()
	assert.Error(t, err)
	assert.Nil(t, cfg)
}
