package mapper

import (
	"github.com/itsLeonB/drex-protos/gen/go/debt/v1"
	"github.com/itsLeonB/drex/internal/appconstant"
	"github.com/itsLeonB/drex/internal/dto"
	"github.com/itsLeonB/ezutil/v2"
	"github.com/itsLeonB/gerpc"
	"golang.org/x/text/currency"
)

func FromProtoTransactionAction(ta debt.TransactionAction) appconstant.DebtTransactionAction {
	switch ta {
	case debt.TransactionAction_TRANSACTION_ACTION_BORROW:
		return appconstant.BorrowAction
	case debt.TransactionAction_TRANSACTION_ACTION_LEND:
		return appconstant.LendAction
	case debt.TransactionAction_TRANSACTION_ACTION_RECEIVE:
		return appconstant.ReceiveAction
	case debt.TransactionAction_TRANSACTION_ACTION_RETURN:
		return appconstant.ReturnAction
	default:
		return ""
	}
}

func ToProtoTransactionAction(ta appconstant.DebtTransactionAction) debt.TransactionAction {
	switch ta {
	case appconstant.BorrowAction:
		return debt.TransactionAction_TRANSACTION_ACTION_BORROW
	case appconstant.LendAction:
		return debt.TransactionAction_TRANSACTION_ACTION_LEND
	case appconstant.ReceiveAction:
		return debt.TransactionAction_TRANSACTION_ACTION_RECEIVE
	case appconstant.ReturnAction:
		return debt.TransactionAction_TRANSACTION_ACTION_RETURN
	default:
		return debt.TransactionAction_TRANSACTION_ACTION_UNSPECIFIED
	}
}

func ToProtoTransactionType(tt appconstant.DebtTransactionType) debt.TransactionType {
	switch tt {
	case appconstant.Lend:
		return debt.TransactionType_TRANSACTION_TYPE_LEND
	case appconstant.Repay:
		return debt.TransactionType_TRANSACTION_TYPE_REPAY
	default:
		return debt.TransactionType_TRANSACTION_TYPE_UNSPECIFIED
	}
}

func ToTransactionProto(debtTrx dto.DebtTransactionResponse) *debt.TransactionResponse {
	return &debt.TransactionResponse{
		Id:             debtTrx.ID.String(),
		ProfileId:      debtTrx.ProfileID.String(),
		Type:           ToProtoTransactionType(debtTrx.Type),
		Action:         ToProtoTransactionAction(debtTrx.Action),
		Amount:         ezutil.DecimalToMoneyRounded(debtTrx.Amount, currency.IDR.String()),
		TransferMethod: debtTrx.TransferMethod,
		Description:    debtTrx.Description,
		CreatedAt:      gerpc.NullableTimeToProto(debtTrx.CreatedAt),
		UpdatedAt:      gerpc.NullableTimeToProto(debtTrx.UpdatedAt),
		DeletedAt:      gerpc.NullableTimeToProto(debtTrx.DeletedAt),
	}
}
