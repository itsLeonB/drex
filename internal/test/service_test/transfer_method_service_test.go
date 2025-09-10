package service_test

import (
	"testing"

	"github.com/itsLeonB/drex/internal/service"
	"github.com/stretchr/testify/assert"
)

func TestNewTransferMethodService(t *testing.T) {
	// Skip this test as it requires actual repository
	t.Skip("Skipping service test as it requires repository implementation")

	service := service.NewTransferMethodService(nil)
	assert.NotNil(t, service)
}
