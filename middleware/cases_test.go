package middleware_test

type validateTest struct {
	desc     string
	s        string
	expected string
}

var validateTests = []validateTest{
	{
		desc:     "empty string is not valid",
		s:        "",
		expected: "invalid string length 0",
	},
	{
		desc:     "string longer than 100 chars is not valid",
		s:        "01234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567891",
		expected: "invalid string length 101",
	},
}
