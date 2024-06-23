package mockfrictional_test

import (
	"testing"

	"github.com/alpacahq/alpacadecimal"
	"github.com/profe-ajedrez/frictional"
	mockfrictional "github.com/profe-ajedrez/frictional/examples/mockfrictional"
)

func TestSomeFunction(t *testing.T) {
	// Create a new instance of MockFrictional with an initial buffer value
	mockFrictional := mockfrictional.NewMockFrictional(udfs("100"))

	// Create a mock Visitor
	mockVisitor := &mockVisitor{}

	// Call the Bind method and handle the error
	err := mockFrictional.Bind(mockVisitor)
	if err == nil {
		t.Error("Expected an error from Bind method, but got nil")
	}

	// Perform other assertions as needed
	// ...
}

func udfs(s string) alpacadecimal.Decimal {
	v, _ := alpacadecimal.NewFromString(s)
	return v
}

type mockVisitor struct{}

func (m *mockVisitor) Do(b frictional.Frictional) {
	// Mock implementation of the Do method
}
