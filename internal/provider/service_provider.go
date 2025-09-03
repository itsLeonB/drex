package provider

import (
	"github.com/itsLeonB/drex/internal/service"
)

type Services struct {
	TransferMethod  service.TransferMethodService
	DebtTransaction service.DebtTransactionService
}

func ProvideServices(repositories *Repositories) *Services {
	if repositories == nil {
		panic("repositories cannot be nil")
	}

	transferMethodService := service.NewTransferMethodService(repositories.TransferMethod)

	debtTransactionService := service.NewDebtTransactionService(
		repositories.Transactor,
		repositories.DebtTransaction,
		transferMethodService,
	)

	return &Services{
		TransferMethod:  transferMethodService,
		DebtTransaction: debtTransactionService,
	}
}
