package frictional

import "github.com/alpacahq/alpacadecimal"

type Tax struct {
	ratio   alpacadecimal.Decimal
	amount  alpacadecimal.Decimal
	taxable alpacadecimal.Decimal
}

type PercentualTax struct {
	Tax
}

func NewPercentualTax(ratio alpacadecimal.Decimal) *PercentualTax {
	return &PercentualTax{
		Tax: Tax{
			ratio:   ratio,
			amount:  alpacadecimal.Decimal{},
			taxable: alpacadecimal.Decimal{},
		},
	}
}

func (pt *PercentualTax) Do(b Frictional) {
	pt.amount = b.Buffer().Mul(pt.ratio.Div(HundredValue))
	pt.taxable = b.Buffer().Copy()
	b.Add(pt.amount)
}

type AmountTax struct {
	Tax
}

func NewAmountTax(amount alpacadecimal.Decimal) *AmountTax {
	return &AmountTax{
		Tax: Tax{
			amount: amount,
		},
	}
}

func (pt *AmountTax) Do(b Frictional) {
	pt.taxable = b.Buffer().Copy()
	b.Add(pt.amount)
	pt.ratio = pt.amount.Mul(HundredValue).Div(pt.taxable)
}
