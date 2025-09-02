package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/itsLeonB/drex/internal/appconstant"
	"github.com/itsLeonB/drex/internal/dto"
	"github.com/itsLeonB/drex/internal/entity"
	"github.com/itsLeonB/drex/internal/mapper"
	"github.com/itsLeonB/drex/internal/repository"
	"github.com/itsLeonB/drex/internal/service/debt"
	"github.com/itsLeonB/ezutil/v2"
	crud "github.com/itsLeonB/go-crud"
	"github.com/itsLeonB/ungerr"
	"github.com/rotisserie/eris"
	"github.com/shopspring/decimal"
)

type debtTransactionServiceImpl struct {
	transactor                      crud.Transactor
	anonymousDebtCalculatorStrategy map[appconstant.DebtTransactionAction]debt.AnonymousDebtCalculator
	debtTransactionRepository       repository.DebtTransactionRepository
	transferMethodService           TransferMethodService
}

func NewDebtTransactionService(
	transactor crud.Transactor,
	debtTransactionRepository repository.DebtTransactionRepository,
	transferMethodService TransferMethodService,
) DebtTransactionService {
	return &debtTransactionServiceImpl{
		transactor,
		debt.NewAnonymousDebtCalculatorStrategies(),
		debtTransactionRepository,
		transferMethodService,
	}
}

func (ds *debtTransactionServiceImpl) RecordNew(ctx context.Context, request dto.NewDebtTransactionRequest) (dto.DebtTransactionResponse, error) {
	var response dto.DebtTransactionResponse

	transferMethod, err := ds.transferMethodService.GetByID(ctx, request.TransferMethodID)
	if err != nil {
		return dto.DebtTransactionResponse{}, err
	}

	if err = ds.Validate(ctx, request); err != nil {
		return dto.DebtTransactionResponse{}, err
	}

	err = ds.transactor.WithinTransaction(ctx, func(ctx context.Context) error {
		debtTransactions, err := ds.debtTransactionRepository.FindAllByProfileID(ctx, request.UserProfileID, request.FriendProfileID)
		if err != nil {
			return err
		}

		calculator, err := ds.selectAnonCalculator(request.Action)
		if err != nil {
			return err
		}

		newDebt := calculator.MapRequestToEntity(request)

		if err = calculator.Validate(newDebt, debtTransactions); err != nil {
			return err
		}

		insertedDebt, err := ds.debtTransactionRepository.Insert(ctx, newDebt)
		if err != nil {
			return err
		}

		insertedDebt.TransferMethod = transferMethod

		response = calculator.MapEntityToResponse(insertedDebt)

		return nil
	})

	if err != nil {
		return dto.DebtTransactionResponse{}, err
	}

	return response, nil
}

func (ds *debtTransactionServiceImpl) GetAll(ctx context.Context, profileID uuid.UUID) ([]dto.DebtTransactionResponse, error) {
	transactions, err := ds.debtTransactionRepository.FindAllByUserProfileID(ctx, profileID)
	if err != nil {
		return nil, err
	}

	mapFunc := func(transaction entity.DebtTransaction) dto.DebtTransactionResponse {
		return mapper.DebtTransactionToResponse(profileID, transaction)
	}

	return ezutil.MapSlice(transactions, mapFunc), nil
}

func (ds *debtTransactionServiceImpl) ProcessConfirmedGroupExpense(ctx context.Context, groupExpense dto.GroupExpenseData) error {
	transferMethod, err := ds.transferMethodService.GetByName(ctx, appconstant.GroupExpenseTransferMethod)
	if err != nil {
		return err
	}

	debtTransactions := mapper.GroupExpenseToDebtTransactions(groupExpense, transferMethod.ID)

	_, err = ds.debtTransactionRepository.BatchInsert(ctx, debtTransactions)

	return err
}

func (ds *debtTransactionServiceImpl) Validate(ctx context.Context, req dto.NewDebtTransactionRequest) error {
	if req.Amount.Compare(decimal.Zero) < 1 {
		return ungerr.ValidationError("amount must be greater than 0")
	}

	return nil
}

func (ds *debtTransactionServiceImpl) selectAnonCalculator(action appconstant.DebtTransactionAction) (debt.AnonymousDebtCalculator, error) {
	calculator, ok := ds.anonymousDebtCalculatorStrategy[action]
	if !ok {
		return nil, eris.Errorf("unsupported anonymous debt calculator action: %s", action)
	}

	return calculator, nil
}
