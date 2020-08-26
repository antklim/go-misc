package enum_test

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/antklim/go-misc/enum"
	"github.com/stretchr/testify/assert"
)

func TestUser_MarshallJSON(t *testing.T) {
	testCases := []struct {
		desc     string
		user     enum.User
		expected string
	}{
		{
			desc:     "User with unknwown gender",
			user:     enum.User{Age: 10, Gender: enum.Unknown},
			expected: `{"age":10,"gender":"unknown"}`,
		},
		{
			desc:     "User male",
			user:     enum.User{Age: 10, Gender: enum.Male},
			expected: `{"age":10,"gender":"male"}`,
		},
		{
			desc:     "User female",
			user:     enum.User{Age: 10, Gender: enum.Female},
			expected: `{"age":10,"gender":"female"}`,
		},
		{
			desc:     "User with other gender",
			user:     enum.User{Age: 10, Gender: enum.Other},
			expected: `{"age":10,"gender":"other"}`,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual, err := json.Marshal(tC.user)
			assert.NoError(t, err)
			assert.Equal(t, tC.expected, string(actual))
		})
	}
}

func TestUser_UnmarshallJSON(t *testing.T) {
	testCases := []struct {
		desc     string
		data     string
		expected enum.User
		err      error
	}{
		{
			desc:     "User with unknwown gender",
			data:     `{"age":10,"gender":"unknown"}`,
			expected: enum.User{Age: 10, Gender: enum.Unknown},
			err:      nil,
		},
		{
			desc:     "User male",
			data:     `{"age":10,"gender":"male"}`,
			expected: enum.User{Age: 10, Gender: enum.Male},
			err:      nil,
		},
		{
			desc:     "User female",
			data:     `{"age":10,"gender":"female"}`,
			expected: enum.User{Age: 10, Gender: enum.Female},
			err:      nil,
		},
		{
			desc:     "User with other gender",
			data:     `{"age":10,"gender":"other"}`,
			expected: enum.User{Age: 10, Gender: enum.Other},
			err:      nil,
		},
		{
			desc:     "User with unsupported gender",
			data:     `{"age":10,"gender":"NA"}`,
			expected: enum.User{Age: 10},
			err:      errors.New("unsupported gender NA"),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := enum.User{}
			err := json.Unmarshal([]byte(tC.data), &actual)
			assert.Equal(t, tC.err, err)
			assert.Equal(t, tC.expected, actual)
		})
	}
}

func TestUser_MarshallYAML(t *testing.T) {
	t.Skip("should marshall User structure to YAML")

	testCases := []struct {
		desc string
	}{
		{
			desc: "",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {

		})
	}
}

func TestUser_UnmarshallYAML(t *testing.T) {
	t.Skip("should unmarshall YAML to User structure")

	testCases := []struct {
		desc string
	}{
		{
			desc: "",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {

		})
	}
}
