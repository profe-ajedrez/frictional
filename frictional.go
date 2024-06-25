// Package frictional lets perform sales calculations considering rules
// as discounts and taxes.
package frictional

import (
	"strings"

	"github.com/alpacahq/alpacadecimal"
)

var _ Frictional = &FromUnitValue{}
var _ Frictional = &FromBrute{}

// Arithmetic is an interface that defines arithmetic operations that can be performed on a Frictional.
type Arithmetic interface {
	// Add adds the given decimal value
	Add(alpacadecimal.Decimal)
	// Sub subtracts the given decimal value
	Sub(alpacadecimal.Decimal)
	// Mul multiplies the buffer by the given value
	Mul(alpacadecimal.Decimal)
	// Div divides the Frictional by the given decimal value.
	Div(alpacadecimal.Decimal)
}

// Frictional provides methods to interact with Visitors which updates its internal value
// to calculate a value appliying rules as discounts or taxes
// If needed you can implement your own Frictional type!
type Frictional interface {

	// Bind allows a Visitor to interact with the Frictional instance.
	Bind(Visitor)
	// Returns a string representation of the Frictional instance, including the current value of the buffer.
	String() string

	// Value returns the current value of the Frictional.
	Value() alpacadecimal.Decimal

	// Reset sets the buffer to zero.
	Reset()
	// Snapshot returns a copy of the current buffer value.
	Snapshot() alpacadecimal.Decimal
	// Restore sets the buffer to the provided decimal value.
	Restore(alpacadecimal.Decimal)

	set(alpacadecimal.Decimal)
	Arithmetic
}

// DefaultFrictional is a concrete implementation of the Frictional
// It holds the current value of the Frictional as an alpacadecimal.Decimal.
// Is used as a common default implementation of the Frictional interface.
// Instead of using it directly you should use [FromUnitValue] or [FromBrute] as needed.
// Also, you could implement your own Frictional type by embedding this struct, to get the basic functionality
type DefaultFrictional struct {
	value alpacadecimal.Decimal
}

// Value returns the current value of the Frictional.
func (b *DefaultFrictional) Value() alpacadecimal.Decimal {
	return b.value
}

// Add adds the given decimal value to the Frictional.
func (b *DefaultFrictional) Add(v alpacadecimal.Decimal) {
	b.value = b.value.Add(v)
}

// Sub subtracts the given decimal value from the Frictional.
func (b *DefaultFrictional) Sub(v alpacadecimal.Decimal) {
	b.value = b.value.Sub(v)
}

// Mul multiplies the Frictional by the given decimal value.
func (b *DefaultFrictional) Mul(v alpacadecimal.Decimal) {
	b.value = b.value.Mul(v)
}

// Div divides the Frictional by the given decimal value.
// This could trigger a division by zero panic because this implementation
// doesn't check if the given value is zero or not.
func (b *DefaultFrictional) Div(v alpacadecimal.Decimal) {
	b.value = b.value.Div(v)
}

// Reset sets the Frictional to zero.
func (b *DefaultFrictional) Reset() {
	b.value = Zero()
}

// String returns a string representation of the Frictional value.
func (b *DefaultFrictional) String() string {
	w := strings.Builder{}

	w.WriteString("buffer: ")
	w.WriteString(b.value.String())

	return w.String()
}

// Bind binds the given Visitor to the defaultFrictional instance.
// The Visitor will be invoked with the defaultFrictional instance
// when the Do method is called on the Visitor.
func (b *DefaultFrictional) Bind(e Visitor) {
	e.Do(b)
}

// Snapshot returns the current value of the Frictional.
func (b *DefaultFrictional) Snapshot() alpacadecimal.Decimal {
	return b.Value()
}

// Restore sets the value of the Frictional instance to the provided decimal value.
func (b *DefaultFrictional) Restore(s alpacadecimal.Decimal) {
	b.set(s)
}

func (b *DefaultFrictional) set(buffer alpacadecimal.Decimal) {
	b.value = buffer
}

// FromUnitValue wraps DefaultDecimal and implements the Frictional interface.
// Should be used when you want to perform calculations over a unit value to get a subtotal
//
// code block:
//
//	package main
//
//	import (
//
//		"fmt"
//
//		"github.com/alpacahq/alpacadecimal"
//		"github.com/profe-ajedrez/frictional"
//
//	)
//
//	func udfs(d string) alpacadecimal.Decimal {
//		return unsafeDecFromStr(d)
//	}
//
//	func unsafeDecFromStr(d string) alpacadecimal.Decimal {
//		dec, _ := alpacadecimal.NewFromString(d)
//		return dec
//	}
//
//	func main() {
//		// Define the values which will be used in the calculations
//
//		unitValue := udfs("1044.543103448276")
//		qty := udfs("35157")
//		percDiscount := udfs("10")
//		amountLineDiscount := udfs("100")
//		percTax := udfs("16")
//		amountLineTax := qty.Div(frictional.HundredValue).Round(0).Mul(udfs("0.04"))
//
//		// instance the calculator as a new FromUnitValue
//		calc := frictional.NewFromUnitValue(unitValue)
//
//		// define the visitors to be used in the calculations
//		qtyVisitor := frictional.WithQTY(qty)
//		percDiscVisitor := frictional.NewPercentualDiscount(percDiscount)
//		amountDiscVisitor := frictional.NewAmountDiscount(amountLineDiscount)
//		percTaxVisitor := frictional.NewUnbufferedPercTax(percTax)
//		amountTaxVisitor := frictional.NewUnbufferedAmountTax(amountLineTax)
//
//		// bind the visitors to the calculator
//		calc.Bind(qtyVisitor)
//		calc.Bind(percDiscVisitor)
//		calc.Bind(amountDiscVisitor)
//		calc.Bind(percTaxVisitor)
//		calc.Bind(amountTaxVisitor)
//
//		// get the net value from the snapshot visitor
//		net := calc.Snapshot()
//
//		calc.Add(percTaxVisitor.Amount())
//		calc.Add(amountTaxVisitor.Amount())
//
//		// get the brute value from the snapshot visitor
//		brute := calc.Snapshot()
//
//		// get the total taxes amount from the percTaxVisitor and amountTaxVisitor visitors
//		totalTaxes := percTaxVisitor.Amount().Add(amountTaxVisitor.Amount())
//
//		// get the total discount amount from the percDiscVisitor and amountDiscVisitor visitors
//		totalDiscounts := percDiscVisitor.Amount().Add(amountDiscVisitor.Amount())
//
//		fmt.Println("net: ", net.String())
//		fmt.Println("brute: ", brute.String())
//		fmt.Println("total discounts: ", totalDiscounts.String())
//		fmt.Println("total taxes: ", totalTaxes.String())
//	}
//
// Output:
//
//	net: 33050601.6991379353988
//	brute: 38338712.051000005062608
//	total discounts: 5288110.351862069663808
//	total taxes: 3672400.1887931039332
type FromUnitValue struct {
	*DefaultFrictional
}

// NewFromUnitValueDefault returns a new instance of FromUnitValue with a zero-valued.
func NewFromUnitValueDefault() *FromUnitValue {
	return &FromUnitValue{
		DefaultFrictional: &DefaultFrictional{},
	}
}

// NewFromUnitValue returns a new instance of FromUnitValue with the provided entry value
func NewFromUnitValue(entry alpacadecimal.Decimal) *FromUnitValue {
	return &FromUnitValue{
		DefaultFrictional: &DefaultFrictional{
			value: entry,
		},
	}
}

// FromBrute is a thing able to be converted from the brute subtotal removing
// elements as discounts and taxes through defined binded visitors
type FromBrute struct {
	*DefaultFrictional
}

// NewFromBruteDefault returns a new instance of FromBrute with a zero-valued.
func NewFromBruteDefault() *FromBrute {
	return &FromBrute{
		DefaultFrictional: &DefaultFrictional{},
	}
}

// NewFromBrute returns a new instance of FromBrute with the provided brute value set as the Frictional.
func NewFromBrute(brute alpacadecimal.Decimal) *FromBrute {
	return &FromBrute{
		DefaultFrictional: &DefaultFrictional{
			value: brute,
		},
	}
}

// WithBrute sets the value of the FromBrute instance to the provided brute value and returns the updated instance.
func (f *FromBrute) WithBrute(brute alpacadecimal.Decimal) *FromBrute {
	f.value = brute
	return f
}
