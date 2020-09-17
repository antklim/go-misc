package fcf

import "errors"

// This package contains examples of the First Class Functions
// Example copied from https://dave.cheney.net/2016/11/13/do-not-fear-first-class-functions

// Calculator ...
type Calculator struct {
	acc float64 // accumalator
}

const (
	OP_ADD = 1 << iota
	OP_SUB
	OP_MUL
)

var unsupportedOperation = errors.New("unsupported operation")

// Do ...
func (c *Calculator) Do(op int, v float64) (float64, error) {
	switch op {
	case OP_ADD:
		c.acc += v
	case OP_SUB:
		c.acc -= v
	case OP_MUL:
		c.acc *= v
	default:
		return 0, unsupportedOperation
	}
	return c.acc, nil
}
