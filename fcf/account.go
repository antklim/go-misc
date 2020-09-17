package fcf

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

var (
	errNegativeDeposit      = errors.New("negative deposit value")
	errNegativeWithdraw     = errors.New("negative withdraw value")
	errNegativeInterestRate = errors.New("negative interest rate")
)

// Account ...
type Account struct {
	mu       sync.Mutex
	number   string       // account number
	balance  float64      // account balance
	ledger   []string     // account operations ledger
	add      chan float64 // deposit to the account
	sub      chan float64 // withdraw from the account
	interest chan float64 // multiply account balance by interest rate
}

// TODO: Refactor operations code duplication

func (a *Account) Deposit(v float64) error {
	if v < 0 {
		return errNegativeDeposit
	}

	a.add <- v
	return nil
}

func (a *Account) Withdraw(v float64) error {
	if v < 0 {
		return errNegativeWithdraw
	}

	a.sub <- v
	return nil
}

func (a *Account) Dividend(v float64) error {
	if v < 0 {
		return errNegativeInterestRate
	}

	a.interest <- v
	return nil
}

func (a *Account) loop() {
	for {
		select {
		case v := <-a.add:
			a.mu.Lock()
			defer a.mu.Unlock()
			a.balance += v
			a.ledger = append(a.ledger, depositLog(v))
		case v := <-a.sub:
			a.mu.Lock()
			defer a.mu.Unlock()
			a.balance -= v
			a.ledger = append(a.ledger, withdrawLog(v))
		case v := <-a.interest:
			a.mu.Lock()
			defer a.mu.Unlock()
			a.balance *= 1 + v/100
			a.ledger = append(a.ledger, dividendLog(v))
		}
	}
}

// TODO: Remove log code duplications
func depositLog(v float64) string {
	ts := time.Now().UTC().Round(time.Microsecond).Format(time.RFC3339)
	return fmt.Sprintf("deposit to account %.3f @%s", v, ts)
}

func withdrawLog(v float64) string {
	ts := time.Now().UTC().Round(time.Microsecond).Format(time.RFC3339)
	return fmt.Sprintf("withdraw from account %.3f @%s", v, ts)
}

func dividendLog(v float64) string {
	ts := time.Now().UTC().Round(time.Microsecond).Format(time.RFC3339)
	return fmt.Sprintf("dividends to account %.3f @%s", v, ts)
}
