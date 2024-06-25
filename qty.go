package frictional

import "github.com/alpacahq/alpacadecimal"

var _ Visitor = &Qty{}

type Qty struct {
	qty alpacadecimal.Decimal
}

func WithQTY(qty alpacadecimal.Decimal) Qty {
	return Qty{qty: qty}
}

func (q Qty) Do(b Frictional) {
	b.Mul(q.qty)
}

type UnitValue struct {
	qty       alpacadecimal.Decimal
	unitValue alpacadecimal.Decimal
}

func NewUnitValue(qty alpacadecimal.Decimal) *UnitValue {
	return &UnitValue{
		qty: qty,
	}
}

func (q *UnitValue) Do(b Frictional) {
	if q.qty.GreaterThan(alpacadecimal.Zero) {
		q.unitValue = b.Value().Div(q.qty)
		b.set(q.unitValue.Copy())
	}
}

func (q *UnitValue) Get() alpacadecimal.Decimal {
	return q.unitValue
}

func (q *UnitValue) Round(sc int32) {
	q.unitValue = q.unitValue.Round(sc)
}
