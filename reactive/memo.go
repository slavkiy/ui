package reactive

import "sync"

// Memo computes a value lazily and caches the result after the first read.
type Memo[T any] struct {
	value   T
	once    sync.Once
	compute func() T
}

// NewMemo creates a lazy memoized computation.
func NewMemo[T any](compute func() T) *Memo[T] { return &Memo[T]{compute: compute} }

// Get computes the value once and returns the cached result.
func (m *Memo[T]) Get() T {
	m.once.Do(func() {
		if m.compute != nil {
			m.value = m.compute()
		}
	})
	return m.value
}
