package frictional

import "github.com/alpacahq/alpacadecimal"

var _ Visitor = &SnapshotVisitor{}

// SnapshotVisitor lets take a snapshot at the current value of a Frictional
type SnapshotVisitor struct {
	buffer alpacadecimal.Decimal
}

func NewSnapshot() *SnapshotVisitor {
	return &SnapshotVisitor{}
}

func (s *SnapshotVisitor) Do(b Frictional) {
	s.buffer = b.Value()
}

func (s *SnapshotVisitor) Get() alpacadecimal.Decimal {
	return s.buffer
}
