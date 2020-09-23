package tax

// Tax agency provides tax strategy
// Tax strategies differ from year to year
// Tax calculator takes year and income and returns tax information

// Period represents tax period.
type Period string

const (
	TP2019 Period = "2019"
	TP2020 Period = "2020"
)

// Income represents payer income for tax period.
type Income map[Period]float64

// Payer represents tax payer.
type Payer struct {
	tfn    string
	income Income
}

// Income returns payer income in period.
func (p Payer) Income(period Period) (float64, bool) {
	income, ok := p.income[period]
	return income, ok
}

// PeriodCalculator represents tax periods and related calculators.
type PeriodCalculator map[Period]Calculator

// Calculator returns the tax value for the payer in the local currency.
type Calculator func(p Payer) float64

// Agency represents tax agency that collects taxes.
type Agency string
