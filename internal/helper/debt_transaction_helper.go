package helper

import (
	"github.com/google/uuid"
	"github.com/itsLeonB/drex/internal/appconstant"
	"github.com/itsLeonB/drex/internal/entity"
	"github.com/shopspring/decimal"
)

func GetDebtAmounts(userProfileID, friendProfileID uuid.UUID, transactions []entity.DebtTransaction) (userOwes, friendOwes decimal.Decimal) {
	userOwes, friendOwes = decimal.Zero, decimal.Zero

	for _, transaction := range transactions {
		if transaction.LenderProfileID == userProfileID && transaction.BorrowerProfileID == friendProfileID {
			// User lent money to friend
			switch transaction.Type {
			case appconstant.Lend:
				friendOwes = friendOwes.Add(transaction.Amount) // Friend owes more
			case appconstant.Repay:
				friendOwes = friendOwes.Sub(transaction.Amount) // Friend owes less
			}
		} else if transaction.LenderProfileID == friendProfileID && transaction.BorrowerProfileID == userProfileID {
			// Friend lent money to user
			switch transaction.Type {
			case appconstant.Lend:
				userOwes = userOwes.Add(transaction.Amount) // User owes more
			case appconstant.Repay:
				userOwes = userOwes.Sub(transaction.Amount) // User owes less
			}
		}
	}

	// Ensure no negative debts
	if userOwes.IsNegative() {
		userOwes = decimal.Zero
	}
	if friendOwes.IsNegative() {
		friendOwes = decimal.Zero
	}

	return userOwes, friendOwes
}
