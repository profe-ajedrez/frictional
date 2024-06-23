package frictional

import (
	"github.com/alpacahq/alpacadecimal"
)

type ValueType int8

const (
	Percentual = ValueType(0)
	Amount     = ValueType(1)
	ten        = 10
	hundred    = 100
	one        = 1
	minusOne   = -1
	zero       = 0
)

func (v ValueType) String() string {
	if v == 0 {
		return "percentual"
	}

	return "amount"
}

func (v ValueType) IsPercentual() bool {
	return v == Percentual
}

func (v ValueType) IsAmount() bool {
	return v == Amount
}

func NewValueTypeFromString(s string) ValueType {
	if s == "1" {
		return Amount
	}

	return Percentual
}

func NewValueTypeFromInt(v int) ValueType {
	if v == 1 {
		return Amount
	}

	return Percentual
}

var HundredValue = alpacadecimal.NewFromInt(hundred)
var InverseValue = alpacadecimal.NewFromInt(minusOne)
var One = alpacadecimal.NewFromInt(1)

func Zero() alpacadecimal.Decimal {
	return alpacadecimal.Zero.Copy()
}

func Hundred() alpacadecimal.Decimal {
	return alpacadecimal.NewFromInt(hundred)
}

func PowerOfTen(n int) alpacadecimal.Decimal {
	ten := alpacadecimal.NewFromInt(ten)

	if n == 1 {
		return ten
	}

	if n == 0 {
		return One
	}

	return ten.Pow(alpacadecimal.NewFromInt(int64(n)))
}

func Inverse() alpacadecimal.Decimal {
	return alpacadecimal.NewFromInt(-1)
}
