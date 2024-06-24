package frictional

import "github.com/alpacahq/alpacadecimal"

type Tax struct {
	ratio   alpacadecimal.Decimal
	amount  alpacadecimal.Decimal
	taxable alpacadecimal.Decimal
}

func (pt *Tax) Amount() alpacadecimal.Decimal {
	return pt.amount
}

func (pt *Tax) Ratio() alpacadecimal.Decimal {
	return pt.ratio
}

func (pt *Tax) Taxable() alpacadecimal.Decimal {
	return pt.taxable
}

type PercTax struct {
	Tax
}

func NewPercTax(ratio alpacadecimal.Decimal) *PercTax {
	return &PercTax{
		Tax: Tax{
			ratio: ratio,
		},
	}
}

func (pt *PercTax) Do(b Frictional) {
	pt.amount = b.Buffer().Mul(pt.ratio.Div(HundredValue))
	pt.taxable = b.Buffer().Copy()
	b.Add(pt.amount)
}

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

func (pt *UnbufferedPercTax) Do(b Frictional) {
	pt.amount = b.Buffer().Mul(pt.ratio.Div(HundredValue))
	pt.taxable = b.Buffer().Copy()
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

func (pt *UnbufferedAmountTax) Do(b Frictional) {
	pt.taxable = b.Buffer().Copy()
	pt.ratio = pt.amount.Mul(HundredValue).Div(pt.taxable)
}
