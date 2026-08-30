package event

// TextCompositionState describes the current composition stage.
type TextCompositionState uint8

const (
	TextCompositionStateUnknown TextCompositionState = iota
	TextCompositionStateStarted
	TextCompositionStateUpdated
	TextCompositionStateEnded
	TextCompositionStateCancelled
)

// TextCompositionEvent is a higher-level composition state machine event.
type TextCompositionEvent struct {
	Event
	Text      string
	CursorPos int
	State     TextCompositionState
	Locale    string
}

// CompositionChange indicates a composition text change.
type CompositionChange struct{ TextCompositionEvent }
