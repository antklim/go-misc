package pattern_test

import "github.com/antklim/go-misc/pattern"

type baristaTest struct {
	desc     string
	brewer   pattern.CoffeeBrewer
	expected string
}

var baristaTests = []baristaTest{
	{
		desc:     "brews v60",
		brewer:   pattern.NewV60Coffee(),
		expected: "V60 Coffee",
	},
	{
		desc:     "brews cold brew",
		brewer:   pattern.NewColdBrew(),
		expected: "Cold Brew",
	},
	{
		desc:     "brews latte",
		brewer:   pattern.NewLatte(),
		expected: "Latte",
	},
}

var textMessageExpected = `recipients: [A B]
subject: Welcome to builder

	Dear A and B
Nice text
With nice footer!`

var jsonMessageExpected = `{"recipients":["A","B"],"subject":"Welcome to builder","greeting":"Dear A and B","body":"Nice text","footer":"With nice footer!"}`
