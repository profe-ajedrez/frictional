package frictional

import (
	"testing"

	"github.com/alpacahq/alpacadecimal"
)

func TestSnapshotVisitorDo(t *testing.T) {
	testCases := []struct {
		name     string
		initial  alpacadecimal.Decimal
		expected alpacadecimal.Decimal
	}{
		{
			name:     "Positive value",
			initial:  udfs("10"),
			expected: udfs("10"),
		},
		{
			name:     "Negative value",
			initial:  udfs("-5.25"),
			expected: udfs("-5.25"),
		},
		{
			name:     "Zero value",
			initial:  udfs("0"),
			expected: udfs("0"),
		},
		{
			name:     "Fractional value",
			initial:  udfs("3.14159"),
			expected: udfs("3.14159"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			b := &DefaultFrictional{value: tc.initial}
			s := &SnapshotVisitor{}
			s.Do(b)
			if !s.buffer.Equal(tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, s.buffer)
			}
		})
	}
}

func TestSnapshotVisitorGet(t *testing.T) {
	testCases := []struct {
		name     string
		initial  alpacadecimal.Decimal
		expected alpacadecimal.Decimal
	}{
		{
			name:     "Positive value",
			initial:  udfs("10"),
			expected: udfs("10"),
		},
		{
			name:     "Negative value",
			initial:  udfs("-5.25"),
			expected: udfs("-5.25"),
		},
		{
			name:     "Zero value",
			initial:  udfs("0"),
			expected: udfs("0"),
		},
		{
			name:     "Fractional value",
			initial:  udfs("3.14159"),
			expected: udfs("3.14159"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s := &SnapshotVisitor{buffer: tc.initial}
			actual := s.Get()
			if !actual.Equal(tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, actual)
			}
		})
	}
}

func BenchmarkSnapshotVisitorDo(b *testing.B) {
	testCases := []struct {
		name    string
		initial alpacadecimal.Decimal
	}{
		{
			name:    "Positive value",
			initial: udfs("10"),
		},
		{
			name:    "Negative value",
			initial: udfs("-5.25"),
		},
		{
			name:    "Zero value",
			initial: udfs("0"),
		},
		{
			name:    "Fractional value",
			initial: udfs("3.14159"),
		},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			bg := &DefaultFrictional{value: tc.initial}
			s := &SnapshotVisitor{}
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				s.Do(bg)
			}
		})
	}
}

func BenchmarkSnapshotVisitorGet(b *testing.B) {
	testCases := []struct {
		name    string
		initial alpacadecimal.Decimal
	}{
		{
			name:    "Positive value",
			initial: udfs("10"),
		},
		{
			name:    "Negative value",
			initial: udfs("-5.25"),
		},
		{
			name:    "Zero value",
			initial: udfs("0"),
		},
		{
			name:    "Fractional value",
			initial: udfs("3.14159"),
		},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			s := &SnapshotVisitor{buffer: tc.initial}
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = s.Get()
			}
		})
	}
}
