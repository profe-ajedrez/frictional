package frictional

import "github.com/alpacahq/alpacadecimal"

var _ Visitor = &SnapshotVisitor{}

type SnapshotVisitor struct {
	buffer alpacadecimal.Decimal
}

func (s *SnapshotVisitor) Do(b Frictional) {
	s.buffer = b.Buffer()
}

func (s *SnapshotVisitor) Get() alpacadecimal.Decimal {
	return s.buffer.Copy()
}
