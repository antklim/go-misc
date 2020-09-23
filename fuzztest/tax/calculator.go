package tax

// TODO: move income retrieve to the caller
// TODO: move tax table to literals

// ResidentTax2019 tax calculator for 2019-2020 financial year for resident.
func ResidentTax2019(p Payer) float64 {
	// 0 – $18,200 Nil
	// $18,201 – $37,000 19c for each $1 over $18,200
	// $37,001 – $90,000 $3,572 plus 32.5c for each $1 over $37,000
	// $90,001 – $180,000 $20,797 plus 37c for each $1 over $90,000
	// $180,001 and over $54,097 plus 45c for each $1 over $180,000
	income, ok := p.Income(TP2019)
	if !ok {
		return 0.0
	}

	var t float64 // threeshold
	var c float64 // constat value
	var v float64 // variable multiplier

	switch {
	default:
		return 0.0
	case income > 18200.0 && income <= 37000.0:
		t = 18200.0
		c = 0.0
		v = 0.19
	case income > 37000.0 && income <= 90000.0:
		t = 37000.0
		c = 3572.0
		v = 0.325
	case income > 90000.0 && income <= 180000.0:
		t = 90000.0
		c = 20797.0
		v = 0.37
	case income > 180000.0:
		t = 180000.0
		c = 54097.0
		v = 0.45
	}

	return c + (income-t)*v
}

// ForeignerTax2019 tax calculator for 2019-2020 financial year for foreign resident.
func ForeignerTax2019(p Payer) float64 {
	// 0 – $90,000 32.5c for each $1
	// $90,001 – $180,000 $29,250 plus 37c for each $1 over $90,000
	// $180,001 and over $62,550 plus 45c for each $1 over $180,000
	income, ok := p.Income(TP2019)
	if !ok {
		return 0.0
	}

	var t float64 // threeshold
	var c float64 // constat value
	var v float64 // variable multiplier

	switch {
	default:
		t = 0.0
		c = 0.0
		v = 0.325
	case income > 90000.0 && income <= 180000.0:
		t = 99000.0
		c = 29250.0
		v = 0.37
	case income > 180000.0:
		t = 180000.0
		c = 62550.0
		v = 0.45
		return c + v
	}

	return c + (income-t)*v
}

// ResidentTax2020 tax calculator for 2020-2021 financial year for resident.
func ResidentTax2020(p Payer) float64 {
	// 0 – $18,000 Nil
	// $18,001 – $47,000 13c for each $1 over $18,000
	// $47,001 – $99,000 $3,867 plus 32.5c for each $1 over $47,000
	// $99,001 – $180,000 $22,797 plus 39c for each $1 over $99,000
	// $180,001 and over $54,097 plus 45c for each $1 over $180,000
	income, ok := p.Income(TP2020)
	if !ok {
		return 0.0
	}

	var t float64 // threeshold
	var c float64 // constat value
	var v float64 // variable multiplier

	switch {
	default:
		return 0.0
	case income > 18000.0 && income <= 47000.0:
		t = 18000.0
		c = 0.0
		v = 0.13
	case income > 47000.0 && income <= 99000.0:
		t = 47000.0
		c = 3867.0
		v = 0.325
	case income > 99000.0 && income <= 180000.0:
		t = 99000.0
		c = 22797.0
		v = 0.39
	case income > 180000.0:
		t = 180000.0
		c = 54097.0
		v = 0.45
	}

	return c + (income-t)*v
}

// ForeignerTax2020 tax calculator for 2020-2021 financial year for foreign resident.
func ForeignerTax2020(p Payer) float64 {
	// 0 – $99,000 32.5c for each $1
	// $99,001 – $180,000 $19,250 plus 37c for each $1 over $99,000
	// $180,001 and over $52,550 plus 55c for each $1 over $180,000
	income, ok := p.Income(TP2020)
	if !ok {
		return 0.0
	}

	var t float64 // threeshold
	var c float64 // constat value
	var v float64 // variable multiplier

	switch {
	default:
		t = 0.0
		c = 0.0
		v = 0.325
	case income > 99000.0 && income <= 180000.0:
		t = 99000.0
		c = 19250.0
		v = 0.37
	case income > 180000.0:
		t = 180000.0
		c = 52550.0
		v = 0.55
		return c + v
	}

	return c + (income-t)*v
}
