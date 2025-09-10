package appconstant_test

import (
	"testing"

	"github.com/itsLeonB/drex/internal/appconstant"
	"github.com/stretchr/testify/assert"
)

func TestDebtTransactionType_Constants(t *testing.T) {
	assert.Equal(t, appconstant.DebtTransactionType("LEND"), appconstant.Lend)
	assert.Equal(t, appconstant.DebtTransactionType("REPAY"), appconstant.Repay)
}

func TestDebtTransactionAction_Constants(t *testing.T) {
	assert.Equal(t, appconstant.DebtTransactionAction("LEND"), appconstant.LendAction)
	assert.Equal(t, appconstant.DebtTransactionAction("BORROW"), appconstant.BorrowAction)
	assert.Equal(t, appconstant.DebtTransactionAction("RECEIVE"), appconstant.ReceiveAction)
	assert.Equal(t, appconstant.DebtTransactionAction("RETURN"), appconstant.ReturnAction)
}

func TestGroupExpenseTransferMethod_Constant(t *testing.T) {
	assert.Equal(t, "GROUP_EXPENSE", appconstant.GroupExpenseTransferMethod)
}
