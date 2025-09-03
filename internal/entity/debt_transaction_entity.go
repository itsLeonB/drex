package entity

import (
	"github.com/google/uuid"
	"github.com/itsLeonB/drex/internal/appconstant"
	"github.com/itsLeonB/go-crud"
	"github.com/shopspring/decimal"
)

type DebtTransaction struct {
	crud.BaseEntity
	LenderProfileID   uuid.UUID
	BorrowerProfileID uuid.UUID
	Type              appconstant.DebtTransactionType
	Action            appconstant.DebtTransactionAction
	Amount            decimal.Decimal
	TransferMethodID  uuid.UUID
	Description       string
	TransferMethod    TransferMethod
}
