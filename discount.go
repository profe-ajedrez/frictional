package frictional

import (
	"strings"

	"github.com/alpacahq/alpacadecimal"
)

var _ Visitor = &PercentualDiscount{}
var _ Visitor = &AmountDiscount{}

// Discount represents a discount that can be applied to a Frictional value.
// The ratio field represents the percentage discount, and the amount field
// represents the fixed amount discount.
type Discount struct {
	ratio  alpacadecimal.Decimal
	amount alpacadecimal.Decimal
}

// Ratio returns the ratio of the discount.
func (d *Discount) Ratio() alpacadecimal.Decimal {
	return d.ratio
}

// Amount returns the fixed amount discount.
func (d *Discount) Amount() alpacadecimal.Decimal {
	return d.amount
}

// String returns a string representation of the Discount, including the ratio and amount.
func (d *Discount) String() string {
	w := strings.Builder{}

	w.WriteString("ratio: ")
	w.WriteString(d.ratio.String())
	w.WriteString(" amount: ")
	w.WriteString(d.amount.String())
	return w.String()
}

// PercentualDiscount represents a discount that is applied as a percentage of the Frictional value.
// It embeds the Discount struct, which contains the ratio and amount fields.
type PercentualDiscount struct {
	Discount
}

// NewPercentualDiscount creates a new PercentualDiscount instance with the given ratio.
// The ratio represents the percentage discount to be applied.
func NewPercentualDiscount(ratio alpacadecimal.Decimal) *PercentualDiscount {
	return &PercentualDiscount{
		Discount: Discount{
			ratio: ratio,
		},
	}
}

// Do applies the percentual discount to the given Frictional value.
// It calculates the discount amount by multiplying the Frictional value's buffer
// by the discount ratio, and then dividing by 100 to get the percentage.
// The calculated discount amount is then subtracted from the Frictional value.
// This implemenetation doesnt check for negative discounts
func (pd *PercentualDiscount) Do(b Frictional) {
	pd.amount = b.Value().Mul(pd.ratio).Div(HundredValue)
	b.Sub(pd.amount)
}

// AmountDiscount represents a discount that is applied as a fixed amount.
// It embeds the Discount struct, which contains the ratio and amount fields.
type AmountDiscount struct {
	Discount
}

// NewAmountDiscount creates a new AmountDiscount instance with the given fixed amount.
// The amount represents the fixed discount to be applied.
func NewAmountDiscount(amount alpacadecimal.Decimal) *AmountDiscount {
	return &AmountDiscount{
		Discount: Discount{
			amount: amount,
		},
	}
}

// Do applies the fixed amount discount to the given Frictional value.
// If the Frictional value's buffer is zero, the discount ratio is set to zero.
// This implemenetation doesnt check for negative discounts
func (pd *AmountDiscount) Do(b Frictional) {
	if b.Value().Equal(alpacadecimal.Zero) {
		pd.ratio = Zero()
		return
	}

	pd.ratio = HundredValue.Mul(pd.amount).Div(b.Value())
	b.Sub(pd.amount)
}
