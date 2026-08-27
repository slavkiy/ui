package reactive

import "sync"

var (
	// Signals stores signals registered by name.
	Signals   = make(map[string]any)
	signalsMu sync.RWMutex
)

// GetSignal returns the named signal when it has the requested type.
func GetSignal[T any](name string) (*Signal[T], bool) {
	signalsMu.RLock()
	defer signalsMu.RUnlock()
	value, ok := Signals[name]
	if !ok {
		return nil, false
	}
	signal, ok := value.(*Signal[T])
	return signal, ok
}

// Signal stores a mutable value and notifies subscribers when it changes.
type Signal[T any] struct {
	mu sync.RWMutex

	value T

	listeners map[uint64]func(T)
	channels  map[uint64]chan T
	effects   map[uint64]func(T)

	nextID uint64
}

// NewSignal creates a signal with an initial value and registers it by name.
func NewSignal[T any](name string, value T) *Signal[T] {
	s := &Signal[T]{value: value, listeners: make(map[uint64]func(T)), channels: make(map[uint64]chan T), effects: make(map[uint64]func(T))}
	signalsMu.Lock()
	Signals[name] = s
	signalsMu.Unlock()
	return s
}

// Get returns the signal's current value.
func (s *Signal[T]) Get() T {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.value
}

// Set replaces the signal's value and notifies its subscribers.
func (s *Signal[T]) Set(value T) {
	s.mu.Lock()
	s.value = value
	listeners := make([]func(T), 0, len(s.listeners))
	for _, listener := range s.listeners {
		listeners = append(listeners, listener)
	}
	effects := make([]func(T), 0, len(s.effects))
	for _, effect := range s.effects {
		effects = append(effects, effect)
	}
	channels := make([]chan T, 0, len(s.channels))
	for _, ch := range s.channels {
		channels = append(channels, ch)
	}
	s.mu.Unlock()

	enqueue(func() {
		for _, effect := range effects {
			effect(value)
		}
		for _, listener := range listeners {
			listener(value)
		}
		for _, ch := range channels {
			select {
			case ch <- value:
			default:
			}
		}
	})
}

// Update replaces the signal's value with the result of applying fn to it.
func (s *Signal[T]) Update(fn func(T) T) {
	if fn == nil {
		return
	}
	s.Set(fn(s.Get()))
}

// Subscribe registers a callback that receives each new signal value.
func (s *Signal[T]) Subscribe(fn func(T)) *Subscription {
	if fn == nil {
		return &Subscription{}
	}
	s.mu.Lock()
	id := s.nextID
	s.nextID++
	s.listeners[id] = fn
	s.mu.Unlock()
	return &Subscription{unsubscribe: func() {
		s.mu.Lock()
		delete(s.listeners, id)
		s.mu.Unlock()
	}}
}

// SubscribeChan returns a channel that receives signal updates and its subscription.
func (s *Signal[T]) SubscribeChan(bufferSize uint) (<-chan T, *Subscription) {
	s.mu.Lock()
	id := s.nextID
	s.nextID++
	ch := make(chan T, bufferSize)
	s.channels[id] = ch
	s.mu.Unlock()
	return ch, &Subscription{unsubscribe: func() {
		s.mu.Lock()
		if ch, ok := s.channels[id]; ok {
			delete(s.channels, id)
			close(ch)
		}
		s.mu.Unlock()
	}}
}
