package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/itsLeonB/drex/internal/appconstant"
	"github.com/shopspring/decimal"
)

type NewDebtTransactionRequest struct {
	UserProfileID    uuid.UUID
	FriendProfileID  uuid.UUID                         `validate:"required"`
	Action           appconstant.DebtTransactionAction `validate:"oneof=LEND BORROW RECEIVE RETURN"`
	Amount           decimal.Decimal                   `validate:"required"`
	TransferMethodID uuid.UUID                         `validate:"required"`
	Description      string
}

type DebtTransactionResponse struct {
	ID             uuid.UUID
	ProfileID      uuid.UUID
	Type           appconstant.DebtTransactionType
	Action         appconstant.DebtTransactionAction
	Amount         decimal.Decimal
	TransferMethod string
	Description    string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      time.Time
}
