package dto

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type GroupExpenseData struct {
	ID               uuid.UUID
	PayerProfileID   uuid.UUID
	CreatorProfileID uuid.UUID
	Description      string
	Participants     []ExpenseParticipantData
}

type ExpenseParticipantData struct {
	ProfileID   uuid.UUID
	ShareAmount decimal.Decimal
}
