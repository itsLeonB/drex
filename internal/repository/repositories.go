package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/itsLeonB/drex/internal/entity"
	"github.com/itsLeonB/go-crud"
)

type DebtTransactionRepository interface {
	crud.CRUDRepository[entity.DebtTransaction]
	FindAllByProfileID(ctx context.Context, userProfileID, friendProfileID uuid.UUID) ([]entity.DebtTransaction, error)
	FindAllByUserProfileID(ctx context.Context, userProfileID uuid.UUID) ([]entity.DebtTransaction, error)
}

type TransferMethodRepository interface {
	crud.CRUDRepository[entity.TransferMethod]
}
