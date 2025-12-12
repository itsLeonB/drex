package debt

import (
	"github.com/itsLeonB/drex/internal/appconstant"
	"github.com/itsLeonB/drex/internal/dto"
	"github.com/itsLeonB/drex/internal/entity"
)

type lendingDebtCalculator struct {
	action appconstant.DebtTransactionAction
}

func newLendingDebtCalculator() DebtCalculator {
	return &lendingDebtCalculator{
		action: appconstant.LendAction,
	}
}

func (dc *lendingDebtCalculator) GetAction() appconstant.DebtTransactionAction {
	return dc.action
}

func (dc *lendingDebtCalculator) MapRequestToEntity(request dto.NewDebtTransactionRequest) entity.DebtTransaction {
	return entity.DebtTransaction{
		LenderProfileID:   request.UserProfileID,
		BorrowerProfileID: request.FriendProfileID,
		Type:              appconstant.Lend,
		Action:            dc.action,
		Amount:            request.Amount,
		TransferMethodID:  request.TransferMethodID,
		Description:       request.Description,
	}
}

func (dc *lendingDebtCalculator) MapEntityToResponse(debtTransaction entity.DebtTransaction) dto.DebtTransactionResponse {
	return dto.DebtTransactionResponse{
		ID:             debtTransaction.ID,
		ProfileID:      debtTransaction.BorrowerProfileID,
		Type:           debtTransaction.Type,
		Action:         debtTransaction.Action,
		Amount:         debtTransaction.Amount,
		TransferMethod: debtTransaction.TransferMethod.Display,
		Description:    debtTransaction.Description,
		CreatedAt:      debtTransaction.CreatedAt,
		UpdatedAt:      debtTransaction.UpdatedAt,
		DeletedAt:      debtTransaction.DeletedAt.Time,
	}
}
