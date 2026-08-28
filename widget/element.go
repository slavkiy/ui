package widget

type ElementInterface interface {
	Align(...Alignment) ElementInterface
}

type Element struct {
	Alignment Alignment
}

func (e *Element) Align(a ...Alignment) *Element {
	e.Alignment = Align(a...)
	return e
}
