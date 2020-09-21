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
	ledger  string
}

var accountTestCases = []accountTestCase{
	{
		desc:    "account deposit increases balance and adds record to ledger",
		ops:     []op{{v: 15.23, t: deposit}},
		balance: 15.23,
		ledger:  []string{"deposit to account 15.230 @"},
	},
	// {
	// 	desc: "zero deposit does not change account balace",
	// 	// op:      []op{{v: 15.23, t: deposit}},
	// 	// balance: 15.23,
	// 	ledger: []string{"deposit to account 15.230 @"},
	// },
	// {
	// 	desc: "after depositing account balance withdraw will decrease account balance",
	// 	// v:       15.23,
	// 	// op:      deposit,
	// 	// balance: 15.23,
	// 	// ledger:  "deposit to account 15.230 @",
	// },
}

// var accountErrorCases = []accountErrorCase{
// 	{
// 		desc: "negative deposit is not valid and does not update account balance and ledger",
// 	},
// 	{
// 		desc: "negative withdraw is not valid and does not update account balance and ledger",
// 	},
// 	{
// 		desc: "negative dividends is not valid and does not update account balance and ledger",
// 	},
// 	{
// 		desc: "zero dividends is not valid and does not update account balance and ledger",
// 	},
// }
