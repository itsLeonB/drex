package debt

import (
	"github.com/itsLeonB/drex/internal/appconstant"
	"github.com/itsLeonB/drex/internal/dto"
	"github.com/itsLeonB/drex/internal/entity"
)

type borrowingDebtCalculator struct {
	action appconstant.DebtTransactionAction
}

func newBorrowingDebtCalculator() DebtCalculator {
	return &borrowingDebtCalculator{
		action: appconstant.BorrowAction,
	}
}

func (dc *borrowingDebtCalculator) GetAction() appconstant.DebtTransactionAction {
	return dc.action
}

func (dc *borrowingDebtCalculator) MapRequestToEntity(request dto.NewDebtTransactionRequest) entity.DebtTransaction {
	return entity.DebtTransaction{
		LenderProfileID:   request.FriendProfileID,
		BorrowerProfileID: request.UserProfileID,
		Type:              appconstant.Lend,
		Action:            dc.action,
		Amount:            request.Amount,
		TransferMethodID:  request.TransferMethodID,
		Description:       request.Description,
	}
}

func (dc *borrowingDebtCalculator) MapEntityToResponse(debtTransaction entity.DebtTransaction) dto.DebtTransactionResponse {
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
