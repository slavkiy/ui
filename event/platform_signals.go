package event

import "github.com/slavkiy/ui/reactive"

var (
	TapSignal           = reactive.GetClientSignal[*TapEvent]("ui/event::Tap")
	DoubleTapSignal     = reactive.GetClientSignal[*DoubleTapEvent]("ui/event::DoubleTap")
	LongPressSignal     = reactive.GetClientSignal[*LongPressEvent]("ui/event::LongPress")
	SwipeSignal         = reactive.GetClientSignal[*SwipeEvent]("ui/event::Swipe")
	PinchSignal         = reactive.GetClientSignal[*PinchEvent]("ui/event::Pinch")
	RotateSignal        = reactive.GetClientSignal[*RotateEvent]("ui/event::Rotate")
	PanSignal           = reactive.GetClientSignal[*PanEvent]("ui/event::Pan")
	DragSignal          = reactive.GetClientSignal[*DragEvent]("ui/event::Drag")
	ClipboardSignal     = reactive.GetClientSignal[*ClipboardChanged]("ui/event::Clipboard")
	GamepadSignal       = reactive.GetClientSignal[*GamepadConnected]("ui/event::Gamepad")
	PenSignal           = reactive.GetClientSignal[*PenDown]("ui/event::Pen")
	AccessibilitySignal = reactive.GetClientSignal[*AccessibilityEvent]("ui/event::Accessibility")
)
