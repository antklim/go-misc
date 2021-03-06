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

	assertLedger := func(t *testing.T, entries []string, ledger string) {
		for _, entry := range entries {
			assert.Contains(t, ledger, entry)
		}
	}

	for _, tC := range accountTestCases {
		t.Run(tC.desc, func(t *testing.T) {
			acc := fcf.NewAccount("123")
			runOps(acc, tC.ops)
			assert.InDelta(t, tC.balance, acc.Balance(), 0.001)
			assertLedger(t, tC.ledger, acc.Ledger())
		})
	}
}

func TestAccountOperationsErrors(t *testing.T) {
	for _, tC := range accountErrorCases {
		t.Run(tC.desc, func(t *testing.T) {
			acc := fcf.NewAccount("123")
			var err error
			switch tC.op.t {
			case deposit:
				err = acc.Deposit(tC.op.v)
			case withdraw:
				err = acc.Withdraw(tC.op.v)
			case dividend:
				err = acc.Dividend(tC.op.v)
			}
			assert.EqualError(t, err, tC.err)
		})
	}
}

func BenchmarkAccountOperations(b *testing.B) {
	acc := fcf.NewAccount("123")
	for i := 0; i < b.N; i++ {
		for _, tC := range accountBenchmarkTestCases {
			var err error
			switch tC.t {
			case deposit:
				err = acc.Deposit(tC.v)
			case withdraw:
				err = acc.Withdraw(tC.v)
			case dividend:
				err = acc.Dividend(tC.v)
			}
			if err != nil {
				b.Error(err)
			}
		}
	}
}
