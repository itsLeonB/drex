package debt

import (
	"fmt"

	"github.com/itsLeonB/drex/internal/appconstant"
	"github.com/itsLeonB/drex/internal/dto"
	"github.com/itsLeonB/drex/internal/entity"
)

type AnonymousDebtCalculator interface {
	GetAction() appconstant.DebtTransactionAction
	MapRequestToEntity(request dto.NewDebtTransactionRequest) entity.DebtTransaction
	MapEntityToResponse(debtTransaction entity.DebtTransaction) dto.DebtTransactionResponse
	Validate(newTransaction entity.DebtTransaction, allTransactions []entity.DebtTransaction) error
}

var initFuncs = []func() AnonymousDebtCalculator{
	newBorrowingAnonDebtCalculator,
	newLendingAnonDebtCalculator,
	newReceivingAnonDebtCalculator,
	newReturningAnonDebtCalculator,
}

func NewAnonymousDebtCalculatorStrategies() map[appconstant.DebtTransactionAction]AnonymousDebtCalculator {
	strategyMap := make(map[appconstant.DebtTransactionAction]AnonymousDebtCalculator)

	for _, initFunc := range initFuncs {
		if initFunc == nil {
			panic("initFunc is nil")
		}

		calculator := initFunc()
		if calculator == nil {
			panic("calculator is nil")
		}

		action := calculator.GetAction()
		if _, exists := strategyMap[action]; exists {
			panic(fmt.Sprintf("duplicate calculator for action: %s", action))
		}

		strategyMap[calculator.GetAction()] = calculator
	}

	return strategyMap
}
