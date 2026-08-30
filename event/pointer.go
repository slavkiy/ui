package event

import "time"

// PointerDevice describes the physical device that produced pointer input.
type PointerDevice uint8

const (
	PointerDeviceUnknown PointerDevice = iota
	PointerDeviceMouse
	PointerDeviceTrackpad
	PointerDeviceStylus
	PointerDeviceTouch
	PointerDevicePen
	PointerDeviceEraser
)

// PointerPhase describes the current phase of a pointer lifecycle.
type PointerPhase uint8

const (
	PointerPhaseUnknown PointerPhase = iota
	PointerPhaseDown
	PointerPhaseMove
	PointerPhaseUp
	PointerPhaseEnter
	PointerPhaseLeave
	PointerPhaseCancel
)

// PointerEvent is the common input event for mouse, trackpad, stylus, touch, and pen.
type PointerEvent struct {
	Event

	Position        Point
	Delta           Point
	Previous        Point
	PointerID       uint64
	Device          PointerDevice
	Button          MouseButton
	Buttons         MouseButtons
	Modifiers       Modifiers
	Pressure        float32
	TiltX           float32
	TiltY           float32
	Twist           float32
	Phase           PointerPhase
	SourceTransform bool
}

// PointerDown is emitted when a pointer becomes active at a location.
type PointerDown struct {
	PointerEvent
}

// PointerUp is emitted when a pointer is released.
type PointerUp struct {
	PointerEvent
}

// PointerMove is emitted while a pointer is moving.
type PointerMove struct {
	PointerEvent
}

// PointerEnter is emitted when a pointer enters a target.
type PointerEnter struct {
	PointerEvent
}

// PointerLeave is emitted when a pointer leaves a target.
type PointerLeave struct {
	PointerEvent
}

// PointerCancel is emitted when pointer input is cancelled.
type PointerCancel struct {
	PointerEvent
}

// PointerCapture indicates that the pointer is captured by a target.
type PointerCapture struct {
	Event
	PointerID uint64
	Captured  bool
	Time      time.Duration
}

// PointerRelease indicates that a pointer capture was released.
type PointerRelease struct {
	Event
	PointerID uint64
	Time      time.Duration
}
