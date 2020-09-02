package generic

import (
	"encoding/json"
)

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

func (n *Name) UnmarshalJSON(data []byte) error {
	var temp interface{}

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	switch v := temp.(type) {
	case string:
		nv1 := NameV1(v)
		n.NameV1 = &nv1
		n.NameV2 = nil
	default:
		var nv2 NameV2
		if err := json.Unmarshal(data, &nv2); err != nil {
			return err
		}
		n.NameV1 = nil
		n.NameV2 = &nv2
	}

	return nil
}

func (n Name) MarshalJSON() ([]byte, error) {
	switch {
	case n.NameV1 != nil:
		return json.Marshal(n.NameV1)
	case n.NameV2 != nil:
		return json.Marshal(n.NameV2)
	default:
		return json.Marshal(nil)
	}
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
