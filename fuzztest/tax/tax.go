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

// PeriodCalculator represents tax periods and related calculators.
type PeriodCalculator map[Period]Calculator

// Calculator returns the tax value for the payer in the local currency.
type Calculator func(p Payer) float64

// Agency represents tax agency that collects taxes.
type Agency string

// TaxCalcuator returns tax calculator for a period.
func (a Agency) TaxCalcuator(p Period) Calculator {
	return nil
}
