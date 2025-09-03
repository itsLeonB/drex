package mapper

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/itsLeonB/drex/internal/appconstant"
	"github.com/itsLeonB/drex/internal/dto"
	"github.com/itsLeonB/drex/internal/entity"
)

func DebtTransactionToResponse(userProfileID uuid.UUID, transaction entity.DebtTransaction) dto.DebtTransactionResponse {
	var profileID uuid.UUID
	if userProfileID == transaction.BorrowerProfileID && userProfileID != transaction.LenderProfileID {
		profileID = transaction.LenderProfileID
	} else if userProfileID == transaction.LenderProfileID && userProfileID != transaction.BorrowerProfileID {
		profileID = transaction.BorrowerProfileID
	}

	return dto.DebtTransactionResponse{
		ID:             transaction.ID,
		ProfileID:      profileID,
		Type:           transaction.Type,
		Action:         transaction.Action,
		Amount:         transaction.Amount,
		TransferMethod: transaction.TransferMethod.Display,
		Description:    transaction.Description,
		CreatedAt:      transaction.CreatedAt,
		UpdatedAt:      transaction.UpdatedAt,
		DeletedAt:      transaction.DeletedAt.Time,
	}
}

func GroupExpenseToDebtTransactions(groupExpense dto.GroupExpenseData, transferMethodID uuid.UUID) []entity.DebtTransaction {
	action := appconstant.BorrowAction
	if groupExpense.PayerProfileID == groupExpense.CreatorProfileID {
		action = appconstant.LendAction
	}

	debtTransactions := make([]entity.DebtTransaction, 0, len(groupExpense.Participants))
	for _, participant := range groupExpense.Participants {
		if groupExpense.PayerProfileID == participant.ProfileID {
			continue
		}
		debtTransactions = append(debtTransactions, entity.DebtTransaction{
			LenderProfileID:   groupExpense.PayerProfileID,
			BorrowerProfileID: participant.ProfileID,
			Type:              appconstant.Lend,
			Action:            action,
			Amount:            participant.ShareAmount,
			TransferMethodID:  transferMethodID,
			Description:       fmt.Sprintf("Share for group expense %s: %s", groupExpense.ID, groupExpense.Description),
		})
	}

	return debtTransactions
}

func GetDebtTransactionSimpleMapper(userProfileID uuid.UUID) func(entity.DebtTransaction) dto.DebtTransactionResponse {
	return func(transaction entity.DebtTransaction) dto.DebtTransactionResponse {
		return DebtTransactionToResponse(userProfileID, transaction)
	}
}
