package fcf

// This package contains examples of the First Class Functions
// Example copied from https://dave.cheney.net/2016/11/13/do-not-fear-first-class-functions

// Calculator ...
type Calculator struct {
	acc float64 // accumalator
}

type opfunc func(float64, float64) float64

// Do ...
func (c *Calculator) Do(op opfunc, v float64) float64 {
	c.acc = op(c.acc, v)
	return c.acc
}

func Add(a, b float64) float64 {
	return a + b
}

func Sub(a, b float64) float64 {
	return a - b
}

func Mul(a, b float64) float64 {
	return a * b
}
