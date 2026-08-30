package event

import "time"

// Direction describes the general direction of a gesture.
type Direction uint8

const (
	DirectionNone Direction = iota
	DirectionLeft
	DirectionRight
	DirectionUp
	DirectionDown
)

// TapEvent is emitted for a tap or click gesture.
type TapEvent struct {
	Event
	Position  Point
	PointerID uint64
	Count     int
	Modifiers Modifiers
}

// DoubleTapEvent is emitted for a double-tap gesture.
type DoubleTapEvent struct {
	TapEvent
}

// LongPressEvent is emitted for a long press gesture.
type LongPressEvent struct {
	Event
	Position  Point
	PointerID uint64
	Duration  time.Duration
	Modifiers Modifiers
}

// SwipeEvent represents a horizontal or vertical swipe.
type SwipeEvent struct {
	Event
	Start     Point
	End       Point
	Delta     Point
	Velocity  Point
	Direction Direction
	Modifiers Modifiers
}

// PinchEvent represents a two-finger pinch or zoom gesture.
type PinchEvent struct {
	Event
	Center    Point
	Scale     float32
	Velocity  float32
	Modifiers Modifiers
}

// RotateEvent represents a two-finger rotation gesture.
type RotateEvent struct {
	Event
	Center    Point
	Rotation  float32
	Velocity  float32
	Modifiers Modifiers
}

// PanEvent is emitted for drag-like pan motion.
type PanEvent struct {
	Event
	Delta     Point
	Velocity  Point
	Modifiers Modifiers
}

// DragEvent represents a drag gesture with raw pointer movement.
type DragEvent struct {
	Event
	Start     Point
	Current   Point
	Delta     Point
	PointerID uint64
	Modifiers Modifiers
}
