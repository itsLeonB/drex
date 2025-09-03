package appconstant

type DebtTransactionType string
type DebtTransactionAction string

const (
	Lend  DebtTransactionType = "LEND"
	Repay DebtTransactionType = "REPAY"

	LendAction    DebtTransactionAction = "LEND"
	BorrowAction  DebtTransactionAction = "BORROW"
	ReceiveAction DebtTransactionAction = "RECEIVE"
	ReturnAction  DebtTransactionAction = "RETURN"

	GroupExpenseTransferMethod = "GROUP_EXPENSE"
)
