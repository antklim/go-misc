package fcf_test

import (
	"math"
	"testing"

	"github.com/antklim/go-misc/fcf"
	"github.com/stretchr/testify/assert"
)

func TestOperations(t *testing.T) {
	c := fcf.Calculator{}

	a := c.Do(fcf.Add(3.5))
	assert.Equal(t, 3.5, a)

	a = c.Do(fcf.Sub(1.5))
	assert.Equal(t, 2.0, a)

	a = c.Do(fcf.Mul(2.5))
	assert.Equal(t, 5.0, a)

	a = c.Do(math.Sqrt)
	assert.InDelta(t, 2.236, a, 0.001)
}
