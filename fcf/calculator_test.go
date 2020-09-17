package fcf_test

import (
	"testing"

	"github.com/antklim/go-misc/fcf"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOperations(t *testing.T) {
	c := fcf.Calculator{}

	a, err := c.Do(fcf.OP_ADD, 3.5)
	require.NoError(t, err)
	assert.Equal(t, 3.5, a)

	a, err = c.Do(fcf.OP_SUB, 1.5)
	require.NoError(t, err)
	assert.Equal(t, 2.0, a)

	a, err = c.Do(fcf.OP_MUL, 2.5)
	require.NoError(t, err)
	assert.Equal(t, 5.0, a)
}
