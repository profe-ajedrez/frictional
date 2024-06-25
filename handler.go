package frictional

import "github.com/alpacahq/alpacadecimal"

var _ Visitor = &TaxHandlerFromUnitValue{}

// TaxHandler lets apply many taxes over the save Frictional value
type TaxHandler struct {
	totalRatio  alpacadecimal.Decimal
	totalAmount alpacadecimal.Decimal
	taxable     alpacadecimal.Decimal
}

func NewTaxHandler() *TaxHandler {
	return &TaxHandler{}
}

func (t *TaxHandler) WithPercentualTax(value alpacadecimal.Decimal) {
	t.totalRatio = t.totalRatio.Add(value)
}

func (t *TaxHandler) WithAmountTax(value alpacadecimal.Decimal) {
	t.totalAmount = t.totalAmount.Add(value)
}

type TaxHandlerFromUnitValue struct {
	*TaxHandler
}

func NewTaxHandlerFromUnitValue() *TaxHandlerFromUnitValue {
	return &TaxHandlerFromUnitValue{
		TaxHandler: NewTaxHandler(),
	}
}

func (t *TaxHandlerFromUnitValue) Do(b Frictional) {
	t.taxable = b.Value()

	t1 := NewPercTax(t.totalRatio)

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

type DiscountHandler struct {
	totalRatio   alpacadecimal.Decimal
	totalAmount  alpacadecimal.Decimal
	discountable alpacadecimal.Decimal
}

func NewDiscountHandler() *DiscountHandler {
	return &DiscountHandler{}
}

func (t *DiscountHandler) WithPercentualDiscount(value alpacadecimal.Decimal) {
	t.totalRatio = t.totalRatio.Add(value)
}

func (t *DiscountHandler) WithAmountDiscount(value alpacadecimal.Decimal) {
	t.totalAmount = t.totalAmount.Add(value)
}

type DiscountHandlerFromUnitValue struct {
	*DiscountHandler
}

func NewDiscHandlerFromUnitValue() *DiscountHandlerFromUnitValue {
	return &DiscountHandlerFromUnitValue{
		DiscountHandler: NewDiscountHandler(),
	}
}

func (t *DiscountHandlerFromUnitValue) Do(b Frictional) {
	t.discountable = b.Value().Copy()

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
