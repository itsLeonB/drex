package debt

import (
	"fmt"

	"github.com/itsLeonB/drex/internal/appconstant"
	"github.com/itsLeonB/drex/internal/dto"
	"github.com/itsLeonB/drex/internal/entity"
)

type DebtCalculator interface {
	GetAction() appconstant.DebtTransactionAction
	MapRequestToEntity(request dto.NewDebtTransactionRequest) entity.DebtTransaction
	MapEntityToResponse(debtTransaction entity.DebtTransaction) dto.DebtTransactionResponse
}

var initFuncs = []func() DebtCalculator{
	newBorrowingDebtCalculator,
	newLendingDebtCalculator,
	newReceivingDebtCalculator,
	newReturningDebtCalculator,
}

func NewDebtCalculatorStrategies() map[appconstant.DebtTransactionAction]DebtCalculator {
	strategyMap := make(map[appconstant.DebtTransactionAction]DebtCalculator)

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
