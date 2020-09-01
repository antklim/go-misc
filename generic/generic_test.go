package generic_test

import (
	"encoding/json"
	"testing"

	"github.com/antklim/go-misc/generic"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnmarshallGenricName(t *testing.T) {
	for _, tC := range genericNameTests {
		t.Run(tC.desc, func(t *testing.T) {
			var actual generic.Name
			err := json.Unmarshal([]byte(tC.payload), &actual)
			require.NoError(t, err)
			assert.Equal(t, tC.expected, actual)
		})
	}
}
