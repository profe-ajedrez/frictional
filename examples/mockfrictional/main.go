package mockfrictional

import (
	"errors"

	"github.com/alpacahq/alpacadecimal"
	"github.com/profe-ajedrez/frictional"
)

// MockFrictional is a mock implementation of the Frictional interface.
type MockFrictional struct {
	buffer alpacadecimal.Decimal
}

// NewMockFrictional creates a new instance of MockFrictional with the given initial buffer value.
func NewMockFrictional(initial alpacadecimal.Decimal) *MockFrictional {
	return &MockFrictional{
		buffer: initial,
	}
}

// Buffer returns the current buffer value.
func (m *MockFrictional) Buffer() alpacadecimal.Decimal {
	return m.buffer
}

// Add adds the given value to the buffer.
func (m *MockFrictional) Add(value alpacadecimal.Decimal) {
	m.buffer = m.buffer.Add(value)
}

// Sub subtracts the given value from the buffer.
func (m *MockFrictional) Sub(value alpacadecimal.Decimal) {
	m.buffer = m.buffer.Sub(value)
}

// Mul multiplies the buffer by the given value.
func (m *MockFrictional) Mul(value alpacadecimal.Decimal) {
	m.buffer = m.buffer.Mul(value)
}

// Div divides the buffer by the given value.
func (m *MockFrictional) Div(value alpacadecimal.Decimal) {
	m.buffer = m.buffer.Div(value)
}

// set sets the buffer to the given value.
//
//lint:ignore U1000 we ignore set unused
func (m *MockFrictional) set(value alpacadecimal.Decimal) {
	m.buffer = value
}

// Bind binds the given Visitor to the MockFrictional instance.
// For the mock implementation, it always returns a custom error.
func (m *MockFrictional) Bind(v frictional.Visitor) error {
	return errors.New("binding operation not supported in mock implementation")
}
