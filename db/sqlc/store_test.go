package db

import (
	"testing"
	"context"
	"github.com/Kira-UIT/simple-bank/util"
	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	// TODO
	store := NewStore(testDB)

	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	arg := TransferTxParams{
		FromAccountID: account1.ID,
		ToAccountID: account2.ID,
		Amount: util.RandomMoney(),
	}

	n := 5
	errs := make(chan error)
	results := make(chan TransferTxResult)

	for i := 0; i < n; i++ {
		go func() {
			result, err := store.TransferTx(context.Background(), arg)
			errs <- err
			results <- result
		}()
	}

	// Check results
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)
		
		result := <-results
		require.NotEmpty(t, result)
	}
}
