package enum_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"github.com/antklim/go-misc/enum"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
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
	testCases := []struct {
		desc     string
		user     enum.User
		expected string
	}{
		{
			desc:     "User with unknwown gender",
			user:     enum.User{Age: 10, Gender: enum.Unknown},
			expected: "age: 10\nsex: unknown\n",
		},
		{
			desc:     "User male",
			user:     enum.User{Age: 10, Gender: enum.Male},
			expected: "age: 10\nsex: male\n",
		},
		{
			desc:     "User female",
			user:     enum.User{Age: 10, Gender: enum.Female},
			expected: "age: 10\nsex: female\n",
		},
		{
			desc:     "User with other gender",
			user:     enum.User{Age: 10, Gender: enum.Other},
			expected: "age: 10\nsex: other\n",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual, err := yaml.Marshal(tC.user)
			assert.NoError(t, err)
			assert.Equal(t, tC.expected, string(actual))
		})
	}
}

func TestUser_UnmarshallYAML(t *testing.T) {
	testCases := []struct {
		desc     string
		data     string
		expected enum.User
		err      error
	}{
		{
			desc:     "User with unknwown gender",
			data:     "age: 10\nsex: unknown\n",
			expected: enum.User{Age: 10, Gender: enum.Unknown},
			err:      nil,
		},
		{
			desc:     "User male",
			data:     "age: 10\nsex: male\n",
			expected: enum.User{Age: 10, Gender: enum.Male},
			err:      nil,
		},
		{
			desc:     "User female",
			data:     "age: 10\nsex: female\n",
			expected: enum.User{Age: 10, Gender: enum.Female},
			err:      nil,
		},
		{
			desc:     "User with other gender",
			data:     "age: 10\nsex: other\n",
			expected: enum.User{Age: 10, Gender: enum.Other},
			err:      nil,
		},
		{
			desc:     "User with unsupported gender",
			data:     "age: 10\nsex: NA\n",
			expected: enum.User{Age: 10},
			err:      errors.New("unsupported gender NA"),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := enum.User{}
			err := yaml.Unmarshal([]byte(tC.data), &actual)
			assert.Equal(t, tC.err, err)
			assert.Equal(t, tC.expected, actual)
		})
	}
}

func TestGenter_String(t *testing.T) {
	testCases := []struct {
		desc     string
		gender   enum.Gender
		expected string
	}{
		{
			desc:     "formats unknwown gender",
			gender:   enum.Unknown,
			expected: "gender: unknown",
		},
		{
			desc:     "formats male gender",
			gender:   enum.Male,
			expected: "gender: male",
		},
		{
			desc:     "formats female gender",
			gender:   enum.Female,
			expected: "gender: female",
		},
		{
			desc:     "formats other gender",
			gender:   enum.Other,
			expected: "gender: other",
		},
		{
			desc:     "formats NA gender",
			gender:   -1,
			expected: "gender: ",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := fmt.Sprintf("gender: %s", tC.gender)
			assert.Equal(t, tC.expected, actual)
		})
	}
}
