package provider_test

import (
	"testing"

	"github.com/itsLeonB/drex/internal/config"
	"github.com/itsLeonB/drex/internal/provider"
	"github.com/stretchr/testify/assert"
)

func TestProvideLogger_DebugMode(t *testing.T) {
	appConfig := config.App{
		Name: "TestApp",
		Env:  "debug",
	}

	logger := provider.ProvideLogger(appConfig)

	assert.NotNil(t, logger)
}

func TestProvideLogger_ReleaseMode(t *testing.T) {
	appConfig := config.App{
		Name: "TestApp",
		Env:  "release",
	}

	logger := provider.ProvideLogger(appConfig)

	assert.NotNil(t, logger)
}
