package tax

// Tax agency provides tax strategy
// Tax strategies differ from year to year
// Tax calculator takes year and income and returns tax information

// Period represents tax period.
type Period string

// Income represents payer income for tax period.
type Income map[Period]float64

// Payer represents tax payer.
type Payer struct {
	tfn    string
	income Income
}

// PeriodCalculator represents tax periods and related calculators.
type PeriodCalculator map[Period]Calculator

// Calculator returns the tax value for the payer in the local currency.
type Calculator func(p Payer) float64

// Agency represents tax agency that collects taxes.
type Agency string
