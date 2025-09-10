package dto_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/itsLeonB/drex/internal/dto"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestGroupExpenseData_Fields(t *testing.T) {
	id := uuid.New()
	payerID := uuid.New()
	creatorID := uuid.New()
	participantID := uuid.New()

	participants := []dto.ExpenseParticipantData{
		{
			ProfileID:   participantID,
			ShareAmount: decimal.NewFromFloat(50.00),
		},
	}

	expense := dto.GroupExpenseData{
		ID:               id,
		PayerProfileID:   payerID,
		CreatorProfileID: creatorID,
		Description:      "Dinner expense",
		Participants:     participants,
	}

	assert.Equal(t, id, expense.ID)
	assert.Equal(t, payerID, expense.PayerProfileID)
	assert.Equal(t, creatorID, expense.CreatorProfileID)
	assert.Equal(t, "Dinner expense", expense.Description)
	assert.Len(t, expense.Participants, 1)
	assert.Equal(t, participantID, expense.Participants[0].ProfileID)
	assert.Equal(t, decimal.NewFromFloat(50.00), expense.Participants[0].ShareAmount)
}

func TestExpenseParticipantData_Fields(t *testing.T) {
	profileID := uuid.New()
	shareAmount := decimal.NewFromFloat(25.50)

	participant := dto.ExpenseParticipantData{
		ProfileID:   profileID,
		ShareAmount: shareAmount,
	}

	assert.Equal(t, profileID, participant.ProfileID)
	assert.Equal(t, shareAmount, participant.ShareAmount)
}
