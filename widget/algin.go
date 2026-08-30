package widget

type Alignment uint8

const (
	Center Alignment = 0
	Left   Alignment = 1 << iota
	Up
	Right
	Down
)

func align(a, b Alignment) Alignment {
	if b == Center {
		return Center
	}

	result := a

	if b&Left != 0 {
		result &^= Right
	}
	if b&Right != 0 {
		result &^= Left
	}
	if b&Up != 0 {
		result &^= Down
	}
	if b&Down != 0 {
		result &^= Up
	}

	return result | b
}

func Align(a ...Alignment) Alignment {
	l := len(a)
	if l == 0 {
		return 0
	}
	result := a[0]
	for i := 1; i < l; i++ {
		result = align(result, a[i])
	}
	return result
}
