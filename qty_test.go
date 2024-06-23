package frictional

import (
	"testing"

	"github.com/alpacahq/alpacadecimal"
)

func TestWithQTY(t *testing.T) {
	testCases := []struct {
		name     string
		qty      alpacadecimal.Decimal
		expected Qty
	}{
		{
			name:     "Positive quantity",
			qty:      udfs("10"),
			expected: Qty{qty: udfs("10")},
		},
		{
			name:     "Negative quantity",
			qty:      udfs("-5"),
			expected: Qty{qty: udfs("-5")},
		},
		{
			name:     "Zero quantity",
			qty:      udfs("0"),
			expected: Qty{qty: udfs("0")},
		},
		{
			name:     "Fractional quantity",
			qty:      udfs("3.14"),
			expected: Qty{qty: udfs("3.14")},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := WithQTY(tc.qty)
			if !actual.qty.Equal(tc.expected.qty) {
				t.Errorf("Expected %v, got %v", tc.expected.qty, actual.qty)
			}
		})
	}
}

func TestQtyDo(t *testing.T) {
	testCases := []struct {
		name     string
		qty      alpacadecimal.Decimal
		initial  alpacadecimal.Decimal
		expected alpacadecimal.Decimal
	}{
		{
			name:     "Positive quantity",
			qty:      udfs("3"),
			initial:  udfs("10"),
			expected: udfs("30"),
		},
		{
			name:     "Negative quantity",
			qty:      udfs("-2"),
			initial:  udfs("10"),
			expected: udfs("-20"),
		},
		{
			name:     "Zero quantity",
			qty:      udfs("0"),
			initial:  udfs("10"),
			expected: udfs("0"),
		},
		{
			name:     "Fractional quantity",
			qty:      udfs("2.5"),
			initial:  udfs("4"),
			expected: udfs("10"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			b := &defaultFrictional{buffer: tc.initial}
			q := WithQTY(tc.qty)
			q.Do(b)
			if !b.buffer.Equal(tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, b.buffer)
			}
		})
	}
}

func TestNewUnitValue(t *testing.T) {
	testCases := []struct {
		name     string
		qty      alpacadecimal.Decimal
		expected *UnitValue
	}{
		{
			name:     "Positive quantity",
			qty:      udfs("10"),
			expected: &UnitValue{qty: udfs("10"), unitValue: udfs("0")},
		},
		{
			name:     "Negative quantity",
			qty:      udfs("-5"),
			expected: &UnitValue{qty: udfs("-5"), unitValue: udfs("0")},
		},
		{
			name:     "Zero quantity",
			qty:      udfs("0"),
			expected: &UnitValue{qty: udfs("0"), unitValue: udfs("0")},
		},
		{
			name:     "Fractional quantity",
			qty:      udfs("3.14"),
			expected: &UnitValue{qty: udfs("3.14"), unitValue: udfs("0")},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := NewUnitValue(tc.qty)
			if !actual.qty.Equal(tc.expected.qty) || !actual.unitValue.Equal(tc.expected.unitValue) {
				t.Errorf("Expected %v, got %v", tc.expected, actual)
			}
		})
	}
}

func TestUnitValueDo(t *testing.T) {
	testCases := []struct {
		name     string
		qty      alpacadecimal.Decimal
		initial  alpacadecimal.Decimal
		expected alpacadecimal.Decimal
	}{
		{
			name:     "Positive quantity",
			qty:      udfs("3"),
			initial:  udfs("30"),
			expected: udfs("10"),
		},
		{
			name:     "Zero quantity",
			qty:      udfs("0"),
			initial:  udfs("0"),
			expected: udfs("0"),
		},
		{
			name:     "Fractional quantity",
			qty:      udfs("2.5"),
			initial:  udfs("10"),
			expected: udfs("4"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			b := &defaultFrictional{buffer: tc.initial}
			q := NewUnitValue(tc.qty)
			q.Do(b)
			if !b.buffer.Equal(tc.expected) || !q.unitValue.Equal(tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, q.unitValue)
			}
		})
	}
}

func TestUnitValueRound(t *testing.T) {
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
			q := &UnitValue{unitValue: tc.initial}
			q.Round(tc.scale)
			if !q.unitValue.Equal(tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, q.unitValue)
			}
		})
	}
}
