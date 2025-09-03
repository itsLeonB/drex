package dto

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type GroupExpenseData struct {
	ID               uuid.UUID `validate:"required"`
	PayerProfileID   uuid.UUID `validate:"required"`
	CreatorProfileID uuid.UUID `validate:"required"`
	Description      string
	Participants     []ExpenseParticipantData `validate:"required,min=1,dive"`
}

type ExpenseParticipantData struct {
	ProfileID   uuid.UUID       `validate:"required"`
	ShareAmount decimal.Decimal `validate:"required"`
}
