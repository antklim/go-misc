package fcf_test

import (
	"testing"

	"github.com/antklim/go-misc/fcf"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccount(t *testing.T) {
	acc := fcf.NewAccount("123")
	err := acc.Deposit(10.0)

	require.NoError(t, err)
	assert.InDelta(t, 10.0, acc.Balance(), 0.001)
	assert.Contains(t, acc.Ledger(), "deposit to account 10.000 @")
}
