package event

// TextInputEvent is emitted for composited or direct text input.
type TextInputEvent struct {
	Event
	Text             string
	ReplacementRange []int
	CursorPos        int
	Locale           string
	Modifiers        Modifiers
}

// IMECompositionEvent describes intermediate text composition state.
type IMECompositionEvent struct {
	Event
	Text      string
	Range     [2]int
	CursorPos int
	Locale    string
	Composing bool
}

// TextInputCommitted is emitted for a final text fragment after composition.
type TextInputCommitted struct{ TextInputEvent }

// CompositionStart is emitted when a composition session begins.
type CompositionStart struct{ IMECompositionEvent }

// CompositionUpdate is emitted while composition updates.
type CompositionUpdate struct{ IMECompositionEvent }

// CompositionEnd is emitted when composition is finalized.
type CompositionEnd struct{ IMECompositionEvent }
