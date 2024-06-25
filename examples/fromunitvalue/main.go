package main

import (
	"fmt"

	"github.com/alpacahq/alpacadecimal"
	"github.com/profe-ajedrez/frictional"
)

func udfs(d string) alpacadecimal.Decimal {
	return unsafeDecFromStr(d)
}

func unsafeDecFromStr(d string) alpacadecimal.Decimal {
	dec, _ := alpacadecimal.NewFromString(d)
	return dec
}

func main() {
	// Define the values which will be used in the calculations

	unitValue := udfs("1044.543103448276")
	qty := udfs("35157")
	percDiscount := udfs("10")
	amountLineDiscount := udfs("100")
	percTax := udfs("16")
	amountLineTax := qty.Div(frictional.HundredValue).Round(0).Mul(udfs("0.04"))

	// instance the calculator as a new FromUnitValue
	calc := frictional.NewFromUnitValue(unitValue)

	// define the visitors to be used in the calculations
	qtyVisitor := frictional.WithQTY(qty)
	percDiscVisitor := frictional.NewPercentualDiscount(percDiscount)
	amountDiscVisitor := frictional.NewAmountDiscount(amountLineDiscount)
	percTaxVisitor := frictional.NewUnbufferedPercTax(percTax)
	amountTaxVisitor := frictional.NewUnbufferedAmountTax(amountLineTax)

	// bind the visitors to the calculator
	calc.Bind(qtyVisitor)
	calc.Bind(percDiscVisitor)
	calc.Bind(amountDiscVisitor)
	calc.Bind(percTaxVisitor)
	calc.Bind(amountTaxVisitor)

	// get the net value from the snapshot visitor
	net := calc.Snapshot()

	calc.Add(percTaxVisitor.Amount())
	calc.Add(amountTaxVisitor.Amount())

	// get the brute value from the snapshot visitor
	brute := calc.Snapshot()

	// get the total taxes amount from the percTaxVisitor and amountTaxVisitor visitors
	totalTaxes := percTaxVisitor.Amount().Add(amountTaxVisitor.Amount())

	// get the total discount amount from the percDiscVisitor and amountDiscVisitor visitors
	totalDiscounts := percDiscVisitor.Amount().Add(amountDiscVisitor.Amount())

	fmt.Println("net: ", net.String())
	fmt.Println("brute: ", brute.String())
	fmt.Println("total discounts: ", totalDiscounts.String())
	fmt.Println("total taxes: ", totalTaxes.String())
}
