package widget

import "github.com/slavkiy/ui/reactive"

type textWidget[T string | *reactive.Signal[string]] struct {
	element      *Element
	text         T
	positionText Alignment
}

func (textWidget[T]) Widget() {}

func Text[T string | *reactive.Signal[string]](value T) *textWidget[T] {
	return &textWidget[T]{
		element:      NewElement(kindText, 0, nil),
		text:         value,
		positionText: Align(Center),
	}
}

func (t *textWidget[T]) Align(a ...Alignment) *textWidget[T] {
	t.element.Align(a...)
	return t
}

func (t *textWidget[T]) AlignText(a ...Alignment) *textWidget[T] {
	t.positionText = Align(a...)
	return t
}

func (t *textWidget[T]) Layer(layer int16) *textWidget[T] {
	t.element.Props.Layer = layer
	return t
}

func (t *textWidget[T]) Visible(v bool) *textWidget[T] {
	t.element.Visible(v)
	return t
}

func (t *textWidget[T]) Clip(v bool) *textWidget[T] {
	t.element.Clip(v)
	return t
}

func (t *textWidget[T]) Opacity(v float32) *textWidget[T] {
	t.element.Opacity(v)
	return t
}

func (t *textWidget[T]) Size(width, height int16) *textWidget[T] {
	t.element.Size(width, height)
	return t
}
