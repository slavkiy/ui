package event

// WindowEvent describes a window lifecycle or geometry change.
type WindowEvent struct {
	Event

	Size        Size
	Position    Point
	ScaleFactor float32
	DisplayID   uint64
}

// DisplayEvent describes monitor or display changes.
type DisplayEvent struct {
	Event
	DisplayID   uint64
	Bounds      Rect
	WorkArea    Rect
	ScaleFactor float32
	RefreshRate float32
}

// WindowCreated is emitted when a window is created.
type WindowCreated struct{ WindowEvent }

// WindowShown is emitted when a window becomes visible.
type WindowShown struct{ WindowEvent }

// WindowHidden is emitted when a window becomes hidden.
type WindowHidden struct{ WindowEvent }

// WindowClosing is emitted when a window closing is requested.
type WindowClosing struct{ WindowEvent }

// WindowClosed is emitted when a window has been closed.
type WindowClosed struct{ WindowEvent }

// WindowMoved is emitted when a window changes position.
type WindowMoved struct{ WindowEvent }

// WindowResized is emitted when a window changes size.
type WindowResized struct{ WindowEvent }

// WindowMinimized is emitted when a window is minimized.
type WindowMinimized struct{ WindowEvent }

// WindowMaximized is emitted when a window is maximized.
type WindowMaximized struct{ WindowEvent }

// WindowRestored is emitted when a window is restored.
type WindowRestored struct{ WindowEvent }

// WindowFocused is emitted when a window gains focus.
type WindowFocused struct{ WindowEvent }

// WindowUnfocused is emitted when a window loses focus.
type WindowUnfocused struct{ WindowEvent }

// WindowActivated is emitted when a window becomes active.
type WindowActivated struct{ WindowEvent }

// WindowDeactivated is emitted when a window ceases to be active.
type WindowDeactivated struct{ WindowEvent }

// DisplayConnected is emitted when a display is connected.
type DisplayConnected struct{ DisplayEvent }

// DisplayDisconnected is emitted when a display is disconnected.
type DisplayDisconnected struct{ DisplayEvent }

// DisplayChanged is emitted when display configuration changes.
type DisplayChanged struct{ DisplayEvent }

// ScaleFactorChanged is emitted when the UI scale changes.
type ScaleFactorChanged struct{ DisplayEvent }
