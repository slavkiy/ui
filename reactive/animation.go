package reactive

import (
	"sync"
	"time"
)

// Easing maps normalized animation progress to an eased progress value.
type Easing func(float64) float64

var (
	// Linear applies no easing.
	Linear Easing = func(t float64) float64 { return t }
	// EaseIn starts slowly and accelerates.
	EaseIn Easing = func(t float64) float64 { return t * t }
	// EaseOut starts quickly and decelerates.
	EaseOut Easing = func(t float64) float64 { return 1 - (1-t)*(1-t) }
	// EaseInOut accelerates at the start and decelerates at the end.
	EaseInOut Easing = func(t float64) float64 {
		if t < 0.5 {
			return 2 * t * t
		}
		return 1 - (-2*t+2)*(-2*t+2)/2
	}
)

// Animate moves a float64 signal to target over duration and returns a cancellable subscription.
func Animate(signal *Signal[float64], target float64, duration time.Duration, easing Easing) *Subscription {
	if easing == nil {
		easing = Linear
	}
	start := signal.Get()
	if duration <= 0 {
		signal.Set(target)
		return &Subscription{}
	}
	ticker := time.NewTicker(time.Millisecond * 16)
	done := make(chan struct{})
	var once sync.Once
	sub := &Subscription{unsubscribe: func() {
		once.Do(func() { close(done); ticker.Stop() })
	}}
	startTime := time.Now()
	go func() {
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				progress := float64(time.Since(startTime)) / float64(duration)
				if progress >= 1 {
					signal.Set(target)
					sub.Unsubscribe()
					return
				}
				signal.Set(start + (target-start)*easing(progress))
			case <-done:
				return
			}
		}
	}()
	return sub
}
