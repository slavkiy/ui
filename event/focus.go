package event

// FocusEvent describes focus transitions for widgets, windows, and text inputs.
type FocusEvent struct {
	Event
	TargetID uint64
	Focused  bool
}

// FocusIn is emitted when focus moves into a target.
type FocusIn struct{ FocusEvent }

// FocusOut is emitted when focus moves out of a target.
type FocusOut struct{ FocusEvent }

// FocusChanged is emitted when focus moves between two targets.
type FocusChanged struct{ FocusEvent }
