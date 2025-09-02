package debt

import (
	"log"

	"github.com/itsLeonB/drex/internal/appconstant"
	"github.com/itsLeonB/drex/internal/dto"
	"github.com/itsLeonB/drex/internal/entity"
)

const namespace = "[AnonymousDebtCalculator]"

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
			log.Fatalf("%s initFunc is nil", namespace)
		}

		calculator := initFunc()
		if calculator == nil {
			log.Fatalf("%s calculator is nil", namespace)
		}

		action := calculator.GetAction()
		if _, exists := strategyMap[action]; exists {
			log.Fatalf("%s duplicate calculator for action: %v", namespace, action)
		}

		strategyMap[calculator.GetAction()] = calculator
	}

	return strategyMap
}
