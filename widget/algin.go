package widget

type Alignment uint8

const (
	Left  Alignment = 1 << iota // 0001
	Up                          // 0010
	Right                       // 0100
	Down                        // 1000
)

const Center Alignment = 0

func align(a, b Alignment) Alignment {
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
