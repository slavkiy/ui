package reactive

import (
	"reflect"
	"strings"
	"sync"
)

var (
	// Signals stores signals registered by package-qualified name and runtime type.
	signals   = make(map[string]any)
	signalsMu sync.RWMutex
)

func normalizeScope(scope string) string {
	scope = strings.TrimSpace(scope)
	scope = strings.ReplaceAll(scope, "\\", "/")
	scope = strings.Trim(scope, "/")
	scope = strings.TrimPrefix(scope, ".")
	scope = strings.TrimSuffix(scope, ".")
	return scope
}

func signalKey(scope, name string) string {
	if name == "" {
		return ""
	}
	scope = normalizeScope(scope)
	if scope == "" {
		return name
	}
	return scope + "::" + name
}

func signalTypeKey(scope, name string, t any) string {
	key := signalKey(scope, name)
	if key == "" {
		return ""
	}
	return key + ":" + reflect.TypeOf(t).String()
}

// GetSignal returns the named signal when it has the requested type.
func GetSignal[T any](name string) *Signal[T] {
	return GetSignalInScope[T]("", name)
}

// GetSignalInScope returns a signal by explicit scope, e.g. app/internal::name.
func GetSignalInScope[T any](scope, name string) *Signal[T] {
	var zero T
	key := signalTypeKey(scope, name, zero)
	signalsMu.RLock()
	defer signalsMu.RUnlock()
	value, ok := signals[key]
	if !ok {
		return nil
	}
	signal, ok := value.(*Signal[T])
	if !ok {
		return nil
	}
	return signal
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
	return NewSignalInScope[T]("", name, value)
}

// NewSignalInScope creates a signal in an explicit scope such as app/internal::name.
func NewSignalInScope[T any](scope, name string, value T) *Signal[T] {
	s := &Signal[T]{value: value, listeners: make(map[uint64]func(T)), channels: make(map[uint64]chan T), effects: make(map[uint64]func(T))}
	if name != "" {
		var zero T
		key := signalTypeKey(scope, name, zero)
		signalsMu.Lock()
		if _, exists := signals[key]; exists {
			signalsMu.Unlock()
			panic("reactive: duplicate signal registration for same scope/name/type: " + key)
		}
		signals[key] = s
		signalsMu.Unlock()
	}
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
	if s == nil {
		return
	}

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
	if s == nil || fn == nil {
		return
	}
	current := s.Get()
	s.Set(fn(current))
}

// Subscribe registers a callback that receives each new signal value.
func (s *Signal[T]) Subscribe(fn func(T)) *Subscription {
	if s == nil || fn == nil {
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
	if s == nil {
		return nil, &Subscription{}
	}
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
