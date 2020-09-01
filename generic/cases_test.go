package generic_test

import "github.com/antklim/go-misc/generic"

/**
Valid JSON payloads

// Student 1
{
	"name": "john doe",
	"year": 1,
	"faculty": "computer science"
}

// Student 2
{
	"name": {
		"firstName": "will",
		"lastName": "smith"
	},
	"year": 1,
	"faculty": "computer science"
}

// Course 1
{
	"students": ["john doe", "will smith"],
	"year": "2020-2021"
}

// Course 2
{
	"students": ["john doe", {
		"firstName": "will",
		"lastName": "smith"
	}],
	"year": "2021-2022"
}

// Course 3
{
	"students": [{
		"firstName": "john",
		"lastName": "doe"
	}, {
		"firstName": "will",
		"lastName": "smith"
	}],
	"year": "2022-2023"
}
*/

type genericNameTest struct {
	desc     string
	payload  string
	expected generic.Name
}

var n1 generic.NameV1 = "john doe"

var genericNameTests = []genericNameTest{
	{
		"unmarshal string to generic name",
		`"john doe"`,
		generic.Name{&n1, nil},
	},
	{
		"unmarshal object to generic name",
		`{"firstName": "john", "lastName": "doe"}`,
		generic.Name{nil, &generic.NameV2{FirstName: "john", LastName: "doe"}},
	},
}
