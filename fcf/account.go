package fcf

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

// TODO: Run benchmark tests to verify there's no race conditions

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

func (l *ledger) Append(f string, v float64) {
	ts := time.Now().UTC().Round(time.Microsecond).Format(time.RFC3339)
	s := fmt.Sprintf(f, v, ts)
	*l = append(*l, s)
}

// balanceop is an operation on balance. It gets an argument and returns new balance.
type balanceop func() float64

// Account ...
type Account struct {
	number  string         // account number
	balance float64        // account balance
	ledger  ledger         // account operations ledger
	bops    chan balanceop // balance operations channel
}

// NewAccount ...
func NewAccount(number string) *Account {
	bopsCh := make(chan balanceop)
	a := &Account{
		number: number,
		bops:   bopsCh,
	}
	go a.loop()
	return a
}

// Deposit ...
func (a *Account) Deposit(v float64) error {
	if v < 0 {
		return errNegativeDeposit
	}

	a.bops <- func() float64 {
		a.ledger.Append(logDeposit, v)
		a.balance += v
		return a.balance
	}

	return nil
}

// Withdraw ...
func (a *Account) Withdraw(v float64) error {
	if v < 0 {
		return errNegativeWithdraw
	}

	a.bops <- func() float64 {
		a.ledger.Append(logWithdraw, v)
		a.balance -= v
		return a.balance
	}

	return nil
}

// Dividend ...
func (a *Account) Dividend(v float64) error {
	if v < 0 {
		return errNegativeInterestRate
	}

	a.bops <- func() float64 {
		a.ledger.Append(logDividend, v)
		a.balance *= 1 + v/100
		return a.balance
	}

	return nil
}

// Balance ...
func (a Account) Balance() float64 {
	return a.balance
}

// Ledger ...
func (a Account) Ledger() string {
	return strings.Join(a.ledger, ",")
}

func (a *Account) loop() {
	for op := range a.bops {
		op()
	}
}
