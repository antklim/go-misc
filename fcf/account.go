package fcf

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

// TODO: Run benchmark tests to verify there's no race conditions
// TODO: l = append(l, log(logWithdraw, v)) should be part of ledger

const (
	logDeposit  = `deposit to account %.3f @%s`
	logWithdraw = `withdraw from account %.3f @%s`
	logDividend = `dividend to account %.3f @%s`
)

var (
	errNegativeDeposit      = errors.New("negative deposit value")
	errNegativeWithdraw     = errors.New("negative withdraw value")
	errNegativeInterestRate = errors.New("negative interest rate")
)

type ledger []string

// balanceop is an operation on balance. It gets an argument and returns new balance.
type balanceop func(float64, *ledger) float64

// Account ...
type Account struct {
	number  string         // account number
	balance float64        // account balance
	ledger  *ledger        // account operations ledger
	bops    chan balanceop // balance operations channel
}

func NewAccount(number string) *Account {
	bopsCh := make(chan balanceop)
	a := &Account{
		number: number,
		bops:   bopsCh,
		ledger: new(ledger),
	}
	go a.loop()
	return a
}

func (a *Account) Deposit(v float64) error {
	if v < 0 {
		return errNegativeDeposit
	}

	a.bops <- func(b float64, l *ledger) float64 {
		*l = append(*l, log(logDeposit, v))
		b += v
		return b
	}

	return nil
}

func (a *Account) Withdraw(v float64) error {
	if v < 0 {
		return errNegativeWithdraw
	}

	a.bops <- func(b float64, l *ledger) float64 {
		*l = append(*l, log(logDeposit, v))
		b -= v
		return b
	}

	return nil
}

func (a *Account) Dividend(v float64) error {
	if v < 0 {
		return errNegativeInterestRate
	}

	a.bops <- func(b float64, l *ledger) float64 {
		*l = append(*l, log(logDeposit, v))
		b *= 1 + v/100
		return b
	}

	return nil
}

func (a Account) Balance() float64 {
	return a.balance
}

func (a Account) Ledger() string {
	return strings.Join(*a.ledger, ",")
}

func (a *Account) loop() {
	for op := range a.bops {
		a.balance = op(a.balance, a.ledger)
	}
}

func log(f string, v float64) string {
	ts := time.Now().UTC().Round(time.Microsecond).Format(time.RFC3339)
	return fmt.Sprintf(f, v, ts)
}
