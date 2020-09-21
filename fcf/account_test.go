package fcf_test

import (
	"testing"
	"time"

	"github.com/antklim/go-misc/fcf"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccountOperations(t *testing.T) {
	runOps := func(acc *fcf.Account, ops []op) {
		for _, op := range ops {
			var err error
			switch op.t {
			case deposit:
				err = acc.Deposit(op.v)
			case withdraw:
				err = acc.Withdraw(op.v)
			case dividend:
				err = acc.Dividend(op.v)
			}
			require.NoError(t, err)
		}
		// Letting account loop to cycle
		time.Sleep(time.Millisecond)
	}

	for _, tC := range accountTestCases {
		t.Run(tC.desc, func(t *testing.T) {
			acc := fcf.NewAccount("123")
			runOps(acc, tC.ops)
			assert.InDelta(t, tC.balance, acc.Balance(), 0.001)
			ledger := acc.Ledger()
			for _, entry := range tC.ledger {
				assert.Contains(t, ledger, entry)
			}
		})
	}
}
