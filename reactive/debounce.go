package reactive

import (
	"sync"
	"time"
)

// Debounce creates a signal that emits the latest source value after a quiet period.
func Debounce[T any](source *Signal[T], duration time.Duration) *Signal[T] {
	if source == nil {
		return NewSignal[T]("", *new(T))
	}
	result := NewSignal("", source.Get())
	var mu sync.Mutex
	var timer *time.Timer
	source.Subscribe(func(value T) {
		mu.Lock()
		defer mu.Unlock()
		if timer != nil {
			timer.Stop()
		}
		timer = time.AfterFunc(duration, func() { result.Set(value) })
	})
	return result
}
