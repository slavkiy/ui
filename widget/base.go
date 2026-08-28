package widget

import "github.com/slavkiy/ui/reactive"

type TextElement struct {
	*Element
	text *reactive.Signal[string]
}

func Text(value string) *TextElement {
	return &TextElement{
		Element: nil, //NewElement(),
		text:    reactive.NewSignal("", value),
	}
}

func (t *TextElement) Text(value string) *TextElement {
	t.text.Set(value)
	return t
}

func (t *TextElement) Align(a ...Alignment) *TextElement {
	t.Alignment = Align(a...)
	return t
}
