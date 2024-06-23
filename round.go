package frictional

var _ Visitor = &Round{}

type Round struct {
	scale int32
}

func NewRound(scale int32) Round {
	return Round{
		scale: scale,
	}
}

func (r Round) Do(b Frictional) {
	b.set(b.Buffer().Round(r.scale))
}
