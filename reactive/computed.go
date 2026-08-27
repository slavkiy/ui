package reactive

import "sync"

// Computed stores a value derived from explicitly supplied dependencies.
type Computed[T any] struct {
	mu           sync.RWMutex
	value        T
	compute      func() T
	listeners    map[uint64]func(T)
	dependencies []*Subscription
	nextID       uint64
}

// Dependency is a value that can notify a computed value about changes.
type Dependency interface{ SubscribeEffect(func()) *Subscription }

// NewComputed creates a computed value and subscribes it to its dependencies.
func NewComputed[T any](compute func() T, dependencies ...Dependency) *Computed[T] {
	c := &Computed[T]{compute: compute, value: compute(), listeners: make(map[uint64]func(T))}
	for _, dependency := range dependencies {
		if dependency != nil {
			c.dependencies = append(c.dependencies, dependency.SubscribeEffect(c.recompute))
		}
	}
	return c
}

// Get returns the computed value.
func (c *Computed[T]) Get() T {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.value
}

func (c *Computed[T]) recompute() {
	value := c.compute()
	c.mu.Lock()
	c.value = value
	listeners := make([]func(T), 0, len(c.listeners))
	for _, listener := range c.listeners {
		listeners = append(listeners, listener)
	}
	c.mu.Unlock()
	enqueue(func() {
		for _, listener := range listeners {
			listener(value)
		}
	})
}

// Subscribe registers a callback that receives each recomputed value.
func (c *Computed[T]) Subscribe(fn func(T)) *Subscription {
	if fn == nil {
		return &Subscription{}
	}
	c.mu.Lock()
	id := c.nextID
	c.nextID++
	c.listeners[id] = fn
	c.mu.Unlock()
	return &Subscription{unsubscribe: func() {
		c.mu.Lock()
		delete(c.listeners, id)
		c.mu.Unlock()
	}}
}

// SubscribeEffect registers a callback that runs after recomputation.
func (c *Computed[T]) SubscribeEffect(fn func()) *Subscription {
	return c.Subscribe(func(T) { fn() })
}
