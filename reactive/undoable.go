package reactive

import "sync"

// Undoable wraps a signal with undo and redo history.
type Undoable[T any] struct {
	signal    *Signal[T]
	mu        sync.Mutex
	history   []T
	future    []T
	index     int
	recording bool
}

// NewUndoable creates an undoable wrapper around a signal.
func NewUndoable[T any](signal *Signal[T]) *Undoable[T] {
	if signal == nil {
		signal = NewSignal[T]("", *new(T))
	}
	return &Undoable[T]{signal: signal, history: []T{signal.Get()}, recording: true}
}

// Get returns the current value.
func (u *Undoable[T]) Get() T { return u.signal.Get() }

// Set changes the value and records it in the undo history.
func (u *Undoable[T]) Set(value T) {
	u.mu.Lock()
	if u.recording {
		u.history = append(u.history[:u.index+1], value)
		u.index++
		u.future = nil
	}
	u.mu.Unlock()
	u.signal.Set(value)
}

// Undo restores the previous value and reports whether a change was made.
func (u *Undoable[T]) Undo() bool {
	u.mu.Lock()
	if u.index <= 0 {
		u.mu.Unlock()
		return false
	}
	u.index--
	value := u.history[u.index]
	u.recording = false
	u.mu.Unlock()
	u.signal.Set(value)
	u.mu.Lock()
	u.recording = true
	u.mu.Unlock()
	return true
}

// Redo reapplies the next value and reports whether a change was made.
func (u *Undoable[T]) Redo() bool {
	u.mu.Lock()
	if u.index >= len(u.history)-1 {
		u.mu.Unlock()
		return false
	}
	u.index++
	value := u.history[u.index]
	u.recording = false
	u.mu.Unlock()
	u.signal.Set(value)
	u.mu.Lock()
	u.recording = true
	u.mu.Unlock()
	return true
}

// Signal returns the wrapped signal.
func (u *Undoable[T]) Signal() *Signal[T] { return u.signal }
