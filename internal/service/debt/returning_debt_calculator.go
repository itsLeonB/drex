package debt

import (
	"github.com/itsLeonB/drex/internal/appconstant"
	"github.com/itsLeonB/drex/internal/dto"
	"github.com/itsLeonB/drex/internal/entity"
)

type returningDebtCalculator struct {
	action appconstant.DebtTransactionAction
}

func newReturningDebtCalculator() DebtCalculator {
	return &returningDebtCalculator{
		action: appconstant.ReturnAction,
	}
}

func (dc *returningDebtCalculator) GetAction() appconstant.DebtTransactionAction {
	return dc.action
}

func (dc *returningDebtCalculator) MapRequestToEntity(request dto.NewDebtTransactionRequest) entity.DebtTransaction {
	return entity.DebtTransaction{
		LenderProfileID:   request.FriendProfileID,
		BorrowerProfileID: request.UserProfileID,
		Action:            dc.action,
		Type:              appconstant.Repay,
		Amount:            request.Amount,
		TransferMethodID:  request.TransferMethodID,
		Description:       request.Description,
	}
}

func (dc *returningDebtCalculator) MapEntityToResponse(debtTransaction entity.DebtTransaction) dto.DebtTransactionResponse {
	return dto.DebtTransactionResponse{
		ID:             debtTransaction.ID,
		ProfileID:      debtTransaction.LenderProfileID,
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
