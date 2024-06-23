package frictional

import (
	"fmt"
	"runtime"

	"github.com/alpacahq/alpacadecimal"
	"golang.org/x/exp/constraints"
)

func NewInvalidScaleErr[T constraints.Integer](scale T) error {
	return newfrictionalErr("[frictional] invalid scale %d %v", scale)
}

func NewNegativeDiscountRatioErr(ratio alpacadecimal.Decimal) error {
	return newfrictionalErr("[frictional] negative discount ratio %v   %v", ratio)
}

func NewOver100DiscountRatioErr(ratio alpacadecimal.Decimal) error {
	return newfrictionalErr("[frictional] over 100 discount ratio %v   %v", ratio)
}

func NewNegativeDiscountAmountErr(ratio alpacadecimal.Decimal) error {
	return newfrictionalErr("[frictional] negative discount amount %v   %v", ratio)
}

func NewOver100DiscountAmountErr(ratio alpacadecimal.Decimal) error {
	return newfrictionalErr("[frictional] over 100 discount amount converted %v  %v", ratio)
}

func NewNegativeTaxErr(ratio alpacadecimal.Decimal) error {
	return newfrictionalErr("[frictional] negative tax ratio %v   %v", ratio)
}

func NewOverTaxableTaxErr(ratio alpacadecimal.Decimal) error {
	return newfrictionalErr("[frictional] over 100% tax ratio %v   %v", ratio)
}

func NewNotFromUnitValue() error {
	return newfrictionalErr("[frictional] not a FromUnitValue instance", "")
}

func newfrictionalErr(e string, info any) error {
	stack := make([]uintptr, maxInfoCallstackSize)
	length := runtime.Callers(fromCaller, stack)
	return fmt.Errorf(e, info, stack[:length])
}

const (
	maxInfoCallstackSize = 12
	fromCaller           = 2
)
