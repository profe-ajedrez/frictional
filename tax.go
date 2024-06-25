package frictional

import "github.com/alpacahq/alpacadecimal"

// Tax struct holds the components necessary for tax calculation on a Frictional value.
// It includes the tax ratio, the tax amount, and the taxable base amount.
// This struct is typically used as a visitor to apply tax calculations to a Frictional value.
type Tax struct {
	ratio   alpacadecimal.Decimal
	amount  alpacadecimal.Decimal
	taxable alpacadecimal.Decimal
}

// Amount returns the amount of tax calculated for the Frictional value.
func (pt *Tax) Amount() alpacadecimal.Decimal {
	return pt.amount
}

// Ratio returns the tax ratio for the Frictional value.
func (pt *Tax) Ratio() alpacadecimal.Decimal {
	return pt.ratio
}

// Taxable returns the taxable value for the Frictional value that this Tax was applied to.
func (pt *Tax) Taxable() alpacadecimal.Decimal {
	return pt.taxable
}

// PercTax is a percentual based tax visitor.
// Wraps over Tax structure and implements the Do method.
type PercTax struct {
	Tax
}

// NewPercTax creates a new PercTax instance with the given tax ratio.
// The PercTax struct wraps over the Tax struct and implements the Do method
// to calculate the tax amount based on the given ratio.
func NewPercTax(ratio alpacadecimal.Decimal) *PercTax {
	return &PercTax{
		Tax: Tax{
			ratio: ratio,
		},
	}
}

// Do applies the percentual tax to the Frictional instance's value and updates the Frictional instance's buffer directly.
// It calculates the tax amount based on the Frictional instance's value and the tax ratio,
// and adds the tax amount to the Frictional instance's buffer.
// It also stores the calculated tax amount and the taxable value in the PercTax struct.
// This implemenetation doesnt check for negative taxes
func (pt *PercTax) Do(b Frictional) {
	pt.amount = b.Value().Mul(pt.ratio.Div(HundredValue))
	pt.taxable = b.Value().Copy()
	b.Add(pt.amount)
}

// UnbufferedPercTax is a Visitor that applies a percentual tax to the Frictional instance's value.
// It does not modify the Frictional instance's buffer directly.
type UnbufferedPercTax struct {
	Tax
}

func NewUnbufferedPercTax(ratio alpacadecimal.Decimal) *UnbufferedPercTax {
	return &UnbufferedPercTax{
		Tax: Tax{
			ratio: ratio,
		},
	}
}

// Do applies the percentual tax to the Frictional instance's value.
// It does not modify the Frictional instance's buffer directly.
// This implemenetation doesnt check for negative taxes
func (pt *UnbufferedPercTax) Do(b Frictional) {
	pt.amount = b.Value().Mul(pt.ratio.Div(HundredValue))
	pt.taxable = b.Value().Copy()
}

// AmountTax is a Tax that applies a fixed amount to the Frictional value.
// It wraps over the Tax struct and implements the Do method to calculate the tax amount.
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
	pt.taxable = b.Value().Copy()
	b.Add(pt.amount)
	pt.ratio = pt.amount.Mul(HundredValue).Div(pt.taxable)
}

// UnbufferedAmountTax is a Tax that applies a fixed amount to the Frictional value.
// It wraps over the Tax struct and implements the Do method to calculate the tax amount,
// but does not modify the Frictional instance's buffer directly.
type UnbufferedAmountTax struct {
	Tax
}

func NewUnbufferedAmountTax(amount alpacadecimal.Decimal) *UnbufferedAmountTax {
	return &UnbufferedAmountTax{
		Tax: Tax{
			amount: amount,
		},
	}
}

// Do applies the fixed amount tax to the Frictional instance's value.
// It does not modify the Frictional instance's buffer directly.
// This implementation calculates the tax ratio based on the fixed amount and the Frictional instance's value.
// This implemenetation doesnt check for negative taxes
func (pt *UnbufferedAmountTax) Do(b Frictional) {
	pt.taxable = b.Value().Copy()
	pt.ratio = pt.amount.Mul(HundredValue).Div(pt.taxable)
}
