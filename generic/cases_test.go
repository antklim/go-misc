package generic_test

import "github.com/antklim/go-misc/generic"

/**
Valid JSON payloads

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

type genericNameUnmarshalTest struct {
	desc     string
	payload  string
	expected generic.Name
}

type genericNameMarshalTest struct {
	desc     string
	name     generic.Name
	expected string
}

type studentUnmarshalTest struct {
	desc     string
	payload  string
	expected generic.Student
}

type studentMarshalTest struct {
	desc     string
	student  generic.Student
	expected string
}

type courseUnmarshalTest struct {
	desc     string
	payload  string
	expected generic.Course
}

type courseMarshalTest struct {
	desc     string
	course   generic.Course
	expected string
}

var n1 generic.NameV1 = "john doe"

var genericNameUnmarshalTests = []genericNameUnmarshalTest{
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

var genericNameMarshalTests = []genericNameMarshalTest{
	{
		"marshal name v1",
		generic.Name{&n1, nil},
		`"john doe"`,
	},
	{
		"marshal name v2",
		generic.Name{nil, &generic.NameV2{FirstName: "john", LastName: "doe"}},
		`{"firstName":"john","lastName":"doe"}`,
	},
}

var studentUnmarshalTests = []studentUnmarshalTest{
	{
		"unmarshal student with simple name",
		`{
			"name": "john doe",
			"year": 1,
			"faculty": "computer science"
		}`,
		generic.Student{generic.Name{&n1, nil}, 1, "computer science"},
	},
	{
		"unmarshal student with extended name",
		`{
			"name": {
				"firstName": "will",
				"lastName": "smith"
			},
			"year": 1,
			"faculty": "computer science"
		}`,
		generic.Student{generic.Name{nil, &generic.NameV2{FirstName: "will", LastName: "smith"}}, 1, "computer science"},
	},
}

var studentMarshalTests = []studentMarshalTest{
	{
		"marshal student with simple name",
		generic.Student{generic.Name{&n1, nil}, 1, "computer science"},
		`{"name":"john doe","year":1,"faculty":"computer science"}`,
	},
	{
		"marshal student with extended name",
		generic.Student{generic.Name{nil, &generic.NameV2{FirstName: "will", LastName: "smith"}}, 1, "computer science"},
		`{"name":{"firstName":"will","lastName":"smith"},"year":1,"faculty":"computer science"}`,
	},
}

var courseUnmarshalTests = []courseUnmarshalTest{
	{
		"unmarshal course with students with simple names",
		`{
			"students": ["john doe"],
			"year": "first year"
		}`,
		generic.Course{[]generic.Name{{&n1, nil}}, "first year"},
	},
	{
		"unmarshal course with students with extended name",
		`{
			"students": [
				{
					"firstName": "will",
					"lastName": "smith"
				}
			],
			"year": "second year"
		}`,
		generic.Course{[]generic.Name{{nil, &generic.NameV2{FirstName: "will", LastName: "smith"}}}, "second year"},
	},
	{
		"unmarshal course with students with mixed names",
		`{
			"students": [
				"john doe",
				{
					"firstName": "will",
					"lastName": "smith"
				}
			],
			"year": "third year"
		}`,
		generic.Course{[]generic.Name{
			{&n1, nil},
			{nil, &generic.NameV2{FirstName: "will", LastName: "smith"}}}, "third year"},
	},
}

var courseMarshalTests = []courseMarshalTest{
	{
		"marshal course with students with simple names",
		generic.Course{[]generic.Name{{&n1, nil}}, "first year"},
		`{"students":["john doe"],"year":"first year"}`,
	},
	{
		"marshal course with students with extended name",
		generic.Course{[]generic.Name{{nil, &generic.NameV2{FirstName: "will", LastName: "smith"}}}, "second year"},
		`{"students":[{"firstName":"will","lastName":"smith"}],"year":"second year"}`,
	},
	{
		"marshal course with students with mixed names",
		generic.Course{[]generic.Name{
			{&n1, nil},
			{nil, &generic.NameV2{FirstName: "will", LastName: "smith"}}}, "third year"},
		`{"students":["john doe",{"firstName":"will","lastName":"smith"}],"year":"third year"}`,
	},
}
