package render

type Path struct {
	commands []PathCommand
}

type PathCommandType uint8

const (
	MoveTo PathCommandType = iota
	LineTo
	QuadTo
	CubicTo
	Close
)

type PathCommand struct {
	Type PathCommandType

	P1 Point
	P2 Point
	P3 Point
}
