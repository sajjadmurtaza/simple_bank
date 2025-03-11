package db

import (
	"context"
	"testing"

	"github.com/sajjadmurtaza/simple_bank/util"
	"github.com/stretchr/testify/require"
)

//=============================================
// Create Random Account

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomString(6),
		Balance:  10,
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	return account
}

//=============================================
// Create Account Test

func TestCreateAccount(t *testing.T) {
	randomCrearedAccount := createRandomAccount(t)

	// delete account
	err := testQueries.DeleteAccount(context.Background(), randomCrearedAccount.ID)
	require.NoError(t, err)

	// try to fetch deleted account and expect error
	account, err := testQueries.GetAccount(context.Background(), randomCrearedAccount.ID)
	require.Error(t, err)
	require.Empty(t, account)
}

// =============================================
// Get Account Test
func TestGetAccount(t *testing.T) {
	randomCrearedAccount := createRandomAccount(t)

	account, err := testQueries.GetAccount(context.Background(), randomCrearedAccount.ID)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, randomCrearedAccount.ID, account.ID)
}
