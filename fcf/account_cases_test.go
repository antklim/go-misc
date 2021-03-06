package fcf_test

type operationType int

const (
	deposit operationType = iota
	withdraw
	dividend
)

type op struct {
	v float64
	t operationType
}

type accountTestCase struct {
	desc    string
	ops     []op
	balance float64
	ledger  []string
}

type accountErrorCase struct {
	desc    string
	op      op
	balance float64
	err     string
}

var accountTestCases = []accountTestCase{
	{
		desc:    "account deposit increases balance and adds record to ledger",
		ops:     []op{{v: 15.23, t: deposit}},
		balance: 15.23,
		ledger:  []string{"deposit to account 15.230 @"},
	},
	{
		desc:    "zero deposit does not change account balace",
		ops:     []op{{v: 0.0, t: deposit}},
		balance: 0.0,
		ledger:  []string{"deposit to account 0.000 @"},
	},
	{
		desc:    "after deposit to account, withdraw decreases account balance",
		ops:     []op{{v: 13.33, t: deposit}, {v: 11.45, t: withdraw}},
		balance: 1.88,
		ledger:  []string{"deposit to account 13.330 @", "withdraw from account 11.450 @"},
	},
	{
		desc:    "zero withdraw does not change account balace",
		ops:     []op{{v: 0.0, t: withdraw}},
		balance: 0.0,
		ledger:  []string{"withdraw from account 0.000 @"},
	},
	{
		desc:    "dividends on zero balance do not change account balace",
		ops:     []op{{v: 1.0, t: dividend}},
		balance: 0.0,
		ledger:  []string{"dividend to account 1.000 @"},
	},
	{
		desc:    "after deposit to account, zero dividends do not change account balace",
		ops:     []op{{v: 10.12, t: deposit}, {v: 0, t: dividend}},
		balance: 10.12,
		ledger:  []string{"deposit to account 10.120 @", "dividend to account 0.000 @"},
	},
	{
		desc:    "after deposit to account, dividends increase account balace",
		ops:     []op{{v: 10.12, t: deposit}, {v: 1.3, t: dividend}},
		balance: 10.251,
		ledger:  []string{"deposit to account 10.120 @", "dividend to account 1.300 @"},
	},
}

var accountErrorCases = []accountErrorCase{
	{
		desc:    "negative deposit is not valid and does not update account balance and ledger",
		op:      op{v: -1.0, t: deposit},
		balance: 0.0,
		err:     "negative deposit value",
	},
	{
		desc:    "negative withdraw is not valid and does not update account balance and ledger",
		op:      op{v: -1.0, t: withdraw},
		balance: 0.0,
		err:     "negative withdraw value",
	},
	{
		desc:    "negative dividends is not valid and does not update account balance and ledger",
		op:      op{v: -1.0, t: dividend},
		balance: 0.0,
		err:     "negative interest rate",
	},
	// {
	// 	desc: "withdraw greater than balance is not valid",
	// },
}

var accountBenchmarkTestCases = []op{
	{v: 1.0, t: deposit},
	{v: 12.23, t: deposit},
	{v: 3.4, t: deposit},
	{v: 6.11, t: deposit},
	{v: 7.23, t: deposit},
	{v: 1.0, t: dividend},
	{v: 13.553, t: withdraw},
	{v: 3.4, t: withdraw},
	{v: 6.11, t: deposit},
	{v: 7.23, t: deposit},
}
