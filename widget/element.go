package widget

type Kind uint8

const (
	kindElement Kind = iota
	kindText
	kindRow
	kindColumn
	kindStack

	kindButton
	kindInput
	kindCheckbox
	kindRadio
	kindSelect
	kindOption
	kindTextarea
	kindImage
	kindCanvas
	kindSVG
)

type Element struct {
	Kind Kind

	Children []Widget

	Alignment Alignment
	Layer     int16
}

func NewElement(kind Kind, Layer int16, children ...Widget) *Element {
	return &Element{
		Kind:     kind,
		Children: children,
		Layer:    Layer,
	}
}

func (e *Element) Align(a ...Alignment) *Element {
	e.Alignment = Align(a...)
	return e
}

type Widget interface {
	Widget()
}
