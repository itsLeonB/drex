package config_test

import (
	"os"
	"testing"
	"time"

	"github.com/itsLeonB/drex/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestConfig_DefaultValues(t *testing.T) {
	// Clear environment variables
	os.Clearenv()

	// Set required DB environment variables
	_ = os.Setenv("DB_HOST", "localhost")
	_ = os.Setenv("DB_PORT", "5432")
	_ = os.Setenv("DB_USER", "test")
	_ = os.Setenv("DB_PASSWORD", "test")
	_ = os.Setenv("DB_NAME", "test")

	cfg := config.Load()

	assert.Equal(t, "Drex", cfg.App.Name)
	assert.Equal(t, "debug", cfg.Env)
	assert.Equal(t, "50051", cfg.App.Port)
	assert.Equal(t, 10*time.Second, cfg.Timeout)

	assert.Equal(t, "postgres", cfg.Driver)
	assert.Equal(t, "localhost", cfg.Host)
	assert.Equal(t, "5432", cfg.DB.Port)
	assert.Equal(t, "test", cfg.User)
	assert.Equal(t, "test", cfg.Password)
	assert.Equal(t, "test", cfg.DB.Name)
}

func TestConfig_CustomValues(t *testing.T) {
	// Clear environment variables
	os.Clearenv()

	// Set custom values
	_ = os.Setenv("APP_NAME", "CustomDrex")
	_ = os.Setenv("APP_ENV", "production")
	_ = os.Setenv("APP_PORT", "8080")
	_ = os.Setenv("APP_TIMEOUT", "30s")

	_ = os.Setenv("DB_DRIVER", "mysql")
	_ = os.Setenv("DB_HOST", "db.example.com")
	_ = os.Setenv("DB_PORT", "3306")
	_ = os.Setenv("DB_USER", "admin")
	_ = os.Setenv("DB_PASSWORD", "secret")
	_ = os.Setenv("DB_NAME", "production_db")

	cfg := config.Load()

	assert.Equal(t, "CustomDrex", cfg.App.Name)
	assert.Equal(t, "production", cfg.Env)
	assert.Equal(t, "8080", cfg.App.Port)
	assert.Equal(t, 30*time.Second, cfg.Timeout)

	assert.Equal(t, "mysql", cfg.Driver)
	assert.Equal(t, "db.example.com", cfg.Host)
	assert.Equal(t, "3306", cfg.DB.Port)
	assert.Equal(t, "admin", cfg.User)
	assert.Equal(t, "secret", cfg.Password)
	assert.Equal(t, "production_db", cfg.DB.Name)
}
