package repository_test

import (
	"testing"

	"github.com/itsLeonB/drex/internal/repository"
	"github.com/stretchr/testify/assert"
)

func TestRepositoryInterfaces(t *testing.T) {
	// Test that the interfaces are properly defined
	var debtRepo repository.DebtTransactionRepository
	var transferRepo repository.TransferMethodRepository
	
	assert.Nil(t, debtRepo)
	assert.Nil(t, transferRepo)
}
