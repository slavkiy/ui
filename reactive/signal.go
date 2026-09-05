package reactive

import (
	"context"
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

func signalTypeOf[T any]() reflect.Type {
	var zero T
	if reflect.TypeOf(zero) != nil {
		return reflect.TypeOf(zero)
	}
	return reflect.TypeOf((*T)(nil)).Elem()
}

func signalTypeKey(scope, name string, t any) string {
	key := signalKey(scope, name)
	if key == "" {
		return ""
	}
	if t == nil {
		return key + ":<nil>"
	}
	if typ, ok := t.(reflect.Type); ok {
		if typ == nil {
			return key + ":<nil>"
		}
		return key + ":" + typ.String()
	}
	typ := reflect.TypeOf(t)
	if typ == nil {
		return key + ":<nil>"
	}
	return key + ":" + typ.String()
}

// GetSignal returns the named signal when it has the requested type.
func GetSignal[T any](name string) *Signal[T] {
	return GetSignalInScope[T]("", name)
}

// GetSignalInScope returns a signal by explicit scope, e.g. app/internal::name.
func GetSignalInScope[T any](scope, name string) *Signal[T] {
	key := signalTypeKey(scope, name, signalTypeOf[T]())
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
func NewSignal[T any](name string, value ...T) *Signal[T] {
	var v T
	if len(value) == 1 {
		v = value[0]
	}
	if len(value) > 1 {
		panic("reactive: NewSignal expects at most one initial value")
	}
	return NewSignalInScope[T]("", name, v)
}

// NewSignalInScope creates a signal in an explicit scope such as app/internal::name.
func NewSignalInScope[T any](scope, name string, value T) *Signal[T] {
	s := &Signal[T]{value: value, listeners: make(map[uint64]func(T)), channels: make(map[uint64]chan T), effects: make(map[uint64]func(T))}
	if name != "" {
		key := signalTypeKey(scope, name, signalTypeOf[T]())
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

// Channel creates a context-bound channel that streams signal values until the context ends.
func (s *Signal[T]) Channel(ctx context.Context, bufferSize uint) <-chan T {
	if s == nil {
		ch := make(chan T)
		close(ch)
		return ch
	}
	if ctx == nil {
		ctx = context.Background()
	}
	out := make(chan T, bufferSize)
	sub := s.Subscribe(func(value T) {
		select {
		case <-ctx.Done():
			return
		default:
			select {
			case out <- value:
			default:
				select {
				case <-ctx.Done():
					return
				default:
					<-out
					out <- value
				}
			}
		}
	})
	go func() {
		<-ctx.Done()
		sub.Unsubscribe()
		close(out)
	}()
	return out
}
