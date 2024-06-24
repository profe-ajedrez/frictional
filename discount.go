package frictional

import (
	"strings"

	"github.com/alpacahq/alpacadecimal"
)

var _ Visitor = &PercentualDiscount{}
var _ Visitor = &AmountDiscount{}

type Discount struct {
	ratio  alpacadecimal.Decimal
	amount alpacadecimal.Decimal
}

func (d *Discount) Ratio() alpacadecimal.Decimal {
	return d.ratio
}

func (d *Discount) Amount() alpacadecimal.Decimal {
	return d.amount
}

func (d *Discount) String() string {
	w := strings.Builder{}

	w.WriteString("ratio: ")
	w.WriteString(d.ratio.String())
	w.WriteString(" amount: ")
	w.WriteString(d.amount.String())
	return w.String()
}

type PercentualDiscount struct {
	Discount
}

func NewPercentualDiscount(ratio alpacadecimal.Decimal) *PercentualDiscount {
	return &PercentualDiscount{
		Discount: Discount{
			ratio: ratio,
		},
	}
}

func (pd *PercentualDiscount) Do(b Frictional) {
	pd.amount = b.Buffer().Mul(pd.ratio).Div(HundredValue)
	b.Sub(pd.amount)
}

type AmountDiscount struct {
	Discount
}

func NewAmountDiscount(amount alpacadecimal.Decimal) *AmountDiscount {
	return &AmountDiscount{
		Discount: Discount{
			amount: amount,
		},
	}
}

func (pd *AmountDiscount) Do(b Frictional) {
	if b.Buffer().Equal(alpacadecimal.Zero) {
		pd.ratio = Zero()
		return
	}

	pd.ratio = HundredValue.Mul(pd.amount).Div(b.Buffer())
	b.Sub(pd.amount)
}
