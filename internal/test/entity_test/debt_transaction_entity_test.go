package entity_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/itsLeonB/drex/internal/appconstant"
	"github.com/itsLeonB/drex/internal/entity"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestDebtTransaction_Fields(t *testing.T) {
	lenderID := uuid.New()
	borrowerID := uuid.New()
	transferMethodID := uuid.New()
	amount := decimal.NewFromFloat(100.50)

	dt := entity.DebtTransaction{
		LenderProfileID:   lenderID,
		BorrowerProfileID: borrowerID,
		Type:              appconstant.Lend,
		Action:            appconstant.LendAction,
		Amount:            amount,
		TransferMethodID:  transferMethodID,
		Description:       "Test transaction",
	}

	assert.Equal(t, lenderID, dt.LenderProfileID)
	assert.Equal(t, borrowerID, dt.BorrowerProfileID)
	assert.Equal(t, appconstant.Lend, dt.Type)
	assert.Equal(t, appconstant.LendAction, dt.Action)
	assert.Equal(t, amount, dt.Amount)
	assert.Equal(t, transferMethodID, dt.TransferMethodID)
	assert.Equal(t, "Test transaction", dt.Description)
}
