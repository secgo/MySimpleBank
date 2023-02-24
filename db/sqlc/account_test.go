package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/secgo/soluation_bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomeAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)
	require.Equal(t, arg.Owner, account.Owner)
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomeAccount(t)
}

func TestGetAccount(t *testing.T) {
	account := createRandomeAccount(t)
	GetAccount, err := testQueries.GetAccount(context.Background(), account.ID)
	require.NoError(t, err)
	require.NotEmpty(t, GetAccount)
	require.Equal(t, account.ID, GetAccount.ID)
	require.Equal(t, account.Balance, GetAccount.Balance)
	require.Equal(t, account.Currency, GetAccount.Currency)
	require.Equal(t, account.Owner, GetAccount.Owner)

	require.WithinDuration(t, account.CreatedAt, GetAccount.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	account := createRandomeAccount(t)
	arg := UpdateAccountParams{
		ID:      account.ID,
		Balance: util.RandomMoney(),
	}
	updateAccount, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updateAccount)
	require.Equal(t, arg.ID, updateAccount.ID)
	require.Equal(t, arg.Balance, updateAccount.Balance)
	require.Equal(t, account.Currency, updateAccount.Currency)
	require.Equal(t, account.Owner, updateAccount.Owner)
	require.WithinDuration(t, account.CreatedAt, updateAccount.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	account := createRandomeAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account.ID)
	require.NoError(t, err)

	getaccount, err := testQueries.GetAccount(context.Background(), account.ID)
	require.Error(t, err, sql.ErrNoRows)
	require.Empty(t, getaccount)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomeAccount(t)
	}
	arg := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}
	listsaccounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, listsaccounts, 5)
	for _, account := range listsaccounts {
		require.NotEmpty(t, account)
	}
}
