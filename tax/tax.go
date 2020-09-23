package tax

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
type Agency struct{}

// FindCalcuator returns tax calculator for a period.
func (a Agency) FindCalcuator(p Period) Calculator {
	switch p {
	default:
		return DefaultCalc
	case TP2019:
		return Calc2019
	case TP2020:
		return Calc2020
	}
}
