package repository_test

import (
	"testing"

	"github.com/itsLeonB/drex/internal/repository"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestNewDebtTransactionRepository(t *testing.T) {
	// Skip this test as it requires actual database connection
	t.Skip("Skipping repository test as it requires database connection")

	var db *gorm.DB // This would be a real DB connection in integration tests
	repo := repository.NewDebtTransactionRepository(db)
	assert.NotNil(t, repo)
}
