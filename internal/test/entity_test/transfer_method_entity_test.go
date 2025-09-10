package entity_test

import (
	"testing"

	"github.com/itsLeonB/drex/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestTransferMethod_Fields(t *testing.T) {
	tm := entity.TransferMethod{
		Name:    "bank_transfer",
		Display: "Bank Transfer",
	}

	assert.Equal(t, "bank_transfer", tm.Name)
	assert.Equal(t, "Bank Transfer", tm.Display)
}
