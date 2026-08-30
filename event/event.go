package event

import "time"

// Type is the normalized event kind emitted by the runtime after native OS events
// are translated by the platform backend.
type Type uint16

const (
	TypeUnknown Type = iota
	TypeKeyDown
	TypeKeyUp
	TypeKeyRepeat
	TypeTextInput
	TypePointerDown
	TypePointerUp
	TypePointerMove
	TypePointerEnter
	TypePointerLeave
	TypePointerCancel
	TypePointerCapture
	TypePointerRelease
	TypeTouchStart
	TypeTouchMove
	TypeTouchEnd
	TypeTouchCancel
	TypeTap
	TypeDoubleTap
	TypeLongPress
	TypeDrag
	TypeSwipe
	TypePinch
	TypeRotate
	TypePan
	TypeScroll
	TypeScrollStart
	TypeScrollUpdate
	TypeScrollEnd
	TypeDragEnter
	TypeDragOver
	TypeDragLeave
	TypeDrop
	TypeClipboardChanged
	TypeCopy
	TypeCut
	TypePaste
	TypeWindowCreated
	TypeWindowShown
	TypeWindowHidden
	TypeWindowClosing
	TypeWindowClosed
	TypeWindowMoved
	TypeWindowResized
	TypeWindowMinimized
	TypeWindowMaximized
	TypeWindowRestored
	TypeWindowFocused
	TypeWindowUnfocused
	TypeWindowActivated
	TypeWindowDeactivated
	TypeDisplayConnected
	TypeDisplayDisconnected
	TypeDisplayChanged
	TypeScaleFactorChanged
	TypeFocusIn
	TypeFocusOut
	TypeFocusChanged
	TypeCompositionStart
	TypeCompositionUpdate
	TypeCompositionEnd
	TypeCursorEntered
	TypeCursorExited
	TypeCursorChanged
	TypeThemeChanged
	TypeAppearanceChanged
	TypeAccessibilityChanged
	TypePowerStateChanged
	TypeOrientationChanged
	TypeSafeAreaChanged
	TypeKeyboardShown
	TypeKeyboardHidden
	TypeKeyboardFrameChanged
	TypeAppLaunched
	TypeAppActivated
	TypeAppInactive
	TypeAppBackgrounded
	TypeAppForegrounded
	TypeAppSuspended
	TypeAppResumed
	TypeAppTerminated
	TypeSurfaceCreated
	TypeSurfaceDestroyed
	TypeSurfaceChanged
	TypeMemoryWarning
)

// Event is the root of the runtime event pipeline. Platform backends normalize
// native OS events into this structure before dispatching them to widgets.
type Event struct {
	Type               Type
	Timestamp          time.Duration
	Handled            bool
	PropagationStopped bool
	DefaultPrevented   bool
}

// StopPropagation prevents the event from bubbling further.
func (e *Event) StopPropagation() {
	if e == nil {
		return
	}
	e.PropagationStopped = true
}

// PreventDefault tells the runtime that the default action should be skipped.
func (e *Event) PreventDefault() {
	if e == nil {
		return
	}
	e.DefaultPrevented = true
}

// Point is a 2D coordinate in local or screen space.
type Point struct {
	X float32
	Y float32
}

// Size represents a width and height.
type Size struct {
	Width  float32
	Height float32
}

// Rect is a rectangle in 2D-space.
type Rect struct {
	X      float32
	Y      float32
	Width  float32
	Height float32
}

// Insets is a common UI inset representation used for safe area and keyboard.
type Insets struct {
	Top    float32
	Right  float32
	Bottom float32
	Left   float32
}
