package fcf

// This package contains examples of the First Class Functions
// Example copied from https://dave.cheney.net/2016/11/13/do-not-fear-first-class-functions

// Calculator ...
type Calculator struct {
	acc float64 // accumalator
}

type opfunc func(float64) float64

// Do ...
func (c *Calculator) Do(op opfunc) float64 {
	c.acc = op(c.acc)
	return c.acc
}

func Add(v float64) opfunc {
	return opfunc(func(acc float64) float64 { return acc + v })
}

func Sub(v float64) opfunc {
	return opfunc(func(acc float64) float64 { return acc - v })
}

func Mul(v float64) opfunc {
	return opfunc(func(acc float64) float64 { return acc * v })
}
