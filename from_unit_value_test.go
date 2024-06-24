package frictional

import (
	"testing"

	"github.com/alpacahq/alpacadecimal"
)

func TestFromUnitValue(t *testing.T) {
	var testCaseFromUnitValue = []struct {
		tester   func() (net, brute, totalTaxes, totalDiscounts alpacadecimal.Decimal)
		expected struct {
			net, brute, totalTaxes, totalDiscounts alpacadecimal.Decimal
		}
	}{
		{
			tester: func() (net, brute, totalTaxes, totalDiscounts alpacadecimal.Decimal) {
				unitValue := udfs("1044.543103448276")
				qty := udfs("35157")
				percDiscount := udfs("10")
				amountLineDiscount := udfs("100")
				percTax := udfs("16")
				amountLineTax := qty.Div(HundredValue).Round(0).Mul(udfs("0.04"))

				calc := NewFromUnitValue(unitValue)

				qtyVisitor := WithQTY(qty)
				percDiscVisitor := NewPercentualDiscount(percDiscount)
				amountDiscVisitor := NewAmountDiscount(amountLineDiscount)
				percTaxVisitor := NewUnbufferedPercTax(percTax)
				amountTaxVisitor := NewUnbufferedAmountTax(amountLineTax)

				calc.Bind(qtyVisitor)
				calc.Bind(percDiscVisitor)
				calc.Bind(amountDiscVisitor)
				calc.Bind(percTaxVisitor)
				calc.Bind(amountTaxVisitor)

				net = calc.Snapshot()

				calc.Add(percTaxVisitor.Amount())
				calc.Add(amountTaxVisitor.Amount())

				brute = calc.Snapshot()

				totalTaxes = percTaxVisitor.Amount().Add(amountTaxVisitor.Amount())
				totalDiscounts = percDiscVisitor.Amount().Add(amountDiscVisitor.Amount())

				return net, brute, totalTaxes, totalDiscounts
			},
			expected: struct{ net, brute, totalTaxes, totalDiscounts alpacadecimal.Decimal }{
				net:            udfs("33050601.6991379353988"),
				brute:          udfs("38338712.051000005062608"),
				totalTaxes:     udfs("5288110.351862069663808"),
				totalDiscounts: udfs("3672400.1887931039332"),
			},
		},
		{
			tester: func() (net, brute, totalTaxes, totalDiscounts alpacadecimal.Decimal) {
				unitValue := udfs("100")
				qty := udfs("1")
				percDiscount := udfs("0")
				amountLineDiscount := udfs("0")
				percTax := udfs("0")
				amountLineTax := udfs("0")

				calc := NewFromUnitValue(unitValue)

				qtyVisitor := WithQTY(qty)
				percDiscVisitor := NewPercentualDiscount(percDiscount)
				amountDiscVisitor := NewAmountDiscount(amountLineDiscount)
				percTaxVisitor := NewUnbufferedPercTax(percTax)
				amountTaxVisitor := NewUnbufferedAmountTax(amountLineTax)

				calc.Bind(qtyVisitor)
				calc.Bind(percDiscVisitor)
				calc.Bind(amountDiscVisitor)
				calc.Bind(percTaxVisitor)
				calc.Bind(amountTaxVisitor)

				net = calc.Snapshot()
				brute = calc.Snapshot()
				totalTaxes = percTaxVisitor.Amount().Add(amountTaxVisitor.Amount())
				totalDiscounts = percDiscVisitor.Amount().Add(amountDiscVisitor.Amount())

				return net, brute, totalTaxes, totalDiscounts
			},
			expected: struct {
				net, brute, totalTaxes, totalDiscounts alpacadecimal.Decimal
			}{
				net:            udfs("100"),
				brute:          udfs("100"),
				totalTaxes:     udfs("0"),
				totalDiscounts: udfs("0"),
			},
		},
	}

	for i, tc := range testCaseFromUnitValue {
		net, brute, totalTaxes, totalDiscounts := tc.tester()

		if !net.Equal(tc.expected.net) {
			t.Errorf("[Test %d] net = %s, expected %s", i, net, tc.expected.net)
		}

		if !brute.Equal(tc.expected.brute) {
			t.Errorf("[Test %d] brute = %s, expected %s", i, brute, tc.expected.brute)
		}

		if !totalTaxes.Equal(tc.expected.totalTaxes) {
			t.Errorf("[Test %d] totalTaxes = %s, expected %s", i, totalTaxes, tc.expected.totalTaxes)
		}

		if !totalDiscounts.Equal(tc.expected.totalDiscounts) {
			t.Errorf("[Test %d] totalDiscounts = %s, expected %s", i, totalDiscounts, tc.expected.totalDiscounts)
		}
	}
}
