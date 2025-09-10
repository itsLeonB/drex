package mapper_test

import (
	"database/sql"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/itsLeonB/drex/internal/appconstant"
	"github.com/itsLeonB/drex/internal/dto"
	"github.com/itsLeonB/drex/internal/entity"
	"github.com/itsLeonB/drex/internal/mapper"
	"github.com/itsLeonB/go-crud"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestDebtTransactionToResponse_UserIsBorrower(t *testing.T) {
	userID := uuid.New()
	lenderID := uuid.New()
	id := uuid.New()
	now := time.Now()

	transaction := entity.DebtTransaction{
		BaseEntity: crud.BaseEntity{
			ID:        id,
			CreatedAt: now,
			UpdatedAt: now,
			DeletedAt: sql.NullTime{},
		},
		LenderProfileID:   lenderID,
		BorrowerProfileID: userID,
		Type:              appconstant.Lend,
		Action:            appconstant.BorrowAction,
		Amount:            decimal.NewFromFloat(100),
		Description:       "Test transaction",
		TransferMethod: entity.TransferMethod{
			Display: "Bank Transfer",
		},
	}

	response := mapper.DebtTransactionToResponse(userID, transaction)

	assert.Equal(t, id, response.ID)
	assert.Equal(t, lenderID, response.ProfileID)
	assert.Equal(t, appconstant.Lend, response.Type)
	assert.Equal(t, appconstant.BorrowAction, response.Action)
	assert.Equal(t, decimal.NewFromFloat(100), response.Amount)
	assert.Equal(t, "Bank Transfer", response.TransferMethod)
	assert.Equal(t, "Test transaction", response.Description)
}

func TestDebtTransactionToResponse_UserIsLender(t *testing.T) {
	userID := uuid.New()
	borrowerID := uuid.New()
	id := uuid.New()
	now := time.Now()

	transaction := entity.DebtTransaction{
		BaseEntity: crud.BaseEntity{
			ID:        id,
			CreatedAt: now,
			UpdatedAt: now,
			DeletedAt: sql.NullTime{},
		},
		LenderProfileID:   userID,
		BorrowerProfileID: borrowerID,
		Type:              appconstant.Lend,
		Action:            appconstant.LendAction,
		Amount:            decimal.NewFromFloat(50),
		Description:       "Test lend",
		TransferMethod: entity.TransferMethod{
			Display: "Cash",
		},
	}

	response := mapper.DebtTransactionToResponse(userID, transaction)

	assert.Equal(t, id, response.ID)
	assert.Equal(t, borrowerID, response.ProfileID)
	assert.Equal(t, appconstant.Lend, response.Type)
	assert.Equal(t, appconstant.LendAction, response.Action)
	assert.Equal(t, decimal.NewFromFloat(50), response.Amount)
	assert.Equal(t, "Cash", response.TransferMethod)
	assert.Equal(t, "Test lend", response.Description)
}

func TestGroupExpenseToDebtTransactions_PayerIsCreator(t *testing.T) {
	payerID := uuid.New()
	participant1ID := uuid.New()
	participant2ID := uuid.New()
	transferMethodID := uuid.New()
	expenseID := uuid.New()

	groupExpense := dto.GroupExpenseData{
		ID:               expenseID,
		PayerProfileID:   payerID,
		CreatorProfileID: payerID,
		Description:      "Dinner",
		Participants: []dto.ExpenseParticipantData{
			{ProfileID: payerID, ShareAmount: decimal.NewFromFloat(30)},
			{ProfileID: participant1ID, ShareAmount: decimal.NewFromFloat(25)},
			{ProfileID: participant2ID, ShareAmount: decimal.NewFromFloat(20)},
		},
	}

	transactions := mapper.GroupExpenseToDebtTransactions(groupExpense, transferMethodID)

	assert.Len(t, transactions, 2)
	
	for _, tx := range transactions {
		assert.Equal(t, payerID, tx.LenderProfileID)
		assert.Equal(t, appconstant.Lend, tx.Type)
		assert.Equal(t, appconstant.LendAction, tx.Action)
		assert.Equal(t, transferMethodID, tx.TransferMethodID)
		assert.Contains(t, tx.Description, expenseID.String())
		assert.Contains(t, tx.Description, "Dinner")
	}
}

func TestGroupExpenseToDebtTransactions_PayerIsNotCreator(t *testing.T) {
	payerID := uuid.New()
	creatorID := uuid.New()
	participant1ID := uuid.New()
	transferMethodID := uuid.New()
	expenseID := uuid.New()

	groupExpense := dto.GroupExpenseData{
		ID:               expenseID,
		PayerProfileID:   payerID,
		CreatorProfileID: creatorID,
		Description:      "Lunch",
		Participants: []dto.ExpenseParticipantData{
			{ProfileID: payerID, ShareAmount: decimal.NewFromFloat(15)},
			{ProfileID: participant1ID, ShareAmount: decimal.NewFromFloat(15)},
		},
	}

	transactions := mapper.GroupExpenseToDebtTransactions(groupExpense, transferMethodID)

	assert.Len(t, transactions, 1)
	assert.Equal(t, payerID, transactions[0].LenderProfileID)
	assert.Equal(t, participant1ID, transactions[0].BorrowerProfileID)
	assert.Equal(t, appconstant.BorrowAction, transactions[0].Action)
}

func TestGetDebtTransactionSimpleMapper(t *testing.T) {
	userID := uuid.New()
	lenderID := uuid.New()
	
	transaction := entity.DebtTransaction{
		LenderProfileID:   lenderID,
		BorrowerProfileID: userID,
		Type:              appconstant.Lend,
		Action:            appconstant.BorrowAction,
		Amount:            decimal.NewFromFloat(75),
		TransferMethod: entity.TransferMethod{
			Display: "Digital Wallet",
		},
	}

	mapperFunc := mapper.GetDebtTransactionSimpleMapper(userID)
	response := mapperFunc(transaction)

	assert.Equal(t, lenderID, response.ProfileID)
	assert.Equal(t, appconstant.Lend, response.Type)
	assert.Equal(t, appconstant.BorrowAction, response.Action)
	assert.Equal(t, decimal.NewFromFloat(75), response.Amount)
	assert.Equal(t, "Digital Wallet", response.TransferMethod)
}
