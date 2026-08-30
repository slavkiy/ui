package event

import (
	"time"

	"github.com/slavkiy/ui/reactive"
)

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

type Events struct {
	KeyDown     *reactive.ClientSignal[*KeyDown]
	KeyUp       *reactive.ClientSignal[*KeyUp]
	MouseMove   *reactive.ClientSignal[*PointerMove]
	MouseDown   *reactive.ClientSignal[*PointerDown]
	MouseUp     *reactive.ClientSignal[*PointerUp]
	PointerDown *reactive.ClientSignal[*PointerDown]
	PointerMove *reactive.ClientSignal[*PointerMove]
	PointerUp   *reactive.ClientSignal[*PointerUp]
	TouchStart  *reactive.ClientSignal[*TouchStart]
	TouchMove   *reactive.ClientSignal[*TouchMove]
	TouchEnd    *reactive.ClientSignal[*TouchEnd]
	Scroll      *reactive.ClientSignal[*ScrollEvent]
	Resize      *reactive.ClientSignal[*WindowResized]
	Window      *reactive.ClientSignal[*WindowEvent]
	Focus       *reactive.ClientSignal[*FocusEvent]
	TextInput   *reactive.ClientSignal[*TextInputCommitted]
}

func NewEvents() *Events {
	return &Events{
		KeyDown:     reactive.GetClientSignal[*KeyDown]("ui/event::KeyDown"),
		KeyUp:       reactive.GetClientSignal[*KeyUp]("ui/event::KeyUp"),
		MouseMove:   reactive.GetClientSignal[*PointerMove]("ui/event::PointerMove"),
		MouseDown:   reactive.GetClientSignal[*PointerDown]("ui/event::PointerDown"),
		MouseUp:     reactive.GetClientSignal[*PointerUp]("ui/event::PointerUp"),
		PointerDown: reactive.GetClientSignal[*PointerDown]("ui/event::PointerDown"),
		PointerMove: reactive.GetClientSignal[*PointerMove]("ui/event::PointerMove"),
		PointerUp:   reactive.GetClientSignal[*PointerUp]("ui/event::PointerUp"),
		TouchStart:  reactive.GetClientSignal[*TouchStart]("ui/event::TouchStart"),
		TouchMove:   reactive.GetClientSignal[*TouchMove]("ui/event::TouchMove"),
		TouchEnd:    reactive.GetClientSignal[*TouchEnd]("ui/event::TouchEnd"),
		Scroll:      reactive.GetClientSignal[*ScrollEvent]("ui/event::Scroll"),
		Resize:      reactive.GetClientSignal[*WindowResized]("ui/event::Window"),
		Window:      reactive.GetClientSignal[*WindowEvent]("ui/event::Window"),
		Focus:       reactive.GetClientSignal[*FocusEvent]("ui/event::Focus"),
		TextInput:   reactive.GetClientSignal[*TextInputCommitted]("ui/event::TextInput"),
	}
}

var (
	// CurrentEvent stores the last event dispatched by the runtime.
	CurrentEvent = reactive.NewSignal[*Event]("ui/event::CurrentEvent", nil)
	// eventBus is the transient event stream used by the runtime to broadcast new events.
	eventBus = reactive.NewEvent[*Event]()
)

// CurrentEventClient exposes the current event as a client-only signal.
var CurrentEventClient = reactive.GetClientSignal[*Event]("ui/event::CurrentEvent")

func init() {
	observer()
}

// Emit sends an event into the runtime event stream.
func Emit(evt *Event) {
	if evt == nil {
		return
	}
	eventBus.Emit(evt)
}

func observer() {
	eventBus.Subscribe(func(evt *Event) {
		if evt == nil {
			return
		}
		CurrentEvent.Set(evt)
	})
}

// KeyDownSignal exposes key-down events to clients as a subscription-only signal.
var KeyDownSignal = reactive.GetClientSignal[*KeyDown]("ui/event::KeyDown")

// KeyUpSignal exposes key-up events to clients as a subscription-only signal.
var KeyUpSignal = reactive.GetClientSignal[*KeyUp]("ui/event::KeyUp")

// PointerMoveSignal exposes pointer movement events.
var PointerMoveSignal = reactive.GetClientSignal[*PointerMove]("ui/event::PointerMove")

// PointerDownSignal exposes pointer down events.
var PointerDownSignal = reactive.GetClientSignal[*PointerDown]("ui/event::PointerDown")

// PointerUpSignal exposes pointer up events.
var PointerUpSignal = reactive.GetClientSignal[*PointerUp]("ui/event::PointerUp")

// TouchMoveSignal exposes touch move events.
var TouchMoveSignal = reactive.GetClientSignal[*TouchMove]("ui/event::TouchMove")

// TouchStartSignal exposes touch start events.
var TouchStartSignal = reactive.GetClientSignal[*TouchStart]("ui/event::TouchStart")

// TouchEndSignal exposes touch end events.
var TouchEndSignal = reactive.GetClientSignal[*TouchEnd]("ui/event::TouchEnd")

// ScrollSignal exposes scroll events.
var ScrollSignal = reactive.GetClientSignal[*ScrollEvent]("ui/event::Scroll")

// WindowSignal exposes window lifecycle events.
var WindowSignal = reactive.GetClientSignal[*WindowEvent]("ui/event::Window")

// FocusSignal exposes focus events.
var FocusSignal = reactive.GetClientSignal[*FocusEvent]("ui/event::Focus")

// TextInputSignal exposes text input events.
var TextInputSignal = reactive.GetClientSignal[*TextInputCommitted]("ui/event::TextInput")

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
