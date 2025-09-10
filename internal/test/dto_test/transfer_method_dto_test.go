package dto_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/itsLeonB/drex/internal/dto"
	"github.com/stretchr/testify/assert"
)

func TestTransferMethodResponse_Fields(t *testing.T) {
	id := uuid.New()
	now := time.Now()

	resp := dto.TransferMethodResponse{
		ID:        id,
		Name:      "bank_transfer",
		Display:   "Bank Transfer",
		CreatedAt: now,
		UpdatedAt: now,
		DeletedAt: time.Time{},
	}

	assert.Equal(t, id, resp.ID)
	assert.Equal(t, "bank_transfer", resp.Name)
	assert.Equal(t, "Bank Transfer", resp.Display)
	assert.Equal(t, now, resp.CreatedAt)
	assert.Equal(t, now, resp.UpdatedAt)
	assert.True(t, resp.DeletedAt.IsZero())
}
