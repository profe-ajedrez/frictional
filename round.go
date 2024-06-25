package frictional

var _ Visitor = &Round{}

// Round is a visitor which performs a rounding operation with a specified scale.
// rounding usually implies a rescale operation, which is costly, use with care.
type Round struct {
	scale int32
}

// NewRound creates a new Round visitor with the specified scale.
// The Round visitor can be used to perform a rounding operation on a Frictional value.
// The scale parameter determines the number of decimal places to round to.
func NewRound(scale int32) Round {
	return Round{
		scale: scale,
	}
}

// Do applies a rounding operation to the given Frictional value, using the scale
// specified when the Round visitor was created. This effectively rescales the
// Frictional value to the desired number of decimal places.
func (r Round) Do(b Frictional) {
	b.set(b.Value().Round(r.scale))
}
