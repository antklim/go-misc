package create_test

import (
	"testing"

	"github.com/antklim/go-misc/create"
	"github.com/stretchr/testify/assert"
)

func TestBarista(t *testing.T) {
	for _, tC := range baristaTests {
		t.Run(tC.desc, func(t *testing.T) {
			barista := create.NewBarista(tC.brewer)
			coffee1 := barista.BrewCoffee()
			coffee2 := barista.BrewCoffee()
			assert.Equal(t, tC.expected, coffee1.Name)
			assert.Equal(t, tC.expected, coffee2.Name)
			assert.NotSame(t, &coffee1, &coffee2)
		})
	}
}

func TestBaristaSetBrewer(t *testing.T) {
	v60Brewer := create.NewV60Coffee()
	cbBrewer := create.NewColdBrew()

	barista := create.NewBarista(v60Brewer)
	coffee := barista.BrewCoffee()
	assert.Equal(t, "V60 Coffee", coffee.Name)

	barista.SetBrewer(cbBrewer)
	coffee = barista.BrewCoffee()
	assert.Equal(t, "Cold Brew", coffee.Name)
}
