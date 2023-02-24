package db

import (
	"context"
	"testing"
	"time"

	"github.com/secgo/soluation_bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomeEntry(t *testing.T, account Account) Entry {
	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}
	newEntry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, newEntry)
	require.Equal(t, arg.AccountID, newEntry.AccountID)
	require.Equal(t, arg.Amount, newEntry.Amount)
	return newEntry
}

func TestCreateEntry(t *testing.T) {
	account := createRandomeAccount(t)
	createRandomeEntry(t, account)
}

func TestGetEntry(t *testing.T) {
	account := createRandomeAccount(t)
	entry := createRandomeEntry(t, account)
	getentry, err := testQueries.GetEntry(context.Background(), entry.ID)
	require.NoError(t, err)
	require.NotEmpty(t, getentry)
	require.Equal(t, entry.AccountID, getentry.AccountID)
	require.Equal(t, entry.Amount, getentry.Amount)
	require.WithinDuration(t, entry.CreatedAt, getentry.CreatedAt, time.Second)
}

func TestListEntry(t *testing.T) {
	account := createRandomeAccount(t)
	for i := 0; i < 10; i++ {
		createRandomeEntry(t, account)
	}
	arg := ListEntriesParams{
		AccountID: account.ID,
		Limit:     5,
		Offset:    5,
	}
	listsEntry, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, listsEntry, 5)

	for _, entry := range listsEntry {
		require.NotEmpty(t, entry)
	}
}
