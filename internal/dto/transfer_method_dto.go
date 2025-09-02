package dto

import (
	"time"

	"github.com/google/uuid"
)

type TransferMethodResponse struct {
	ID        uuid.UUID
	Name      string
	Display   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
