package mapper

import (
	"github.com/itsLeonB/drex-protos/gen/go/debt/v1"
	"github.com/itsLeonB/drex/internal/appconstant"
	"github.com/itsLeonB/drex/internal/dto"
	"github.com/itsLeonB/ezutil/v2"
	"github.com/itsLeonB/gerpc"
	"github.com/rotisserie/eris"
	"golang.org/x/text/currency"
)

func FromProtoTransactionAction(ta debt.TransactionAction) (appconstant.DebtTransactionAction, error) {
	switch ta {
	case debt.TransactionAction_TRANSACTION_ACTION_BORROW:
		return appconstant.BorrowAction, nil
	case debt.TransactionAction_TRANSACTION_ACTION_LEND:
		return appconstant.LendAction, nil
	case debt.TransactionAction_TRANSACTION_ACTION_RECEIVE:
		return appconstant.ReceiveAction, nil
	case debt.TransactionAction_TRANSACTION_ACTION_RETURN:
		return appconstant.ReturnAction, nil
	default:
		return "", eris.Errorf("undefined TransactionAction enum: %s", ta)
	}
}

func ToProtoTransactionAction(ta appconstant.DebtTransactionAction) (debt.TransactionAction, error) {
	switch ta {
	case appconstant.BorrowAction:
		return debt.TransactionAction_TRANSACTION_ACTION_BORROW, nil
	case appconstant.LendAction:
		return debt.TransactionAction_TRANSACTION_ACTION_LEND, nil
	case appconstant.ReceiveAction:
		return debt.TransactionAction_TRANSACTION_ACTION_RECEIVE, nil
	case appconstant.ReturnAction:
		return debt.TransactionAction_TRANSACTION_ACTION_RETURN, nil
	default:
		return debt.TransactionAction_TRANSACTION_ACTION_UNSPECIFIED, eris.Errorf("undefined TransactionAction constant: %s", ta)
	}
}

func ToProtoTransactionType(tt appconstant.DebtTransactionType) (debt.TransactionType, error) {
	switch tt {
	case appconstant.Lend:
		return debt.TransactionType_TRANSACTION_TYPE_LEND, nil
	case appconstant.Repay:
		return debt.TransactionType_TRANSACTION_TYPE_REPAY, nil
	default:
		return debt.TransactionType_TRANSACTION_TYPE_UNSPECIFIED, eris.Errorf("undefined TransactionType constant: %s", tt)
	}
}

func ToTransactionProto(debtTrx dto.DebtTransactionResponse) (*debt.TransactionResponse, error) {
	trxType, err := ToProtoTransactionType(debtTrx.Type)
	if err != nil {
		return nil, err
	}

	trxAction, err := ToProtoTransactionAction(debtTrx.Action)
	if err != nil {
		return nil, err
	}

	return &debt.TransactionResponse{
		Id:             debtTrx.ID.String(),
		ProfileId:      debtTrx.ProfileID.String(),
		Type:           trxType,
		Action:         trxAction,
		Amount:         ezutil.DecimalToMoneyRounded(debtTrx.Amount, currency.IDR.String()),
		TransferMethod: debtTrx.TransferMethod,
		Description:    debtTrx.Description,
		CreatedAt:      gerpc.NullableTimeToProto(debtTrx.CreatedAt),
		UpdatedAt:      gerpc.NullableTimeToProto(debtTrx.UpdatedAt),
		DeletedAt:      gerpc.NullableTimeToProto(debtTrx.DeletedAt),
	}, nil
}
