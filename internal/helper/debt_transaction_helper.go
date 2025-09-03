package helper

import (
	"github.com/google/uuid"
	"github.com/itsLeonB/drex/internal/appconstant"
	"github.com/itsLeonB/drex/internal/entity"
	"github.com/shopspring/decimal"
)

func GetAmountSumsFromDebtTransactions(
	userProfileID, friendProfileID uuid.UUID,
	transactions []entity.DebtTransaction,
) (decimal.Decimal, decimal.Decimal) {
	userAmount, friendAmount := decimal.Zero, decimal.Zero

	for _, transaction := range transactions {
		if transaction.LenderProfileID == userProfileID && transaction.BorrowerProfileID == friendProfileID {
			switch transaction.Type {
			case appconstant.Lend:
				friendAmount = friendAmount.Add(transaction.Amount)
				userAmount = userAmount.Sub(transaction.Amount)
			case appconstant.Repay:
				friendAmount = friendAmount.Sub(transaction.Amount)
				userAmount = userAmount.Add(transaction.Amount)
			}
		} else if transaction.LenderProfileID == friendProfileID && transaction.BorrowerProfileID == userProfileID {
			switch transaction.Type {
			case appconstant.Lend:
				userAmount = userAmount.Add(transaction.Amount)
				friendAmount = friendAmount.Sub(transaction.Amount)
			case appconstant.Repay:
				userAmount = userAmount.Sub(transaction.Amount)
				friendAmount = friendAmount.Add(transaction.Amount)
			}
		}
	}

	return userAmount, friendAmount
}
