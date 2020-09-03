package generic_test

import (
	"encoding/json"
	"testing"

	"github.com/antklim/go-misc/generic"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnmarshallGenricName(t *testing.T) {
	for _, tC := range genericNameUnmarshalTests {
		t.Run(tC.desc, func(t *testing.T) {
			var actual generic.Name
			err := json.Unmarshal([]byte(tC.payload), &actual)
			require.NoError(t, err)
			assert.Equal(t, tC.expected, actual)
		})
	}
}

func TestMarshalGenericName(t *testing.T) {
	for _, tC := range genericNameMarshalTests {
		t.Run(tC.desc, func(t *testing.T) {
			actual, err := json.Marshal(tC.name)
			require.NoError(t, err)
			assert.Equal(t, tC.expected, string(actual))
		})
	}
}

func TestUnmarshalStudent(t *testing.T) {
	for _, tC := range studentUnmarshalTests {
		t.Run(tC.desc, func(t *testing.T) {
			var actual generic.Student
			err := json.Unmarshal([]byte(tC.payload), &actual)
			require.NoError(t, err)
			assert.Equal(t, tC.expected, actual)
		})
	}
}

func TestMarshalStudent(t *testing.T) {
	for _, tC := range studentMarshalTests {
		t.Run(tC.desc, func(t *testing.T) {
			actual, err := json.Marshal(tC.student)
			require.NoError(t, err)
			assert.Equal(t, tC.expected, string(actual))
		})
	}
}
