package event

// MouseButton represents a mouse button or a generic pointer button.
type MouseButton uint8

const (
	MouseButtonNone MouseButton = iota
	MouseButtonLeft
	MouseButtonRight
	MouseButtonMiddle
	MouseButtonBack
	MouseButtonForward
	MouseButtonExtra1
	MouseButtonExtra2
	MouseButtonExtra3
	MouseButtonExtra4
	MouseButtonExtra5
)

// MouseButtons is a bitmask for the currently pressed mouse buttons.
type MouseButtons uint32

const (
	MouseButtonMaskNone MouseButtons = 0
	MouseButtonMaskLeft MouseButtons = 1 << iota
	MouseButtonMaskRight
	MouseButtonMaskMiddle
	MouseButtonMaskBack
	MouseButtonMaskForward
	MouseButtonMaskExtra1
	MouseButtonMaskExtra2
	MouseButtonMaskExtra3
	MouseButtonMaskExtra4
	MouseButtonMaskExtra5
)

// Has reports whether a mouse button bit is present.
func (b MouseButtons) Has(flag MouseButtons) bool {
	return b&flag == flag
}

// MouseEvent is a typed pointer event specialized for current mouse state.
type MouseEvent struct {
	PointerEvent
}
