package widget

import "math"

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
	Width  int16 // percentage of viewport width, e.g. 50 => 50% of screen width
	Height int16 // percentage of viewport height, e.g. 25 => 25% of screen height

	OffsetX int32
	OffsetY int32

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

func ResolvePercent(percent int16, viewport int) int {
	if viewport <= 0 {
		return 0
	}
	if percent <= 0 {
		return 0
	}
	if percent >= 100 {
		return viewport
	}
	return int(math.Round(float64(percent) / 100.0 * float64(viewport)))
}

func (e *Element) ResolveSize(viewportWidth, viewportHeight int) (int, int) {
	if e == nil {
		return 0, 0
	}
	return ResolvePercent(e.Props.Width, viewportWidth), ResolvePercent(e.Props.Height, viewportHeight)
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

func (e *Element) Offset(x, y int32) *Element {
	e.Props.OffsetX = x
	e.Props.OffsetY = y
	return e
}

func (e *Element) Layer(layer int16) *Element {
	e.Props.Layer = layer
	return e
}

type Widget interface {
	Widget()
}
