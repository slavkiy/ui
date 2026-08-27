package reactive

import "sync"

// Event broadcasts transient values to its subscribers.
type Event[T any] struct {
	mu        sync.RWMutex
	listeners map[uint64]func(T)
	nextID    uint64
}

// NewEvent creates an empty event stream.
func NewEvent[T any]() *Event[T] { return &Event[T]{listeners: make(map[uint64]func(T))} }

// Emit sends a value to all current subscribers.
func (e *Event[T]) Emit(value T) {
	e.mu.RLock()
	listeners := make([]func(T), 0, len(e.listeners))
	for _, listener := range e.listeners {
		listeners = append(listeners, listener)
	}
	e.mu.RUnlock()
	enqueue(func() {
		for _, listener := range listeners {
			listener(value)
		}
	})
}

// Subscribe registers a callback for emitted values.
func (e *Event[T]) Subscribe(fn func(T)) *Subscription {
	if fn == nil {
		return &Subscription{}
	}
	e.mu.Lock()
	id := e.nextID
	e.nextID++
	e.listeners[id] = fn
	e.mu.Unlock()
	return &Subscription{unsubscribe: func() {
		e.mu.Lock()
		delete(e.listeners, id)
		e.mu.Unlock()
	}}
}
