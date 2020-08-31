package generic

type Student struct {
	Name    Name   `json:"name"`    // generic name
	Year    int    `json:"year"`    // year of study
	Faculty string `json:"faculty"` // faculty name
}

// Name generic student name that hadles two name formats.
type Name struct {
	*NameV1
	*NameV2
}

type NameV1 string
type NameV2 struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Course struct {
	Students []Name `json:"students"`
	Year     string `json:"year"`
}

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
