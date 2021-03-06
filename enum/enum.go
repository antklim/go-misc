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

	gg, err := strToGender(s)
	if err != nil {
		return err
	}

	*g = gg

	return nil
}

func (g Gender) MarshalYAML() (interface{}, error) {
	val, ok := lookup(g)
	if !ok {
		return nil, fmt.Errorf("no value for %q", g)
	}
	return val, nil
}

func (g *Gender) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	gg, err := strToGender(s)
	if err != nil {
		return err
	}

	*g = gg

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

func strToGender(s string) (Gender, error) {
	switch s {
	case "unknown":
		return Unknown, nil
	case "male":
		return Male, nil
	case "female":
		return Female, nil
	case "other":
		return Other, nil
	default:
		return -1, fmt.Errorf("unsupported gender %s", s)
	}
}
