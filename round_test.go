package frictional

import (
	"fmt"
	"testing"

	"github.com/alpacahq/alpacadecimal"
)

func TestNewRound(t *testing.T) {
	testCases := []struct {
		name     string
		scale    int32
		expected Round
	}{
		{
			name:     "Positive scale",
			scale:    2,
			expected: Round{scale: 2},
		},
		{
			name:     "Negative scale",
			scale:    -1,
			expected: Round{scale: -1},
		},
		{
			name:     "Zero scale",
			scale:    0,
			expected: Round{scale: 0},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := NewRound(tc.scale)
			if actual.scale != tc.expected.scale {
				t.Errorf("Expected %v, got %v", tc.expected.scale, actual.scale)
			}
		})
	}
}

func TestRoundDo(t *testing.T) {
	testCases := []struct {
		name     string
		initial  alpacadecimal.Decimal
		scale    int32
		expected alpacadecimal.Decimal
	}{
		{
			name:     "Round up",
			initial:  udfs("3.14159"),
			scale:    2,
			expected: udfs("3.14"),
		},
		{
			name:     "Round down",
			initial:  udfs("3.14159"),
			scale:    3,
			expected: udfs("3.142"),
		},
		{
			name:     "No rounding",
			initial:  udfs("3.14"),
			scale:    2,
			expected: udfs("3.14"),
		},
		{
			name:     "Negative scale",
			initial:  udfs("3.14159"),
			scale:    -1,
			expected: udfs("0"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			b := &defaultFrictional{buffer: tc.initial}
			r := NewRound(tc.scale)
			r.Do(b)
			if !b.buffer.Equal(tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, b.buffer)
			}
		})
	}
}

// BenchmarkRoundDo function benchmarks the Do method of the Round struct by creating a defaultFrictional instance with an initial value, applying the Round visitor repeatedly in a loop, and measuring the time taken for the loop to complete.
func BenchmarkRoundDo(b *testing.B) {
	testCases := []struct {
		name    string
		initial alpacadecimal.Decimal
		scale   int32
	}{
		{
			name:    "Round up",
			initial: udfs("3.14159"),
			scale:   2,
		},
		{
			name:    "Round down",
			initial: udfs("3.14159"),
			scale:   3,
		},
		{
			name:    "No rounding",
			initial: udfs("3.14"),
			scale:   2,
		},
		{
			name:    "Negative scale",
			initial: udfs("3.14159"),
			scale:   -1,
		},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			bg := &defaultFrictional{buffer: tc.initial}
			r := NewRound(tc.scale)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				r.Do(bg)
			}
		})
	}
}

// BenchmarkNewRound function benchmarks the NewRound function by creating a new Round instance repeatedly in a loop for different scales ranging from 2 to 10, and measuring the time taken for the loop to complete.
func BenchmarkNewRound(b *testing.B) {
	scales := []int32{2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, scale := range scales {
		b.Run(fmt.Sprint(scale), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = NewRound(scale)
			}
		})
	}
}
