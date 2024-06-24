package frictional

import (
	"strings"

	"github.com/alpacahq/alpacadecimal"
)

var _ Frictional = &FromUnitValue{}
var _ Frictional = &FromBrute{}

// Arithmetic is an interface that defines arithmetic operations that can be performed on the buffer of a Frictional instance.
type Arithmetic interface {
	// Add adds the given decimal value to the buffer.
	Add(alpacadecimal.Decimal)
	// Sub subtracts the given decimal value from the buffer.
	Sub(alpacadecimal.Decimal)
	// Mul multiplies the buffer by the given decimal value.
	Mul(alpacadecimal.Decimal)
	Div(alpacadecimal.Decimal)
}

// Frictional is an interface that defines the behavior of a Frictional instance.
// It provides methods for managing the buffer, performing arithmetic operations,
// and interacting with a Visitor.
//
// The Buffer method returns the current value of the buffer.
//
// The Reset method sets the buffer to zero.
//
// The Snapshot and Restore methods allow for saving and restoring the state of the buffer.
//
// The Bind method allows a Visitor to interact with the Frictional instance.
type Frictional interface {

	// Bind allows a Visitor to interact with the Frictional instance.
	Bind(Visitor)
	// Returns a string representation of the Frictional instance, including the current value of the buffer.
	String() string

	// Buffer returns the current value of the buffer.
	Buffer() alpacadecimal.Decimal

	// Reset sets the buffer to zero.
	Reset()
	// Snapshot returns a copy of the current buffer value.
	Snapshot() alpacadecimal.Decimal
	// Restore sets the buffer to the provided decimal value.
	Restore(alpacadecimal.Decimal)

	set(alpacadecimal.Decimal)

	Arithmetic
}

// defaultFrictional is a concrete implementation of the Frictional interface.
// It holds the current value of the buffer as an alpacadecimal.Decimal.
type defaultFrictional struct {
	buffer alpacadecimal.Decimal
}

func (b *defaultFrictional) Buffer() alpacadecimal.Decimal {
	return b.buffer
}

func (b *defaultFrictional) Add(v alpacadecimal.Decimal) {
	b.buffer = b.buffer.Add(v)
}

func (b *defaultFrictional) Sub(v alpacadecimal.Decimal) {
	b.buffer = b.buffer.Sub(v)
}

func (b *defaultFrictional) Mul(v alpacadecimal.Decimal) {
	b.buffer = b.buffer.Mul(v)
}

func (b *defaultFrictional) Div(v alpacadecimal.Decimal) {
	b.buffer = b.buffer.Div(v)
}

func (b *defaultFrictional) Reset() {
	b.buffer = Zero()
}

func (b *defaultFrictional) String() string {
	w := strings.Builder{}

	w.WriteString("buffer: ")
	w.WriteString(b.buffer.String())

	return w.String()
}

func (b *defaultFrictional) Bind(e Visitor) {
	e.Do(b)
}

func (b *defaultFrictional) Snapshot() alpacadecimal.Decimal {
	return b.Buffer()
}

func (b *defaultFrictional) Restore(s alpacadecimal.Decimal) {
	b.set(s)
}

func (b *defaultFrictional) set(buffer alpacadecimal.Decimal) {
	b.buffer = buffer
}

// FromUnitValue is a struct that embeds the defaultFrictional struct.
// It provides a way to create a new FromUnitValue instance with a default or specified buffer value.
type FromUnitValue struct {
	*defaultFrictional
}

func NewFromUnitValueDefault() *FromUnitValue {
	return &FromUnitValue{
		defaultFrictional: &defaultFrictional{},
	}
}

func NewFromUnitValue(entry alpacadecimal.Decimal) *FromUnitValue {
	return &FromUnitValue{
		defaultFrictional: &defaultFrictional{
			buffer: entry,
		},
	}
}

// FromBrute is a struct that embeds the defaultFrictional struct.
// It provides a way to create a new FromBrute instance with a default or specified buffer value.
type FromBrute struct {
	*defaultFrictional
}

func NewFromBruteDefault() *FromBrute {
	return &FromBrute{
		defaultFrictional: &defaultFrictional{},
	}
}

func NewFromBrute(brute alpacadecimal.Decimal) *FromBrute {
	return &FromBrute{
		defaultFrictional: &defaultFrictional{
			buffer: brute,
		},
	}
}

// WithBrute sets the buffer of the FromBrute instance to the provided brute value and returns the updated instance.
func (f *FromBrute) WithBrute(brute alpacadecimal.Decimal) *FromBrute {
	f.buffer = brute
	return f
}
