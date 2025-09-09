package helper_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/itsLeonB/drex/internal/appconstant"
	"github.com/itsLeonB/drex/internal/entity"
	"github.com/itsLeonB/drex/internal/helper"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestGetDebtAmounts(t *testing.T) {
	userID := uuid.New()
	friendID := uuid.New()
	otherID := uuid.New()

	tests := []struct {
		name           string
		transactions   []entity.DebtTransaction
		wantUserOwes   decimal.Decimal // How much user owes to friend
		wantFriendOwes decimal.Decimal // How much friend owes to user
	}{
		{
			name:           "Empty transactions",
			transactions:   []entity.DebtTransaction{},
			wantUserOwes:   decimal.Zero,
			wantFriendOwes: decimal.Zero,
		},
		{
			name: "Unrelated transactions",
			transactions: []entity.DebtTransaction{
				{
					LenderProfileID:   otherID,
					BorrowerProfileID: otherID,
					Type:              appconstant.Lend,
					Amount:            decimal.NewFromInt(100),
				},
			},
			wantUserOwes:   decimal.Zero,
			wantFriendOwes: decimal.Zero,
		},
		{
			name: "User lends to friend and gets partial repayment",
			transactions: []entity.DebtTransaction{
				{
					LenderProfileID:   userID,
					BorrowerProfileID: friendID,
					Type:              appconstant.Lend,
					Amount:            decimal.NewFromInt(100),
				},
				{
					LenderProfileID:   userID,
					BorrowerProfileID: friendID,
					Type:              appconstant.Repay,
					Amount:            decimal.NewFromInt(50),
				},
			},
			wantUserOwes:   decimal.Zero,           // User doesn't owe anything
			wantFriendOwes: decimal.NewFromInt(50), // Friend still owes 50
		},
		{
			name: "Friend lends to user and gets partial repayment",
			transactions: []entity.DebtTransaction{
				{
					LenderProfileID:   friendID,
					BorrowerProfileID: userID,
					Type:              appconstant.Lend,
					Amount:            decimal.NewFromInt(70),
				},
				{
					LenderProfileID:   friendID,
					BorrowerProfileID: userID,
					Type:              appconstant.Repay,
					Amount:            decimal.NewFromInt(30),
				},
			},
			wantUserOwes:   decimal.NewFromInt(40), // User still owes 40
			wantFriendOwes: decimal.Zero,           // Friend doesn't owe anything
		},
		{
			name: "Cross lending - user owes more",
			transactions: []entity.DebtTransaction{
				{
					LenderProfileID:   userID,
					BorrowerProfileID: friendID,
					Type:              appconstant.Lend,
					Amount:            decimal.NewFromInt(20),
				},
				{
					LenderProfileID:   friendID,
					BorrowerProfileID: userID,
					Type:              appconstant.Lend,
					Amount:            decimal.NewFromInt(60),
				},
			},
			wantUserOwes:   decimal.NewFromInt(60), // User owes 60 (independent tracking)
			wantFriendOwes: decimal.NewFromInt(20), // Friend owes 20 (independent tracking)
		},
		{
			name: "Cross lending - friend owes more",
			transactions: []entity.DebtTransaction{
				{
					LenderProfileID:   userID,
					BorrowerProfileID: friendID,
					Type:              appconstant.Lend,
					Amount:            decimal.NewFromInt(60),
				},
				{
					LenderProfileID:   friendID,
					BorrowerProfileID: userID,
					Type:              appconstant.Lend,
					Amount:            decimal.NewFromInt(20),
				},
			},
			wantUserOwes:   decimal.NewFromInt(20), // User owes 20 (independent tracking)
			wantFriendOwes: decimal.NewFromInt(60), // Friend owes 60 (independent tracking)
		},
		{
			name: "Complex scenario with multiple transactions",
			transactions: []entity.DebtTransaction{
				{
					LenderProfileID:   userID,
					BorrowerProfileID: friendID,
					Type:              appconstant.Lend,
					Amount:            decimal.NewFromInt(100),
				},
				{
					LenderProfileID:   friendID,
					BorrowerProfileID: userID,
					Type:              appconstant.Lend,
					Amount:            decimal.NewFromInt(50),
				},
				{
					LenderProfileID:   userID,
					BorrowerProfileID: friendID,
					Type:              appconstant.Repay,
					Amount:            decimal.NewFromInt(30),
				},
			},
			wantUserOwes:   decimal.NewFromInt(50), // User owes 50 (independent tracking)
			wantFriendOwes: decimal.NewFromInt(70), // Friend owes 100-30=70 (independent tracking)
		},
		{
			name: "All repay transactions (unusual but possible)",
			transactions: []entity.DebtTransaction{
				{
					LenderProfileID:   userID,
					BorrowerProfileID: friendID,
					Type:              appconstant.Repay,
					Amount:            decimal.NewFromInt(25),
				},
				{
					LenderProfileID:   friendID,
					BorrowerProfileID: userID,
					Type:              appconstant.Repay,
					Amount:            decimal.NewFromInt(35),
				},
			},
			wantUserOwes:   decimal.Zero, // No previous lends, so no valid debts (protected by negative check)
			wantFriendOwes: decimal.Zero, // No previous lends, so no valid debts (protected by negative check)
		},
		{
			name: "Multiple lends and repays - user side",
			transactions: []entity.DebtTransaction{
				{
					LenderProfileID:   userID,
					BorrowerProfileID: friendID,
					Type:              appconstant.Lend,
					Amount:            decimal.NewFromInt(100),
				},
				{
					LenderProfileID:   userID,
					BorrowerProfileID: friendID,
					Type:              appconstant.Lend,
					Amount:            decimal.NewFromInt(50),
				},
				{
					LenderProfileID:   userID,
					BorrowerProfileID: friendID,
					Type:              appconstant.Repay,
					Amount:            decimal.NewFromInt(80),
				},
			},
			wantUserOwes:   decimal.Zero,           // User doesn't owe anything
			wantFriendOwes: decimal.NewFromInt(70), // Friend owes 100+50-80=70
		},
		{
			name: "Multiple lends and repays - friend side",
			transactions: []entity.DebtTransaction{
				{
					LenderProfileID:   friendID,
					BorrowerProfileID: userID,
					Type:              appconstant.Lend,
					Amount:            decimal.NewFromInt(80),
				},
				{
					LenderProfileID:   friendID,
					BorrowerProfileID: userID,
					Type:              appconstant.Lend,
					Amount:            decimal.NewFromInt(40),
				},
				{
					LenderProfileID:   friendID,
					BorrowerProfileID: userID,
					Type:              appconstant.Repay,
					Amount:            decimal.NewFromInt(90),
				},
			},
			wantUserOwes:   decimal.NewFromInt(30), // User owes 80+40-90=30
			wantFriendOwes: decimal.Zero,           // Friend doesn't owe anything
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUserOwes, gotFriendOwes := helper.GetDebtAmounts(userID, friendID, tt.transactions)

			assert.True(t, gotUserOwes.Equal(tt.wantUserOwes), "UserOwes mismatch: got %v want %v", gotUserOwes, tt.wantUserOwes)
			assert.True(t, gotFriendOwes.Equal(tt.wantFriendOwes), "FriendOwes mismatch: got %v want %v", gotFriendOwes, tt.wantFriendOwes)

			// Additional assertions to ensure debt amounts are never negative
			assert.True(t, gotUserOwes.GreaterThanOrEqual(decimal.Zero), "UserOwes should never be negative: got %v", gotUserOwes)
			assert.True(t, gotFriendOwes.GreaterThanOrEqual(decimal.Zero), "FriendOwes should never be negative: got %v", gotFriendOwes)
		})
	}
}
