package frictional

import "github.com/alpacahq/alpacadecimal"

var _ Visitor = &TaxHandlerFromUnitValue{}

type taxHandler struct {
	totalRatio  alpacadecimal.Decimal
	totalAmount alpacadecimal.Decimal
	taxable     alpacadecimal.Decimal
}

func NewTaxHandler() *taxHandler {
	return &taxHandler{}
}

func (t *taxHandler) WithPercentualTax(value alpacadecimal.Decimal) {
	t.totalRatio = t.totalRatio.Add(value)
}

func (t *taxHandler) WithAmountTax(value alpacadecimal.Decimal) {
	t.totalAmount = t.totalAmount.Add(value)
}

type TaxHandlerFromUnitValue struct {
	*taxHandler
}

func NewTaxHandlerFromUnitValue() *TaxHandlerFromUnitValue {
	return &TaxHandlerFromUnitValue{
		taxHandler: NewTaxHandler(),
	}
}

func (t *TaxHandlerFromUnitValue) Do(b Frictional) {
	t.taxable = b.Buffer().Copy()

	t1 := NewPercentualTax(t.totalRatio)

	t2 := NewAmountTax(t.totalAmount)

	do(b, t1, t2)

	t.totalRatio = t1.ratio.Add(t2.ratio)
	t.totalAmount = t1.amount.Add(t2.amount)
}

func (t *TaxHandlerFromUnitValue) Taxable() alpacadecimal.Decimal {
	return t.taxable.Copy()
}

func (t *TaxHandlerFromUnitValue) TotalRatio() alpacadecimal.Decimal {
	return t.totalRatio.Copy()
}

func (t *TaxHandlerFromUnitValue) TotalAmount() alpacadecimal.Decimal {
	return t.totalAmount.Copy()
}

type discountHandler struct {
	totalRatio   alpacadecimal.Decimal
	totalAmount  alpacadecimal.Decimal
	discountable alpacadecimal.Decimal
}

func NewDiscountHandler() *discountHandler {
	return &discountHandler{}
}

func (t *discountHandler) WithPercentualDiscount(value alpacadecimal.Decimal) {
	t.totalRatio = t.totalRatio.Add(value)
}

func (t *discountHandler) WithAmountDiscount(value alpacadecimal.Decimal) {
	t.totalAmount = t.totalAmount.Add(value)
}

type DiscountHandlerFromUnitValue struct {
	*discountHandler
}

func NewDiscHandlerFromUnitValue() *DiscountHandlerFromUnitValue {
	return &DiscountHandlerFromUnitValue{
		discountHandler: NewDiscountHandler(),
	}
}

func (t *DiscountHandlerFromUnitValue) Do(b Frictional) {
	t.discountable = b.Buffer().Copy()

	t1 := NewPercentualDiscount(t.totalRatio)

	t2 := NewAmountDiscount(t.totalAmount)

	do(b, t1, t2)

	t.totalRatio = t1.ratio.Add(t2.ratio)
	t.totalAmount = t1.amount.Add(t2.amount)
}

func (t *DiscountHandlerFromUnitValue) Discountable() alpacadecimal.Decimal {
	return t.discountable.Copy()
}

func (t *DiscountHandlerFromUnitValue) TotalRatio() alpacadecimal.Decimal {
	return t.totalRatio.Copy()
}

func (t *DiscountHandlerFromUnitValue) TotalAmount() alpacadecimal.Decimal {
	return t.totalAmount.Copy()
}

func do(b Frictional, e1, e2 Visitor) {
	b.Bind(e1)
	b.Bind(e2)
}
