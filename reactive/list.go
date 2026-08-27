package reactive

import "sync"

// List stores an ordered collection and notifies subscribers when it changes.
type List[T any] struct {
	mu        sync.RWMutex
	items     []T
	listeners map[uint64]func([]T)
	nextID    uint64
}

// NewList creates an empty list.
func NewList[T any]() *List[T] {
	return &List[T]{items: make([]T, 0), listeners: make(map[uint64]func([]T))}
}

// Get returns a copy of the current items.
func (l *List[T]) Get() []T {
	l.mu.RLock()
	defer l.mu.RUnlock()
	result := make([]T, len(l.items))
	copy(result, l.items)
	return result
}

// Set replaces all items and notifies subscribers.
func (l *List[T]) Set(items []T) {
	l.mu.Lock()
	l.items = make([]T, len(items))
	copy(l.items, items)
	value := make([]T, len(l.items))
	copy(value, l.items)
	listeners := make([]func([]T), 0, len(l.listeners))
	for _, listener := range l.listeners {
		listeners = append(listeners, listener)
	}
	l.mu.Unlock()
	enqueue(func() {
		for _, listener := range listeners {
			listener(value)
		}
	})
}

// Append adds items to the end of the list and notifies subscribers.
func (l *List[T]) Append(items ...T) {
	l.mu.Lock()
	l.items = append(l.items, items...)
	value := make([]T, len(l.items))
	copy(value, l.items)
	listeners := make([]func([]T), 0, len(l.listeners))
	for _, listener := range l.listeners {
		listeners = append(listeners, listener)
	}
	l.mu.Unlock()
	enqueue(func() {
		for _, listener := range listeners {
			listener(value)
		}
	})
}

// Clear removes all items from the list.
func (l *List[T]) Clear() { l.Set(nil) }

// Subscribe registers a callback that receives a copy of the items after changes.
func (l *List[T]) Subscribe(fn func([]T)) *Subscription {
	if fn == nil {
		return &Subscription{}
	}
	l.mu.Lock()
	id := l.nextID
	l.nextID++
	l.listeners[id] = fn
	l.mu.Unlock()
	return &Subscription{unsubscribe: func() {
		l.mu.Lock()
		delete(l.listeners, id)
		l.mu.Unlock()
	}}
}
