package create_test

import "github.com/antklim/go-misc/create"

type baristaTest struct {
	desc     string
	brewer   create.CoffeeBrewer
	expected string
}

var baristaTests = []baristaTest{
	{
		desc:     "brews v60",
		brewer:   create.NewV60Coffee(),
		expected: "V60 Coffee",
	},
	{
		desc:     "brews cold brew",
		brewer:   create.NewColdBrew(),
		expected: "Cold Brew",
	},
	{
		desc:     "brews latte",
		brewer:   create.NewLatte(),
		expected: "Latte",
	},
}

var textMessageExpected = `recipients: [A B]
subject: Welcome to builder

	Dear A and B
Nice text
With nice footer!`

var jsonMessageExpected = `{"recipients":["A","B"],"subject":"Welcome to builder","greeting":"Dear A and B","body":"Nice text","footer":"With nice footer!"}`
