package frictional

import "github.com/alpacahq/alpacadecimal"

var _ Visitor = &PercentualUndiscount{}
var _ Visitor = &AmountUndiscount{}

type PercentualUndiscount struct {
	*Discount
}

func NewPercentualUnDiscount(ratio alpacadecimal.Decimal) *PercentualUndiscount {
	return &PercentualUndiscount{
		Discount: &Discount{
			ratio: ratio,
		},
	}
}

func (u *PercentualUndiscount) Do(b Frictional) {
	if u.ratio.Equal(alpacadecimal.Zero) {
		return
	}

	b.set(b.Buffer().Div(HundredValue.Sub(u.ratio)).Mul(HundredValue))
	u.amount = b.Buffer().Mul(u.ratio.Div(HundredValue))
}

type AmountUndiscount struct {
	*Discount
}

func NewAmountUnDiscount(amount alpacadecimal.Decimal) *AmountUndiscount {
	return &AmountUndiscount{
		Discount: &Discount{
			amount: amount,
		},
	}
}

func (u *AmountUndiscount) Do(b Frictional) {
	b.Add(u.amount)
	u.ratio = u.amount.Mul(HundredValue).Div(b.Buffer())
}
