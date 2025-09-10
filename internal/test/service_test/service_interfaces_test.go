package service_test

import (
	"testing"

	"github.com/itsLeonB/drex/internal/service"
	"github.com/stretchr/testify/assert"
)

func TestServiceInterfaces(t *testing.T) {
	// Test that the interfaces are properly defined
	var debtService service.DebtTransactionService
	var transferService service.TransferMethodService
	
	assert.Nil(t, debtService)
	assert.Nil(t, transferService)
}
