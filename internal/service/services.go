package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/itsLeonB/drex/internal/dto"
	"github.com/itsLeonB/drex/internal/entity"
)

type DebtTransactionService interface {
	RecordNew(ctx context.Context, request dto.NewDebtTransactionRequest) (dto.DebtTransactionResponse, error)
	GetAll(ctx context.Context, userProfileID uuid.UUID) ([]dto.DebtTransactionResponse, error)
	ProcessConfirmedGroupExpense(ctx context.Context, groupExpense dto.GroupExpenseData) error
	FindAllByProfileIDs(ctx context.Context, userProfileID, friendProfileID uuid.UUID) ([]dto.DebtTransactionResponse, error)
}

type TransferMethodService interface {
	GetAll(ctx context.Context) ([]dto.TransferMethodResponse, error)
	GetByID(ctx context.Context, id uuid.UUID) (entity.TransferMethod, error)
	GetByName(ctx context.Context, name string) (entity.TransferMethod, error)
}
