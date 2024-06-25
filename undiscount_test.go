package frictional

import (
	"testing"

	"github.com/alpacahq/alpacadecimal"
)

func TestPercentualUndiscount(t *testing.T) {
	testCases := []struct {
		name     string
		initial  alpacadecimal.Decimal
		ratio    alpacadecimal.Decimal
		expected alpacadecimal.Decimal
	}{
		{
			name:     "Positive ratio",
			initial:  udfs("100"),
			ratio:    udfs("10"),
			expected: udfs("111.11111111111111"),
		},
		{
			name:     "Negative ratio",
			initial:  udfs("100"),
			ratio:    udfs("-10"),
			expected: udfs("90.90909090909091"),
		},
		{
			name:     "Zero ratio",
			initial:  udfs("100"),
			ratio:    udfs("0"),
			expected: udfs("100"),
		},
		{
			name:     "Fractional ratio",
			initial:  udfs("100"),
			ratio:    udfs("5.5"),
			expected: udfs("105.82010582010582"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			b := &DefaultFrictional{value: tc.initial}
			u := NewPercentualUnDiscount(tc.ratio)
			u.Do(b)
			if !b.value.Equal(tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, b.value)
			}
		})
	}
}

func TestAmountUndiscount(t *testing.T) {
	testCases := []struct {
		name     string
		initial  alpacadecimal.Decimal
		amount   alpacadecimal.Decimal
		expected alpacadecimal.Decimal
	}{
		{
			name:     "Positive amount",
			initial:  udfs("100"),
			amount:   udfs("10"),
			expected: udfs("110"),
		},
		{
			name:     "Negative amount",
			initial:  udfs("100"),
			amount:   udfs("-10"),
			expected: udfs("90"),
		},
		{
			name:     "Zero amount",
			initial:  udfs("100"),
			amount:   udfs("0"),
			expected: udfs("100"),
		},
		{
			name:     "Fractional amount",
			initial:  udfs("100"),
			amount:   udfs("5.5"),
			expected: udfs("105.5"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			b := &DefaultFrictional{value: tc.initial}
			u := NewAmountUnDiscount(tc.amount)
			u.Do(b)
			if !b.value.Equal(tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, b.value)
			}
		})
	}
}

// BenchmarkPercentualUndiscount: Benchmarks the Do method of the PercentualUndiscount struct with positive, negative, zero, and fractional ratios.
func BenchmarkPercentualUndiscount(b *testing.B) {
	testCases := []struct {
		name    string
		initial alpacadecimal.Decimal
		ratio   alpacadecimal.Decimal
	}{
		{
			name:    "Positive ratio",
			initial: udfs("100"),
			ratio:   udfs("10"),
		},
		{
			name:    "Zero ratio",
			initial: udfs("100"),
			ratio:   udfs("0"),
		},
		{
			name:    "Fractional ratio",
			initial: udfs("100"),
			ratio:   udfs("5.5"),
		},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				u := NewPercentualUnDiscount(tc.ratio)
				bg := &DefaultFrictional{value: tc.initial}
				u.Do(bg)
			}
		})
	}
}

// BenchmarkAmountUndiscount: Benchmarks the Do method of the AmountUndiscount struct with positive, negative, zero, and fractional amounts.
func BenchmarkAmountUndiscount(b *testing.B) {
	testCases := []struct {
		name    string
		initial alpacadecimal.Decimal
		amount  alpacadecimal.Decimal
	}{
		{
			name:    "Positive amount",
			initial: udfs("100"),
			amount:  udfs("10"),
		},
		{
			name:    "Zero amount",
			initial: udfs("100"),
			amount:  udfs("0"),
		},
		{
			name:    "Fractional amount",
			initial: udfs("100"),
			amount:  udfs("5.5"),
		},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			bg := &DefaultFrictional{value: tc.initial}
			u := NewAmountUnDiscount(tc.amount)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				u.Do(bg)
			}
		})
	}
}
