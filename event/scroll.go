package event

// ScrollEvent represents a scroll amount that may be discrete or continuous.
type ScrollEvent struct {
	Event

	DeltaX    float32
	DeltaY    float32
	Position  Point
	Precise   bool
	Modifiers Modifiers
}

// ScrollStart is emitted when a scroll sequence begins.
type ScrollStart struct {
	ScrollEvent
}

// ScrollUpdate is emitted while a scroll sequence is active.
type ScrollUpdate struct {
	ScrollEvent
}

// ScrollEnd is emitted when a scroll sequence ends.
type ScrollEnd struct {
	ScrollEvent
}
