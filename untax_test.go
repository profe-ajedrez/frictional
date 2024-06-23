package frictional

import (
	"testing"

	"github.com/alpacahq/alpacadecimal"
)

func TestPercentualUntax(t *testing.T) {
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
			expected: udfs("90.9090909090909091"),
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
			expected: udfs("94.7867298578199052"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			b := &defaultFrictional{buffer: tc.initial}
			u := NewPercentualUnTax(tc.ratio)
			u.Do(b)
			if !b.buffer.Equal(tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, b.buffer)
			}
		})
	}
}

func TestAmountUntax(t *testing.T) {
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
			expected: udfs("94.5"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			b := &defaultFrictional{buffer: tc.initial}
			u := NewAmountUnTax(tc.amount)
			u.Do(b)
			if !b.buffer.Equal(tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, b.buffer)
			}
		})
	}
}

func BenchmarkPercentualUntax(b *testing.B) {
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

	b.ResetTimer()

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			bg := NewFromBrute(tc.initial)
			u := NewPercentualUnTax(tc.ratio)

			for i := 0; i < b.N; i++ {
				u.Do(bg)
			}
		})
	}
}

func BenchmarkAmountUntax(b *testing.B) {
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

	b.ResetTimer()

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				bg := NewFromBrute(tc.initial)
				u := NewAmountUnTax(tc.amount)
				u.Do(bg)
			}
		})
	}
}
