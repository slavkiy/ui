package render

type Path struct {
	commands []PathCommand
}

func NewPath(commands ...PathCommand) Path {
	if len(commands) == 0 {
		return Path{}
	}
	out := make([]PathCommand, len(commands))
	copy(out, commands)
	return Path{commands: out}
}

func (p Path) Commands() []PathCommand {
	if len(p.commands) == 0 {
		return nil
	}
	out := make([]PathCommand, len(p.commands))
	copy(out, p.commands)
	return out
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
