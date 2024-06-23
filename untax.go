package frictional

import "github.com/alpacahq/alpacadecimal"

var _ Visitor = &PercentualUntax{}
var _ Visitor = &AmountUntax{}

type PercentualUntax struct {
	Tax
}

func NewPercentualUnTax(ratio alpacadecimal.Decimal) *PercentualUntax {
	return &PercentualUntax{
		Tax: Tax{
			ratio: ratio,
		},
	}
}

func (pu *PercentualUntax) Do(b Frictional) {
	ratio := pu.ratio.Div(HundredValue)
	b.Div(One.Add(ratio))
	pu.amount = b.Buffer().Mul(ratio)
}

type AmountUntax struct {
	Tax
}

func NewAmountUnTax(amount alpacadecimal.Decimal) *AmountUntax {
	return &AmountUntax{
		Tax: Tax{
			amount: amount,
		},
	}
}

func (pu *AmountUntax) Do(b Frictional) {
	b.Sub(pu.amount)
	pu.ratio = pu.amount.Mul(HundredValue).Div(b.Buffer())
}
