package main

import (
	"fmt"

	"github.com/antklim/go-misc/fuzztest/tax"
)

func main() {
	fmt.Println("Starting tax payments")

	p1 := &tax.Payer{TFN: "123", Residency: tax.AU}
	p1.WithIncomeEarned(tax.TP2019, 95000.0).WithIncomeEarned(tax.TP2020, 105000.0).
		WithTaxesPaid(tax.TP2019, 24200.0)

	p2 := &tax.Payer{TFN: "ABC", Residency: tax.NZ}
	p2.WithIncomeEarned(tax.TP2019, 87000.0).WithIncomeEarned(tax.TP2020, 120000.0)

	ta := tax.Agency{}
	var tc tax.Calculator

	fmt.Println("2019 Tax period")
	tc = ta.FindCalcuator(tax.TP2019)
	p1.PayTaxes(tax.TP2019, tc(*p1))
	p2.PayTaxes(tax.TP2019, tc(*p2))

	fmt.Println("2020 Tax period")
	tc = ta.FindCalcuator(tax.TP2020)
	p1.PayTaxes(tax.TP2020, tc(*p1))
	p2.PayTaxes(tax.TP2020, tc(*p2))

	fmt.Println("Taxes paid")
}
