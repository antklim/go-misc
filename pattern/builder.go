package pattern

// TODO: Add second type of builder: text/mesage builder that takes different
// values and output text/message in different formats

// CoffeeBrewer is a coffee builder interface.
type CoffeeBrewer interface {
	AddCoffee() CoffeeBrewer
	AddWater() CoffeeBrewer
	AddMilk() CoffeeBrewer
	GetCoffee() CoffeeDrink
}

// CoffeeDrink is a coffee drink.
type CoffeeDrink struct {
	Name string // coffee drink name
	cw   int    // coffee weght in grams
	cs   int    // coffee grind size in units 1-10 (1 fine, 10 coarse)
	wv   int    // water volume in ml
	wt   int    // water temperature in celcius
	mv   int    // milk volume
	mt   int    // milk temperature
}

// Barista is the brew process director.
type Barista struct {
	brewer CoffeeBrewer
}

// NewBarista creates a new barista.
func NewBarista(b CoffeeBrewer) *Barista {
	return &Barista{b}
}

// BrewCoffee coffee brewing.
func (b *Barista) BrewCoffee() CoffeeDrink {
	b.brewer.AddCoffee().AddWater().AddMilk()
	return b.brewer.GetCoffee()
}

// SetBrewer changes coffee brew type.
func (b *Barista) SetBrewer(cb CoffeeBrewer) {
	b.brewer = cb
}

// V60Coffee is a concrete coffee product brew.
type V60Coffee struct {
	coffee CoffeeDrink
}

func NewV60Coffee() CoffeeBrewer {
	return &V60Coffee{}
}

func (b *V60Coffee) AddCoffee() CoffeeBrewer {
	b.coffee.cw = 15
	b.coffee.cs = 5
	return b
}

func (b *V60Coffee) AddWater() CoffeeBrewer {
	b.coffee.wv = 250
	b.coffee.wt = 96
	return b
}

func (b *V60Coffee) AddMilk() CoffeeBrewer {
	return b
}

func (b *V60Coffee) GetCoffee() CoffeeDrink {
	b.coffee.Name = "V60 Coffee"
	return b.coffee
}

// ColdBrew is a concrete coffee product brew.
type ColdBrew struct {
	coffee CoffeeDrink
}

func NewColdBrew() CoffeeBrewer {
	return &ColdBrew{}
}

func (b *ColdBrew) AddCoffee() CoffeeBrewer {
	b.coffee.cw = 15
	b.coffee.cs = 7
	return b
}

func (b *ColdBrew) AddWater() CoffeeBrewer {
	b.coffee.wv = 250
	b.coffee.wt = 20
	return b
}

func (b *ColdBrew) AddMilk() CoffeeBrewer {
	return b
}

func (b *ColdBrew) GetCoffee() CoffeeDrink {
	b.coffee.Name = "Cold Brew"
	return b.coffee
}

// Latte is a concrete coffee product brew.
type Latte struct {
	coffee CoffeeDrink
}

func NewLatte() CoffeeBrewer {
	return &Latte{}
}

func (b *Latte) AddCoffee() CoffeeBrewer {
	b.coffee.cw = 18
	b.coffee.cs = 2
	return b
}

func (b *Latte) AddWater() CoffeeBrewer {
	b.coffee.wv = 36
	b.coffee.wt = 100
	return b
}

func (b *Latte) AddMilk() CoffeeBrewer {
	b.coffee.mv = 100
	b.coffee.mt = 70
	return b
}

func (b *Latte) GetCoffee() CoffeeDrink {
	b.coffee.Name = "Latte"
	return b.coffee
}
