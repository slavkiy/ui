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

	Position Alignment
	Props    Props
}

type Props struct {
	Width  int16 // 100% relative to the parent
	Height int16 // 100% too

	Visible bool
	Clip    bool
	Opacity float32 // 0.0 to 1.0

	Layer int16
}

func NewElement(kind Kind, Layer int16, children ...Widget) *Element {
	return &Element{
		Kind:     kind,
		Children: children,
		Props: Props{
			Width:   10,
			Height:  10,
			Visible: true,
			Clip:    false,
			Opacity: 1.0,
			Layer:   Layer,
		},
	}
}

func (e *Element) Align(a ...Alignment) *Element {
	e.Position = Align(a...)
	return e
}

func (e *Element) Visible(v bool) *Element {
	e.Props.Visible = v
	return e
}

func (e *Element) Clip(v bool) *Element {
	e.Props.Clip = v
	return e
}

func (e *Element) Opacity(v float32) *Element {
	e.Props.Opacity = v
	return e
}

func (e *Element) Size(width, height int16) *Element {
	e.Props.Width = width
	e.Props.Height = height
	return e
}

type Widget interface {
	Widget()
}
