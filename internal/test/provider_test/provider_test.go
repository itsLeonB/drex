package provider_test

import (
	"testing"

	"github.com/itsLeonB/drex/internal/config"
	"github.com/itsLeonB/drex/internal/provider"
	"github.com/stretchr/testify/assert"
)

func TestProvider_Fields(t *testing.T) {
	// Skip this test as it requires actual database connection
	t.Skip("Skipping provider test as it requires database connection")

	configs := config.Config{
		App: config.App{
			Name: "TestApp",
			Env:  "debug",
		},
		DB: config.DB{
			Driver:   "postgres",
			Host:     "localhost",
			Port:     "5432",
			User:     "test",
			Password: "test",
			Name:     "test",
		},
	}

	provider := provider.All(configs)

	assert.NotNil(t, provider.Logger)
	assert.NotNil(t, provider.DBs)
	assert.NotNil(t, provider.Repositories)
	assert.NotNil(t, provider.Services)

	err := provider.Shutdown()
	assert.NoError(t, err)
}
