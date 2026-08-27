package ui

import "sync"

var (
	Signals   = make(map[string]any)
	signalsMu sync.RWMutex
)

type Signal[T any] struct {
	mu sync.RWMutex

	value T

	listeners map[uint64]func(T)
	channels  map[uint64]chan T

	nextID uint64
}

type Subscription struct {
	unsubscribe func()
	once        sync.Once
}

func NewSignal[T any](name string, value T) *Signal[T] {
	s := &Signal[T]{
		value:     value,
		listeners: make(map[uint64]func(T)),
		channels:  make(map[uint64]chan T),
	}

	signalsMu.Lock()
	Signals[name] = s
	signalsMu.Unlock()

	return s
}

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

func (s *Signal[T]) Get() T {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.value
}

func (s *Signal[T]) Set(value T) {
	s.mu.Lock()

	s.value = value

	listeners := make([]func(T), 0, len(s.listeners))
	for _, listener := range s.listeners {
		listeners = append(listeners, listener)
	}

	channels := make([]chan T, 0, len(s.channels))
	for _, ch := range s.channels {
		channels = append(channels, ch)
	}

	s.mu.Unlock()

	for _, listener := range listeners {
		listener(value)
	}

	for _, ch := range channels {
		select {
		case ch <- value:
		default:
		}
	}
}

func (s *Signal[T]) Subscribe(fn func(T)) *Subscription {
	if fn == nil {
		return &Subscription{}
	}

	s.mu.Lock()

	if s.listeners == nil {
		s.listeners = make(map[uint64]func(T))
	}

	id := s.nextID
	s.nextID++

	s.listeners[id] = fn

	s.mu.Unlock()

	return &Subscription{
		unsubscribe: func() {
			s.mu.Lock()
			delete(s.listeners, id)
			s.mu.Unlock()
		},
	}
}

func (s *Signal[T]) SubscribeChan(bufferSize uint) (<-chan T, *Subscription) {
	s.mu.Lock()

	if s.channels == nil {
		s.channels = make(map[uint64]chan T)
	}

	id := s.nextID
	s.nextID++

	ch := make(chan T, bufferSize)

	s.channels[id] = ch

	s.mu.Unlock()

	return ch, &Subscription{
		unsubscribe: func() {
			s.mu.Lock()

			if ch, ok := s.channels[id]; ok {
				delete(s.channels, id)
				close(ch)
			}

			s.mu.Unlock()
		},
	}
}

func (s *Subscription) Unsubscribe() {
	if s == nil || s.unsubscribe == nil {
		return
	}

	s.once.Do(s.unsubscribe)
}
