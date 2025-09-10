package helper_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/itsLeonB/drex/internal/appconstant"
	"github.com/itsLeonB/drex/internal/entity"
	"github.com/itsLeonB/drex/internal/helper"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestGetDebtAmounts_UserLendsToFriend(t *testing.T) {
	userID := uuid.New()
	friendID := uuid.New()

	transactions := []entity.DebtTransaction{
		{
			LenderProfileID:   userID,
			BorrowerProfileID: friendID,
			Type:              appconstant.Lend,
			Amount:            decimal.NewFromFloat(100),
		},
	}

	userOwes, friendOwes := helper.GetDebtAmounts(userID, friendID, transactions)

	assert.True(t, userOwes.IsZero())
	assert.True(t, friendOwes.Equal(decimal.NewFromFloat(100)))
}

func TestGetDebtAmounts_FriendLendsToUser(t *testing.T) {
	userID := uuid.New()
	friendID := uuid.New()

	transactions := []entity.DebtTransaction{
		{
			LenderProfileID:   friendID,
			BorrowerProfileID: userID,
			Type:              appconstant.Lend,
			Amount:            decimal.NewFromFloat(50),
		},
	}

	userOwes, friendOwes := helper.GetDebtAmounts(userID, friendID, transactions)

	assert.True(t, userOwes.Equal(decimal.NewFromFloat(50)))
	assert.True(t, friendOwes.IsZero())
}

func TestGetDebtAmounts_RepaymentReducesDebt(t *testing.T) {
	userID := uuid.New()
	friendID := uuid.New()

	transactions := []entity.DebtTransaction{
		{
			LenderProfileID:   userID,
			BorrowerProfileID: friendID,
			Type:              appconstant.Lend,
			Amount:            decimal.NewFromFloat(100),
		},
		{
			LenderProfileID:   userID,
			BorrowerProfileID: friendID,
			Type:              appconstant.Repay,
			Amount:            decimal.NewFromFloat(30),
		},
	}

	userOwes, friendOwes := helper.GetDebtAmounts(userID, friendID, transactions)

	assert.True(t, userOwes.IsZero())
	assert.True(t, friendOwes.Equal(decimal.NewFromFloat(70)))
}

func TestGetDebtAmounts_NoNegativeDebts(t *testing.T) {
	userID := uuid.New()
	friendID := uuid.New()

	transactions := []entity.DebtTransaction{
		{
			LenderProfileID:   userID,
			BorrowerProfileID: friendID,
			Type:              appconstant.Lend,
			Amount:            decimal.NewFromFloat(50),
		},
		{
			LenderProfileID:   userID,
			BorrowerProfileID: friendID,
			Type:              appconstant.Repay,
			Amount:            decimal.NewFromFloat(100),
		},
	}

	userOwes, friendOwes := helper.GetDebtAmounts(userID, friendID, transactions)

	assert.True(t, userOwes.IsZero())
	assert.True(t, friendOwes.IsZero())
}

func TestGetDebtAmounts_EmptyTransactions(t *testing.T) {
	userID := uuid.New()
	friendID := uuid.New()

	userOwes, friendOwes := helper.GetDebtAmounts(userID, friendID, []entity.DebtTransaction{})

	assert.True(t, userOwes.IsZero())
	assert.True(t, friendOwes.IsZero())
}
