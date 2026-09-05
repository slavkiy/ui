package marker

import (
	"context"

	. "github.com/slavkiy/ui/reactive"
)

var curWin = NewSignal[WindowSize]("ui/marker::Window", WindowSize{})

// WindowSignal is the reactive signal storing the current window state.
var WindowSignal = GetClientSignal[WindowSize]("ui/marker::Window")

// Current returns the current window snapshot.
func Current() WindowSize {
	if WindowSignal == nil {
		return WindowSize{}
	}
	return WindowSignal.Get()
}

// SetCurrent updates the current window state and notifies subscribers.
func SetCurrent(win WindowSize) {
	if curWin == nil {
		return
	}
	curWin.Set(win)
}

// Subscribe registers a callback for every window change.
func Subscribe(fn func(WindowSize)) *Subscription {
	if curWin == nil || fn == nil {
		return &Subscription{}
	}
	return curWin.Subscribe(fn)
}

// Channel streams current window updates until the context is done.
func Channel(ctx context.Context, bufferSize uint) <-chan WindowSize {
	if curWin == nil {
		ch := make(chan WindowSize)
		close(ch)
		return ch
	}
	return curWin.Channel(ctx, bufferSize)
}

// UpdateWindow updates the current window state using the last-known state.
func UpdateWindow(win WindowSize) {
	SetCurrent(win)
}
