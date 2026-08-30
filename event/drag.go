package event

// DragKind describes whether a drag is internal or external to the app.
type DragKind uint8

const (
	DragKindUnknown DragKind = iota
	DragKindInternal
	DragKindExternal
)

// DragEventData is an opaque payload for the drag operation.
type DragEventData interface{}

// DragEnter contains a drag payload entering a target.
type DragEnter struct {
	Event
	Position  Point
	Kind      DragKind
	Data      DragEventData
	Modifiers Modifiers
}

// DragOver describes drag motion over a target.
type DragOver struct {
	Event
	Position  Point
	Kind      DragKind
	Data      DragEventData
	Modifiers Modifiers
}

// DragLeave indicates a drag has left a target.
type DragLeave struct {
	Event
	Position  Point
	Kind      DragKind
	Data      DragEventData
	Modifiers Modifiers
}

// Drop indicates a drag payload was dropped.
type Drop struct {
	Event
	Position  Point
	Kind      DragKind
	Data      DragEventData
	Modifiers Modifiers
}
