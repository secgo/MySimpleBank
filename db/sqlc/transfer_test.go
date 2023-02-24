package db

import (
	"context"
	"testing"
	"time"

	"github.com/secgo/soluation_bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T, account1, account2 Account) Transfer {
	arg := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomMoney(),
	}
	newTransfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, newTransfer)
	require.Equal(t, arg.FromAccountID, newTransfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, newTransfer.ToAccountID)
	require.Equal(t, arg.Amount, newTransfer.Amount)
	require.NotZero(t, newTransfer.ID)
	require.NotZero(t, newTransfer.CreatedAt)
	return newTransfer
}

func TestCreateTransfer(t *testing.T) {
	account1 := createRandomeAccount(t)
	account2 := createRandomeAccount(t)
	createRandomTransfer(t, account1, account2)
}

func TestGetTransfer(t *testing.T) {
	account1 := createRandomeAccount(t)
	account2 := createRandomeAccount(t)
	tr := createRandomTransfer(t, account1, account2)
	getTr, err := testQueries.GetTransfer(context.Background(), tr.ID)
	require.NoError(t, err)
	require.NotEmpty(t, getTr)
	require.Equal(t, tr.Amount, getTr.Amount)
	require.Equal(t, tr.FromAccountID, getTr.FromAccountID)
	require.Equal(t, tr.ToAccountID, getTr.ToAccountID)
	require.Equal(t, tr.ID, getTr.ID)
	require.WithinDuration(t, tr.CreatedAt, getTr.CreatedAt, time.Second)
}

func TestListTransfer(t *testing.T) {
	account1 := createRandomeAccount(t)
	account2 := createRandomeAccount(t)
	for i := 0; i < 10; i++ {

		createRandomTransfer(t, account1, account2)
	}
	arg := listTransfersParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Limit:         5,
		Offset:        5,
	}
	listTr, err := testQueries.listTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, listTr, 5)

	for _, transfer := range listTr {
		require.NotEmpty(t, transfer)
		require.True(t, transfer.FromAccountID == account1.ID || transfer.ToAccountID == account1.ID)
	}
}
