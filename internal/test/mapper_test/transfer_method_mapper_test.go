package mapper_test

import (
	"database/sql"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/itsLeonB/drex/internal/entity"
	"github.com/itsLeonB/drex/internal/mapper"
	"github.com/itsLeonB/go-crud"
	"github.com/stretchr/testify/assert"
)

func TestTransferMethodToResponse(t *testing.T) {
	id := uuid.New()
	now := time.Now()

	transferMethod := entity.TransferMethod{
		BaseEntity: crud.BaseEntity{
			ID:        id,
			CreatedAt: now,
			UpdatedAt: now,
			DeletedAt: sql.NullTime{},
		},
		Name:    "bank_transfer",
		Display: "Bank Transfer",
	}

	response := mapper.TransferMethodToResponse(transferMethod)

	assert.Equal(t, id, response.ID)
	assert.Equal(t, "bank_transfer", response.Name)
	assert.Equal(t, "Bank Transfer", response.Display)
	assert.Equal(t, now, response.CreatedAt)
	assert.Equal(t, now, response.UpdatedAt)
	assert.True(t, response.DeletedAt.IsZero())
}
