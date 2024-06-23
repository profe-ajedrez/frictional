package frictional

// Visitor is an interface that defines a method for performing an operation
// on a Frictional.
// The Do method takes a Frictional as an argument and performs some operation on it.
type Visitor interface {
	Do(Frictional)
}
