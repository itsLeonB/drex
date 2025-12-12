package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/itsLeonB/drex/internal/appconstant"
	"github.com/itsLeonB/drex/internal/dto"
	"github.com/itsLeonB/drex/internal/mapper"
	"github.com/itsLeonB/drex/internal/repository"
	"github.com/itsLeonB/drex/internal/service/debt"
	"github.com/itsLeonB/ezutil/v2"
	"github.com/itsLeonB/ungerr"
	"github.com/rotisserie/eris"
)

type debtTransactionServiceImpl struct {
	debtCalculatorStrategy    map[appconstant.DebtTransactionAction]debt.DebtCalculator
	debtTransactionRepository repository.DebtTransactionRepository
	transferMethodService     TransferMethodService
}

func NewDebtTransactionService(
	debtTransactionRepository repository.DebtTransactionRepository,
	transferMethodService TransferMethodService,
) DebtTransactionService {
	return &debtTransactionServiceImpl{
		debt.NewDebtCalculatorStrategies(),
		debtTransactionRepository,
		transferMethodService,
	}
}

func (ds *debtTransactionServiceImpl) RecordNew(ctx context.Context, request dto.NewDebtTransactionRequest) (dto.DebtTransactionResponse, error) {
	if !request.Amount.IsPositive() {
		return dto.DebtTransactionResponse{}, ungerr.ValidationError("amount must be greater than 0")
	}

	transferMethod, err := ds.transferMethodService.GetByID(ctx, request.TransferMethodID)
	if err != nil {
		return dto.DebtTransactionResponse{}, err
	}

	calculator, err := ds.selectCalculator(request.Action)
	if err != nil {
		return dto.DebtTransactionResponse{}, err
	}

	insertedDebt, err := ds.debtTransactionRepository.Insert(ctx, calculator.MapRequestToEntity(request))
	if err != nil {
		return dto.DebtTransactionResponse{}, err
	}

	insertedDebt.TransferMethod = transferMethod
	return calculator.MapEntityToResponse(insertedDebt), nil
}

func (ds *debtTransactionServiceImpl) GetAll(ctx context.Context, profileID uuid.UUID) ([]dto.DebtTransactionResponse, error) {
	transactions, err := ds.debtTransactionRepository.FindAllByUserProfileID(ctx, profileID)
	if err != nil {
		return nil, err
	}

	return ezutil.MapSlice(transactions, mapper.GetDebtTransactionSimpleMapper(profileID)), nil
}

func (ds *debtTransactionServiceImpl) ProcessConfirmedGroupExpense(ctx context.Context, groupExpense dto.GroupExpenseData) error {
	transferMethod, err := ds.transferMethodService.GetByName(ctx, appconstant.GroupExpenseTransferMethod)
	if err != nil {
		return err
	}

	debtTransactions := mapper.GroupExpenseToDebtTransactions(groupExpense, transferMethod.ID)

	_, err = ds.debtTransactionRepository.InsertMany(ctx, debtTransactions)

	return err
}

func (ds *debtTransactionServiceImpl) FindAllByProfileIDs(ctx context.Context, userProfileID, friendProfileID uuid.UUID) ([]dto.DebtTransactionResponse, error) {
	transactions, err := ds.debtTransactionRepository.FindAllByProfileIDs(ctx, userProfileID, friendProfileID)
	if err != nil {
		return nil, err
	}

	return ezutil.MapSlice(transactions, mapper.GetDebtTransactionSimpleMapper(userProfileID)), nil
}

func (ds *debtTransactionServiceImpl) selectCalculator(action appconstant.DebtTransactionAction) (debt.DebtCalculator, error) {
	calculator, ok := ds.debtCalculatorStrategy[action]
	if !ok {
		return nil, eris.Errorf("unsupported debt calculator action: %s", action)
	}

	return calculator, nil
}
