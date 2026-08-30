package event

// TouchPhase defines the lifecycle of a touch contact.
type TouchPhase uint8

const (
	TouchPhaseUnknown TouchPhase = iota
	TouchPhaseStart
	TouchPhaseMove
	TouchPhaseEnd
	TouchPhaseCancel
)

// TouchEvent is a multi-touch event normalized across mobile and desktop devices.
type TouchEvent struct {
	Event

	ID               uint64
	Position         Point
	PreviousPosition Point
	Pressure         float32
	Device           PointerDevice
	Modifiers        Modifiers
	Phase            TouchPhase
}

// TouchStart is emitted when a new touch contact begins.
type TouchStart struct {
	TouchEvent
}

// TouchMove is emitted for a touch contact that is moving.
type TouchMove struct {
	TouchEvent
}

// TouchEnd is emitted when a touch contact ends.
type TouchEnd struct {
	TouchEvent
}

// TouchCancel is emitted when a touch sequence is cancelled by the platform.
type TouchCancel struct {
	TouchEvent
}
