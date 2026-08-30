package event

// PenEvent describes stylus input with pressure and tilt data.
type PenEvent struct {
	PointerEvent
	BarrelButton bool
	Eraser       bool
}

// PenDown is emitted when a pen touches the surface.
type PenDown struct{ PenEvent }

// PenMove is emitted when a pen moves while touching.
type PenMove struct{ PenEvent }

// PenUp is emitted when a pen lifts.
type PenUp struct{ PenEvent }

// PenCancel is emitted when a pen sequence is canceled.
type PenCancel struct{ PenEvent }
