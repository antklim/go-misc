package enum

import (
	"encoding/json"
	"fmt"
)

const (
	Unknown Gender = iota
	Male
	Female
	Other
)

var genderName = [...]string{
	"unknown",
	"male",
	"female",
	"other"}

type Gender int

func (g Gender) MarshalJSON() ([]byte, error) {
	val, ok := lookup(g)
	if !ok {
		return nil, fmt.Errorf("no value for %q", g)
	}
	return json.Marshal(val)
}

func (g *Gender) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	switch s {
	case "unknown":
		*g = Unknown
		break
	case "male":
		*g = Male
		break
	case "female":
		*g = Female
		break
	case "other":
		*g = Other
		break
	default:
		return fmt.Errorf("unsupported gender %s", s)
	}

	return nil
}

func (g Gender) String() string {
	val, ok := lookup(g)
	if !ok {
		return ""
	}
	return val
}

type User struct {
	Age    int    `json:"age,omitempty" yaml:"age,omitempty"`
	Gender Gender `json:"gender" yaml:"sex"`
}

func lookup(g Gender) (string, bool) {
	if g < Unknown || g > Other {
		return "", false
	}
	return genderName[g], true
}
