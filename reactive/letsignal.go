package reactive

import "context"

// ServerSignal is the write-only side of a reactive channel.
// Only the server can send values into it.
type ServerSignal[T any] struct {
	signal *Signal[T]
}

// ClientSignal is the read-only side of the same value.
// Clients can subscribe to updates and read the latest value, but cannot write.
type ClientSignal[T any] struct {
	signal *Signal[T]
}

// NewServerSignal creates a server-side signal for push updates.
func NewServerSignal[T any](name string, value T) *ServerSignal[T] {
	return &ServerSignal[T]{signal: NewSignal[T](name, value)}
}

// NewClientSignal creates a client-side view for a server signal.
func NewClientSignal[T any](server *ServerSignal[T]) *ClientSignal[T] {
	if server == nil || server.signal == nil {
		return &ClientSignal[T]{signal: NewSignal[T]("", zeroValue[T]())}
	}
	return &ClientSignal[T]{signal: server.signal}
}

// GetClientSignal returns an existing signal in the global registry as a read-only client view.
func GetClientSignal[T any](name string) *ClientSignal[T] {
	if name == "" {
		return &ClientSignal[T]{signal: NewSignal[T]("", zeroValue[T]())}
	}
	if s := GetSignal[T](name); s != nil {
		return &ClientSignal[T]{signal: s}
	}
	return &ClientSignal[T]{signal: NewSignal[T](name, zeroValue[T]())}
}

func zeroValue[T any]() T {
	var v T
	return v
}

// Set sends a new value from the server.
func (s *ServerSignal[T]) Set(value T) {
	if s == nil || s.signal == nil {
		return
	}
	s.signal.Set(value)
}

// Get returns the current value from the server side.
func (s *ServerSignal[T]) Get() T {
	if s == nil || s.signal == nil {
		return zeroValue[T]()
	}
	return s.signal.Get()
}

// Subscribe registers a consumer on the client side.
func (c *ClientSignal[T]) Subscribe(fn func(T)) *Subscription {
	if c == nil || c.signal == nil || fn == nil {
		return &Subscription{}
	}
	return c.signal.Subscribe(fn)
}

// SubscribeChan registers a client listener on a channel.
func (c *ClientSignal[T]) SubscribeChan(bufferSize uint) (<-chan T, *Subscription) {
	if c == nil || c.signal == nil {
		return nil, &Subscription{}
	}
	return c.signal.SubscribeChan(bufferSize)
}

// Channel gives clients a cancelable channel stream of values.
func (c *ClientSignal[T]) Channel(ctx context.Context, bufferSize uint) <-chan T {
	if c == nil || c.signal == nil {
		ch := make(chan T)
		close(ch)
		return ch
	}
	return c.signal.Channel(ctx, bufferSize)
}

// Get returns the current value from the client side.
func (c *ClientSignal[T]) Get() T {
	if c == nil || c.signal == nil {
		return zeroValue[T]()
	}
	return c.signal.Get()
}

// Server returns the server-side signal from a client view.
func (c *ClientSignal[T]) Server() *ServerSignal[T] {
	if c == nil || c.signal == nil {
		return nil
	}
	return &ServerSignal[T]{signal: c.signal}
}
