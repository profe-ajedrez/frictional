package frictional

import (
	"github.com/alpacahq/alpacadecimal"
)

// ValueType is an enum that represents wether a thing is a percentual or an amount.
// Percentual and Amount are the two possible ValueType values, representing whether a value is a percentage or an absolute amount.
//
// The ten, hundred, one, minusOne, and zero constants represent common numeric values used in the package.
//
// The percentualValue and amountValue strings represent the string representations of the
// Percentual and Amount ValueType values, respectively.
type ValueType int8

const (
	// Percentual represents a percentage value
	Percentual = ValueType(0)
	// Amount represents an absolute amount value.
	Amount          = ValueType(1)
	ten             = 10
	hundred         = 100
	one             = 1
	minusOne        = -1
	zero            = 0
	percentualValue = "percentual"
	amountValue     = "amount"
	oneStrValue     = "1"
)

// String returns the string representation of ValueType
func (v ValueType) String() string {
	if v == 0 {
		return percentualValue
	}

	return amountValue
}

// IsPercentual returns true if the ValueType is Percentual, false otherwise.
func (v ValueType) IsPercentual() bool {
	return v == Percentual
}

// IsAmount returns true if the ValueType is Amount, false otherwise.
func (v ValueType) IsAmount() bool {
	return v == Amount
}

// NewValueTypeFromString returns a ValueType based on the given string.
// If the string is "1", it returns Amount. Otherwise, it returns Percentual.
func NewValueTypeFromString(s string) ValueType {
	if s == oneStrValue {
		return Amount
	}

	return Percentual
}

// NewValueTypeFromInt returns a ValueType based on the given integer.
// If the integer is 1, it returns Amount. Otherwise, it returns Percentual.
func NewValueTypeFromInt(v int) ValueType {
	if v == one {
		return Amount
	}

	return Percentual
}

// HundredValue returns an alpacadecimal.Decimal representing the value 100.
var HundredValue = alpacadecimal.NewFromInt(hundred)

// InverseValue is an alpacadecimal.Decimal representing the value -1.
var InverseValue = alpacadecimal.NewFromInt(minusOne)

// One is an alpacadecimal.Decimal representing the value 1.
var One = alpacadecimal.NewFromInt(one)

// Zero returns an alpacadecimal.Decimal representing the value 0. Allocates a new value
func Zero() alpacadecimal.Decimal {
	return alpacadecimal.Zero.Copy()
}

// Hundred returns an alpacadecimal.Decimal representing the value 100. Allocates a new value

func Hundred() alpacadecimal.Decimal {
	return alpacadecimal.NewFromInt(hundred)
}

// PowerOfTen returns an alpacadecimal.Decimal representing the value of 10 raised to the power of n.
// If n is 1, it returns 10.
// If n is 0, it returns 1.
// Otherwise, it returns 10 raised to the power of n.
// allocates one new value
func PowerOfTen(n int) alpacadecimal.Decimal {
	if n == 1 {
		tenn := alpacadecimal.NewFromInt(ten)
		return tenn
	}

	if n == 0 {
		return One.Copy()
	}

	return alpacadecimal.NewFromInt(ten).Pow(alpacadecimal.NewFromInt(int64(n)))
}

// Inverse returns an alpacadecimal.Decimal representing the value -1. Allocates a new value
func Inverse() alpacadecimal.Decimal {
	return alpacadecimal.NewFromInt(minusOne)
}
