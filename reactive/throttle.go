package reactive

import (
	"sync"
	"time"
)

// Throttle creates a signal that emits source values at most once per duration.
func Throttle[T any](source *Signal[T], duration time.Duration) *Signal[T] {
	result := NewSignal("", source.Get())
	var mu sync.Mutex
	var lastTime time.Time
	source.Subscribe(func(value T) {
		mu.Lock()
		defer mu.Unlock()
		now := time.Now()
		if !lastTime.IsZero() && now.Sub(lastTime) < duration {
			return
		}
		lastTime = now
		result.Set(value)
	})
	return result
}
