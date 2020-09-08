package pattern_test

import "github.com/antklim/go-misc/pattern"

type brewTest struct {
	desc     string
	brewer   pattern.CoffeeBrewer
	expected string
}

var brewTests = []brewTest{
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
