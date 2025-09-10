package dto_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/itsLeonB/drex/internal/appconstant"
	"github.com/itsLeonB/drex/internal/dto"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestNewDebtTransactionRequest_Fields(t *testing.T) {
	userID := uuid.New()
	friendID := uuid.New()
	transferMethodID := uuid.New()
	amount := decimal.NewFromFloat(100.50)

	req := dto.NewDebtTransactionRequest{
		UserProfileID:    userID,
		FriendProfileID:  friendID,
		Action:           appconstant.LendAction,
		Amount:           amount,
		TransferMethodID: transferMethodID,
		Description:      "Test transaction",
	}

	assert.Equal(t, userID, req.UserProfileID)
	assert.Equal(t, friendID, req.FriendProfileID)
	assert.Equal(t, appconstant.LendAction, req.Action)
	assert.Equal(t, amount, req.Amount)
	assert.Equal(t, transferMethodID, req.TransferMethodID)
	assert.Equal(t, "Test transaction", req.Description)
}

func TestDebtTransactionResponse_Fields(t *testing.T) {
	id := uuid.New()
	profileID := uuid.New()
	amount := decimal.NewFromFloat(75.25)
	now := time.Now()

	resp := dto.DebtTransactionResponse{
		ID:             id,
		ProfileID:      profileID,
		Type:           appconstant.Lend,
		Action:         appconstant.BorrowAction,
		Amount:         amount,
		TransferMethod: "Bank Transfer",
		Description:    "Test response",
		CreatedAt:      now,
		UpdatedAt:      now,
		DeletedAt:      time.Time{},
	}

	assert.Equal(t, id, resp.ID)
	assert.Equal(t, profileID, resp.ProfileID)
	assert.Equal(t, appconstant.Lend, resp.Type)
	assert.Equal(t, appconstant.BorrowAction, resp.Action)
	assert.Equal(t, amount, resp.Amount)
	assert.Equal(t, "Bank Transfer", resp.TransferMethod)
	assert.Equal(t, "Test response", resp.Description)
	assert.Equal(t, now, resp.CreatedAt)
	assert.Equal(t, now, resp.UpdatedAt)
	assert.True(t, resp.DeletedAt.IsZero())
}
