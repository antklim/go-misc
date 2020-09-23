package tax

import "fmt"

type Country int

const (
	AU Country = iota + 1
	NZ
)

// IncomeEarned represents payer income for tax period.
type IncomeEarned map[Period]float64

// TaxesPaid represents amount of taxes paid in period.
type TaxesPaid map[Period]float64

// Payer represents tax payer.
type Payer struct {
	TFN          string
	Residency    Country
	incomeEarned IncomeEarned
	taxesPaid    TaxesPaid
}

// WithIncomeEarned adds income to payer.
func (p *Payer) WithIncomeEarned(period Period, income float64) *Payer {
	if p.incomeEarned == nil {
		p.incomeEarned = make(map[Period]float64)
	}
	p.incomeEarned[period] = income
	return p
}

// WithTaxesPaid adds taxes records to payer.
func (p *Payer) WithTaxesPaid(period Period, taxesPaid float64) *Payer {
	if p.taxesPaid == nil {
		p.taxesPaid = make(map[Period]float64)
	}
	p.taxesPaid[period] = taxesPaid
	return p
}

// PayTaxes updates payer state with tax payment information.
func (p *Payer) PayTaxes(period Period, v float64) {
	if _, ok := p.taxesPaid[period]; ok {
		return
	}

	fmt.Printf("payer with TFN %s paid %.3f in %s tax period\n", p.TFN, v, period)
	p.taxesPaid[period] = v
	return
}

// Income returns payer income for period.
func (p Payer) Income(period Period) (float64, bool) {
	income, ok := p.incomeEarned[period]
	return income, ok
}

// TaxesPaid returns amount of taxes paid in period.
func (p Payer) TaxesPaid(period Period) (float64, bool) {
	taxesPaid, ok := p.taxesPaid[period]
	return taxesPaid, ok
}
