package widget

import "github.com/slavkiy/ui/reactive"

type textWidget[T string | *reactive.Signal[string]] struct {
	element *Element
	text    T
}

func (textWidget[T]) Widget() {}

func Text[T string | *reactive.Signal[string]](value T) *textWidget[T] {
	return &textWidget[T]{
		element: NewElement(kindText, 0, nil),
		text:    value,
	}
}

func (t *textWidget[T]) Align(a ...Alignment) *textWidget[T] {
	t.element.Align(a...)
	return t
}

func (t *textWidget[T]) Layer(layer int16) *textWidget[T] {
	t.element.Layer = layer
	return t
}
